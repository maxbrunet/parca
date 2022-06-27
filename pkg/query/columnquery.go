// Copyright 2022 The Parca Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package query

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/apache/arrow/go/v8/arrow"
	"github.com/apache/arrow/go/v8/arrow/array"
	"github.com/go-kit/log"
	"github.com/polarsignals/frostdb/query"
	"github.com/polarsignals/frostdb/query/logicalplan"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/timestamp"
	"github.com/prometheus/prometheus/promql/parser"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	metastorepb "github.com/parca-dev/parca/gen/proto/go/parca/metastore/v1alpha1"
	profilestorepb "github.com/parca-dev/parca/gen/proto/go/parca/profilestore/v1alpha1"
	pb "github.com/parca-dev/parca/gen/proto/go/parca/query/v1alpha1"
	"github.com/parca-dev/parca/pkg/parcacol"
	"github.com/parca-dev/parca/pkg/profile"
)

type Engine interface {
	ScanTable(name string) query.Builder
	ScanSchema(name string) query.Builder
}

// ColumnQueryAPI is the read api interface for parca
// It implements the proto/query/query.proto APIServer interface.
type ColumnQueryAPI struct {
	pb.UnimplementedQueryServiceServer

	logger    log.Logger
	tracer    trace.Tracer
	engine    Engine
	tableName string
	metastore metastorepb.MetastoreServiceClient
}

func NewColumnQueryAPI(
	logger log.Logger,
	tracer trace.Tracer,
	metastore metastorepb.MetastoreServiceClient,
	engine Engine,
	tableName string,
) *ColumnQueryAPI {
	return &ColumnQueryAPI{
		logger:    logger,
		tracer:    tracer,
		engine:    engine,
		tableName: tableName,
		metastore: metastore,
	}
}

// Labels issues a labels request against the storage.
func (q *ColumnQueryAPI) Labels(ctx context.Context, req *pb.LabelsRequest) (*pb.LabelsResponse, error) {
	seen := map[string]struct{}{}

	err := q.engine.ScanSchema(q.tableName).
		Distinct(logicalplan.Col("name")).
		Filter(logicalplan.Col("name").RegexMatch("^labels\\..+$")).
		Execute(ctx, func(ar arrow.Record) error {
			if ar.NumCols() != 1 {
				return fmt.Errorf("expected 1 column, got %d", ar.NumCols())
			}

			col := ar.Column(0)
			stringCol, ok := col.(*array.String)
			if !ok {
				return fmt.Errorf("expected string column, got %T", col)
			}

			for i := 0; i < stringCol.Len(); i++ {
				val := stringCol.Value(i)
				seen[strings.TrimPrefix(val, "labels.")] = struct{}{}
			}

			return nil
		})
	if err != nil {
		return nil, err
	}

	vals := make([]string, 0, len(seen))
	for val := range seen {
		vals = append(vals, val)
	}

	sort.Strings(vals)

	return &pb.LabelsResponse{
		LabelNames: vals,
	}, nil
}

// Values issues a values request against the storage.
func (q *ColumnQueryAPI) Values(ctx context.Context, req *pb.ValuesRequest) (*pb.ValuesResponse, error) {
	name := req.LabelName
	vals := []string{}

	err := q.engine.ScanTable(q.tableName).
		Distinct(logicalplan.Col("labels."+name)).
		Execute(ctx, func(ar arrow.Record) error {
			if ar.NumCols() != 1 {
				return fmt.Errorf("expected 1 column, got %d", ar.NumCols())
			}

			col := ar.Column(0)
			stringCol, ok := col.(*array.Binary)
			if !ok {
				return fmt.Errorf("expected string column, got %T", col)
			}

			for i := 0; i < stringCol.Len(); i++ {
				val := stringCol.Value(i)
				vals = append(vals, string(val))
			}

			return nil
		})
	if err != nil {
		return nil, err
	}

	sort.Strings(vals)

	return &pb.ValuesResponse{
		LabelValues: vals,
	}, nil
}

