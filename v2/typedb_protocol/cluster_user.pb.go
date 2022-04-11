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
// source: v2/protobuf/cluster/cluster_user.proto

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

type ClusterUserManager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterUserManager) Reset() {
	*x = ClusterUserManager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUserManager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUserManager) ProtoMessage() {}

func (x *ClusterUserManager) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUserManager.ProtoReflect.Descriptor instead.
func (*ClusterUserManager) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{0}
}

type ClusterUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterUser) Reset() {
	*x = ClusterUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUser) ProtoMessage() {}

func (x *ClusterUser) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUser.ProtoReflect.Descriptor instead.
func (*ClusterUser) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{1}
}

type ClusterUserManager_Contains struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterUserManager_Contains) Reset() {
	*x = ClusterUserManager_Contains{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUserManager_Contains) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUserManager_Contains) ProtoMessage() {}

func (x *ClusterUserManager_Contains) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUserManager_Contains.ProtoReflect.Descriptor instead.
func (*ClusterUserManager_Contains) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{0, 0}
}

type ClusterUserManager_Create struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterUserManager_Create) Reset() {
	*x = ClusterUserManager_Create{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUserManager_Create) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUserManager_Create) ProtoMessage() {}

func (x *ClusterUserManager_Create) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUserManager_Create.ProtoReflect.Descriptor instead.
func (*ClusterUserManager_Create) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{0, 1}
}

type ClusterUserManager_All struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterUserManager_All) Reset() {
	*x = ClusterUserManager_All{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUserManager_All) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUserManager_All) ProtoMessage() {}

func (x *ClusterUserManager_All) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUserManager_All.ProtoReflect.Descriptor instead.
func (*ClusterUserManager_All) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{0, 2}
}

type ClusterUserManager_Contains_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ClusterUserManager_Contains_Req) Reset() {
	*x = ClusterUserManager_Contains_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUserManager_Contains_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUserManager_Contains_Req) ProtoMessage() {}

func (x *ClusterUserManager_Contains_Req) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUserManager_Contains_Req.ProtoReflect.Descriptor instead.
func (*ClusterUserManager_Contains_Req) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ClusterUserManager_Contains_Req) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ClusterUserManager_Contains_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contains bool `protobuf:"varint,1,opt,name=contains,proto3" json:"contains,omitempty"`
}

func (x *ClusterUserManager_Contains_Res) Reset() {
	*x = ClusterUserManager_Contains_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUserManager_Contains_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUserManager_Contains_Res) ProtoMessage() {}

func (x *ClusterUserManager_Contains_Res) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUserManager_Contains_Res.ProtoReflect.Descriptor instead.
func (*ClusterUserManager_Contains_Res) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *ClusterUserManager_Contains_Res) GetContains() bool {
	if x != nil {
		return x.Contains
	}
	return false
}

type ClusterUserManager_Create_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ClusterUserManager_Create_Req) Reset() {
	*x = ClusterUserManager_Create_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUserManager_Create_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUserManager_Create_Req) ProtoMessage() {}

func (x *ClusterUserManager_Create_Req) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUserManager_Create_Req.ProtoReflect.Descriptor instead.
func (*ClusterUserManager_Create_Req) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *ClusterUserManager_Create_Req) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ClusterUserManager_Create_Req) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ClusterUserManager_Create_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterUserManager_Create_Res) Reset() {
	*x = ClusterUserManager_Create_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUserManager_Create_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUserManager_Create_Res) ProtoMessage() {}

func (x *ClusterUserManager_Create_Res) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUserManager_Create_Res.ProtoReflect.Descriptor instead.
func (*ClusterUserManager_Create_Res) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{0, 1, 1}
}

type ClusterUserManager_All_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterUserManager_All_Req) Reset() {
	*x = ClusterUserManager_All_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUserManager_All_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUserManager_All_Req) ProtoMessage() {}

func (x *ClusterUserManager_All_Req) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUserManager_All_Req.ProtoReflect.Descriptor instead.
func (*ClusterUserManager_All_Req) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{0, 2, 0}
}

type ClusterUserManager_All_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *ClusterUserManager_All_Res) Reset() {
	*x = ClusterUserManager_All_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUserManager_All_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUserManager_All_Res) ProtoMessage() {}

func (x *ClusterUserManager_All_Res) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUserManager_All_Res.ProtoReflect.Descriptor instead.
func (*ClusterUserManager_All_Res) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{0, 2, 1}
}

