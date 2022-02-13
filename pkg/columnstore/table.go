package columnstore

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/apache/arrow/go/v7/arrow"
	"github.com/apache/arrow/go/v7/arrow/memory"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/google/btree"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var ErrNoSchema = fmt.Errorf("no schema")

type Table struct {
	db      *DB
	metrics *tableMetrics
	logger  log.Logger

	schema Schema
	index  *btree.BTree

	// compactor settings
	work chan *Granule
	sync.WaitGroup
}

type tableMetrics struct {
	granulesCreated  prometheus.Counter
	granulesSplits   prometheus.Counter
	rowsInserted     prometheus.Counter
	zeroRowsInserted prometheus.Counter
	rowInsertSize    prometheus.Histogram
}

func newTable(
	db *DB,
	name string,
	schema Schema,
	reg prometheus.Registerer,
	logger log.Logger,
) *Table {
	reg = prometheus.WrapRegistererWith(prometheus.Labels{"table": name}, reg)

	t := &Table{
		db:     db,
		schema: schema,
		index:  btree.New(2), // TODO make the degree a setting
		metrics: &tableMetrics{
			granulesCreated: promauto.With(reg).NewCounter(prometheus.CounterOpts{
				Name: "granules_created",
				Help: "Number of granules created.",
			}),
			granulesSplits: promauto.With(reg).NewCounter(prometheus.CounterOpts{
				Name: "granules_splits",
				Help: "Number of granules splits executed.",
			}),
			rowsInserted: promauto.With(reg).NewCounter(prometheus.CounterOpts{
				Name: "rows_inserted",
				Help: "Number of rows inserted into table.",
			}),
			zeroRowsInserted: promauto.With(reg).NewCounter(prometheus.CounterOpts{
				Name: "zero_rows_inserted",
				Help: "Number of times it was attempted to insert zero rows into the table.",
			}),
			rowInsertSize: promauto.With(reg).NewHistogram(prometheus.HistogramOpts{
				Name:    "row_insert_size",
				Help:    "Size of batch inserts into table.",
				Buckets: prometheus.ExponentialBuckets(1, 2, 10),
			}),
		},
		work: make(chan *Granule, 1024), // TODO buffered or unbuffered? expose as setting if buffered
	}

	promauto.With(reg).NewGaugeFunc(prometheus.GaugeOpts{
		Name: "index_size",
		Help: "Number of granules in the table index currently.",
	}, func() float64 {
		return float64(t.index.Len())
	})

	g := NewGranule(t.metrics.granulesCreated, &t.schema, []*Part{}...)
	t.index.ReplaceOrInsert(g)

	// Start the background compactor go routine
	t.Add(1)
	go t.compactor()

	return t
}

func (t *Table) Insert(rows []Row) error {
	defer func() {
		t.metrics.rowsInserted.Add(float64(len(rows)))
		t.metrics.rowInsertSize.Observe(float64(len(rows)))
	}()

	if len(rows) == 0 {
		t.metrics.zeroRowsInserted.Add(float64(len(rows)))
		return nil
	}

	tx, commit := t.db.begin()
	defer commit()

	rowsToInsertPerGranule := t.splitRowsByGranule(rows)
	for granule, rows := range rowsToInsertPerGranule {
		p, err := NewPart(tx, &t.schema, rows)
		if err != nil {
			return err
		}

		// TODO check if granule pruned; if so, follow new granule links
		granule.AddPart(p)
		if granule.Cardinality() >= t.schema.GranuleSize {
			// Schedule the granule for compaction
			t.work <- granule
		}
	}

	return nil
}

