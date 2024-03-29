//
// Copyright (C) 2022 Vaticle
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
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: v2/protobuf/cluster/cluster_database.proto

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

type ClusterDatabaseManager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterDatabaseManager) Reset() {
	*x = ClusterDatabaseManager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterDatabaseManager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterDatabaseManager) ProtoMessage() {}

func (x *ClusterDatabaseManager) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterDatabaseManager.ProtoReflect.Descriptor instead.
func (*ClusterDatabaseManager) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_database_proto_rawDescGZIP(), []int{0}
}

type ClusterDatabase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Replicas []*ClusterDatabase_Replica `protobuf:"bytes,2,rep,name=replicas,proto3" json:"replicas,omitempty"`
}

func (x *ClusterDatabase) Reset() {
	*x = ClusterDatabase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterDatabase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterDatabase) ProtoMessage() {}

func (x *ClusterDatabase) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterDatabase.ProtoReflect.Descriptor instead.
func (*ClusterDatabase) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_database_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterDatabase) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClusterDatabase) GetReplicas() []*ClusterDatabase_Replica {
	if x != nil {
		return x.Replicas
	}
	return nil
}

type ClusterDatabaseManager_Get struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterDatabaseManager_Get) Reset() {
	*x = ClusterDatabaseManager_Get{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterDatabaseManager_Get) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterDatabaseManager_Get) ProtoMessage() {}

func (x *ClusterDatabaseManager_Get) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterDatabaseManager_Get.ProtoReflect.Descriptor instead.
func (*ClusterDatabaseManager_Get) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_database_proto_rawDescGZIP(), []int{0, 0}
}

type ClusterDatabaseManager_All struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterDatabaseManager_All) Reset() {
	*x = ClusterDatabaseManager_All{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterDatabaseManager_All) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterDatabaseManager_All) ProtoMessage() {}

func (x *ClusterDatabaseManager_All) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterDatabaseManager_All.ProtoReflect.Descriptor instead.
func (*ClusterDatabaseManager_All) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_database_proto_rawDescGZIP(), []int{0, 1}
}

type ClusterDatabaseManager_Get_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ClusterDatabaseManager_Get_Req) Reset() {
	*x = ClusterDatabaseManager_Get_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterDatabaseManager_Get_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterDatabaseManager_Get_Req) ProtoMessage() {}

