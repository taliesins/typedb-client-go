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
// 	protoc        v3.20.0
// source: v2/protobuf/common/session.proto

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

type Session_Type int32

const (
	Session_DATA   Session_Type = 0
	Session_SCHEMA Session_Type = 1
)

// Enum value maps for Session_Type.
var (
	Session_Type_name = map[int32]string{
		0: "DATA",
		1: "SCHEMA",
	}
	Session_Type_value = map[string]int32{
		"DATA":   0,
		"SCHEMA": 1,
	}
)

func (x Session_Type) Enum() *Session_Type {
	p := new(Session_Type)
	*p = x
	return p
}

func (x Session_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Session_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_v2_protobuf_common_session_proto_enumTypes[0].Descriptor()
}

func (Session_Type) Type() protoreflect.EnumType {
	return &file_v2_protobuf_common_session_proto_enumTypes[0]
}

func (x Session_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Session_Type.Descriptor instead.
func (Session_Type) EnumDescriptor() ([]byte, []int) {
	return file_v2_protobuf_common_session_proto_rawDescGZIP(), []int{0, 0}
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_session_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_session_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_session_proto_rawDescGZIP(), []int{0}
}

type Session_Open struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Session_Open) Reset() {
	*x = Session_Open{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_session_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session_Open) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session_Open) ProtoMessage() {}

func (x *Session_Open) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_session_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session_Open.ProtoReflect.Descriptor instead.
func (*Session_Open) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_session_proto_rawDescGZIP(), []int{0, 0}
}

type Session_Close struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Session_Close) Reset() {
	*x = Session_Close{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_session_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session_Close) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session_Close) ProtoMessage() {}

func (x *Session_Close) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_session_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session_Close.ProtoReflect.Descriptor instead.
func (*Session_Close) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_session_proto_rawDescGZIP(), []int{0, 1}
}

type Session_Pulse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Session_Pulse) Reset() {
	*x = Session_Pulse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_session_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session_Pulse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session_Pulse) ProtoMessage() {}

func (x *Session_Pulse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_session_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session_Pulse.ProtoReflect.Descriptor instead.
func (*Session_Pulse) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_session_proto_rawDescGZIP(), []int{0, 2}
}

type Session_Open_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database string       `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Type     Session_Type `protobuf:"varint,2,opt,name=type,proto3,enum=typedb.protocol.Session_Type" json:"type,omitempty"`
	Options  *Options     `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *Session_Open_Req) Reset() {
	*x = Session_Open_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_session_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session_Open_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session_Open_Req) ProtoMessage() {}

func (x *Session_Open_Req) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_session_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session_Open_Req.ProtoReflect.Descriptor instead.
func (*Session_Open_Req) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_session_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Session_Open_Req) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *Session_Open_Req) GetType() Session_Type {
	if x != nil {
		return x.Type
	}
	return Session_DATA
}

func (x *Session_Open_Req) GetOptions() *Options {
	if x != nil {
		return x.Options
	}
	return nil
}

type Session_Open_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId            []byte `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	ServerDurationMillis int32  `protobuf:"varint,2,opt,name=server_duration_millis,json=serverDurationMillis,proto3" json:"server_duration_millis,omitempty"`
}

func (x *Session_Open_Res) Reset() {
	*x = Session_Open_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_session_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session_Open_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session_Open_Res) ProtoMessage() {}

func (x *Session_Open_Res) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_session_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session_Open_Res.ProtoReflect.Descriptor instead.
func (*Session_Open_Res) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_session_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *Session_Open_Res) GetSessionId() []byte {
	if x != nil {
		return x.SessionId
	}
	return nil
}

func (x *Session_Open_Res) GetServerDurationMillis() int32 {
	if x != nil {
		return x.ServerDurationMillis
	}
	return 0
}

type Session_Close_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId []byte `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *Session_Close_Req) Reset() {
	*x = Session_Close_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_session_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session_Close_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session_Close_Req) ProtoMessage() {}

func (x *Session_Close_Req) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_session_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session_Close_Req.ProtoReflect.Descriptor instead.
func (*Session_Close_Req) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_session_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *Session_Close_Req) GetSessionId() []byte {
	if x != nil {
		return x.SessionId
	}
	return nil
}

type Session_Close_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Session_Close_Res) Reset() {
	*x = Session_Close_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_session_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session_Close_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session_Close_Res) ProtoMessage() {}

func (x *Session_Close_Res) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_session_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session_Close_Res.ProtoReflect.Descriptor instead.
func (*Session_Close_Res) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_session_proto_rawDescGZIP(), []int{0, 1, 1}
}

