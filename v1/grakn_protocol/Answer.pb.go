//
// Copyright (C) 2020 Grakn Labs
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
// source: v1/protobuf/Answer.proto

package grakn_protocol

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

type Answer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Answer:
	//	*Answer_AnswerGroup
	//	*Answer_ConceptMap
	//	*Answer_ConceptList
	//	*Answer_ConceptSet
	//	*Answer_ConceptSetMeasure
	//	*Answer_Value
	//	*Answer_Void
	Answer isAnswer_Answer `protobuf_oneof:"answer"`
}

func (x *Answer) Reset() {
	*x = Answer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_protobuf_Answer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Answer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Answer) ProtoMessage() {}

func (x *Answer) ProtoReflect() protoreflect.Message {
	mi := &file_v1_protobuf_Answer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Answer.ProtoReflect.Descriptor instead.
func (*Answer) Descriptor() ([]byte, []int) {
	return file_v1_protobuf_Answer_proto_rawDescGZIP(), []int{0}
}

func (m *Answer) GetAnswer() isAnswer_Answer {
	if m != nil {
		return m.Answer
	}
	return nil
}

func (x *Answer) GetAnswerGroup() *AnswerGroup {
	if x, ok := x.GetAnswer().(*Answer_AnswerGroup); ok {
		return x.AnswerGroup
	}
	return nil
}

func (x *Answer) GetConceptMap() *ConceptMap {
	if x, ok := x.GetAnswer().(*Answer_ConceptMap); ok {
		return x.ConceptMap
	}
	return nil
}

func (x *Answer) GetConceptList() *ConceptList {
	if x, ok := x.GetAnswer().(*Answer_ConceptList); ok {
		return x.ConceptList
	}
	return nil
}

func (x *Answer) GetConceptSet() *ConceptSet {
	if x, ok := x.GetAnswer().(*Answer_ConceptSet); ok {
		return x.ConceptSet
	}
	return nil
}

func (x *Answer) GetConceptSetMeasure() *ConceptSetMeasure {
	if x, ok := x.GetAnswer().(*Answer_ConceptSetMeasure); ok {
		return x.ConceptSetMeasure
	}
	return nil
}

func (x *Answer) GetValue() *Value {
	if x, ok := x.GetAnswer().(*Answer_Value); ok {
		return x.Value
	}
	return nil
}

func (x *Answer) GetVoid() *Void {
	if x, ok := x.GetAnswer().(*Answer_Void); ok {
		return x.Void
	}
	return nil
}

type isAnswer_Answer interface {
	isAnswer_Answer()
}

type Answer_AnswerGroup struct {
	AnswerGroup *AnswerGroup `protobuf:"bytes,1,opt,name=answerGroup,proto3,oneof"`
}

type Answer_ConceptMap struct {
	ConceptMap *ConceptMap `protobuf:"bytes,2,opt,name=conceptMap,proto3,oneof"`
}

type Answer_ConceptList struct {
	ConceptList *ConceptList `protobuf:"bytes,3,opt,name=conceptList,proto3,oneof"`
}

type Answer_ConceptSet struct {
	ConceptSet *ConceptSet `protobuf:"bytes,4,opt,name=conceptSet,proto3,oneof"`
}

type Answer_ConceptSetMeasure struct {
	ConceptSetMeasure *ConceptSetMeasure `protobuf:"bytes,5,opt,name=conceptSetMeasure,proto3,oneof"`
}

type Answer_Value struct {
	Value *Value `protobuf:"bytes,6,opt,name=value,proto3,oneof"`
}

type Answer_Void struct {
	Void *Void `protobuf:"bytes,7,opt,name=void,proto3,oneof"`
}

func (*Answer_AnswerGroup) isAnswer_Answer() {}

func (*Answer_ConceptMap) isAnswer_Answer() {}

func (*Answer_ConceptList) isAnswer_Answer() {}

func (*Answer_ConceptSet) isAnswer_Answer() {}

func (*Answer_ConceptSetMeasure) isAnswer_Answer() {}

func (*Answer_Value) isAnswer_Answer() {}

func (*Answer_Void) isAnswer_Answer() {}

type Explanation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Explanation) Reset() {
	*x = Explanation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_protobuf_Answer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Explanation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Explanation) ProtoMessage() {}