func matcherToBooleanExpression(matcher *labels.Matcher) (logicalplan.Expr, error) {
	ref := logicalplan.Col("labels." + matcher.Name)
	switch matcher.Type {
	case labels.MatchEqual:
		return ref.Eq(logicalplan.Literal(matcher.Value)), nil
	case labels.MatchNotEqual:
		return ref.NotEq(logicalplan.Literal(matcher.Value)), nil
	case labels.MatchRegexp:
		return ref.RegexMatch(matcher.Value), nil
	case labels.MatchNotRegexp:
		return ref.RegexNotMatch(matcher.Value), nil
	default:
		return nil, fmt.Errorf("unsupported matcher type %v", matcher.Type.String())
	}
}

func matchersToBooleanExpressions(matchers []*labels.Matcher) ([]logicalplan.Expr, error) {
	exprs := make([]logicalplan.Expr, 0, len(matchers))

	for _, matcher := range matchers {
		expr, err := matcherToBooleanExpression(matcher)
		if err != nil {
			return nil, err
		}

		exprs = append(exprs, expr)
	}

	return exprs, nil
}

var (
	ErrTimestampColumnNotFound = errors.New("timestamp column not found")
	ErrValueColumnNotFound     = errors.New("value column not found")
)

func queryToFilterExprs(query string) (profile.Meta, []logicalplan.Expr, error) {
	parsedSelector, err := parser.ParseMetricSelector(query)
	if err != nil {
		return profile.Meta{}, nil, status.Error(codes.InvalidArgument, "failed to parse query")
	}

	sel := make([]*labels.Matcher, 0, len(parsedSelector))
	var nameLabel *labels.Matcher
	for _, matcher := range parsedSelector {
		if matcher.Name == labels.MetricName {
			nameLabel = matcher
		} else {
			sel = append(sel, matcher)
		}
	}
	if nameLabel == nil {
		return profile.Meta{}, nil, status.Error(codes.InvalidArgument, "query must contain a profile-type selection")
	}

	parts := strings.Split(nameLabel.Value, ":")
	if len(parts) != 5 && len(parts) != 6 {
		return profile.Meta{}, nil, status.Errorf(codes.InvalidArgument, "profile-type selection must be of the form <name>:<sample-type>:<sample-unit>:<period-type>:<period-unit>(:delta), got(%d): %q", len(parts), nameLabel.Value)
	}
	name, sampleType, sampleUnit, periodType, periodUnit, delta := parts[0], parts[1], parts[2], parts[3], parts[4], false
	if len(parts) == 6 && parts[5] == "delta" {
		delta = true
	}

	labelFilterExpressions, err := matchersToBooleanExpressions(sel)
	if err != nil {
		return profile.Meta{}, nil, status.Error(codes.InvalidArgument, "failed to build query")
	}

	exprs := append([]logicalplan.Expr{
		logicalplan.Col("name").Eq(logicalplan.Literal(name)),
		logicalplan.Col("sample_type").Eq(logicalplan.Literal(sampleType)),
		logicalplan.Col("sample_unit").Eq(logicalplan.Literal(sampleUnit)),
		logicalplan.Col("period_type").Eq(logicalplan.Literal(periodType)),
		logicalplan.Col("period_unit").Eq(logicalplan.Literal(periodUnit)),
	}, labelFilterExpressions...)

	deltaPlan := logicalplan.Col("duration").Eq(logicalplan.Literal(0))
	if delta {
		deltaPlan = logicalplan.Col("duration").NotEq(logicalplan.Literal(0))
	}

	exprs = append(exprs, deltaPlan)

	return profile.Meta{
		Name:       name,
		SampleType: profile.ValueType{Type: sampleType, Unit: sampleUnit},
		PeriodType: profile.ValueType{Type: periodType, Unit: periodUnit},
	}, exprs, nil
}