type Session_Pulse_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId []byte `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
}

func (x *Session_Pulse_Req) Reset() {
	*x = Session_Pulse_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_session_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session_Pulse_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session_Pulse_Req) ProtoMessage() {}

func (x *Session_Pulse_Req) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_session_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session_Pulse_Req.ProtoReflect.Descriptor instead.
func (*Session_Pulse_Req) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_session_proto_rawDescGZIP(), []int{0, 2, 0}
}

func (x *Session_Pulse_Req) GetSessionId() []byte {
	if x != nil {
		return x.SessionId
	}
	return nil
}

type Session_Pulse_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alive bool `protobuf:"varint,1,opt,name=alive,proto3" json:"alive,omitempty"`
}

func (x *Session_Pulse_Res) Reset() {
	*x = Session_Pulse_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_common_session_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session_Pulse_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session_Pulse_Res) ProtoMessage() {}

func (x *Session_Pulse_Res) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_common_session_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session_Pulse_Res.ProtoReflect.Descriptor instead.
func (*Session_Pulse_Res) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_common_session_proto_rawDescGZIP(), []int{0, 2, 1}
}

func (x *Session_Pulse_Res) GetAlive() bool {
	if x != nil {
		return x.Alive
	}
	return false
}

var File_v2_protobuf_common_session_proto protoreflect.FileDescriptor

var file_v2_protobuf_common_session_proto_rawDesc = []byte{
	0x0a, 0x20, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x1a, 0x20, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x03, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x1a, 0xed, 0x01, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x1a, 0x88, 0x01, 0x0a, 0x03, 0x52,
	0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x32, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x5a, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6c, 0x6c, 0x69,
	0x73, 0x1a, 0x34, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x1a, 0x24, 0x0a, 0x03, 0x52, 0x65,
	0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x1a, 0x05, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x1a, 0x4a, 0x0a, 0x05, 0x50, 0x75, 0x6c, 0x73, 0x65,
	0x1a, 0x24, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x1b, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x6c,
	0x69, 0x76, 0x65, 0x22, 0x1c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x44,
	0x41, 0x54, 0x41, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x10,
	0x01, 0x42, 0x51, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x42, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x24,
	0x2e, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_protobuf_common_session_proto_rawDescOnce sync.Once
	file_v2_protobuf_common_session_proto_rawDescData = file_v2_protobuf_common_session_proto_rawDesc
)

func file_v2_protobuf_common_session_proto_rawDescGZIP() []byte {
	file_v2_protobuf_common_session_proto_rawDescOnce.Do(func() {
		file_v2_protobuf_common_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_protobuf_common_session_proto_rawDescData)
	})
	return file_v2_protobuf_common_session_proto_rawDescData
}

var file_v2_protobuf_common_session_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v2_protobuf_common_session_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_v2_protobuf_common_session_proto_goTypes = []interface{}{
	(Session_Type)(0),         // 0: typedb.protocol.Session.Type
	(*Session)(nil),           // 1: typedb.protocol.Session
	(*Session_Open)(nil),      // 2: typedb.protocol.Session.Open
	(*Session_Close)(nil),     // 3: typedb.protocol.Session.Close
	(*Session_Pulse)(nil),     // 4: typedb.protocol.Session.Pulse
	(*Session_Open_Req)(nil),  // 5: typedb.protocol.Session.Open.Req
	(*Session_Open_Res)(nil),  // 6: typedb.protocol.Session.Open.Res
	(*Session_Close_Req)(nil), // 7: typedb.protocol.Session.Close.Req
	(*Session_Close_Res)(nil), // 8: typedb.protocol.Session.Close.Res
	(*Session_Pulse_Req)(nil), // 9: typedb.protocol.Session.Pulse.Req
	(*Session_Pulse_Res)(nil), // 10: typedb.protocol.Session.Pulse.Res
	(*Options)(nil),           // 11: typedb.protocol.Options
}
var file_v2_protobuf_common_session_proto_depIdxs = []int32{
	0,  // 0: typedb.protocol.Session.Open.Req.type:type_name -> typedb.protocol.Session.Type
	11, // 1: typedb.protocol.Session.Open.Req.options:type_name -> typedb.protocol.Options
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_v2_protobuf_common_session_proto_init() }
func file_v2_protobuf_common_session_proto_init() {
	if File_v2_protobuf_common_session_proto != nil {
		return
	}
	file_v2_protobuf_common_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v2_protobuf_common_session_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_v2_protobuf_common_session_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session_Open); i {
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
		file_v2_protobuf_common_session_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session_Close); i {
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
		file_v2_protobuf_common_session_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session_Pulse); i {
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
		file_v2_protobuf_common_session_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session_Open_Req); i {
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
		file_v2_protobuf_common_session_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session_Open_Res); i {
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
		file_v2_protobuf_common_session_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session_Close_Req); i {
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
		file_v2_protobuf_common_session_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session_Close_Res); i {
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
		file_v2_protobuf_common_session_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session_Pulse_Req); i {
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
		file_v2_protobuf_common_session_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session_Pulse_Res); i {
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
			RawDescriptor: file_v2_protobuf_common_session_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_protobuf_common_session_proto_goTypes,
		DependencyIndexes: file_v2_protobuf_common_session_proto_depIdxs,
		EnumInfos:         file_v2_protobuf_common_session_proto_enumTypes,
		MessageInfos:      file_v2_protobuf_common_session_proto_msgTypes,
	}.Build()
	File_v2_protobuf_common_session_proto = out.File
	file_v2_protobuf_common_session_proto_rawDesc = nil
	file_v2_protobuf_common_session_proto_goTypes = nil
	file_v2_protobuf_common_session_proto_depIdxs = nil
}