// splitGranule is the main function of the compactor. It will merge all the parts from completed tx's in a Granule (which sorts all rows), and then split that Granule into new Granules.
// It will create a copy of the sparse index, delete the old granule from that index, and insert the new Granules into the new copy of the index. After these steps, it will swap the pointer
// of the old index with the new one.
// The deleted Granule will be marked as purged, and pointers to the newly create Granules will be added to the old Granule. This is important because there may be read or write actions that are
// waiting on the Granule lock while this split is happening. Upon obtaining the lock a read operation will continue as normal reading the old Granule and index. A write operation that obtains the lock
// of a granule markes as purged will duplicate it's writes both to the old granule, as well as the new granule. This will ensure that reads that are still pending, will be able to find data that should
// have existed in the old granule since a merge is destructive of tx ids.
func (t *Table) splitGranule(granule *Granule) {
	granule.Lock()
	defer granule.Unlock()

	// Recheck to ensure the granule still needs to be split
	if granule.pruned || granule.cardinality() < t.schema.GranuleSize {
		return
	}

	// Get a new tx id for this merge/split operation
	tx, commit := t.db.begin()
	defer commit()

	// TODO: There's a bug here, we need to copy non-completed tx writes from this Granule into the new ones.
	newpart, err := Merge(tx, t.db.txCompleted, &t.schema, granule.parts...) // need to merge all parts in a granule before splitting
	if err != nil {
		level.Error(t.logger).Log("msg", "failed to merge parts", "error", err)
	}
	granule.parts = []*Part{newpart}

	granules, err := granule.split(t.schema.GranuleSize / 2) // TODO magic numbers
	if err != nil {
		level.Error(t.logger).Log("msg", "granule split failed after add part", "error", err)
	}

	// Clone the index
	index := t.index.Clone()

	deleted := index.Delete(granule)
	if deleted == nil {
		level.Error(t.logger).Log("msg", "failed to delete granule during split")
	}

	// mark this granule as having been pruned
	// Add the new granules into the old one for pending reads/writes
	granule.pruned = true
	granule.newGranules = granules

	for _, g := range granules {
		index.ReplaceOrInsert(g)
	}

	// Point to the new index
	atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&t.index)), unsafe.Pointer(index))
}

// Iterator iterates in order over all granules in the table. It stops iterating when the iterator function returns false.
func (t *Table) Iterator(pool memory.Allocator, iterator func(r arrow.Record) error) error {
	tx := t.db.beginRead()

	var err error
	t.granuleIterator(func(g *Granule) bool {
		var r arrow.Record
		r, err = g.ArrowRecord(tx, t.db.txCompleted, pool)
		if err != nil {
			return false
		}
		err = iterator(r)
		r.Release()
		return err == nil
	})
	return err
}

func (t *Table) granuleIterator(iterator func(g *Granule) bool) {
	t.index.Ascend(func(i btree.Item) bool {
		g := i.(*Granule)
		return iterator(g)
	})
}

func (t *Table) splitRowsByGranule(rows []Row) map[*Granule][]Row {
	rowsByGranule := map[*Granule][]Row{}

	// Special case: if there is only one granule, insert parts into it until full.
	if t.index.Len() == 1 {
		rowsByGranule[t.index.Min().(*Granule)] = rows
		return rowsByGranule
	}

	// TODO: we might be able to do ascend less than or ascend greater than here?
	j := 0
	var prev *Granule
	t.index.Ascend(func(i btree.Item) bool {
		g := i.(*Granule)
		g.RLock()
		defer g.RUnlock()

		for ; j < len(rows); j++ {
			if rows[j].Less(g.least, t.schema.ordered) {
				if prev != nil {
					rowsByGranule[prev] = append(rowsByGranule[prev], rows[j])
					continue
				}
			}

			// stop at the first granule where this is not the least
			// this might be the correct granule, but we need to check that it isn't the next granule
			prev = g
			return true // continue btree iteration
		}

		// All rows accounted for
		return false
	})

	// Save any remaining rows that belong into prev
	for ; j < len(rows); j++ {
		rowsByGranule[prev] = append(rowsByGranule[prev], rows[j])
	}

	return rowsByGranule
}

// compactor is the background routine responsible for compacting and splitting granules. Only one compactor should ever be running at a time.
func (t *Table) compactor() {
	defer t.Done()
	for granule := range t.work {
		t.splitGranule(granule)
	}
}
