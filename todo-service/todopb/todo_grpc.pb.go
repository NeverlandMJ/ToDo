// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: todopb/todo.proto

package todopb

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

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreateTodo(ctx context.Context, in *RequestTodo, opts ...grpc.CallOption) (*ResponseTodo, error)
	GetTodoByID(ctx context.Context, in *RequestGetTodo, opts ...grpc.CallOption) (*ResponseTodo, error)
	MarkAsDone(ctx context.Context, in *RequestMarkAsDone, opts ...grpc.CallOption) (*Empty, error)
	DeleteTodoByID(ctx context.Context, in *RequestDeleteTodo, opts ...grpc.CallOption) (*Empty, error)
	GetAllTodos(ctx context.Context, in *RequestUserID, opts ...grpc.CallOption) (*ResponseAllTodos, error)
	UpdateTodosBody(ctx context.Context, in *RequestUpdateTodosBody, opts ...grpc.CallOption) (*Empty, error)
	UpdateTodosDeadline(ctx context.Context, in *RequestUpdateTodosDeadline, opts ...grpc.CallOption) (*Empty, error)
	DeleteDoneTodos(ctx context.Context, in *RequestUserID, opts ...grpc.CallOption) (*Empty, error)
	DeletePassedDeadline(ctx context.Context, in *RequestUserID, opts ...grpc.CallOption) (*Empty, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *RequestTodo, opts ...grpc.CallOption) (*ResponseTodo, error) {
	out := new(ResponseTodo)
	err := c.cc.Invoke(ctx, "/todopb.TodoService/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetTodoByID(ctx context.Context, in *RequestGetTodo, opts ...grpc.CallOption) (*ResponseTodo, error) {
	out := new(ResponseTodo)
	err := c.cc.Invoke(ctx, "/todopb.TodoService/GetTodoByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) MarkAsDone(ctx context.Context, in *RequestMarkAsDone, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/todopb.TodoService/MarkAsDone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteTodoByID(ctx context.Context, in *RequestDeleteTodo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/todopb.TodoService/DeleteTodoByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetAllTodos(ctx context.Context, in *RequestUserID, opts ...grpc.CallOption) (*ResponseAllTodos, error) {
	out := new(ResponseAllTodos)
	err := c.cc.Invoke(ctx, "/todopb.TodoService/GetAllTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateTodosBody(ctx context.Context, in *RequestUpdateTodosBody, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/todopb.TodoService/UpdateTodosBody", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) UpdateTodosDeadline(ctx context.Context, in *RequestUpdateTodosDeadline, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/todopb.TodoService/UpdateTodosDeadline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteDoneTodos(ctx context.Context, in *RequestUserID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/todopb.TodoService/DeleteDoneTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeletePassedDeadline(ctx context.Context, in *RequestUserID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/todopb.TodoService/DeletePassedDeadline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	CreateTodo(context.Context, *RequestTodo) (*ResponseTodo, error)
	GetTodoByID(context.Context, *RequestGetTodo) (*ResponseTodo, error)
	MarkAsDone(context.Context, *RequestMarkAsDone) (*Empty, error)
	DeleteTodoByID(context.Context, *RequestDeleteTodo) (*Empty, error)
	GetAllTodos(context.Context, *RequestUserID) (*ResponseAllTodos, error)
	UpdateTodosBody(context.Context, *RequestUpdateTodosBody) (*Empty, error)
	UpdateTodosDeadline(context.Context, *RequestUpdateTodosDeadline) (*Empty, error)
	DeleteDoneTodos(context.Context, *RequestUserID) (*Empty, error)
	DeletePassedDeadline(context.Context, *RequestUserID) (*Empty, error)
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) CreateTodo(context.Context, *RequestTodo) (*ResponseTodo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedTodoServiceServer) GetTodoByID(context.Context, *RequestGetTodo) (*ResponseTodo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTodoByID not implemented")
}
func (UnimplementedTodoServiceServer) MarkAsDone(context.Context, *RequestMarkAsDone) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkAsDone not implemented")
}
func (UnimplementedTodoServiceServer) DeleteTodoByID(context.Context, *RequestDeleteTodo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTodoByID not implemented")
}
func (UnimplementedTodoServiceServer) GetAllTodos(context.Context, *RequestUserID) (*ResponseAllTodos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTodos not implemented")
}
func (UnimplementedTodoServiceServer) UpdateTodosBody(context.Context, *RequestUpdateTodosBody) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodosBody not implemented")
}
func (UnimplementedTodoServiceServer) UpdateTodosDeadline(context.Context, *RequestUpdateTodosDeadline) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTodosDeadline not implemented")
}
func (UnimplementedTodoServiceServer) DeleteDoneTodos(context.Context, *RequestUserID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDoneTodos not implemented")
}
func (UnimplementedTodoServiceServer) DeletePassedDeadline(context.Context, *RequestUserID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePassedDeadline not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestTodo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todopb.TodoService/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*RequestTodo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetTodoByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetTodo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetTodoByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todopb.TodoService/GetTodoByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetTodoByID(ctx, req.(*RequestGetTodo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_MarkAsDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMarkAsDone)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).MarkAsDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todopb.TodoService/MarkAsDone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).MarkAsDone(ctx, req.(*RequestMarkAsDone))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteTodoByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeleteTodo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteTodoByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todopb.TodoService/DeleteTodoByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteTodoByID(ctx, req.(*RequestDeleteTodo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetAllTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetAllTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todopb.TodoService/GetAllTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetAllTodos(ctx, req.(*RequestUserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateTodosBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUpdateTodosBody)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateTodosBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todopb.TodoService/UpdateTodosBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateTodosBody(ctx, req.(*RequestUpdateTodosBody))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_UpdateTodosDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUpdateTodosDeadline)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).UpdateTodosDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todopb.TodoService/UpdateTodosDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).UpdateTodosDeadline(ctx, req.(*RequestUpdateTodosDeadline))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteDoneTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteDoneTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todopb.TodoService/DeleteDoneTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteDoneTodos(ctx, req.(*RequestUserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeletePassedDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeletePassedDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todopb.TodoService/DeletePassedDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeletePassedDeadline(ctx, req.(*RequestUserID))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todopb.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
		{
			MethodName: "GetTodoByID",
			Handler:    _TodoService_GetTodoByID_Handler,
		},
		{
			MethodName: "MarkAsDone",
			Handler:    _TodoService_MarkAsDone_Handler,
		},
		{
			MethodName: "DeleteTodoByID",
			Handler:    _TodoService_DeleteTodoByID_Handler,
		},
		{
			MethodName: "GetAllTodos",
			Handler:    _TodoService_GetAllTodos_Handler,
		},
		{
			MethodName: "UpdateTodosBody",
			Handler:    _TodoService_UpdateTodosBody_Handler,
		},
		{
			MethodName: "UpdateTodosDeadline",
			Handler:    _TodoService_UpdateTodosDeadline_Handler,
		},
		{
			MethodName: "DeleteDoneTodos",
			Handler:    _TodoService_DeleteDoneTodos_Handler,
		},
		{
			MethodName: "DeletePassedDeadline",
			Handler:    _TodoService_DeletePassedDeadline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todopb/todo.proto",
}
