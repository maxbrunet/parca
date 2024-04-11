// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: parca/profilestore/v1alpha1/profilestore.proto

package profilestorev1alpha1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// WriteRawRequest writes a pprof profile for a given tenant
type WriteRawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tenant is the given tenant to store the pprof profile under
	//
	// Deprecated: Marked as deprecated in parca/profilestore/v1alpha1/profilestore.proto.
	Tenant string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	// series is a set raw pprof profiles and accompanying labels
	Series []*RawProfileSeries `protobuf:"bytes,2,rep,name=series,proto3" json:"series,omitempty"`
	// normalized is a flag indicating if the addresses in the profile is normalized for position independent code
	Normalized bool `protobuf:"varint,3,opt,name=normalized,proto3" json:"normalized,omitempty"`
}

func (x *WriteRawRequest) Reset() {
	*x = WriteRawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRawRequest) ProtoMessage() {}

func (x *WriteRawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRawRequest.ProtoReflect.Descriptor instead.
func (*WriteRawRequest) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in parca/profilestore/v1alpha1/profilestore.proto.
func (x *WriteRawRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *WriteRawRequest) GetSeries() []*RawProfileSeries {
	if x != nil {
		return x.Series
	}
	return nil
}

func (x *WriteRawRequest) GetNormalized() bool {
	if x != nil {
		return x.Normalized
	}
	return false
}

// WriteRawResponse is the empty response
type WriteRawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteRawResponse) Reset() {
	*x = WriteRawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRawResponse) ProtoMessage() {}

func (x *WriteRawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRawResponse.ProtoReflect.Descriptor instead.
func (*WriteRawResponse) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{1}
}

// RawProfileSeries represents the pprof profile and its associated labels
type RawProfileSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LabelSet is the key value pairs to identify the corresponding profile
	Labels *LabelSet `protobuf:"bytes,1,opt,name=labels,proto3" json:"labels,omitempty"`
	// samples are the set of profile bytes
	Samples []*RawSample `protobuf:"bytes,2,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *RawProfileSeries) Reset() {
	*x = RawProfileSeries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawProfileSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawProfileSeries) ProtoMessage() {}

func (x *RawProfileSeries) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawProfileSeries.ProtoReflect.Descriptor instead.
func (*RawProfileSeries) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{2}
}

func (x *RawProfileSeries) GetLabels() *LabelSet {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *RawProfileSeries) GetSamples() []*RawSample {
	if x != nil {
		return x.Samples
	}
	return nil
}

// Label is a key value pair of identifiers
type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the label name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// value is the value for the label name
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{3}
}

func (x *Label) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Label) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// LabelSet is a group of labels
type LabelSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// labels are the grouping of labels
	Labels []*Label `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *LabelSet) Reset() {
	*x = LabelSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelSet) ProtoMessage() {}

func (x *LabelSet) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelSet.ProtoReflect.Descriptor instead.
func (*LabelSet) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{4}
}

func (x *LabelSet) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

// RawSample is the set of bytes that correspond to a pprof profile
type RawSample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// raw_profile is the set of bytes of the pprof profile
	RawProfile []byte `protobuf:"bytes,1,opt,name=raw_profile,json=rawProfile,proto3" json:"raw_profile,omitempty"`
	// information about the executable and executable section for normalizaton
	// purposes.
	ExecutableInfo []*ExecutableInfo `protobuf:"bytes,2,rep,name=executable_info,json=executableInfo,proto3" json:"executable_info,omitempty"`
}

func (x *RawSample) Reset() {
	*x = RawSample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawSample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawSample) ProtoMessage() {}

func (x *RawSample) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawSample.ProtoReflect.Descriptor instead.
func (*RawSample) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{5}
}

func (x *RawSample) GetRawProfile() []byte {
	if x != nil {
		return x.RawProfile
	}
	return nil
}

func (x *RawSample) GetExecutableInfo() []*ExecutableInfo {
	if x != nil {
		return x.ExecutableInfo
	}
	return nil
}

// ExecutableInfo is the information about the executable and executable
// section for normalizaton purposes before symbolization.
type ExecutableInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// elf_type is the type of the elf executable. Technically the elf type is a
	// 16 bit integer, but protobuf's smallest unsigned integer is 32 bits.
	ElfType uint32 `protobuf:"varint,1,opt,name=elf_type,json=elfType,proto3" json:"elf_type,omitempty"`
	// load_segment is the load segment of the executable.
	LoadSegment *LoadSegment `protobuf:"bytes,2,opt,name=load_segment,json=loadSegment,proto3" json:"load_segment,omitempty"`
}

func (x *ExecutableInfo) Reset() {
	*x = ExecutableInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutableInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutableInfo) ProtoMessage() {}

func (x *ExecutableInfo) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutableInfo.ProtoReflect.Descriptor instead.
func (*ExecutableInfo) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{6}
}

func (x *ExecutableInfo) GetElfType() uint32 {
	if x != nil {
		return x.ElfType
	}
	return 0
}