func (x *ClusterDatabaseManager_Get_Req) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterDatabaseManager_Get_Req.ProtoReflect.Descriptor instead.
func (*ClusterDatabaseManager_Get_Req) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_database_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ClusterDatabaseManager_Get_Req) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ClusterDatabaseManager_Get_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database *ClusterDatabase `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
}

func (x *ClusterDatabaseManager_Get_Res) Reset() {
	*x = ClusterDatabaseManager_Get_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterDatabaseManager_Get_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterDatabaseManager_Get_Res) ProtoMessage() {}

func (x *ClusterDatabaseManager_Get_Res) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterDatabaseManager_Get_Res.ProtoReflect.Descriptor instead.
func (*ClusterDatabaseManager_Get_Res) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_database_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *ClusterDatabaseManager_Get_Res) GetDatabase() *ClusterDatabase {
	if x != nil {
		return x.Database
	}
	return nil
}

type ClusterDatabaseManager_All_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClusterDatabaseManager_All_Req) Reset() {
	*x = ClusterDatabaseManager_All_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterDatabaseManager_All_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterDatabaseManager_All_Req) ProtoMessage() {}

func (x *ClusterDatabaseManager_All_Req) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterDatabaseManager_All_Req.ProtoReflect.Descriptor instead.
func (*ClusterDatabaseManager_All_Req) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_database_proto_rawDescGZIP(), []int{0, 1, 0}
}

type ClusterDatabaseManager_All_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Databases []*ClusterDatabase `protobuf:"bytes,1,rep,name=databases,proto3" json:"databases,omitempty"`
}

func (x *ClusterDatabaseManager_All_Res) Reset() {
	*x = ClusterDatabaseManager_All_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterDatabaseManager_All_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterDatabaseManager_All_Res) ProtoMessage() {}

func (x *ClusterDatabaseManager_All_Res) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterDatabaseManager_All_Res.ProtoReflect.Descriptor instead.
func (*ClusterDatabaseManager_All_Res) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_database_proto_rawDescGZIP(), []int{0, 1, 1}
}

func (x *ClusterDatabaseManager_All_Res) GetDatabases() []*ClusterDatabase {
	if x != nil {
		return x.Databases
	}
	return nil
}

type ClusterDatabase_Replica struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Primary   bool   `protobuf:"varint,2,opt,name=primary,proto3" json:"primary,omitempty"`
	Preferred bool   `protobuf:"varint,3,opt,name=preferred,proto3" json:"preferred,omitempty"`
	Term      int64  `protobuf:"varint,4,opt,name=term,proto3" json:"term,omitempty"`
}

func (x *ClusterDatabase_Replica) Reset() {
	*x = ClusterDatabase_Replica{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterDatabase_Replica) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterDatabase_Replica) ProtoMessage() {}

func (x *ClusterDatabase_Replica) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_database_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterDatabase_Replica.ProtoReflect.Descriptor instead.
func (*ClusterDatabase_Replica) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_database_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ClusterDatabase_Replica) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ClusterDatabase_Replica) GetPrimary() bool {
	if x != nil {
		return x.Primary
	}
	return false
}

func (x *ClusterDatabase_Replica) GetPreferred() bool {
	if x != nil {
		return x.Preferred
	}
	return false
}

func (x *ClusterDatabase_Replica) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

var File_v2_protobuf_cluster_cluster_database_proto protoreflect.FileDescriptor

var file_v2_protobuf_cluster_cluster_database_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x79,
	0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0xd4, 0x01,
	0x0a, 0x16, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a, 0x65, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x1a,
	0x19, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x43, 0x0a, 0x03, 0x52, 0x65,
	0x73, 0x12, 0x3c, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x1a,
	0x53, 0x0a, 0x03, 0x41, 0x6c, 0x6c, 0x1a, 0x05, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x1a, 0x45, 0x0a,
	0x03, 0x52, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x73, 0x22, 0xdc, 0x01, 0x0a, 0x0f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x08,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x73, 0x1a, 0x6f, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74,
	0x65, 0x72, 0x6d, 0x42, 0x59, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x61, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x42, 0x14, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x24, 0x2e, 0x2f, 0x76, 0x32, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x64, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x3b, 0x74,
	0x79, 0x70, 0x65, 0x64, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_protobuf_cluster_cluster_database_proto_rawDescOnce sync.Once
	file_v2_protobuf_cluster_cluster_database_proto_rawDescData = file_v2_protobuf_cluster_cluster_database_proto_rawDesc
)

func file_v2_protobuf_cluster_cluster_database_proto_rawDescGZIP() []byte {
	file_v2_protobuf_cluster_cluster_database_proto_rawDescOnce.Do(func() {
		file_v2_protobuf_cluster_cluster_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_protobuf_cluster_cluster_database_proto_rawDescData)
	})
	return file_v2_protobuf_cluster_cluster_database_proto_rawDescData
}

var file_v2_protobuf_cluster_cluster_database_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_v2_protobuf_cluster_cluster_database_proto_goTypes = []interface{}{
	(*ClusterDatabaseManager)(nil),         // 0: typedb.protocol.ClusterDatabaseManager
	(*ClusterDatabase)(nil),                // 1: typedb.protocol.ClusterDatabase
	(*ClusterDatabaseManager_Get)(nil),     // 2: typedb.protocol.ClusterDatabaseManager.Get
	(*ClusterDatabaseManager_All)(nil),     // 3: typedb.protocol.ClusterDatabaseManager.All
	(*ClusterDatabaseManager_Get_Req)(nil), // 4: typedb.protocol.ClusterDatabaseManager.Get.Req
	(*ClusterDatabaseManager_Get_Res)(nil), // 5: typedb.protocol.ClusterDatabaseManager.Get.Res
	(*ClusterDatabaseManager_All_Req)(nil), // 6: typedb.protocol.ClusterDatabaseManager.All.Req
	(*ClusterDatabaseManager_All_Res)(nil), // 7: typedb.protocol.ClusterDatabaseManager.All.Res
	(*ClusterDatabase_Replica)(nil),        // 8: typedb.protocol.ClusterDatabase.Replica
}
var file_v2_protobuf_cluster_cluster_database_proto_depIdxs = []int32{
	8, // 0: typedb.protocol.ClusterDatabase.replicas:type_name -> typedb.protocol.ClusterDatabase.Replica
	1, // 1: typedb.protocol.ClusterDatabaseManager.Get.Res.database:type_name -> typedb.protocol.ClusterDatabase
	1, // 2: typedb.protocol.ClusterDatabaseManager.All.Res.databases:type_name -> typedb.protocol.ClusterDatabase
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v2_protobuf_cluster_cluster_database_proto_init() }
func file_v2_protobuf_cluster_cluster_database_proto_init() {
	if File_v2_protobuf_cluster_cluster_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_protobuf_cluster_cluster_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterDatabaseManager); i {
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
		file_v2_protobuf_cluster_cluster_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterDatabase); i {
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
		file_v2_protobuf_cluster_cluster_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterDatabaseManager_Get); i {
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
		file_v2_protobuf_cluster_cluster_database_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterDatabaseManager_All); i {
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
		file_v2_protobuf_cluster_cluster_database_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterDatabaseManager_Get_Req); i {
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
		file_v2_protobuf_cluster_cluster_database_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterDatabaseManager_Get_Res); i {
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
		file_v2_protobuf_cluster_cluster_database_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterDatabaseManager_All_Req); i {
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
		file_v2_protobuf_cluster_cluster_database_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterDatabaseManager_All_Res); i {
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
		file_v2_protobuf_cluster_cluster_database_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterDatabase_Replica); i {
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
			RawDescriptor: file_v2_protobuf_cluster_cluster_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_protobuf_cluster_cluster_database_proto_goTypes,
		DependencyIndexes: file_v2_protobuf_cluster_cluster_database_proto_depIdxs,
		MessageInfos:      file_v2_protobuf_cluster_cluster_database_proto_msgTypes,
	}.Build()
	File_v2_protobuf_cluster_cluster_database_proto = out.File
	file_v2_protobuf_cluster_cluster_database_proto_rawDesc = nil
	file_v2_protobuf_cluster_cluster_database_proto_goTypes = nil
	file_v2_protobuf_cluster_cluster_database_proto_depIdxs = nil
}
