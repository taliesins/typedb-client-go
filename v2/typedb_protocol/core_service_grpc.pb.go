// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: v2/protobuf/core/core_service.proto

package typedb_protocol

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TypeDBClient is the client API for TypeDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TypeDBClient interface {
	// Database Manager API
	DatabasesContains(ctx context.Context, in *CoreDatabaseManager_Contains_Req, opts ...grpc.CallOption) (*CoreDatabaseManager_Contains_Res, error)
	DatabasesCreate(ctx context.Context, in *CoreDatabaseManager_Create_Req, opts ...grpc.CallOption) (*CoreDatabaseManager_Create_Res, error)
	DatabasesAll(ctx context.Context, in *CoreDatabaseManager_All_Req, opts ...grpc.CallOption) (*CoreDatabaseManager_All_Res, error)
	// Database API
	DatabaseSchema(ctx context.Context, in *CoreDatabase_Schema_Req, opts ...grpc.CallOption) (*CoreDatabase_Schema_Res, error)
	DatabaseDelete(ctx context.Context, in *CoreDatabase_Delete_Req, opts ...grpc.CallOption) (*CoreDatabase_Delete_Res, error)
	// Session API
	SessionOpen(ctx context.Context, in *Session_Open_Req, opts ...grpc.CallOption) (*Session_Open_Res, error)
	SessionClose(ctx context.Context, in *Session_Close_Req, opts ...grpc.CallOption) (*Session_Close_Res, error)
	// Checks with the server that the session is still alive, and informs it that it should be kept alive.
	SessionPulse(ctx context.Context, in *Session_Pulse_Req, opts ...grpc.CallOption) (*Session_Pulse_Res, error)
	// Transaction Streaming API
	// Opens a bi-directional stream representing a stateful transaction, streaming
	// requests and responses back-and-forth. The first transaction client message must
	// be {Transaction.Open.Req}. Closing the stream closes the transaction.
	Transaction(ctx context.Context, opts ...grpc.CallOption) (TypeDB_TransactionClient, error)
}

type typeDBClient struct {
	cc grpc.ClientConnInterface
}

func NewTypeDBClient(cc grpc.ClientConnInterface) TypeDBClient {
	return &typeDBClient{cc}
}

