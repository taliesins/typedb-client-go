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
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: v2/protobuf/common/answer.proto

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

type ConceptMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Map          map[string]*Concept `protobuf:"bytes,1,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Explainables *Explainables       `protobuf:"bytes,2,opt,name=explainables,proto3" json:"explainables,omitempty"`
}

func (x *ConceptMap) Reset() {
	*x = ConceptMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_answer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConceptMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConceptMap) ProtoMessage() {}

func (x *ConceptMap) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_answer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConceptMap.ProtoReflect.Descriptor instead.
func (*ConceptMap) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_answer_proto_rawDescGZIP(), []int{0}
}

func (x *ConceptMap) GetMap() map[string]*Concept {
	if x != nil {
		return x.Map
	}
	return nil
}

func (x *ConceptMap) GetExplainables() *Explainables {
	if x != nil {
		return x.Explainables
	}
	return nil
}

type Explainables struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Relations  map[string]*Explainable        `protobuf:"bytes,1,rep,name=relations,proto3" json:"relations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Attributes map[string]*Explainable        `protobuf:"bytes,2,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Ownerships map[string]*Explainables_Owned `protobuf:"bytes,3,rep,name=ownerships,proto3" json:"ownerships,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Explainables) Reset() {
	*x = Explainables{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_answer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Explainables) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Explainables) ProtoMessage() {}

func (x *Explainables) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_answer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Explainables.ProtoReflect.Descriptor instead.
func (*Explainables) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_answer_proto_rawDescGZIP(), []int{1}
}

func (x *Explainables) GetRelations() map[string]*Explainable {
	if x != nil {
		return x.Relations
	}
	return nil
}

func (x *Explainables) GetAttributes() map[string]*Explainable {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Explainables) GetOwnerships() map[string]*Explainables_Owned {
	if x != nil {
		return x.Ownerships
	}
	return nil
}

type Explainable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conjunction string `protobuf:"bytes,1,opt,name=conjunction,proto3" json:"conjunction,omitempty"`
	Id          int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Explainable) Reset() {
	*x = Explainable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_answer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Explainable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Explainable) ProtoMessage() {}

func (x *Explainable) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_answer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Explainable.ProtoReflect.Descriptor instead.
func (*Explainable) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_answer_proto_rawDescGZIP(), []int{2}
}

func (x *Explainable) GetConjunction() string {
	if x != nil {
		return x.Conjunction
	}
	return ""
}

func (x *Explainable) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ConceptMapGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner       *Concept      `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	ConceptMaps []*ConceptMap `protobuf:"bytes,2,rep,name=concept_maps,json=conceptMaps,proto3" json:"concept_maps,omitempty"`
}

func (x *ConceptMapGroup) Reset() {
	*x = ConceptMapGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_answer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConceptMapGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConceptMapGroup) ProtoMessage() {}

func (x *ConceptMapGroup) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_answer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConceptMapGroup.ProtoReflect.Descriptor instead.
func (*ConceptMapGroup) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_answer_proto_rawDescGZIP(), []int{3}
}

func (x *ConceptMapGroup) GetOwner() *Concept {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *ConceptMapGroup) GetConceptMaps() []*ConceptMap {
	if x != nil {
		return x.ConceptMaps
	}
	return nil
}

type Numeric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*Numeric_LongValue
	//	*Numeric_DoubleValue
	//	*Numeric_Nan
	Value isNumeric_Value `protobuf_oneof:"value"`
}

func (x *Numeric) Reset() {
	*x = Numeric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_answer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Numeric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Numeric) ProtoMessage() {}

func (x *Numeric) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_answer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Numeric.ProtoReflect.Descriptor instead.
func (*Numeric) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_answer_proto_rawDescGZIP(), []int{4}
}

func (m *Numeric) GetValue() isNumeric_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Numeric) GetLongValue() int64 {
	if x, ok := x.GetValue().(*Numeric_LongValue); ok {
		return x.LongValue
	}
	return 0
}

