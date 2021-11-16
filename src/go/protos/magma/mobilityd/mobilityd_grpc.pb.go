// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mobilityd

import (
	context "context"
	orc8r "github.com/magma/magma/src/go/protos/magma/orc8r"
	subscriberdb "github.com/magma/magma/src/go/protos/magma/subscriberdb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MobilityServiceClient is the client API for MobilityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MobilityServiceClient interface {
	// Add a range of IP addresses into the free IP pool
	// Throws INVALID_ARGUMENT if IPBlock is invalid
	// Throws FAILED_PRECONDITION if IPBlock overlaps with existing ones
	//
	AddIPBlock(ctx context.Context, in *IPBlock, opts ...grpc.CallOption) (*orc8r.Void, error)
	// Return a list of assigned IP blocks
	//
	ListAddedIPv4Blocks(ctx context.Context, in *orc8r.Void, opts ...grpc.CallOption) (*ListAddedIPBlocksResponse, error)
	// Return a list of assigned IPv6 blocks
	//
	ListAddedIPv6Blocks(ctx context.Context, in *orc8r.Void, opts ...grpc.CallOption) (*ListAddedIPBlocksResponse, error)
	// Return a list of allocated IPs inside a IP block
	// Throws INVALID_ARGUMENT if IPBlock is invalid
	// Throws FAILED_PRECONDITION if IPBlock is not previously assigned
	//
	ListAllocatedIPs(ctx context.Context, in *IPBlock, opts ...grpc.CallOption) (*ListAllocatedIPsResponse, error)
	// Allocate an IP address from the free IP pool
	// Throws RESOURCE_EXHAUSTED if no free IP available
	// Throws ALREADY_EXISTS if an IP has been allocated for the subscriber
	//
	AllocateIPAddress(ctx context.Context, in *AllocateIPRequest, opts ...grpc.CallOption) (*AllocateIPAddressResponse, error)
	// Release and recycle an allocated IP address
	// Throws NOT_FOUND if the requested (SID, IP) pair is not found
	//
	ReleaseIPAddress(ctx context.Context, in *ReleaseIPRequest, opts ...grpc.CallOption) (*orc8r.Void, error)
	// Gets subscriber's IP address. Throws NOT_FOUND if it doesn't exist
	GetIPForSubscriber(ctx context.Context, in *IPLookupRequest, opts ...grpc.CallOption) (*IPAddress, error)
	// Gets subscriber's ID from an IP Address.
	// Throws NOT_FOUND if it doesn't exist
	GetSubscriberIDFromIP(ctx context.Context, in *IPAddress, opts ...grpc.CallOption) (*subscriberdb.SubscriberID, error)
	// Get the full subscriber table
	GetSubscriberIPTable(ctx context.Context, in *orc8r.Void, opts ...grpc.CallOption) (*SubscriberIPTable, error)
	// Remove allocated IP blocks
	// Default behavior is to only remove all IP blocks that have no IP addresses
	// allocated from them. If force is set, then will remove all IP blocks,
	// regardless of whether any IPs have been allocated.
	RemoveIPBlock(ctx context.Context, in *RemoveIPBlockRequest, opts ...grpc.CallOption) (*RemoveIPBlockResponse, error)
	// Get Internet Gateway and  mac addresss
	ListGatewayInfo(ctx context.Context, in *orc8r.Void, opts ...grpc.CallOption) (*ListGWInfoResponse, error)
	// Set ip and mac address of def Internet Gateway
	SetGatewayInfo(ctx context.Context, in *GWInfo, opts ...grpc.CallOption) (*orc8r.Void, error)
}

type mobilityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMobilityServiceClient(cc grpc.ClientConnInterface) MobilityServiceClient {
	return &mobilityServiceClient{cc}
}