// QueryRange issues a range query against the storage.
func (q *ColumnQueryAPI) QueryRange(ctx context.Context, req *pb.QueryRangeRequest) (*pb.QueryRangeResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, selectorExprs, err := queryToFilterExprs(req.Query)
	if err != nil {
		return nil, err
	}

	start := timestamp.FromTime(req.Start.AsTime())
	end := timestamp.FromTime(req.End.AsTime())

	exprs := append(
		selectorExprs,
		logicalplan.Col("timestamp").GT(logicalplan.Literal(start)),
		logicalplan.Col("timestamp").LT(logicalplan.Literal(end)),
	)

	filterExpr := logicalplan.And(exprs...)

	res := &pb.QueryRangeResponse{}
	labelsetToIndex := map[string]int{}

	labelSet := labels.Labels{}

	var ar arrow.Record
	err = q.engine.ScanTable(q.tableName).
		Filter(filterExpr).
		Aggregate(
			logicalplan.Sum(logicalplan.Col("value")),
			logicalplan.DynCol("labels"),
			logicalplan.Col("timestamp"),
		).
		Execute(ctx, func(r arrow.Record) error {
			r.Retain()
			ar = r
			return nil
		})
	if err != nil {
		return nil, err
	}
	if ar.NumRows() == 0 {
		return nil, status.Error(
			codes.NotFound,
			"No data found for the query, try a different query or time range or no data has been written to be queried yet.",
		)
	}

	timestampColumnIndex := 0
	timestampColumnFound := false
	valueColumnIndex := 0
	valueColumnFound := false
	labelColumnIndices := []int{}

	fields := ar.Schema().Fields()
	for i, field := range fields {
		if field.Name == "timestamp" {
			timestampColumnIndex = i
			timestampColumnFound = true
			continue
		}
		if field.Name == "sum(value)" {
			valueColumnIndex = i
			valueColumnFound = true
			continue
		}

		if strings.HasPrefix(field.Name, "labels.") {
			labelColumnIndices = append(labelColumnIndices, i)
		}
	}

	if !timestampColumnFound {
		return nil, ErrTimestampColumnNotFound
	}

	if !valueColumnFound {
		return nil, ErrValueColumnNotFound
	}

	for i := 0; i < int(ar.NumRows()); i++ {
		labelSet = labelSet[:0]
		for _, labelColumnIndex := range labelColumnIndices {
			col := ar.Column(labelColumnIndex).(*array.Binary)
			if col.IsNull(i) {
				continue
			}

			v := col.Value(i)
			if len(v) > 0 {
				labelSet = append(labelSet, labels.Label{Name: strings.TrimPrefix(fields[labelColumnIndex].Name, "labels."), Value: string(v)})
			}
		}

		sort.Sort(labelSet)
		s := labelSet.String()
		index, ok := labelsetToIndex[s]
		if !ok {
			pbLabelSet := make([]*profilestorepb.Label, 0, len(labelSet))
			for _, l := range labelSet {
				pbLabelSet = append(pbLabelSet, &profilestorepb.Label{
					Name:  l.Name,
					Value: l.Value,
				})
			}
			res.Series = append(res.Series, &pb.MetricsSeries{Labelset: &profilestorepb.LabelSet{Labels: pbLabelSet}})
			index = len(res.Series) - 1
			labelsetToIndex[s] = index
		}

		series := res.Series[index]
		series.Samples = append(series.Samples, &pb.MetricsSample{
			Timestamp: timestamppb.New(timestamp.Time(ar.Column(timestampColumnIndex).(*array.Int64).Value(i))),
			Value:     ar.Column(valueColumnIndex).(*array.Int64).Value(i),
		})
	}

	// This is horrible and should be fixed. The data is sorted in the storage, we should not have to sort it here.
	for _, series := range res.Series {
		sort.Slice(series.Samples, func(i, j int) bool {
			return series.Samples[i].Timestamp.AsTime().Before(series.Samples[j].Timestamp.AsTime())
		})
	}

	return res, nil
}