func (x *Numeric) GetDoubleValue() float64 {
	if x, ok := x.GetValue().(*Numeric_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *Numeric) GetNan() bool {
	if x, ok := x.GetValue().(*Numeric_Nan); ok {
		return x.Nan
	}
	return false
}

type isNumeric_Value interface {
	isNumeric_Value()
}

type Numeric_LongValue struct {
	LongValue int64 `protobuf:"varint,1,opt,name=long_value,json=longValue,proto3,oneof"`
}

type Numeric_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,2,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Numeric_Nan struct {
	Nan bool `protobuf:"varint,3,opt,name=nan,proto3,oneof"`
}

func (*Numeric_LongValue) isNumeric_Value() {}

func (*Numeric_DoubleValue) isNumeric_Value() {}

func (*Numeric_Nan) isNumeric_Value() {}

type NumericGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner  *Concept `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Number *Numeric `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *NumericGroup) Reset() {
	*x = NumericGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_answer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumericGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumericGroup) ProtoMessage() {}

func (x *NumericGroup) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_answer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumericGroup.ProtoReflect.Descriptor instead.
func (*NumericGroup) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_answer_proto_rawDescGZIP(), []int{5}
}

func (x *NumericGroup) GetOwner() *Concept {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *NumericGroup) GetNumber() *Numeric {
	if x != nil {
		return x.Number
	}
	return nil
}