func (c *mobilityServiceClient) AddIPBlock(ctx context.Context, in *IPBlock, opts ...grpc.CallOption) (*orc8r.Void, error) {
	out := new(orc8r.Void)
	err := c.cc.Invoke(ctx, "/magma.lte.MobilityService/AddIPBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobilityServiceClient) ListAddedIPv4Blocks(ctx context.Context, in *orc8r.Void, opts ...grpc.CallOption) (*ListAddedIPBlocksResponse, error) {
	out := new(ListAddedIPBlocksResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.MobilityService/ListAddedIPv4Blocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobilityServiceClient) ListAddedIPv6Blocks(ctx context.Context, in *orc8r.Void, opts ...grpc.CallOption) (*ListAddedIPBlocksResponse, error) {
	out := new(ListAddedIPBlocksResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.MobilityService/ListAddedIPv6Blocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobilityServiceClient) ListAllocatedIPs(ctx context.Context, in *IPBlock, opts ...grpc.CallOption) (*ListAllocatedIPsResponse, error) {
	out := new(ListAllocatedIPsResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.MobilityService/ListAllocatedIPs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobilityServiceClient) AllocateIPAddress(ctx context.Context, in *AllocateIPRequest, opts ...grpc.CallOption) (*AllocateIPAddressResponse, error) {
	out := new(AllocateIPAddressResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.MobilityService/AllocateIPAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobilityServiceClient) ReleaseIPAddress(ctx context.Context, in *ReleaseIPRequest, opts ...grpc.CallOption) (*orc8r.Void, error) {
	out := new(orc8r.Void)
	err := c.cc.Invoke(ctx, "/magma.lte.MobilityService/ReleaseIPAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobilityServiceClient) GetIPForSubscriber(ctx context.Context, in *IPLookupRequest, opts ...grpc.CallOption) (*IPAddress, error) {
	out := new(IPAddress)
	err := c.cc.Invoke(ctx, "/magma.lte.MobilityService/GetIPForSubscriber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobilityServiceClient) GetSubscriberIDFromIP(ctx context.Context, in *IPAddress, opts ...grpc.CallOption) (*subscriberdb.SubscriberID, error) {
	out := new(subscriberdb.SubscriberID)
	err := c.cc.Invoke(ctx, "/magma.lte.MobilityService/GetSubscriberIDFromIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobilityServiceClient) GetSubscriberIPTable(ctx context.Context, in *orc8r.Void, opts ...grpc.CallOption) (*SubscriberIPTable, error) {
	out := new(SubscriberIPTable)
	err := c.cc.Invoke(ctx, "/magma.lte.MobilityService/GetSubscriberIPTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobilityServiceClient) RemoveIPBlock(ctx context.Context, in *RemoveIPBlockRequest, opts ...grpc.CallOption) (*RemoveIPBlockResponse, error) {
	out := new(RemoveIPBlockResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.MobilityService/RemoveIPBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobilityServiceClient) ListGatewayInfo(ctx context.Context, in *orc8r.Void, opts ...grpc.CallOption) (*ListGWInfoResponse, error) {
	out := new(ListGWInfoResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.MobilityService/ListGatewayInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobilityServiceClient) SetGatewayInfo(ctx context.Context, in *GWInfo, opts ...grpc.CallOption) (*orc8r.Void, error) {
	out := new(orc8r.Void)
	err := c.cc.Invoke(ctx, "/magma.lte.MobilityService/SetGatewayInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MobilityServiceServer is the server API for MobilityService service.
// All implementations must embed UnimplementedMobilityServiceServer
// for forward compatibility
type MobilityServiceServer interface {
	// Add a range of IP addresses into the free IP pool
	// Throws INVALID_ARGUMENT if IPBlock is invalid
	// Throws FAILED_PRECONDITION if IPBlock overlaps with existing ones
	//
	AddIPBlock(context.Context, *IPBlock) (*orc8r.Void, error)
	// Return a list of assigned IP blocks
	//
	ListAddedIPv4Blocks(context.Context, *orc8r.Void) (*ListAddedIPBlocksResponse, error)
	// Return a list of assigned IPv6 blocks
	//
	ListAddedIPv6Blocks(context.Context, *orc8r.Void) (*ListAddedIPBlocksResponse, error)
	// Return a list of allocated IPs inside a IP block
	// Throws INVALID_ARGUMENT if IPBlock is invalid
	// Throws FAILED_PRECONDITION if IPBlock is not previously assigned
	//
	ListAllocatedIPs(context.Context, *IPBlock) (*ListAllocatedIPsResponse, error)
	// Allocate an IP address from the free IP pool
	// Throws RESOURCE_EXHAUSTED if no free IP available
	// Throws ALREADY_EXISTS if an IP has been allocated for the subscriber
	//
	AllocateIPAddress(context.Context, *AllocateIPRequest) (*AllocateIPAddressResponse, error)
	// Release and recycle an allocated IP address
	// Throws NOT_FOUND if the requested (SID, IP) pair is not found
	//
	ReleaseIPAddress(context.Context, *ReleaseIPRequest) (*orc8r.Void, error)
	// Gets subscriber's IP address. Throws NOT_FOUND if it doesn't exist
	GetIPForSubscriber(context.Context, *IPLookupRequest) (*IPAddress, error)
	// Gets subscriber's ID from an IP Address.
	// Throws NOT_FOUND if it doesn't exist
	GetSubscriberIDFromIP(context.Context, *IPAddress) (*subscriberdb.SubscriberID, error)
	// Get the full subscriber table
	GetSubscriberIPTable(context.Context, *orc8r.Void) (*SubscriberIPTable, error)
	// Remove allocated IP blocks
	// Default behavior is to only remove all IP blocks that have no IP addresses
	// allocated from them. If force is set, then will remove all IP blocks,
	// regardless of whether any IPs have been allocated.
	RemoveIPBlock(context.Context, *RemoveIPBlockRequest) (*RemoveIPBlockResponse, error)
	// Get Internet Gateway and  mac addresss
	ListGatewayInfo(context.Context, *orc8r.Void) (*ListGWInfoResponse, error)
	// Set ip and mac address of def Internet Gateway
	SetGatewayInfo(context.Context, *GWInfo) (*orc8r.Void, error)
	mustEmbedUnimplementedMobilityServiceServer()
}

// UnimplementedMobilityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMobilityServiceServer struct {
}

func (UnimplementedMobilityServiceServer) AddIPBlock(context.Context, *IPBlock) (*orc8r.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddIPBlock not implemented")
}
func (UnimplementedMobilityServiceServer) ListAddedIPv4Blocks(context.Context, *orc8r.Void) (*ListAddedIPBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAddedIPv4Blocks not implemented")
}
func (UnimplementedMobilityServiceServer) ListAddedIPv6Blocks(context.Context, *orc8r.Void) (*ListAddedIPBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAddedIPv6Blocks not implemented")
}
func (UnimplementedMobilityServiceServer) ListAllocatedIPs(context.Context, *IPBlock) (*ListAllocatedIPsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllocatedIPs not implemented")
}
func (UnimplementedMobilityServiceServer) AllocateIPAddress(context.Context, *AllocateIPRequest) (*AllocateIPAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocateIPAddress not implemented")
}
func (UnimplementedMobilityServiceServer) ReleaseIPAddress(context.Context, *ReleaseIPRequest) (*orc8r.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseIPAddress not implemented")
}
func (UnimplementedMobilityServiceServer) GetIPForSubscriber(context.Context, *IPLookupRequest) (*IPAddress, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIPForSubscriber not implemented")
}
func (UnimplementedMobilityServiceServer) GetSubscriberIDFromIP(context.Context, *IPAddress) (*subscriberdb.SubscriberID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriberIDFromIP not implemented")
}
func (UnimplementedMobilityServiceServer) GetSubscriberIPTable(context.Context, *orc8r.Void) (*SubscriberIPTable, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriberIPTable not implemented")
}
func (UnimplementedMobilityServiceServer) RemoveIPBlock(context.Context, *RemoveIPBlockRequest) (*RemoveIPBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveIPBlock not implemented")
}
func (UnimplementedMobilityServiceServer) ListGatewayInfo(context.Context, *orc8r.Void) (*ListGWInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGatewayInfo not implemented")
}
func (UnimplementedMobilityServiceServer) SetGatewayInfo(context.Context, *GWInfo) (*orc8r.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGatewayInfo not implemented")
}
func (UnimplementedMobilityServiceServer) mustEmbedUnimplementedMobilityServiceServer() {}

// UnsafeMobilityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MobilityServiceServer will
// result in compilation errors.
type UnsafeMobilityServiceServer interface {
	mustEmbedUnimplementedMobilityServiceServer()
}

func RegisterMobilityServiceServer(s grpc.ServiceRegistrar, srv MobilityServiceServer) {
	s.RegisterService(&MobilityService_ServiceDesc, srv)
}

func _MobilityService_AddIPBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobilityServiceServer).AddIPBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.MobilityService/AddIPBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobilityServiceServer).AddIPBlock(ctx, req.(*IPBlock))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobilityService_ListAddedIPv4Blocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(orc8r.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobilityServiceServer).ListAddedIPv4Blocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.MobilityService/ListAddedIPv4Blocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobilityServiceServer).ListAddedIPv4Blocks(ctx, req.(*orc8r.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobilityService_ListAddedIPv6Blocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(orc8r.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobilityServiceServer).ListAddedIPv6Blocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.MobilityService/ListAddedIPv6Blocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobilityServiceServer).ListAddedIPv6Blocks(ctx, req.(*orc8r.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobilityService_ListAllocatedIPs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPBlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobilityServiceServer).ListAllocatedIPs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.MobilityService/ListAllocatedIPs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobilityServiceServer).ListAllocatedIPs(ctx, req.(*IPBlock))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobilityService_AllocateIPAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocateIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobilityServiceServer).AllocateIPAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.MobilityService/AllocateIPAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobilityServiceServer).AllocateIPAddress(ctx, req.(*AllocateIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobilityService_ReleaseIPAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobilityServiceServer).ReleaseIPAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.MobilityService/ReleaseIPAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobilityServiceServer).ReleaseIPAddress(ctx, req.(*ReleaseIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobilityService_GetIPForSubscriber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobilityServiceServer).GetIPForSubscriber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.MobilityService/GetIPForSubscriber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobilityServiceServer).GetIPForSubscriber(ctx, req.(*IPLookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobilityService_GetSubscriberIDFromIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobilityServiceServer).GetSubscriberIDFromIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.MobilityService/GetSubscriberIDFromIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobilityServiceServer).GetSubscriberIDFromIP(ctx, req.(*IPAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobilityService_GetSubscriberIPTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(orc8r.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobilityServiceServer).GetSubscriberIPTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.MobilityService/GetSubscriberIPTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobilityServiceServer).GetSubscriberIPTable(ctx, req.(*orc8r.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobilityService_RemoveIPBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveIPBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobilityServiceServer).RemoveIPBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.MobilityService/RemoveIPBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobilityServiceServer).RemoveIPBlock(ctx, req.(*RemoveIPBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobilityService_ListGatewayInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(orc8r.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobilityServiceServer).ListGatewayInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.MobilityService/ListGatewayInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobilityServiceServer).ListGatewayInfo(ctx, req.(*orc8r.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _MobilityService_SetGatewayInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GWInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobilityServiceServer).SetGatewayInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.MobilityService/SetGatewayInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobilityServiceServer).SetGatewayInfo(ctx, req.(*GWInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// MobilityService_ServiceDesc is the grpc.ServiceDesc for MobilityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MobilityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "magma.lte.MobilityService",
	HandlerType: (*MobilityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddIPBlock",
			Handler:    _MobilityService_AddIPBlock_Handler,
		},
		{
			MethodName: "ListAddedIPv4Blocks",
			Handler:    _MobilityService_ListAddedIPv4Blocks_Handler,
		},
		{
			MethodName: "ListAddedIPv6Blocks",
			Handler:    _MobilityService_ListAddedIPv6Blocks_Handler,
		},
		{
			MethodName: "ListAllocatedIPs",
			Handler:    _MobilityService_ListAllocatedIPs_Handler,
		},
		{
			MethodName: "AllocateIPAddress",
			Handler:    _MobilityService_AllocateIPAddress_Handler,
		},
		{
			MethodName: "ReleaseIPAddress",
			Handler:    _MobilityService_ReleaseIPAddress_Handler,
		},
		{
			MethodName: "GetIPForSubscriber",
			Handler:    _MobilityService_GetIPForSubscriber_Handler,
		},
		{
			MethodName: "GetSubscriberIDFromIP",
			Handler:    _MobilityService_GetSubscriberIDFromIP_Handler,
		},
		{
			MethodName: "GetSubscriberIPTable",
			Handler:    _MobilityService_GetSubscriberIPTable_Handler,
		},
		{
			MethodName: "RemoveIPBlock",
			Handler:    _MobilityService_RemoveIPBlock_Handler,
		},
		{
			MethodName: "ListGatewayInfo",
			Handler:    _MobilityService_ListGatewayInfo_Handler,
		},
		{
			MethodName: "SetGatewayInfo",
			Handler:    _MobilityService_SetGatewayInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mobilityd.proto",
}