func (c *typeDBClient) DatabasesContains(ctx context.Context, in *CoreDatabaseManager_Contains_Req, opts ...grpc.CallOption) (*CoreDatabaseManager_Contains_Res, error) {
	out := new(CoreDatabaseManager_Contains_Res)
	err := c.cc.Invoke(ctx, "/typedb.protocol.TypeDB/databases_contains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typeDBClient) DatabasesCreate(ctx context.Context, in *CoreDatabaseManager_Create_Req, opts ...grpc.CallOption) (*CoreDatabaseManager_Create_Res, error) {
	out := new(CoreDatabaseManager_Create_Res)
	err := c.cc.Invoke(ctx, "/typedb.protocol.TypeDB/databases_create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typeDBClient) DatabasesAll(ctx context.Context, in *CoreDatabaseManager_All_Req, opts ...grpc.CallOption) (*CoreDatabaseManager_All_Res, error) {
	out := new(CoreDatabaseManager_All_Res)
	err := c.cc.Invoke(ctx, "/typedb.protocol.TypeDB/databases_all", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typeDBClient) DatabaseSchema(ctx context.Context, in *CoreDatabase_Schema_Req, opts ...grpc.CallOption) (*CoreDatabase_Schema_Res, error) {
	out := new(CoreDatabase_Schema_Res)
	err := c.cc.Invoke(ctx, "/typedb.protocol.TypeDB/database_schema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typeDBClient) DatabaseDelete(ctx context.Context, in *CoreDatabase_Delete_Req, opts ...grpc.CallOption) (*CoreDatabase_Delete_Res, error) {
	out := new(CoreDatabase_Delete_Res)
	err := c.cc.Invoke(ctx, "/typedb.protocol.TypeDB/database_delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typeDBClient) SessionOpen(ctx context.Context, in *Session_Open_Req, opts ...grpc.CallOption) (*Session_Open_Res, error) {
	out := new(Session_Open_Res)
	err := c.cc.Invoke(ctx, "/typedb.protocol.TypeDB/session_open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typeDBClient) SessionClose(ctx context.Context, in *Session_Close_Req, opts ...grpc.CallOption) (*Session_Close_Res, error) {
	out := new(Session_Close_Res)
	err := c.cc.Invoke(ctx, "/typedb.protocol.TypeDB/session_close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typeDBClient) SessionPulse(ctx context.Context, in *Session_Pulse_Req, opts ...grpc.CallOption) (*Session_Pulse_Res, error) {
	out := new(Session_Pulse_Res)
	err := c.cc.Invoke(ctx, "/typedb.protocol.TypeDB/session_pulse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *typeDBClient) Transaction(ctx context.Context, opts ...grpc.CallOption) (TypeDB_TransactionClient, error) {
	stream, err := c.cc.NewStream(ctx, &TypeDB_ServiceDesc.Streams[0], "/typedb.protocol.TypeDB/transaction", opts...)
	if err != nil {
		return nil, err
	}
	x := &typeDBTransactionClient{stream}
	return x, nil
}

type TypeDB_TransactionClient interface {
	Send(*Transaction_Client) error
	Recv() (*Transaction_Server, error)
	grpc.ClientStream
}

type typeDBTransactionClient struct {
	grpc.ClientStream
}

func (x *typeDBTransactionClient) Send(m *Transaction_Client) error {
	return x.ClientStream.SendMsg(m)
}

func (x *typeDBTransactionClient) Recv() (*Transaction_Server, error) {
	m := new(Transaction_Server)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TypeDBServer is the server API for TypeDB service.
// All implementations must embed UnimplementedTypeDBServer
// for forward compatibility
type TypeDBServer interface {
	// Database Manager API
	DatabasesContains(context.Context, *CoreDatabaseManager_Contains_Req) (*CoreDatabaseManager_Contains_Res, error)
	DatabasesCreate(context.Context, *CoreDatabaseManager_Create_Req) (*CoreDatabaseManager_Create_Res, error)
	DatabasesAll(context.Context, *CoreDatabaseManager_All_Req) (*CoreDatabaseManager_All_Res, error)
	// Database API
	DatabaseSchema(context.Context, *CoreDatabase_Schema_Req) (*CoreDatabase_Schema_Res, error)
	DatabaseDelete(context.Context, *CoreDatabase_Delete_Req) (*CoreDatabase_Delete_Res, error)
	// Session API
	SessionOpen(context.Context, *Session_Open_Req) (*Session_Open_Res, error)
	SessionClose(context.Context, *Session_Close_Req) (*Session_Close_Res, error)
	// Checks with the server that the session is still alive, and informs it that it should be kept alive.
	SessionPulse(context.Context, *Session_Pulse_Req) (*Session_Pulse_Res, error)
	// Transaction Streaming API
	// Opens a bi-directional stream representing a stateful transaction, streaming
	// requests and responses back-and-forth. The first transaction client message must
	// be {Transaction.Open.Req}. Closing the stream closes the transaction.
	Transaction(TypeDB_TransactionServer) error
	mustEmbedUnimplementedTypeDBServer()
}

// UnimplementedTypeDBServer must be embedded to have forward compatible implementations.
type UnimplementedTypeDBServer struct {
}

func (UnimplementedTypeDBServer) DatabasesContains(context.Context, *CoreDatabaseManager_Contains_Req) (*CoreDatabaseManager_Contains_Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DatabasesContains not implemented")
}
func (UnimplementedTypeDBServer) DatabasesCreate(context.Context, *CoreDatabaseManager_Create_Req) (*CoreDatabaseManager_Create_Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DatabasesCreate not implemented")
}
func (UnimplementedTypeDBServer) DatabasesAll(context.Context, *CoreDatabaseManager_All_Req) (*CoreDatabaseManager_All_Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DatabasesAll not implemented")
}
func (UnimplementedTypeDBServer) DatabaseSchema(context.Context, *CoreDatabase_Schema_Req) (*CoreDatabase_Schema_Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DatabaseSchema not implemented")
}
func (UnimplementedTypeDBServer) DatabaseDelete(context.Context, *CoreDatabase_Delete_Req) (*CoreDatabase_Delete_Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DatabaseDelete not implemented")
}
func (UnimplementedTypeDBServer) SessionOpen(context.Context, *Session_Open_Req) (*Session_Open_Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionOpen not implemented")
}
func (UnimplementedTypeDBServer) SessionClose(context.Context, *Session_Close_Req) (*Session_Close_Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionClose not implemented")
}
func (UnimplementedTypeDBServer) SessionPulse(context.Context, *Session_Pulse_Req) (*Session_Pulse_Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SessionPulse not implemented")
}
func (UnimplementedTypeDBServer) Transaction(TypeDB_TransactionServer) error {
	return status.Errorf(codes.Unimplemented, "method Transaction not implemented")
}
func (UnimplementedTypeDBServer) mustEmbedUnimplementedTypeDBServer() {}

// UnsafeTypeDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TypeDBServer will
// result in compilation errors.
type UnsafeTypeDBServer interface {
	mustEmbedUnimplementedTypeDBServer()
}

func RegisterTypeDBServer(s grpc.ServiceRegistrar, srv TypeDBServer) {
	s.RegisterService(&TypeDB_ServiceDesc, srv)
}

func _TypeDB_DatabasesContains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoreDatabaseManager_Contains_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeDBServer).DatabasesContains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/typedb.protocol.TypeDB/databases_contains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeDBServer).DatabasesContains(ctx, req.(*CoreDatabaseManager_Contains_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypeDB_DatabasesCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoreDatabaseManager_Create_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeDBServer).DatabasesCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/typedb.protocol.TypeDB/databases_create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeDBServer).DatabasesCreate(ctx, req.(*CoreDatabaseManager_Create_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypeDB_DatabasesAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoreDatabaseManager_All_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeDBServer).DatabasesAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/typedb.protocol.TypeDB/databases_all",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeDBServer).DatabasesAll(ctx, req.(*CoreDatabaseManager_All_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypeDB_DatabaseSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoreDatabase_Schema_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeDBServer).DatabaseSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/typedb.protocol.TypeDB/database_schema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeDBServer).DatabaseSchema(ctx, req.(*CoreDatabase_Schema_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypeDB_DatabaseDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoreDatabase_Delete_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeDBServer).DatabaseDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/typedb.protocol.TypeDB/database_delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeDBServer).DatabaseDelete(ctx, req.(*CoreDatabase_Delete_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypeDB_SessionOpen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session_Open_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeDBServer).SessionOpen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/typedb.protocol.TypeDB/session_open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeDBServer).SessionOpen(ctx, req.(*Session_Open_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypeDB_SessionClose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session_Close_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeDBServer).SessionClose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/typedb.protocol.TypeDB/session_close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeDBServer).SessionClose(ctx, req.(*Session_Close_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypeDB_SessionPulse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Session_Pulse_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TypeDBServer).SessionPulse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/typedb.protocol.TypeDB/session_pulse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TypeDBServer).SessionPulse(ctx, req.(*Session_Pulse_Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _TypeDB_Transaction_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TypeDBServer).Transaction(&typeDBTransactionServer{stream})
}

type TypeDB_TransactionServer interface {
	Send(*Transaction_Server) error
	Recv() (*Transaction_Client, error)
	grpc.ServerStream
}

type typeDBTransactionServer struct {
	grpc.ServerStream
}

func (x *typeDBTransactionServer) Send(m *Transaction_Server) error {
	return x.ServerStream.SendMsg(m)
}

func (x *typeDBTransactionServer) Recv() (*Transaction_Client, error) {
	m := new(Transaction_Client)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TypeDB_ServiceDesc is the grpc.ServiceDesc for TypeDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TypeDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "typedb.protocol.TypeDB",
	HandlerType: (*TypeDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "databases_contains",
			Handler:    _TypeDB_DatabasesContains_Handler,
		},
		{
			MethodName: "databases_create",
			Handler:    _TypeDB_DatabasesCreate_Handler,
		},
		{
			MethodName: "databases_all",
			Handler:    _TypeDB_DatabasesAll_Handler,
		},
		{
			MethodName: "database_schema",
			Handler:    _TypeDB_DatabaseSchema_Handler,
		},
		{
			MethodName: "database_delete",
			Handler:    _TypeDB_DatabaseDelete_Handler,
		},
		{
			MethodName: "session_open",
			Handler:    _TypeDB_SessionOpen_Handler,
		},
		{
			MethodName: "session_close",
			Handler:    _TypeDB_SessionClose_Handler,
		},
		{
			MethodName: "session_pulse",
			Handler:    _TypeDB_SessionPulse_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "transaction",
			Handler:       _TypeDB_Transaction_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "v2/protobuf/core/core_service.proto",
}