type Explainables_Owned struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owned map[string]*Explainable `protobuf:"bytes,1,rep,name=owned,proto3" json:"owned,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Explainables_Owned) Reset() {
	*x = Explainables_Owned{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_answer_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Explainables_Owned) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Explainables_Owned) ProtoMessage() {}

func (x *Explainables_Owned) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_answer_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Explainables_Owned.ProtoReflect.Descriptor instead.
func (*Explainables_Owned) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_answer_proto_rawDescGZIP(), []int{1, 3}
}

func (x *Explainables_Owned) GetOwned() map[string]*Explainable {
	if x != nil {
		return x.Owned
	}
	return nil
}

var File_v2_protobuf_common_answer_proto protoreflect.FileDescriptor

var file_v2_protobuf_common_answer_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x1a, 0x20, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74,
	0x4d, 0x61, 0x70, 0x12, 0x36, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x4d, 0x61, 0x70, 0x2e, 0x4d, 0x61,
	0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x12, 0x41, 0x0a, 0x0c, 0x65,
	0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x52, 0x0c, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x1a, 0x50,
	0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x6f,
	0x6e, 0x63, 0x65, 0x70, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xbd, 0x05, 0x0a, 0x0c, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x12, 0x4a, 0x0a, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x09, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4d, 0x0a,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x0a,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x1a, 0x5a, 0x0a, 0x0e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5b, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x78,
	0x70, 0x6c, 0x61, 0x69, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x62, 0x0a, 0x0f, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xa5, 0x01, 0x0a, 0x05, 0x4f, 0x77, 0x6e,
	0x65, 0x64, 0x12, 0x44, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x2e, 0x4f, 0x77, 0x6e, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x64, 0x1a, 0x56, 0x0a, 0x0a, 0x4f, 0x77, 0x6e, 0x65,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x3f, 0x0a, 0x0b, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x81, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x4d, 0x61, 0x70,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2e, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74,
	0x5f, 0x6d, 0x61, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x6f,
	0x6e, 0x63, 0x65, 0x70, 0x74, 0x4d, 0x61, 0x70, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70,
	0x74, 0x4d, 0x61, 0x70, 0x73, 0x22, 0x6c, 0x0a, 0x07, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63,
	0x12, 0x1f, 0x0a, 0x0a, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x03, 0x6e, 0x61, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x03, 0x6e, 0x61, 0x6e, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x70, 0x0a, 0x0c, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x2e, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x52, 0x05, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x50, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x61, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x0b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5a, 0x24, 0x2e, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_protobuf_common_answer_proto_rawDescOnce sync.Once
	file_v2_protobuf_common_answer_proto_rawDescData = file_v2_protobuf_common_answer_proto_rawDesc
)

func file_v2_protobuf_common_answer_proto_rawDescGZIP() []byte {
	file_v2_protobuf_common_answer_proto_rawDescOnce.Do(func() {
		file_v2_protobuf_common_answer_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_protobuf_common_answer_proto_rawDescData)
	})
	return file_v2_protobuf_common_answer_proto_rawDescData
}

var file_v2_protobuf_common_answer_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_v2_protobuf_common_answer_proto_goTypes = []interface{}{
	(*ConceptMap)(nil),         // 0: typedb.protocol.ConceptMap
	(*Explainables)(nil),       // 1: typedb.protocol.Explainables
	(*Explainable)(nil),        // 2: typedb.protocol.Explainable
	(*ConceptMapGroup)(nil),    // 3: typedb.protocol.ConceptMapGroup
	(*Numeric)(nil),            // 4: typedb.protocol.Numeric
	(*NumericGroup)(nil),       // 5: typedb.protocol.NumericGroup
	nil,                        // 6: typedb.protocol.ConceptMap.MapEntry
	nil,                        // 7: typedb.protocol.Explainables.RelationsEntry
	nil,                        // 8: typedb.protocol.Explainables.AttributesEntry
	nil,                        // 9: typedb.protocol.Explainables.OwnershipsEntry
	(*Explainables_Owned)(nil), // 10: typedb.protocol.Explainables.Owned
	nil,                        // 11: typedb.protocol.Explainables.Owned.OwnedEntry
	(*Concept)(nil),            // 12: typedb.protocol.Concept
}
var file_v2_protobuf_common_answer_proto_depIdxs = []int32{
	6,  // 0: typedb.protocol.ConceptMap.map:type_name -> typedb.protocol.ConceptMap.MapEntry
	1,  // 1: typedb.protocol.ConceptMap.explainables:type_name -> typedb.protocol.Explainables
	7,  // 2: typedb.protocol.Explainables.relations:type_name -> typedb.protocol.Explainables.RelationsEntry
	8,  // 3: typedb.protocol.Explainables.attributes:type_name -> typedb.protocol.Explainables.AttributesEntry
	9,  // 4: typedb.protocol.Explainables.ownerships:type_name -> typedb.protocol.Explainables.OwnershipsEntry
	12, // 5: typedb.protocol.ConceptMapGroup.owner:type_name -> typedb.protocol.Concept
	0,  // 6: typedb.protocol.ConceptMapGroup.concept_maps:type_name -> typedb.protocol.ConceptMap
	12, // 7: typedb.protocol.NumericGroup.owner:type_name -> typedb.protocol.Concept
	4,  // 8: typedb.protocol.NumericGroup.number:type_name -> typedb.protocol.Numeric
	12, // 9: typedb.protocol.ConceptMap.MapEntry.value:type_name -> typedb.protocol.Concept
	2,  // 10: typedb.protocol.Explainables.RelationsEntry.value:type_name -> typedb.protocol.Explainable
	2,  // 11: typedb.protocol.Explainables.AttributesEntry.value:type_name -> typedb.protocol.Explainable
	10, // 12: typedb.protocol.Explainables.OwnershipsEntry.value:type_name -> typedb.protocol.Explainables.Owned
	11, // 13: typedb.protocol.Explainables.Owned.owned:type_name -> typedb.protocol.Explainables.Owned.OwnedEntry
	2,  // 14: typedb.protocol.Explainables.Owned.OwnedEntry.value:type_name -> typedb.protocol.Explainable
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_v2_protobuf_common_answer_proto_init() }
func file_v2_protobuf_common_answer_proto_init() {
	if File_v2_protobuf_common_answer_proto != nil {
		return
	}
	file_v2_protobuf_common_concept_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v2_protobuf_common_answer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConceptMap); i {
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
		file_v2_protobuf_common_answer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Explainables); i {
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
		file_v2_protobuf_common_answer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Explainable); i {
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
		file_v2_protobuf_common_answer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConceptMapGroup); i {
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
		file_v2_protobuf_common_answer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Numeric); i {
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
		file_v2_protobuf_common_answer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumericGroup); i {
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
		file_v2_protobuf_common_answer_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Explainables_Owned); i {
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
	file_v2_protobuf_common_answer_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Numeric_LongValue)(nil),
		(*Numeric_DoubleValue)(nil),
		(*Numeric_Nan)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v2_protobuf_common_answer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_protobuf_common_answer_proto_goTypes,
		DependencyIndexes: file_v2_protobuf_common_answer_proto_depIdxs,
		MessageInfos:      file_v2_protobuf_common_answer_proto_msgTypes,
	}.Build()
	File_v2_protobuf_common_answer_proto = out.File
	file_v2_protobuf_common_answer_proto_rawDesc = nil
	file_v2_protobuf_common_answer_proto_goTypes = nil
	file_v2_protobuf_common_answer_proto_depIdxs = nil
}
