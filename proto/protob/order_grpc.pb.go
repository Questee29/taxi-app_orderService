// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/order.proto

package protob

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

// OrderGrpcServiceClient is the client API for OrderGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderGrpcServiceClient interface {
	OrderTaxi(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	FindDriver(ctx context.Context, in *FindDriverRequest, opts ...grpc.CallOption) (OrderGrpcService_FindDriverClient, error)
	GetAllOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error)
}

type orderGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderGrpcServiceClient(cc grpc.ClientConnInterface) OrderGrpcServiceClient {
	return &orderGrpcServiceClient{cc}
}

func (c *orderGrpcServiceClient) OrderTaxi(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/protob.OrderGrpcService/OrderTaxi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderGrpcServiceClient) FindDriver(ctx context.Context, in *FindDriverRequest, opts ...grpc.CallOption) (OrderGrpcService_FindDriverClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderGrpcService_ServiceDesc.Streams[0], "/protob.OrderGrpcService/FindDriver", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderGrpcServiceFindDriverClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderGrpcService_FindDriverClient interface {
	Recv() (*FindDriverResponse, error)
	grpc.ClientStream
}

type orderGrpcServiceFindDriverClient struct {
	grpc.ClientStream
}

func (x *orderGrpcServiceFindDriverClient) Recv() (*FindDriverResponse, error) {
	m := new(FindDriverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *orderGrpcServiceClient) GetAllOrders(ctx context.Context, in *GetOrdersRequest, opts ...grpc.CallOption) (*GetOrdersResponse, error) {
	out := new(GetOrdersResponse)
	err := c.cc.Invoke(ctx, "/protob.OrderGrpcService/GetAllOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderGrpcServiceServer is the server API for OrderGrpcService service.
// All implementations must embed UnimplementedOrderGrpcServiceServer
// for forward compatibility
type OrderGrpcServiceServer interface {
	OrderTaxi(context.Context, *OrderRequest) (*OrderResponse, error)
	FindDriver(*FindDriverRequest, OrderGrpcService_FindDriverServer) error
	GetAllOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error)
	mustEmbedUnimplementedOrderGrpcServiceServer()
}

// UnimplementedOrderGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderGrpcServiceServer struct {
}

func (UnimplementedOrderGrpcServiceServer) OrderTaxi(context.Context, *OrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderTaxi not implemented")
}
func (UnimplementedOrderGrpcServiceServer) FindDriver(*FindDriverRequest, OrderGrpcService_FindDriverServer) error {
	return status.Errorf(codes.Unimplemented, "method FindDriver not implemented")
}
func (UnimplementedOrderGrpcServiceServer) GetAllOrders(context.Context, *GetOrdersRequest) (*GetOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllOrders not implemented")
}
func (UnimplementedOrderGrpcServiceServer) mustEmbedUnimplementedOrderGrpcServiceServer() {}

// UnsafeOrderGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderGrpcServiceServer will
// result in compilation errors.
type UnsafeOrderGrpcServiceServer interface {
	mustEmbedUnimplementedOrderGrpcServiceServer()
}

func RegisterOrderGrpcServiceServer(s grpc.ServiceRegistrar, srv OrderGrpcServiceServer) {
	s.RegisterService(&OrderGrpcService_ServiceDesc, srv)
}

func _OrderGrpcService_OrderTaxi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderGrpcServiceServer).OrderTaxi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protob.OrderGrpcService/OrderTaxi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderGrpcServiceServer).OrderTaxi(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderGrpcService_FindDriver_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindDriverRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderGrpcServiceServer).FindDriver(m, &orderGrpcServiceFindDriverServer{stream})
}

type OrderGrpcService_FindDriverServer interface {
	Send(*FindDriverResponse) error
	grpc.ServerStream
}

type orderGrpcServiceFindDriverServer struct {
	grpc.ServerStream
}

func (x *orderGrpcServiceFindDriverServer) Send(m *FindDriverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _OrderGrpcService_GetAllOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderGrpcServiceServer).GetAllOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protob.OrderGrpcService/GetAllOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderGrpcServiceServer).GetAllOrders(ctx, req.(*GetOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderGrpcService_ServiceDesc is the grpc.ServiceDesc for OrderGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protob.OrderGrpcService",
	HandlerType: (*OrderGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderTaxi",
			Handler:    _OrderGrpcService_OrderTaxi_Handler,
		},
		{
			MethodName: "GetAllOrders",
			Handler:    _OrderGrpcService_GetAllOrders_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindDriver",
			Handler:       _OrderGrpcService_FindDriver_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/order.proto",
}

// RatingGrpcServiceClient is the client API for RatingGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingGrpcServiceClient interface {
	RateLastOrder(ctx context.Context, in *RateOrderRequest, opts ...grpc.CallOption) (*RateOrderResponse, error)
}

type ratingGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingGrpcServiceClient(cc grpc.ClientConnInterface) RatingGrpcServiceClient {
	return &ratingGrpcServiceClient{cc}
}

func (c *ratingGrpcServiceClient) RateLastOrder(ctx context.Context, in *RateOrderRequest, opts ...grpc.CallOption) (*RateOrderResponse, error) {
	out := new(RateOrderResponse)
	err := c.cc.Invoke(ctx, "/protob.RatingGrpcService/RateLastOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingGrpcServiceServer is the server API for RatingGrpcService service.
// All implementations must embed UnimplementedRatingGrpcServiceServer
// for forward compatibility
type RatingGrpcServiceServer interface {
	RateLastOrder(context.Context, *RateOrderRequest) (*RateOrderResponse, error)
	mustEmbedUnimplementedRatingGrpcServiceServer()
}

// UnimplementedRatingGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRatingGrpcServiceServer struct {
}

func (UnimplementedRatingGrpcServiceServer) RateLastOrder(context.Context, *RateOrderRequest) (*RateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RateLastOrder not implemented")
}
func (UnimplementedRatingGrpcServiceServer) mustEmbedUnimplementedRatingGrpcServiceServer() {}

// UnsafeRatingGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingGrpcServiceServer will
// result in compilation errors.
type UnsafeRatingGrpcServiceServer interface {
	mustEmbedUnimplementedRatingGrpcServiceServer()
}

func RegisterRatingGrpcServiceServer(s grpc.ServiceRegistrar, srv RatingGrpcServiceServer) {
	s.RegisterService(&RatingGrpcService_ServiceDesc, srv)
}

func _RatingGrpcService_RateLastOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingGrpcServiceServer).RateLastOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protob.RatingGrpcService/RateLastOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingGrpcServiceServer).RateLastOrder(ctx, req.(*RateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingGrpcService_ServiceDesc is the grpc.ServiceDesc for RatingGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protob.RatingGrpcService",
	HandlerType: (*RatingGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RateLastOrder",
			Handler:    _RatingGrpcService_RateLastOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/order.proto",
}