func (x *ClusterUserManager_All_Res) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type ClusterUser_Password struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterUser_Password) Reset() {
	*x = ClusterUser_Password{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUser_Password) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUser_Password) ProtoMessage() {}

func (x *ClusterUser_Password) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUser_Password.ProtoReflect.Descriptor instead.
func (*ClusterUser_Password) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{1, 0}
}

type ClusterUser_Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterUser_Token) Reset() {
	*x = ClusterUser_Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUser_Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUser_Token) ProtoMessage() {}

func (x *ClusterUser_Token) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUser_Token.ProtoReflect.Descriptor instead.
func (*ClusterUser_Token) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{1, 1}
}

type ClusterUser_Delete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterUser_Delete) Reset() {
	*x = ClusterUser_Delete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUser_Delete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUser_Delete) ProtoMessage() {}

func (x *ClusterUser_Delete) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUser_Delete.ProtoReflect.Descriptor instead.
func (*ClusterUser_Delete) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{1, 2}
}

type ClusterUser_Password_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ClusterUser_Password_Req) Reset() {
	*x = ClusterUser_Password_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUser_Password_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUser_Password_Req) ProtoMessage() {}

func (x *ClusterUser_Password_Req) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUser_Password_Req.ProtoReflect.Descriptor instead.
func (*ClusterUser_Password_Req) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *ClusterUser_Password_Req) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ClusterUser_Password_Req) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ClusterUser_Password_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterUser_Password_Res) Reset() {
	*x = ClusterUser_Password_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUser_Password_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUser_Password_Res) ProtoMessage() {}

func (x *ClusterUser_Password_Res) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUser_Password_Res.ProtoReflect.Descriptor instead.
func (*ClusterUser_Password_Res) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{1, 0, 1}
}

type ClusterUser_Token_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ClusterUser_Token_Req) Reset() {
	*x = ClusterUser_Token_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUser_Token_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUser_Token_Req) ProtoMessage() {}

func (x *ClusterUser_Token_Req) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUser_Token_Req.ProtoReflect.Descriptor instead.
func (*ClusterUser_Token_Req) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{1, 1, 0}
}

func (x *ClusterUser_Token_Req) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ClusterUser_Token_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ClusterUser_Token_Res) Reset() {
	*x = ClusterUser_Token_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUser_Token_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUser_Token_Res) ProtoMessage() {}

func (x *ClusterUser_Token_Res) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUser_Token_Res.ProtoReflect.Descriptor instead.
func (*ClusterUser_Token_Res) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{1, 1, 1}
}

func (x *ClusterUser_Token_Res) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ClusterUser_Delete_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *ClusterUser_Delete_Req) Reset() {
	*x = ClusterUser_Delete_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUser_Delete_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUser_Delete_Req) ProtoMessage() {}

func (x *ClusterUser_Delete_Req) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUser_Delete_Req.ProtoReflect.Descriptor instead.
func (*ClusterUser_Delete_Req) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{1, 2, 0}
}

func (x *ClusterUser_Delete_Req) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ClusterUser_Delete_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterUser_Delete_Res) Reset() {
	*x = ClusterUser_Delete_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUser_Delete_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUser_Delete_Res) ProtoMessage() {}

func (x *ClusterUser_Delete_Res) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_user_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUser_Delete_Res.ProtoReflect.Descriptor instead.
func (*ClusterUser_Delete_Res) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP(), []int{1, 2, 1}
}

var File_v2_protobuf_cluster_cluster_user_proto protoreflect.FileDescriptor

var file_v2_protobuf_cluster_cluster_user_proto_rawDesc = []byte{
	0x0a, 0x26, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0xe1, 0x01, 0x0a, 0x12, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x1a, 0x50, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x73, 0x1a, 0x21, 0x0a, 0x03,
	0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a,
	0x21, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x73, 0x1a, 0x4e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x3d, 0x0a, 0x03,
	0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x05, 0x0a, 0x03, 0x52,
	0x65, 0x73, 0x1a, 0x29, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x1a, 0x05, 0x0a, 0x03, 0x52, 0x65, 0x71,
	0x1a, 0x1b, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0xdc, 0x01,
	0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x50, 0x0a,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x3d, 0x0a, 0x03, 0x52, 0x65, 0x71,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x05, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x1a,
	0x47, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x21, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x1b, 0x0a, 0x03, 0x52,
	0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x32, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x1a, 0x21, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x05, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x42, 0x55, 0x0a, 0x1b,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x10, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x24, 0x2e,
	0x2f, 0x76, 0x32, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_protobuf_cluster_cluster_user_proto_rawDescOnce sync.Once
	file_v2_protobuf_cluster_cluster_user_proto_rawDescData = file_v2_protobuf_cluster_cluster_user_proto_rawDesc
)