// Types returns the available types of profiles.
func (q *ColumnQueryAPI) ProfileTypes(ctx context.Context, req *pb.ProfileTypesRequest) (*pb.ProfileTypesResponse, error) {
	res := &pb.ProfileTypesResponse{}

	seen := map[string]struct{}{}

	err := q.engine.ScanTable(q.tableName).
		Distinct(
			logicalplan.Col(parcacol.ColumnName),
			logicalplan.Col(parcacol.ColumnSampleType),
			logicalplan.Col(parcacol.ColumnSampleUnit),
			logicalplan.Col(parcacol.ColumnPeriodType),
			logicalplan.Col(parcacol.ColumnPeriodUnit),
			logicalplan.Col(parcacol.ColumnDuration).GT(logicalplan.Literal(0)),
		).
		Execute(ctx, func(ar arrow.Record) error {
			if ar.NumCols() != 6 {
				return fmt.Errorf("expected 6 column, got %d", ar.NumCols())
			}

			nameColumn, err := binaryFieldFromRecord(ar, parcacol.ColumnName)
			if err != nil {
				return err
			}

			sampleTypeColumn, err := binaryFieldFromRecord(ar, parcacol.ColumnSampleType)
			if err != nil {
				return err
			}

			sampleUnitColumn, err := binaryFieldFromRecord(ar, parcacol.ColumnSampleUnit)
			if err != nil {
				return err
			}

			periodTypeColumn, err := binaryFieldFromRecord(ar, parcacol.ColumnPeriodType)
			if err != nil {
				return err
			}

			periodUnitColumn, err := binaryFieldFromRecord(ar, parcacol.ColumnPeriodUnit)
			if err != nil {
				return err
			}

			deltaColumn, err := booleanFieldFromRecord(ar, "duration > 0")
			if err != nil {
				return err
			}

			for i := 0; i < int(ar.NumRows()); i++ {
				name := string(nameColumn.Value(i))
				sampleType := string(sampleTypeColumn.Value(i))
				sampleUnit := string(sampleUnitColumn.Value(i))
				periodType := string(periodTypeColumn.Value(i))
				periodUnit := string(periodUnitColumn.Value(i))
				delta := deltaColumn.Value(i)

				key := fmt.Sprintf("%s:%s:%s:%s:%s", name, sampleType, sampleUnit, periodType, periodUnit)
				if delta {
					key = fmt.Sprintf("%s:delta", key)
				}

				if _, ok := seen[key]; ok {
					continue
				}
				seen[key] = struct{}{}

				res.Types = append(res.Types, &pb.ProfileType{
					Name:       name,
					SampleType: sampleType,
					SampleUnit: sampleUnit,
					PeriodType: periodType,
					PeriodUnit: periodUnit,
					Delta:      delta,
				})
			}

			return nil
		})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func binaryFieldFromRecord(ar arrow.Record, name string) (*array.Binary, error) {
	indices := ar.Schema().FieldIndices(name)
	if len(indices) != 1 {
		return nil, fmt.Errorf("expected 1 column named %q, got %d", name, len(indices))
	}

	col, ok := ar.Column(indices[0]).(*array.Binary)
	if !ok {
		return nil, fmt.Errorf("expected column %q to be a binary column, got %T", name, ar.Column(indices[0]))
	}

	return col, nil
}

func booleanFieldFromRecord(ar arrow.Record, name string) (*array.Boolean, error) {
	indices := ar.Schema().FieldIndices(name)
	if len(indices) != 1 {
		return nil, fmt.Errorf("expected 1 column named %q, got %d", name, len(indices))
	}

	col, ok := ar.Column(indices[0]).(*array.Boolean)
	if !ok {
		return nil, fmt.Errorf("expected column %q to be a boolean column, got %T", name, ar.Column(indices[0]))
	}

	return col, nil
}

// Query issues a instant query against the storage.
func (q *ColumnQueryAPI) Query(ctx context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	switch req.Mode {
	case pb.QueryRequest_MODE_SINGLE_UNSPECIFIED:
		return q.singleRequest(ctx, req.GetSingle(), req.GetReportType())
	case pb.QueryRequest_MODE_MERGE:
		return q.mergeRequest(ctx, req.GetMerge(), req.GetReportType())
	case pb.QueryRequest_MODE_DIFF:
		return q.diffRequest(ctx, req.GetDiff(), req.GetReportType())
	default:
		return nil, status.Error(codes.InvalidArgument, "unknown query mode")
	}
}

func (q *ColumnQueryAPI) renderReport(ctx context.Context, p *profile.Profile, typ pb.QueryRequest_ReportType) (*pb.QueryResponse, error) {
	switch typ {
	case pb.QueryRequest_REPORT_TYPE_FLAMEGRAPH_UNSPECIFIED:
		fg, err := GenerateFlamegraphFlat(ctx, q.tracer, p)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to generate flamegraph: %v", err.Error())
		}
		return &pb.QueryResponse{
			Report: &pb.QueryResponse_Flamegraph{
				Flamegraph: fg,
			},
		}, nil
	case pb.QueryRequest_REPORT_TYPE_PPROF:
		pp, err := GenerateFlatPprof(ctx, p)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to generate pprof: %v", err.Error())
		}

		var buf bytes.Buffer
		if err := pp.Write(&buf); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to generate pprof: %v", err.Error())
		}

		return &pb.QueryResponse{
			Report: &pb.QueryResponse_Pprof{Pprof: buf.Bytes()},
		}, nil
	case pb.QueryRequest_REPORT_TYPE_TOP:
		top, err := GenerateTopTable(ctx, p)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to generate pprof: %v", err.Error())
		}

		return &pb.QueryResponse{
			Report: &pb.QueryResponse_Top{Top: top},
		}, nil
	default:
		return nil, status.Error(codes.InvalidArgument, "requested report type does not exist")
	}
}