func (x *ExecutableInfo) GetLoadSegment() *LoadSegment {
	if x != nil {
		return x.LoadSegment
	}
	return nil
}

// LoadSegment is the load segment of the executable
type LoadSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The offset from the beginning of the file at which the first byte of the segment resides.
	Offset uint64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	// The virtual address at which the first byte of the segment resides in memory.
	Vaddr uint64 `protobuf:"varint,2,opt,name=vaddr,proto3" json:"vaddr,omitempty"`
}

func (x *LoadSegment) Reset() {
	*x = LoadSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadSegment) ProtoMessage() {}

func (x *LoadSegment) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadSegment.ProtoReflect.Descriptor instead.
func (*LoadSegment) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{7}
}

func (x *LoadSegment) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *LoadSegment) GetVaddr() uint64 {
	if x != nil {
		return x.Vaddr
	}
	return 0
}

// AgentsRequest is the request to retrieve a list of agents
type AgentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AgentsRequest) Reset() {
	*x = AgentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentsRequest) ProtoMessage() {}

func (x *AgentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentsRequest.ProtoReflect.Descriptor instead.
func (*AgentsRequest) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{8}
}

// AgentsResponse is the request to retrieve a list of agents
type AgentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// agents is a list of agents
	Agents []*Agent `protobuf:"bytes,1,rep,name=agents,proto3" json:"agents,omitempty"`
}

func (x *AgentsResponse) Reset() {
	*x = AgentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentsResponse) ProtoMessage() {}

func (x *AgentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentsResponse.ProtoReflect.Descriptor instead.
func (*AgentsResponse) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{9}
}

func (x *AgentsResponse) GetAgents() []*Agent {
	if x != nil {
		return x.Agents
	}
	return nil
}

// Agent is the agent representation
type Agent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id is the agent identity that either represent by the node name or the IP address.
	// When node name is not found, this will fallback to IP address.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// last_error is the error message most recently received from a push attempt
	LastError string `protobuf:"bytes,2,opt,name=last_error,json=lastError,proto3" json:"last_error,omitempty"`
	// last_push is the time stamp the last push request was performed
	LastPush *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=last_push,json=lastPush,proto3" json:"last_push,omitempty"`
	// last_push_duration is the duration of the last push request
	LastPushDuration *durationpb.Duration `protobuf:"bytes,4,opt,name=last_push_duration,json=lastPushDuration,proto3" json:"last_push_duration,omitempty"`
}

func (x *Agent) Reset() {
	*x = Agent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Agent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Agent) ProtoMessage() {}

func (x *Agent) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Agent.ProtoReflect.Descriptor instead.
func (*Agent) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{10}
}

func (x *Agent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Agent) GetLastError() string {
	if x != nil {
		return x.LastError
	}
	return ""
}

func (x *Agent) GetLastPush() *timestamppb.Timestamp {
	if x != nil {
		return x.LastPush
	}
	return nil
}

func (x *Agent) GetLastPushDuration() *durationpb.Duration {
	if x != nil {
		return x.LastPushDuration
	}
	return nil
}

var File_parca_profilestore_v1alpha1_profilestore_proto protoreflect.FileDescriptor

var file_parca_profilestore_v1alpha1_profilestore_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x01, 0x0a,
	0x0f, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x02, 0x18, 0x01, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x06,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70,
	0x61, 0x72, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x06, 0x73, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x61, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x10, 0x52, 0x61, 0x77, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70,
	0x61, 0x72, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x53, 0x65, 0x74, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x40, 0x0a, 0x07, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70,
	0x61, 0x72, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x22, 0x31, 0x0a,
	0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x46, 0x0a, 0x08, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x74, 0x12, 0x3a, 0x0a, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70,
	0x61, 0x72, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x09, 0x52, 0x61, 0x77,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x61, 0x77, 0x5f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x61, 0x77,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x54, 0x0a, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x78, 0x0a,
	0x0e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x6c, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x65, 0x6c, 0x66, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c,
	0x6f, 0x61, 0x64, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x6c, 0x6f, 0x61, 0x64,
	0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x3b, 0x0a, 0x0b, 0x4c, 0x6f, 0x61, 0x64, 0x53,
	0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76,
	0x61, 0x64, 0x64, 0x72, 0x22, 0x0f, 0x0a, 0x0d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x0e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0xb8, 0x01, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70, 0x75, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x50, 0x75, 0x73, 0x68, 0x12, 0x47, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x70, 0x75,
	0x73, 0x68, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x6c, 0x61,
	0x73, 0x74, 0x50, 0x75, 0x73, 0x68, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x9e,
	0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x08, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x61, 0x77, 0x12, 0x2c, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2d, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x61, 0x77, 0x32,
	0x83, 0x01, 0x0a, 0x0d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x72, 0x0a, 0x06, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2a, 0x2e, 0x70, 0x61,
	0x72, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x9c, 0x02, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61,
	0x72, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x11, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x72, 0x63, 0x61,
	0x2d, 0x64, 0x65, 0x76, 0x2f, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x50, 0x58, 0xaa, 0x02,
	0x1b, 0x50, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1b, 0x50,
	0x61, 0x72, 0x63, 0x61, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x27, 0x50, 0x61, 0x72,
	0x63, 0x61, 0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x50, 0x61, 0x72, 0x63, 0x61, 0x3a, 0x3a, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_parca_profilestore_v1alpha1_profilestore_proto_rawDescOnce sync.Once
	file_parca_profilestore_v1alpha1_profilestore_proto_rawDescData = file_parca_profilestore_v1alpha1_profilestore_proto_rawDesc
)

