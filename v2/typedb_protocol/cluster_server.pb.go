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
// source: v2/protobuf/cluster/cluster_server.proto

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

type ServerManager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServerManager) Reset() {
	*x = ServerManager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerManager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerManager) ProtoMessage() {}

func (x *ServerManager) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerManager.ProtoReflect.Descriptor instead.
func (*ServerManager) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_server_proto_rawDescGZIP(), []int{0}
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_server_proto_rawDescGZIP(), []int{1}
}

func (x *Server) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ServerManager_All struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServerManager_All) Reset() {
	*x = ServerManager_All{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerManager_All) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerManager_All) ProtoMessage() {}

func (x *ServerManager_All) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerManager_All.ProtoReflect.Descriptor instead.
func (*ServerManager_All) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_server_proto_rawDescGZIP(), []int{0, 0}
}

type ServerManager_All_Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServerManager_All_Req) Reset() {
	*x = ServerManager_All_Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerManager_All_Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerManager_All_Req) ProtoMessage() {}

func (x *ServerManager_All_Req) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerManager_All_Req.ProtoReflect.Descriptor instead.
func (*ServerManager_All_Req) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_server_proto_rawDescGZIP(), []int{0, 0, 0}
}

type ServerManager_All_Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servers []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
}

func (x *ServerManager_All_Res) Reset() {
	*x = ServerManager_All_Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_protobuf_cluster_cluster_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerManager_All_Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerManager_All_Res) ProtoMessage() {}

func (x *ServerManager_All_Res) ProtoReflect() protoreflect.Message {
	mi := &file_v2_protobuf_cluster_cluster_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerManager_All_Res.ProtoReflect.Descriptor instead.
func (*ServerManager_All_Res) Descriptor() ([]byte, []int) {
	return file_v2_protobuf_cluster_cluster_server_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *ServerManager_All_Res) GetServers() []*Server {
	if x != nil {
		return x.Servers
	}
	return nil
}

var File_v2_protobuf_cluster_cluster_server_proto protoreflect.FileDescriptor

var file_v2_protobuf_cluster_cluster_server_proto_rawDesc = []byte{
	0x0a, 0x28, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74, 0x79, 0x70, 0x65,
	0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x57, 0x0a, 0x0d, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a, 0x46, 0x0a, 0x03,
	0x41, 0x6c, 0x6c, 0x1a, 0x05, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x1a, 0x38, 0x0a, 0x03, 0x52, 0x65,
	0x73, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x22, 0x22, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x57, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e,
	0x76, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x42, 0x12, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x24, 0x2e, 0x2f, 0x76,
	0x32, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x64, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_protobuf_cluster_cluster_server_proto_rawDescOnce sync.Once
	file_v2_protobuf_cluster_cluster_server_proto_rawDescData = file_v2_protobuf_cluster_cluster_server_proto_rawDesc
)

func file_v2_protobuf_cluster_cluster_server_proto_rawDescGZIP() []byte {
	file_v2_protobuf_cluster_cluster_server_proto_rawDescOnce.Do(func() {
		file_v2_protobuf_cluster_cluster_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_protobuf_cluster_cluster_server_proto_rawDescData)
	})
	return file_v2_protobuf_cluster_cluster_server_proto_rawDescData
}

var file_v2_protobuf_cluster_cluster_server_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v2_protobuf_cluster_cluster_server_proto_goTypes = []interface{}{
	(*ServerManager)(nil),         // 0: typedb.protocol.ServerManager
	(*Server)(nil),                // 1: typedb.protocol.Server
	(*ServerManager_All)(nil),     // 2: typedb.protocol.ServerManager.All
	(*ServerManager_All_Req)(nil), // 3: typedb.protocol.ServerManager.All.Req
	(*ServerManager_All_Res)(nil), // 4: typedb.protocol.ServerManager.All.Res
}
var file_v2_protobuf_cluster_cluster_server_proto_depIdxs = []int32{
	1, // 0: typedb.protocol.ServerManager.All.Res.servers:type_name -> typedb.protocol.Server
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_v2_protobuf_cluster_cluster_server_proto_init() }
func file_v2_protobuf_cluster_cluster_server_proto_init() {
	if File_v2_protobuf_cluster_cluster_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_protobuf_cluster_cluster_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerManager); i {
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
		file_v2_protobuf_cluster_cluster_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_v2_protobuf_cluster_cluster_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerManager_All); i {
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
		file_v2_protobuf_cluster_cluster_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerManager_All_Req); i {
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
		file_v2_protobuf_cluster_cluster_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerManager_All_Res); i {
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
			RawDescriptor: file_v2_protobuf_cluster_cluster_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_protobuf_cluster_cluster_server_proto_goTypes,
		DependencyIndexes: file_v2_protobuf_cluster_cluster_server_proto_depIdxs,
		MessageInfos:      file_v2_protobuf_cluster_cluster_server_proto_msgTypes,
	}.Build()
	File_v2_protobuf_cluster_cluster_server_proto = out.File
	file_v2_protobuf_cluster_cluster_server_proto_rawDesc = nil
	file_v2_protobuf_cluster_cluster_server_proto_goTypes = nil
	file_v2_protobuf_cluster_cluster_server_proto_depIdxs = nil
}