func (q *ColumnQueryAPI) singleRequest(ctx context.Context, s *pb.SingleProfile, reportType pb.QueryRequest_ReportType) (*pb.QueryResponse, error) {
	p, err := q.selectSingle(ctx, s)
	if err != nil {
		return nil, err
	}

	return q.renderReport(ctx, p, reportType)
}

func (q *ColumnQueryAPI) selectSingle(ctx context.Context, s *pb.SingleProfile) (*profile.Profile, error) {
	t := s.Time.AsTime()
	p, err := q.findSingle(ctx, s.Query, t)
	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, status.Error(codes.NotFound, "could not find profile at requested time and selectors")
	}

	return p, nil
}

func (q *ColumnQueryAPI) findSingle(ctx context.Context, query string, t time.Time) (*profile.Profile, error) {
	requestedTime := timestamp.FromTime(t)

	ctx, span := q.tracer.Start(ctx, "findSingle")
	span.SetAttributes(attribute.String("query", query))
	span.SetAttributes(attribute.Int64("time", t.Unix()))
	defer span.End()

	meta, selectorExprs, err := queryToFilterExprs(query)
	if err != nil {
		return nil, err
	}

	filterExpr := logicalplan.And(
		append(
			selectorExprs,
			logicalplan.Col("timestamp").Eq(logicalplan.Literal(requestedTime)),
		)...,
	)

	var ar arrow.Record
	err = q.engine.ScanTable(q.tableName).
		Filter(filterExpr).
		Aggregate(
			logicalplan.Sum(logicalplan.Col("value")),
			logicalplan.Col("stacktrace"),
			logicalplan.DynCol("pprof_labels"),
			logicalplan.DynCol("pprof_num_labels"),
		).
		Execute(ctx, func(r arrow.Record) error {
			r.Retain()
			ar = r
			return nil
		})
	if err != nil {
		return nil, fmt.Errorf("execute query: %w", err)
	}
	defer ar.Release()

	return parcacol.ArrowRecordToStacktraceSamples(
		ctx,
		q.metastore,
		ar,
		"sum(value)",
		profile.Meta{
			Name:       meta.Name,
			SampleType: meta.SampleType,
			PeriodType: meta.PeriodType,
			Timestamp:  requestedTime,
		},
	)
}

func (q *ColumnQueryAPI) mergeRequest(ctx context.Context, m *pb.MergeProfile, reportType pb.QueryRequest_ReportType) (*pb.QueryResponse, error) {
	ctx, span := q.tracer.Start(ctx, "mergeRequest")
	defer span.End()

	p, err := q.selectMerge(ctx, m)
	if err != nil {
		return nil, err
	}

	return q.renderReport(ctx, p, reportType)
}