func (x *Explanation) ProtoReflect() protoreflect.Message {
	mi := &file_v1_protobuf_Answer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Explanation.ProtoReflect.Descriptor instead.
func (*Explanation) Descriptor() ([]byte, []int) {
	return file_v1_protobuf_Answer_proto_rawDescGZIP(), []int{1}
}

type AnswerGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner   *Concept  `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Answers []*Answer `protobuf:"bytes,2,rep,name=answers,proto3" json:"answers,omitempty"`
}

func (x *AnswerGroup) Reset() {
	*x = AnswerGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_protobuf_Answer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnswerGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnswerGroup) ProtoMessage() {}

func (x *AnswerGroup) ProtoReflect() protoreflect.Message {
	mi := &file_v1_protobuf_Answer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnswerGroup.ProtoReflect.Descriptor instead.
func (*AnswerGroup) Descriptor() ([]byte, []int) {
	return file_v1_protobuf_Answer_proto_rawDescGZIP(), []int{2}
}

func (x *AnswerGroup) GetOwner() *Concept {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *AnswerGroup) GetAnswers() []*Answer {
	if x != nil {
		return x.Answers
	}
	return nil
}

type ConceptMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Map            map[string]*Concept `protobuf:"bytes,1,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Pattern        string              `protobuf:"bytes,2,opt,name=pattern,proto3" json:"pattern,omitempty"`
	HasExplanation bool                `protobuf:"varint,3,opt,name=hasExplanation,proto3" json:"hasExplanation,omitempty"`
}

func (x *ConceptMap) Reset() {
	*x = ConceptMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_protobuf_Answer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConceptMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConceptMap) ProtoMessage() {}

func (x *ConceptMap) ProtoReflect() protoreflect.Message {
	mi := &file_v1_protobuf_Answer_proto_msgTypes[3]
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
	return file_v1_protobuf_Answer_proto_rawDescGZIP(), []int{3}
}

func (x *ConceptMap) GetMap() map[string]*Concept {
	if x != nil {
		return x.Map
	}
	return nil
}

func (x *ConceptMap) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

func (x *ConceptMap) GetHasExplanation() bool {
	if x != nil {
		return x.HasExplanation
	}
	return false
}

type ConceptList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List *ConceptIds `protobuf:"bytes,1,opt,name=list,proto3" json:"list,omitempty"`
}

func (x *ConceptList) Reset() {
	*x = ConceptList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_protobuf_Answer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConceptList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConceptList) ProtoMessage() {}

func (x *ConceptList) ProtoReflect() protoreflect.Message {
	mi := &file_v1_protobuf_Answer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConceptList.ProtoReflect.Descriptor instead.
func (*ConceptList) Descriptor() ([]byte, []int) {
	return file_v1_protobuf_Answer_proto_rawDescGZIP(), []int{4}
}

func (x *ConceptList) GetList() *ConceptIds {
	if x != nil {
		return x.List
	}
	return nil
}

type ConceptSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set *ConceptIds `protobuf:"bytes,1,opt,name=set,proto3" json:"set,omitempty"`
}

func (x *ConceptSet) Reset() {
	*x = ConceptSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_protobuf_Answer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConceptSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConceptSet) ProtoMessage() {}

func (x *ConceptSet) ProtoReflect() protoreflect.Message {
	mi := &file_v1_protobuf_Answer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConceptSet.ProtoReflect.Descriptor instead.
func (*ConceptSet) Descriptor() ([]byte, []int) {
	return file_v1_protobuf_Answer_proto_rawDescGZIP(), []int{5}
}

func (x *ConceptSet) GetSet() *ConceptIds {
	if x != nil {
		return x.Set
	}
	return nil
}

type ConceptSetMeasure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Set         *ConceptIds `protobuf:"bytes,1,opt,name=set,proto3" json:"set,omitempty"`
	Measurement *Number     `protobuf:"bytes,2,opt,name=measurement,proto3" json:"measurement,omitempty"`
}

