//
// Copyright (C) 2021 Vaticle
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: v2/protobuf/common/options.proto

package typedb_protocol

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to InferOpt:
	//	*Options_Infer
	InferOpt isOptions_InferOpt `protobuf_oneof:"infer_opt"`
	// Types that are assignable to TraceInferenceOpt:
	//	*Options_TraceInference
	TraceInferenceOpt isOptions_TraceInferenceOpt `protobuf_oneof:"trace_inference_opt"`
	// Types that are assignable to ExplainOpt:
	//	*Options_Explain
	ExplainOpt isOptions_ExplainOpt `protobuf_oneof:"explain_opt"`
	// Types that are assignable to ParallelOpt:
	//	*Options_Parallel
	ParallelOpt isOptions_ParallelOpt `protobuf_oneof:"parallel_opt"`
	// Types that are assignable to PrefetchSizeOpt:
	//	*Options_PrefetchSize
	PrefetchSizeOpt isOptions_PrefetchSizeOpt `protobuf_oneof:"prefetch_size_opt"`
	// Types that are assignable to PrefetchOpt:
	//	*Options_Prefetch
	PrefetchOpt isOptions_PrefetchOpt `protobuf_oneof:"prefetch_opt"`
	// Types that are assignable to SessionIdleTimeoutOpt:
	//	*Options_SessionIdleTimeoutMillis
	SessionIdleTimeoutOpt isOptions_SessionIdleTimeoutOpt `protobuf_oneof:"session_idle_timeout_opt"`
	// Types that are assignable to TransactionTimeoutOpt:
	//	*Options_TransactionTimeoutMillis
	TransactionTimeoutOpt isOptions_TransactionTimeoutOpt `protobuf_oneof:"transaction_timeout_opt"`
	// Types that are assignable to SchemaLockAcquireTimeoutOpt:
	//	*Options_SchemaLockAcquireTimeoutMillis
	SchemaLockAcquireTimeoutOpt isOptions_SchemaLockAcquireTimeoutOpt `protobuf_oneof:"schema_lock_acquire_timeout_opt"`
	// Types that are assignable to ReadAnyReplicaOpt:
	//	*Options_ReadAnyReplica
	ReadAnyReplicaOpt isOptions_ReadAnyReplicaOpt `protobuf_oneof:"read_any_replica_opt"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_options_proto_rawDescGZIP(), []int{0}
}

func (m *Options) GetInferOpt() isOptions_InferOpt {
	if m != nil {
		return m.InferOpt
	}
	return nil
}

func (x *Options) GetInfer() bool {
	if x, ok := x.GetInferOpt().(*Options_Infer); ok {
		return x.Infer
	}
	return false
}

func (m *Options) GetTraceInferenceOpt() isOptions_TraceInferenceOpt {
	if m != nil {
		return m.TraceInferenceOpt
	}
	return nil
}

func (x *Options) GetTraceInference() bool {
	if x, ok := x.GetTraceInferenceOpt().(*Options_TraceInference); ok {
		return x.TraceInference
	}
	return false
}

func (m *Options) GetExplainOpt() isOptions_ExplainOpt {
	if m != nil {
		return m.ExplainOpt
	}
	return nil
}

func (x *Options) GetExplain() bool {
	if x, ok := x.GetExplainOpt().(*Options_Explain); ok {
		return x.Explain
	}
	return false
}

func (m *Options) GetParallelOpt() isOptions_ParallelOpt {
	if m != nil {
		return m.ParallelOpt
	}
	return nil
}

func (x *Options) GetParallel() bool {
	if x, ok := x.GetParallelOpt().(*Options_Parallel); ok {
		return x.Parallel
	}
	return false
}

func (m *Options) GetPrefetchSizeOpt() isOptions_PrefetchSizeOpt {
	if m != nil {
		return m.PrefetchSizeOpt
	}
	return nil
}

func (x *Options) GetPrefetchSize() int32 {
	if x, ok := x.GetPrefetchSizeOpt().(*Options_PrefetchSize); ok {
		return x.PrefetchSize
	}
	return 0
}

func (m *Options) GetPrefetchOpt() isOptions_PrefetchOpt {
	if m != nil {
		return m.PrefetchOpt
	}
	return nil
}

func (x *Options) GetPrefetch() bool {
	if x, ok := x.GetPrefetchOpt().(*Options_Prefetch); ok {
		return x.Prefetch
	}
	return false
}

func (m *Options) GetSessionIdleTimeoutOpt() isOptions_SessionIdleTimeoutOpt {
	if m != nil {
		return m.SessionIdleTimeoutOpt
	}
	return nil
}

func (x *Options) GetSessionIdleTimeoutMillis() int32 {
	if x, ok := x.GetSessionIdleTimeoutOpt().(*Options_SessionIdleTimeoutMillis); ok {
		return x.SessionIdleTimeoutMillis
	}
	return 0
}

func (m *Options) GetTransactionTimeoutOpt() isOptions_TransactionTimeoutOpt {
	if m != nil {
		return m.TransactionTimeoutOpt
	}
	return nil
}

func (x *Options) GetTransactionTimeoutMillis() int32 {
	if x, ok := x.GetTransactionTimeoutOpt().(*Options_TransactionTimeoutMillis); ok {
		return x.TransactionTimeoutMillis
	}
	return 0
}

func (m *Options) GetSchemaLockAcquireTimeoutOpt() isOptions_SchemaLockAcquireTimeoutOpt {
	if m != nil {
		return m.SchemaLockAcquireTimeoutOpt
	}
	return nil
}

func (x *Options) GetSchemaLockAcquireTimeoutMillis() int32 {
	if x, ok := x.GetSchemaLockAcquireTimeoutOpt().(*Options_SchemaLockAcquireTimeoutMillis); ok {
		return x.SchemaLockAcquireTimeoutMillis
	}
	return 0
}

func (m *Options) GetReadAnyReplicaOpt() isOptions_ReadAnyReplicaOpt {
	if m != nil {
		return m.ReadAnyReplicaOpt
	}
	return nil
}

func (x *Options) GetReadAnyReplica() bool {
	if x, ok := x.GetReadAnyReplicaOpt().(*Options_ReadAnyReplica); ok {
		return x.ReadAnyReplica
	}
	return false
}

type isOptions_InferOpt interface {
	isOptions_InferOpt()
}

type Options_Infer struct {
	Infer bool `protobuf:"varint,1,opt,name=infer,proto3,oneof"`
}

func (*Options_Infer) isOptions_InferOpt() {}

type isOptions_TraceInferenceOpt interface {
	isOptions_TraceInferenceOpt()
}

type Options_TraceInference struct {
	TraceInference bool `protobuf:"varint,2,opt,name=trace_inference,json=traceInference,proto3,oneof"`
}

func (*Options_TraceInference) isOptions_TraceInferenceOpt() {}

type isOptions_ExplainOpt interface {
	isOptions_ExplainOpt()
}

type Options_Explain struct {
	Explain bool `protobuf:"varint,3,opt,name=explain,proto3,oneof"`
}

func (*Options_Explain) isOptions_ExplainOpt() {}

type isOptions_ParallelOpt interface {
	isOptions_ParallelOpt()
}

type Options_Parallel struct {
	Parallel bool `protobuf:"varint,4,opt,name=parallel,proto3,oneof"`
}

func (*Options_Parallel) isOptions_ParallelOpt() {}

type isOptions_PrefetchSizeOpt interface {
	isOptions_PrefetchSizeOpt()
}

type Options_PrefetchSize struct {
	PrefetchSize int32 `protobuf:"varint,5,opt,name=prefetch_size,json=prefetchSize,proto3,oneof"`
}

func (*Options_PrefetchSize) isOptions_PrefetchSizeOpt() {}

type isOptions_PrefetchOpt interface {
	isOptions_PrefetchOpt()
}

type Options_Prefetch struct {
	Prefetch bool `protobuf:"varint,6,opt,name=prefetch,proto3,oneof"`
}

func (*Options_Prefetch) isOptions_PrefetchOpt() {}

type isOptions_SessionIdleTimeoutOpt interface {
	isOptions_SessionIdleTimeoutOpt()
}

type Options_SessionIdleTimeoutMillis struct {
	SessionIdleTimeoutMillis int32 `protobuf:"varint,7,opt,name=session_idle_timeout_millis,json=sessionIdleTimeoutMillis,proto3,oneof"`
}

func (*Options_SessionIdleTimeoutMillis) isOptions_SessionIdleTimeoutOpt() {}

type isOptions_TransactionTimeoutOpt interface {
	isOptions_TransactionTimeoutOpt()
}

type Options_TransactionTimeoutMillis struct {
	TransactionTimeoutMillis int32 `protobuf:"varint,8,opt,name=transaction_timeout_millis,json=transactionTimeoutMillis,proto3,oneof"`
}

func (*Options_TransactionTimeoutMillis) isOptions_TransactionTimeoutOpt() {}

type isOptions_SchemaLockAcquireTimeoutOpt interface {
	isOptions_SchemaLockAcquireTimeoutOpt()
}

type Options_SchemaLockAcquireTimeoutMillis struct {
	SchemaLockAcquireTimeoutMillis int32 `protobuf:"varint,9,opt,name=schema_lock_acquire_timeout_millis,json=schemaLockAcquireTimeoutMillis,proto3,oneof"`
}

func (*Options_SchemaLockAcquireTimeoutMillis) isOptions_SchemaLockAcquireTimeoutOpt() {}

type isOptions_ReadAnyReplicaOpt interface {
	isOptions_ReadAnyReplicaOpt()
}

type Options_ReadAnyReplica struct {
	ReadAnyReplica bool `protobuf:"varint,10,opt,name=read_any_replica,json=readAnyReplica,proto3,oneof"`
}

func (*Options_ReadAnyReplica) isOptions_ReadAnyReplicaOpt() {}

var File_v2_protobuf_common_options_proto protoreflect.FileDescriptor

var file_v2_protobuf_common_options_proto_rawDesc = []byte{
	0x0a, 0x20, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x22, 0xa0, 0x05, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x16, 0x0a, 0x05, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x05, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x5f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x01, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x07, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x12, 0x1c,
	0x0a, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x12, 0x25, 0x0a, 0x0d,
	0x70, 0x72, 0x65, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x04, 0x52, 0x0c, 0x70, 0x72, 0x65, 0x66, 0x65, 0x74, 0x63, 0x68, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x66, 0x65, 0x74, 0x63, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x05, 0x52, 0x08, 0x70, 0x72, 0x65, 0x66, 0x65, 0x74, 0x63,
	0x68, 0x12, 0x3f, 0x0a, 0x1b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x6c,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48, 0x06, 0x52, 0x18, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x69, 0x6c, 0x6c,
	0x69, 0x73, 0x12, 0x3e, 0x0a, 0x1a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x48, 0x07, 0x52, 0x18, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x69, 0x6c, 0x6c,
	0x69, 0x73, 0x12, 0x4c, 0x0a, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x61, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x48, 0x08,
	0x52, 0x1e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4c, 0x6f, 0x63, 0x6b, 0x41, 0x63, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73,
	0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x61, 0x6e, 0x79, 0x5f, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x48, 0x09, 0x52, 0x0e, 0x72, 0x65,
	0x61, 0x64, 0x41, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x42, 0x0b, 0x0a, 0x09,
	0x69, 0x6e, 0x66, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x42, 0x15, 0x0a, 0x13, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x74,
	0x42, 0x0d, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x42,
	0x0e, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x42,
	0x13, 0x0a, 0x11, 0x70, 0x72, 0x65, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x5f, 0x6f, 0x70, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x70, 0x72, 0x65, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x5f, 0x6f, 0x70, 0x74, 0x42, 0x1a, 0x0a, 0x18, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x6f, 0x70, 0x74,
	0x42, 0x19, 0x0a, 0x17, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x42, 0x21, 0x0a, 0x1f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x61, 0x63, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x42, 0x16,
	0x0a, 0x14, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x61, 0x6e, 0x79, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x5f, 0x6f, 0x70, 0x74, 0x42, 0x51, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x61,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x0c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5a, 0x24, 0x2e, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v2_protobuf_common_options_proto_rawDescOnce sync.Once
	file_v2_protobuf_common_options_proto_rawDescData = file_v2_protobuf_common_options_proto_rawDesc
)

func file_v2_protobuf_common_options_proto_rawDescGZIP() []byte {
	file_v2_protobuf_common_options_proto_rawDescOnce.Do(func() {
		file_v2_protobuf_common_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_protobuf_common_options_proto_rawDescData)
	})
	return file_v2_protobuf_common_options_proto_rawDescData
}

var file_v2_protobuf_common_options_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v2_protobuf_common_options_proto_goTypes = []interface{}{
	(*Options)(nil), // 0: typedb.protocol.Options
}
var file_v2_protobuf_common_options_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v2_protobuf_common_options_proto_init() }
func file_v2_protobuf_common_options_proto_init() {
	if File_v2_protobuf_common_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_protobuf_common_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
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
	file_v2_protobuf_common_options_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Options_Infer)(nil),
		(*Options_TraceInference)(nil),
		(*Options_Explain)(nil),
		(*Options_Parallel)(nil),
		(*Options_PrefetchSize)(nil),
		(*Options_Prefetch)(nil),
		(*Options_SessionIdleTimeoutMillis)(nil),
		(*Options_TransactionTimeoutMillis)(nil),
		(*Options_SchemaLockAcquireTimeoutMillis)(nil),
		(*Options_ReadAnyReplica)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v2_protobuf_common_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_protobuf_common_options_proto_goTypes,
		DependencyIndexes: file_v2_protobuf_common_options_proto_depIdxs,
		MessageInfos:      file_v2_protobuf_common_options_proto_msgTypes,
	}.Build()
	File_v2_protobuf_common_options_proto = out.File
	file_v2_protobuf_common_options_proto_rawDesc = nil
	file_v2_protobuf_common_options_proto_goTypes = nil
	file_v2_protobuf_common_options_proto_depIdxs = nil
}