func file_v2_protobuf_cluster_cluster_user_proto_rawDescGZIP() []byte {
	file_v2_protobuf_cluster_cluster_user_proto_rawDescOnce.Do(func() {
		file_v2_protobuf_cluster_cluster_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_protobuf_cluster_cluster_user_proto_rawDescData)
	})
	return file_v2_protobuf_cluster_cluster_user_proto_rawDescData
}

var file_v2_protobuf_cluster_cluster_user_proto_msgTypes = make([]protoimpl.MessageInfo, 20)
var file_v2_protobuf_cluster_cluster_user_proto_goTypes = []interface{}{
	(*ClusterUserManager)(nil),              // 0: typedb.protocol.ClusterUserManager
	(*ClusterUser)(nil),                     // 1: typedb.protocol.ClusterUser
	(*ClusterUserManager_Contains)(nil),     // 2: typedb.protocol.ClusterUserManager.Contains
	(*ClusterUserManager_Create)(nil),       // 3: typedb.protocol.ClusterUserManager.Create
	(*ClusterUserManager_All)(nil),          // 4: typedb.protocol.ClusterUserManager.All
	(*ClusterUserManager_Contains_Req)(nil), // 5: typedb.protocol.ClusterUserManager.Contains.Req
	(*ClusterUserManager_Contains_Res)(nil), // 6: typedb.protocol.ClusterUserManager.Contains.Res
	(*ClusterUserManager_Create_Req)(nil),   // 7: typedb.protocol.ClusterUserManager.Create.Req
	(*ClusterUserManager_Create_Res)(nil),   // 8: typedb.protocol.ClusterUserManager.Create.Res
	(*ClusterUserManager_All_Req)(nil),      // 9: typedb.protocol.ClusterUserManager.All.Req
	(*ClusterUserManager_All_Res)(nil),      // 10: typedb.protocol.ClusterUserManager.All.Res
	(*ClusterUser_Password)(nil),            // 11: typedb.protocol.ClusterUser.Password
	(*ClusterUser_Token)(nil),               // 12: typedb.protocol.ClusterUser.Token
	(*ClusterUser_Delete)(nil),              // 13: typedb.protocol.ClusterUser.Delete
	(*ClusterUser_Password_Req)(nil),        // 14: typedb.protocol.ClusterUser.Password.Req
	(*ClusterUser_Password_Res)(nil),        // 15: typedb.protocol.ClusterUser.Password.Res
	(*ClusterUser_Token_Req)(nil),           // 16: typedb.protocol.ClusterUser.Token.Req
	(*ClusterUser_Token_Res)(nil),           // 17: typedb.protocol.ClusterUser.Token.Res
	(*ClusterUser_Delete_Req)(nil),          // 18: typedb.protocol.ClusterUser.Delete.Req
	(*ClusterUser_Delete_Res)(nil),          // 19: typedb.protocol.ClusterUser.Delete.Res
}
var file_v2_protobuf_cluster_cluster_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v2_protobuf_cluster_cluster_user_proto_init() }
func file_v2_protobuf_cluster_cluster_user_proto_init() {
	if File_v2_protobuf_cluster_cluster_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUserManager); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUser); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUserManager_Contains); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUserManager_Create); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUserManager_All); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUserManager_Contains_Req); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUserManager_Contains_Res); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUserManager_Create_Req); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUserManager_Create_Res); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUserManager_All_Req); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUserManager_All_Res); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUser_Password); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUser_Token); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUser_Delete); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUser_Password_Req); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUser_Password_Res); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUser_Token_Req); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUser_Token_Res); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUser_Delete_Req); i {
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
		file_v2_protobuf_cluster_cluster_user_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUser_Delete_Res); i {
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
			RawDescriptor: file_v2_protobuf_cluster_cluster_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   20,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_protobuf_cluster_cluster_user_proto_goTypes,
		DependencyIndexes: file_v2_protobuf_cluster_cluster_user_proto_depIdxs,
		MessageInfos:      file_v2_protobuf_cluster_cluster_user_proto_msgTypes,
	}.Build()
	File_v2_protobuf_cluster_cluster_user_proto = out.File
	file_v2_protobuf_cluster_cluster_user_proto_rawDesc = nil
	file_v2_protobuf_cluster_cluster_user_proto_goTypes = nil
	file_v2_protobuf_cluster_cluster_user_proto_depIdxs = nil
}