func (x *ConceptSetMeasure) Reset() {
	*x = ConceptSetMeasure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_protobuf_Answer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConceptSetMeasure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConceptSetMeasure) ProtoMessage() {}

func (x *ConceptSetMeasure) ProtoReflect() protoreflect.Message {
	mi := &file_v1_protobuf_Answer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConceptSetMeasure.ProtoReflect.Descriptor instead.
func (*ConceptSetMeasure) Descriptor() ([]byte, []int) {
	return file_v1_protobuf_Answer_proto_rawDescGZIP(), []int{6}
}

func (x *ConceptSetMeasure) GetSet() *ConceptIds {
	if x != nil {
		return x.Set
	}
	return nil
}

func (x *ConceptSetMeasure) GetMeasurement() *Number {
	if x != nil {
		return x.Measurement
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number *Number `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_protobuf_Answer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_v1_protobuf_Answer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_v1_protobuf_Answer_proto_rawDescGZIP(), []int{7}
}

func (x *Value) GetNumber() *Number {
	if x != nil {
		return x.Number
	}
	return nil
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_protobuf_Answer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_v1_protobuf_Answer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_v1_protobuf_Answer_proto_rawDescGZIP(), []int{8}
}

func (x *Void) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ConceptIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ConceptIds) Reset() {
	*x = ConceptIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_protobuf_Answer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConceptIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConceptIds) ProtoMessage() {}

func (x *ConceptIds) ProtoReflect() protoreflect.Message {
	mi := &file_v1_protobuf_Answer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConceptIds.ProtoReflect.Descriptor instead.
func (*ConceptIds) Descriptor() ([]byte, []int) {
	return file_v1_protobuf_Answer_proto_rawDescGZIP(), []int{9}
}

func (x *ConceptIds) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type Number struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"` // We use a string to contain the full representation of a number
}

func (x *Number) Reset() {
	*x = Number{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_protobuf_Answer_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Number) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Number) ProtoMessage() {}

func (x *Number) ProtoReflect() protoreflect.Message {
	mi := &file_v1_protobuf_Answer_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Number.ProtoReflect.Descriptor instead.
func (*Number) Descriptor() ([]byte, []int) {
	return file_v1_protobuf_Answer_proto_rawDescGZIP(), []int{10}
}

func (x *Number) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Explanation_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Explainable *ConceptMap `protobuf:"bytes,1,opt,name=explainable,proto3" json:"explainable,omitempty"`
}

func (x *Explanation_Req) Reset() {
	*x = Explanation_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_protobuf_Answer_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Explanation_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Explanation_Req) ProtoMessage() {}

func (x *Explanation_Req) ProtoReflect() protoreflect.Message {
	mi := &file_v1_protobuf_Answer_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Explanation_Req.ProtoReflect.Descriptor instead.
func (*Explanation_Req) Descriptor() ([]byte, []int) {
	return file_v1_protobuf_Answer_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Explanation_Req) GetExplainable() *ConceptMap {
	if x != nil {
		return x.Explainable
	}
	return nil
}

type Explanation_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Explanation []*ConceptMap `protobuf:"bytes,1,rep,name=explanation,proto3" json:"explanation,omitempty"`
	Rule        *Concept      `protobuf:"bytes,2,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (x *Explanation_Res) Reset() {
	*x = Explanation_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_protobuf_Answer_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Explanation_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Explanation_Res) ProtoMessage() {}

func (x *Explanation_Res) ProtoReflect() protoreflect.Message {
	mi := &file_v1_protobuf_Answer_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Explanation_Res.ProtoReflect.Descriptor instead.
func (*Explanation_Res) Descriptor() ([]byte, []int) {
	return file_v1_protobuf_Answer_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Explanation_Res) GetExplanation() []*ConceptMap {
	if x != nil {
		return x.Explanation
	}
	return nil
}

func (x *Explanation_Res) GetRule() *Concept {
	if x != nil {
		return x.Rule
	}
	return nil
}

var File_v1_protobuf_Answer_proto protoreflect.FileDescriptor

var file_v1_protobuf_Answer_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d,
	0x03, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0b, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x35, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x4d, 0x61,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x4d, 0x61, 0x70, 0x48, 0x00, 0x52, 0x0a,
	0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x4d, 0x61, 0x70, 0x12, 0x38, 0x0a, 0x0b, 0x63, 0x6f,
	0x6e, 0x63, 0x65, 0x70, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x53,
	0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x53, 0x65, 0x74, 0x48, 0x00, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x53, 0x65, 0x74, 0x12, 0x4a, 0x0a, 0x11, 0x63,
	0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x48, 0x00, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x53, 0x65, 0x74,
	0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x23, 0x0a, 0x04, 0x76, 0x6f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x48, 0x00, 0x52, 0x04,
	0x76, 0x6f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0xaf,
	0x01, 0x0a, 0x0b, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3c,
	0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x35, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x4d, 0x61, 0x70, 0x52,
	0x0b, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x1a, 0x62, 0x0a, 0x03,
	0x52, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x4d, 0x61, 0x70, 0x52, 0x0b, 0x65,
	0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x04, 0x72, 0x75,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65,
	0x22, 0x60, 0x0a, 0x0b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x26, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x07, 0x61, 0x6e, 0x73, 0x77, 0x65,
	0x72, 0x73, 0x22, 0xc8, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x4d, 0x61,
	0x70, 0x12, 0x2e, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74,
	0x4d, 0x61, 0x70, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6d, 0x61,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x68,
	0x61, 0x73, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x68, 0x61, 0x73, 0x45, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x48, 0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65,
	0x70, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x36, 0x0a,
	0x0b, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x49, 0x64, 0x73, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x33, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74,
	0x53, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65,
	0x70, 0x74, 0x49, 0x64, 0x73, 0x52, 0x03, 0x73, 0x65, 0x74, 0x22, 0x6d, 0x0a, 0x11, 0x43, 0x6f,
	0x6e, 0x63, 0x65, 0x70, 0x74, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x12,
	0x25, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x49, 0x64,
	0x73, 0x52, 0x03, 0x73, 0x65, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0b, 0x6d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x30, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x20, 0x0a, 0x04, 0x56,
	0x6f, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1e, 0x0a,
	0x0a, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x70, 0x74, 0x49, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x1e, 0x0a,
	0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x49, 0x0a,
	0x16, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x22, 0x2e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x61, 0x6b, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3b, 0x67, 0x72, 0x61, 0x6b, 0x6e, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_protobuf_Answer_proto_rawDescOnce sync.Once
	file_v1_protobuf_Answer_proto_rawDescData = file_v1_protobuf_Answer_proto_rawDesc
)

func file_v1_protobuf_Answer_proto_rawDescGZIP() []byte {
	file_v1_protobuf_Answer_proto_rawDescOnce.Do(func() {
		file_v1_protobuf_Answer_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_protobuf_Answer_proto_rawDescData)
	})
	return file_v1_protobuf_Answer_proto_rawDescData
}

var file_v1_protobuf_Answer_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_v1_protobuf_Answer_proto_goTypes = []interface{}{
	(*Answer)(nil),            // 0: session.Answer
	(*Explanation)(nil),       // 1: session.Explanation
	(*AnswerGroup)(nil),       // 2: session.AnswerGroup
	(*ConceptMap)(nil),        // 3: session.ConceptMap
	(*ConceptList)(nil),       // 4: session.ConceptList
	(*ConceptSet)(nil),        // 5: session.ConceptSet
	(*ConceptSetMeasure)(nil), // 6: session.ConceptSetMeasure
	(*Value)(nil),             // 7: session.Value
	(*Void)(nil),              // 8: session.Void
	(*ConceptIds)(nil),        // 9: session.ConceptIds
	(*Number)(nil),            // 10: session.Number
	(*Explanation_Req)(nil),   // 11: session.Explanation.Req
	(*Explanation_Res)(nil),   // 12: session.Explanation.Res
	nil,                       // 13: session.ConceptMap.MapEntry
	(*Concept)(nil),           // 14: session.Concept
}
var file_v1_protobuf_Answer_proto_depIdxs = []int32{
	2,  // 0: session.Answer.answerGroup:type_name -> session.AnswerGroup
	3,  // 1: session.Answer.conceptMap:type_name -> session.ConceptMap
	4,  // 2: session.Answer.conceptList:type_name -> session.ConceptList
	5,  // 3: session.Answer.conceptSet:type_name -> session.ConceptSet
	6,  // 4: session.Answer.conceptSetMeasure:type_name -> session.ConceptSetMeasure
	7,  // 5: session.Answer.value:type_name -> session.Value
	8,  // 6: session.Answer.void:type_name -> session.Void
	14, // 7: session.AnswerGroup.owner:type_name -> session.Concept
	0,  // 8: session.AnswerGroup.answers:type_name -> session.Answer
	13, // 9: session.ConceptMap.map:type_name -> session.ConceptMap.MapEntry
	9,  // 10: session.ConceptList.list:type_name -> session.ConceptIds
	9,  // 11: session.ConceptSet.set:type_name -> session.ConceptIds
	9,  // 12: session.ConceptSetMeasure.set:type_name -> session.ConceptIds
	10, // 13: session.ConceptSetMeasure.measurement:type_name -> session.Number
	10, // 14: session.Value.number:type_name -> session.Number
	3,  // 15: session.Explanation.Req.explainable:type_name -> session.ConceptMap
	3,  // 16: session.Explanation.Res.explanation:type_name -> session.ConceptMap
	14, // 17: session.Explanation.Res.rule:type_name -> session.Concept
	14, // 18: session.ConceptMap.MapEntry.value:type_name -> session.Concept
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_v1_protobuf_Answer_proto_init() }
func file_v1_protobuf_Answer_proto_init() {
	if File_v1_protobuf_Answer_proto != nil {
		return
	}
	file_v1_protobuf_Concept_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_protobuf_Answer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Answer); i {
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
		file_v1_protobuf_Answer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Explanation); i {
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
		file_v1_protobuf_Answer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnswerGroup); i {
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
		file_v1_protobuf_Answer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_v1_protobuf_Answer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConceptList); i {
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
		file_v1_protobuf_Answer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConceptSet); i {
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
		file_v1_protobuf_Answer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConceptSetMeasure); i {
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
		file_v1_protobuf_Answer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_v1_protobuf_Answer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_v1_protobuf_Answer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConceptIds); i {
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
		file_v1_protobuf_Answer_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Number); i {
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
		file_v1_protobuf_Answer_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Explanation_Req); i {
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
		file_v1_protobuf_Answer_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Explanation_Res); i {
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
	file_v1_protobuf_Answer_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Answer_AnswerGroup)(nil),
		(*Answer_ConceptMap)(nil),
		(*Answer_ConceptList)(nil),
		(*Answer_ConceptSet)(nil),
		(*Answer_ConceptSetMeasure)(nil),
		(*Answer_Value)(nil),
		(*Answer_Void)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_protobuf_Answer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_protobuf_Answer_proto_goTypes,
		DependencyIndexes: file_v1_protobuf_Answer_proto_depIdxs,
		MessageInfos:      file_v1_protobuf_Answer_proto_msgTypes,
	}.Build()
	File_v1_protobuf_Answer_proto = out.File
	file_v1_protobuf_Answer_proto_rawDesc = nil
	file_v1_protobuf_Answer_proto_goTypes = nil
	file_v1_protobuf_Answer_proto_depIdxs = nil
}