func (q *ColumnQueryAPI) selectMerge(ctx context.Context, m *pb.MergeProfile) (*profile.Profile, error) {
	ctx, span := q.tracer.Start(ctx, "selectMerge")
	defer span.End()

	meta, selectorExprs, err := queryToFilterExprs(m.Query)
	if err != nil {
		return nil, err
	}

	start := timestamp.FromTime(m.Start.AsTime())
	end := timestamp.FromTime(m.End.AsTime())

	filterExpr := logicalplan.And(
		append(
			selectorExprs,
			logicalplan.Col("timestamp").GT(logicalplan.Literal(start)),
			logicalplan.Col("timestamp").LT(logicalplan.Literal(end)),
		)...,
	)

	var ar arrow.Record
	err = q.engine.ScanTable(q.tableName).
		Filter(filterExpr).
		Aggregate(
			logicalplan.Sum(logicalplan.Col("value")),
			logicalplan.Col("stacktrace"),
		).
		Execute(ctx, func(r arrow.Record) error {
			r.Retain()
			ar = r
			return nil
		})
	if err != nil {
		return nil, err
	}
	defer ar.Release()

	return parcacol.ArrowRecordToStacktraceSamples(
		ctx,
		q.metastore,
		ar,
		"sum(value)",
		profile.Meta{
			Name:       meta.Name,
			SampleType: meta.SampleType,
			PeriodType: meta.PeriodType,
			Timestamp:  start,
		},
	)
}

func (q *ColumnQueryAPI) diffRequest(ctx context.Context, d *pb.DiffProfile, reportType pb.QueryRequest_ReportType) (*pb.QueryResponse, error) {
	ctx, span := q.tracer.Start(ctx, "diffRequest")
	defer span.End()

	if d == nil {
		return nil, status.Error(codes.InvalidArgument, "requested diff mode, but did not provide parameters for diff")
	}

	base, err := q.selectProfileForDiff(ctx, d.A)
	if err != nil {
		return nil, fmt.Errorf("reading base profile: %w", err)
	}

	compare, err := q.selectProfileForDiff(ctx, d.B)
	if err != nil {
		return nil, fmt.Errorf("reading compared profile: %w", err)
	}

	// TODO: This is cheating a bit. This should be done with a sub-query in the columnstore.
	diff := &profile.Profile{}

	// TODO: Use parcacol.Sample for comparing these
	for i := range compare.Samples {
		diff.Samples = append(diff.Samples, &profile.SymbolizedSample{
			Locations: compare.Samples[i].Locations,
			Value:     compare.Samples[i].Value,
			DiffValue: compare.Samples[i].Value,
			Label:     compare.Samples[i].Label,
			NumLabel:  compare.Samples[i].NumLabel,
		})
	}

	for i := range base.Samples {
		diff.Samples = append(diff.Samples, &profile.SymbolizedSample{
			Locations: base.Samples[i].Locations,
			DiffValue: -base.Samples[i].Value,
			Label:     base.Samples[i].Label,
			NumLabel:  base.Samples[i].NumLabel,
		})
	}

	return q.renderReport(ctx, diff, reportType)
}

func (q *ColumnQueryAPI) selectProfileForDiff(ctx context.Context, s *pb.ProfileDiffSelection) (*profile.Profile, error) {
	switch s.Mode {
	case pb.ProfileDiffSelection_MODE_SINGLE_UNSPECIFIED:
		p, err := q.selectSingle(ctx, s.GetSingle())
		if err != nil {
			return nil, fmt.Errorf("selecting single profile: %w", err)
		}
		return p, err
	case pb.ProfileDiffSelection_MODE_MERGE:
		p, err := q.selectMerge(ctx, s.GetMerge())
		if err != nil {
			return nil, fmt.Errorf("selecting merged profile: %w", err)
		}
		return p, err
	default:
		return nil, status.Error(codes.InvalidArgument, "unknown mode for diff profile selection")
	}
}