func file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP() []byte {
	file_parca_profilestore_v1alpha1_profilestore_proto_rawDescOnce.Do(func() {
		file_parca_profilestore_v1alpha1_profilestore_proto_rawDescData = protoimpl.X.CompressGZIP(file_parca_profilestore_v1alpha1_profilestore_proto_rawDescData)
	})
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescData
}

var file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_parca_profilestore_v1alpha1_profilestore_proto_goTypes = []interface{}{
	(*WriteRawRequest)(nil),       // 0: parca.profilestore.v1alpha1.WriteRawRequest
	(*WriteRawResponse)(nil),      // 1: parca.profilestore.v1alpha1.WriteRawResponse
	(*RawProfileSeries)(nil),      // 2: parca.profilestore.v1alpha1.RawProfileSeries
	(*Label)(nil),                 // 3: parca.profilestore.v1alpha1.Label
	(*LabelSet)(nil),              // 4: parca.profilestore.v1alpha1.LabelSet
	(*RawSample)(nil),             // 5: parca.profilestore.v1alpha1.RawSample
	(*ExecutableInfo)(nil),        // 6: parca.profilestore.v1alpha1.ExecutableInfo
	(*LoadSegment)(nil),           // 7: parca.profilestore.v1alpha1.LoadSegment
	(*AgentsRequest)(nil),         // 8: parca.profilestore.v1alpha1.AgentsRequest
	(*AgentsResponse)(nil),        // 9: parca.profilestore.v1alpha1.AgentsResponse
	(*Agent)(nil),                 // 10: parca.profilestore.v1alpha1.Agent
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 12: google.protobuf.Duration
}
var file_parca_profilestore_v1alpha1_profilestore_proto_depIdxs = []int32{
	2,  // 0: parca.profilestore.v1alpha1.WriteRawRequest.series:type_name -> parca.profilestore.v1alpha1.RawProfileSeries
	4,  // 1: parca.profilestore.v1alpha1.RawProfileSeries.labels:type_name -> parca.profilestore.v1alpha1.LabelSet
	5,  // 2: parca.profilestore.v1alpha1.RawProfileSeries.samples:type_name -> parca.profilestore.v1alpha1.RawSample
	3,  // 3: parca.profilestore.v1alpha1.LabelSet.labels:type_name -> parca.profilestore.v1alpha1.Label
	6,  // 4: parca.profilestore.v1alpha1.RawSample.executable_info:type_name -> parca.profilestore.v1alpha1.ExecutableInfo
	7,  // 5: parca.profilestore.v1alpha1.ExecutableInfo.load_segment:type_name -> parca.profilestore.v1alpha1.LoadSegment
	10, // 6: parca.profilestore.v1alpha1.AgentsResponse.agents:type_name -> parca.profilestore.v1alpha1.Agent
	11, // 7: parca.profilestore.v1alpha1.Agent.last_push:type_name -> google.protobuf.Timestamp
	12, // 8: parca.profilestore.v1alpha1.Agent.last_push_duration:type_name -> google.protobuf.Duration
	0,  // 9: parca.profilestore.v1alpha1.ProfileStoreService.WriteRaw:input_type -> parca.profilestore.v1alpha1.WriteRawRequest
	8,  // 10: parca.profilestore.v1alpha1.AgentsService.Agents:input_type -> parca.profilestore.v1alpha1.AgentsRequest
	1,  // 11: parca.profilestore.v1alpha1.ProfileStoreService.WriteRaw:output_type -> parca.profilestore.v1alpha1.WriteRawResponse
	9,  // 12: parca.profilestore.v1alpha1.AgentsService.Agents:output_type -> parca.profilestore.v1alpha1.AgentsResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_parca_profilestore_v1alpha1_profilestore_proto_init() }
func file_parca_profilestore_v1alpha1_profilestore_proto_init() {
	if File_parca_profilestore_v1alpha1_profilestore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRawRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRawResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawProfileSeries); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Label); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawSample); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutableInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoadSegment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Agent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_parca_profilestore_v1alpha1_profilestore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_parca_profilestore_v1alpha1_profilestore_proto_goTypes,
		DependencyIndexes: file_parca_profilestore_v1alpha1_profilestore_proto_depIdxs,
		MessageInfos:      file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes,
	}.Build()
	File_parca_profilestore_v1alpha1_profilestore_proto = out.File
	file_parca_profilestore_v1alpha1_profilestore_proto_rawDesc = nil
	file_parca_profilestore_v1alpha1_profilestore_proto_goTypes = nil
	file_parca_profilestore_v1alpha1_profilestore_proto_depIdxs = nil
}
