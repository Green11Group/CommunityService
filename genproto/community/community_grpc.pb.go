// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protos/community.proto

package community

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

// CommunityServiceClient is the client API for CommunityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommunityServiceClient interface {
	CreateCommunity(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	GetCommunityByID(ctx context.Context, in *GetCommunityRequest, opts ...grpc.CallOption) (*GetCommunityResponse, error)
	UpdateCommunityByID(ctx context.Context, in *UpdateCommunityRequest, opts ...grpc.CallOption) (*UpdateCommunityResponse, error)
	DeleteCommunityByID(ctx context.Context, in *DeleteCommunityRequest, opts ...grpc.CallOption) (*DeleteCommunityResponse, error)
	GetAllCommunities(ctx context.Context, in *GetCommunitiesRequest, opts ...grpc.CallOption) (*GetCommunitiesResponse, error)
	JoinCommunityMember(ctx context.Context, in *JoinMemberRequest, opts ...grpc.CallOption) (*JoinMemberResponse, error)
	LeftCommunityMember(ctx context.Context, in *LeftMemberRequest, opts ...grpc.CallOption) (*LeftMemberResponse, error)
	JoinCommunityEvent(ctx context.Context, in *JoinEventRequest, opts ...grpc.CallOption) (*JoinEventResponse, error)
	GetEventByID(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error)
	JoinCommunityForum(ctx context.Context, in *JoinForumRequest, opts ...grpc.CallOption) (*JoinForumResponse, error)
	GetForumByID(ctx context.Context, in *GetForumRequest, opts ...grpc.CallOption) (*GetForumResponse, error)
	AddCommentForum(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*AddCommentResponse, error)
	GetForumCommentByID(ctx context.Context, in *GetForumCommentRequest, opts ...grpc.CallOption) (*GetForumCommentResponse, error)
}

type communityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunityServiceClient(cc grpc.ClientConnInterface) CommunityServiceClient {
	return &communityServiceClient{cc}
}

func (c *communityServiceClient) CreateCommunity(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/community.CommunityService/CreateCommunity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) GetCommunityByID(ctx context.Context, in *GetCommunityRequest, opts ...grpc.CallOption) (*GetCommunityResponse, error) {
	out := new(GetCommunityResponse)
	err := c.cc.Invoke(ctx, "/community.CommunityService/GetCommunityByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) UpdateCommunityByID(ctx context.Context, in *UpdateCommunityRequest, opts ...grpc.CallOption) (*UpdateCommunityResponse, error) {
	out := new(UpdateCommunityResponse)
	err := c.cc.Invoke(ctx, "/community.CommunityService/UpdateCommunityByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) DeleteCommunityByID(ctx context.Context, in *DeleteCommunityRequest, opts ...grpc.CallOption) (*DeleteCommunityResponse, error) {
	out := new(DeleteCommunityResponse)
	err := c.cc.Invoke(ctx, "/community.CommunityService/DeleteCommunityByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) GetAllCommunities(ctx context.Context, in *GetCommunitiesRequest, opts ...grpc.CallOption) (*GetCommunitiesResponse, error) {
	out := new(GetCommunitiesResponse)
	err := c.cc.Invoke(ctx, "/community.CommunityService/GetAllCommunities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) JoinCommunityMember(ctx context.Context, in *JoinMemberRequest, opts ...grpc.CallOption) (*JoinMemberResponse, error) {
	out := new(JoinMemberResponse)
	err := c.cc.Invoke(ctx, "/community.CommunityService/JoinCommunityMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) LeftCommunityMember(ctx context.Context, in *LeftMemberRequest, opts ...grpc.CallOption) (*LeftMemberResponse, error) {
	out := new(LeftMemberResponse)
	err := c.cc.Invoke(ctx, "/community.CommunityService/LeftCommunityMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) JoinCommunityEvent(ctx context.Context, in *JoinEventRequest, opts ...grpc.CallOption) (*JoinEventResponse, error) {
	out := new(JoinEventResponse)
	err := c.cc.Invoke(ctx, "/community.CommunityService/JoinCommunityEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) GetEventByID(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error) {
	out := new(GetEventsResponse)
	err := c.cc.Invoke(ctx, "/community.CommunityService/GetEventByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) JoinCommunityForum(ctx context.Context, in *JoinForumRequest, opts ...grpc.CallOption) (*JoinForumResponse, error) {
	out := new(JoinForumResponse)
	err := c.cc.Invoke(ctx, "/community.CommunityService/JoinCommunityForum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) GetForumByID(ctx context.Context, in *GetForumRequest, opts ...grpc.CallOption) (*GetForumResponse, error) {
	out := new(GetForumResponse)
	err := c.cc.Invoke(ctx, "/community.CommunityService/GetForumByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) AddCommentForum(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*AddCommentResponse, error) {
	out := new(AddCommentResponse)
	err := c.cc.Invoke(ctx, "/community.CommunityService/AddCommentForum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communityServiceClient) GetForumCommentByID(ctx context.Context, in *GetForumCommentRequest, opts ...grpc.CallOption) (*GetForumCommentResponse, error) {
	out := new(GetForumCommentResponse)
	err := c.cc.Invoke(ctx, "/community.CommunityService/GetForumCommentByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunityServiceServer is the server API for CommunityService service.
// All implementations must embed UnimplementedCommunityServiceServer
// for forward compatibility
type CommunityServiceServer interface {
	CreateCommunity(context.Context, *CreateRequest) (*CreateResponse, error)
	GetCommunityByID(context.Context, *GetCommunityRequest) (*GetCommunityResponse, error)
	UpdateCommunityByID(context.Context, *UpdateCommunityRequest) (*UpdateCommunityResponse, error)
	DeleteCommunityByID(context.Context, *DeleteCommunityRequest) (*DeleteCommunityResponse, error)
	GetAllCommunities(context.Context, *GetCommunitiesRequest) (*GetCommunitiesResponse, error)
	JoinCommunityMember(context.Context, *JoinMemberRequest) (*JoinMemberResponse, error)
	LeftCommunityMember(context.Context, *LeftMemberRequest) (*LeftMemberResponse, error)
	JoinCommunityEvent(context.Context, *JoinEventRequest) (*JoinEventResponse, error)
	GetEventByID(context.Context, *GetEventsRequest) (*GetEventsResponse, error)
	JoinCommunityForum(context.Context, *JoinForumRequest) (*JoinForumResponse, error)
	GetForumByID(context.Context, *GetForumRequest) (*GetForumResponse, error)
	AddCommentForum(context.Context, *AddCommentRequest) (*AddCommentResponse, error)
	GetForumCommentByID(context.Context, *GetForumCommentRequest) (*GetForumCommentResponse, error)
	mustEmbedUnimplementedCommunityServiceServer()
}

// UnimplementedCommunityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommunityServiceServer struct {
}

func (UnimplementedCommunityServiceServer) CreateCommunity(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommunity not implemented")
}
func (UnimplementedCommunityServiceServer) GetCommunityByID(context.Context, *GetCommunityRequest) (*GetCommunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommunityByID not implemented")
}
func (UnimplementedCommunityServiceServer) UpdateCommunityByID(context.Context, *UpdateCommunityRequest) (*UpdateCommunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCommunityByID not implemented")
}
func (UnimplementedCommunityServiceServer) DeleteCommunityByID(context.Context, *DeleteCommunityRequest) (*DeleteCommunityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommunityByID not implemented")
}
func (UnimplementedCommunityServiceServer) GetAllCommunities(context.Context, *GetCommunitiesRequest) (*GetCommunitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCommunities not implemented")
}
func (UnimplementedCommunityServiceServer) JoinCommunityMember(context.Context, *JoinMemberRequest) (*JoinMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinCommunityMember not implemented")
}
func (UnimplementedCommunityServiceServer) LeftCommunityMember(context.Context, *LeftMemberRequest) (*LeftMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeftCommunityMember not implemented")
}
func (UnimplementedCommunityServiceServer) JoinCommunityEvent(context.Context, *JoinEventRequest) (*JoinEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinCommunityEvent not implemented")
}
func (UnimplementedCommunityServiceServer) GetEventByID(context.Context, *GetEventsRequest) (*GetEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventByID not implemented")
}
func (UnimplementedCommunityServiceServer) JoinCommunityForum(context.Context, *JoinForumRequest) (*JoinForumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinCommunityForum not implemented")
}
func (UnimplementedCommunityServiceServer) GetForumByID(context.Context, *GetForumRequest) (*GetForumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForumByID not implemented")
}
func (UnimplementedCommunityServiceServer) AddCommentForum(context.Context, *AddCommentRequest) (*AddCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCommentForum not implemented")
}
func (UnimplementedCommunityServiceServer) GetForumCommentByID(context.Context, *GetForumCommentRequest) (*GetForumCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForumCommentByID not implemented")
}
func (UnimplementedCommunityServiceServer) mustEmbedUnimplementedCommunityServiceServer() {}

// UnsafeCommunityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommunityServiceServer will
// result in compilation errors.
type UnsafeCommunityServiceServer interface {
	mustEmbedUnimplementedCommunityServiceServer()
}

func RegisterCommunityServiceServer(s grpc.ServiceRegistrar, srv CommunityServiceServer) {
	s.RegisterService(&CommunityService_ServiceDesc, srv)
}

func _CommunityService_CreateCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).CreateCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/community.CommunityService/CreateCommunity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).CreateCommunity(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_GetCommunityByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).GetCommunityByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/community.CommunityService/GetCommunityByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).GetCommunityByID(ctx, req.(*GetCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_UpdateCommunityByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).UpdateCommunityByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/community.CommunityService/UpdateCommunityByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).UpdateCommunityByID(ctx, req.(*UpdateCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_DeleteCommunityByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommunityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).DeleteCommunityByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/community.CommunityService/DeleteCommunityByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).DeleteCommunityByID(ctx, req.(*DeleteCommunityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_GetAllCommunities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommunitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).GetAllCommunities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/community.CommunityService/GetAllCommunities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).GetAllCommunities(ctx, req.(*GetCommunitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_JoinCommunityMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).JoinCommunityMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/community.CommunityService/JoinCommunityMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).JoinCommunityMember(ctx, req.(*JoinMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_LeftCommunityMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeftMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).LeftCommunityMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/community.CommunityService/LeftCommunityMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).LeftCommunityMember(ctx, req.(*LeftMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_JoinCommunityEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).JoinCommunityEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/community.CommunityService/JoinCommunityEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).JoinCommunityEvent(ctx, req.(*JoinEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_GetEventByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).GetEventByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/community.CommunityService/GetEventByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).GetEventByID(ctx, req.(*GetEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_JoinCommunityForum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinForumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).JoinCommunityForum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/community.CommunityService/JoinCommunityForum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).JoinCommunityForum(ctx, req.(*JoinForumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_GetForumByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetForumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).GetForumByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/community.CommunityService/GetForumByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).GetForumByID(ctx, req.(*GetForumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_AddCommentForum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).AddCommentForum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/community.CommunityService/AddCommentForum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).AddCommentForum(ctx, req.(*AddCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunityService_GetForumCommentByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetForumCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunityServiceServer).GetForumCommentByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/community.CommunityService/GetForumCommentByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunityServiceServer).GetForumCommentByID(ctx, req.(*GetForumCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommunityService_ServiceDesc is the grpc.ServiceDesc for CommunityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommunityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "community.CommunityService",
	HandlerType: (*CommunityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCommunity",
			Handler:    _CommunityService_CreateCommunity_Handler,
		},
		{
			MethodName: "GetCommunityByID",
			Handler:    _CommunityService_GetCommunityByID_Handler,
		},
		{
			MethodName: "UpdateCommunityByID",
			Handler:    _CommunityService_UpdateCommunityByID_Handler,
		},
		{
			MethodName: "DeleteCommunityByID",
			Handler:    _CommunityService_DeleteCommunityByID_Handler,
		},
		{
			MethodName: "GetAllCommunities",
			Handler:    _CommunityService_GetAllCommunities_Handler,
		},
		{
			MethodName: "JoinCommunityMember",
			Handler:    _CommunityService_JoinCommunityMember_Handler,
		},
		{
			MethodName: "LeftCommunityMember",
			Handler:    _CommunityService_LeftCommunityMember_Handler,
		},
		{
			MethodName: "JoinCommunityEvent",
			Handler:    _CommunityService_JoinCommunityEvent_Handler,
		},
		{
			MethodName: "GetEventByID",
			Handler:    _CommunityService_GetEventByID_Handler,
		},
		{
			MethodName: "JoinCommunityForum",
			Handler:    _CommunityService_JoinCommunityForum_Handler,
		},
		{
			MethodName: "GetForumByID",
			Handler:    _CommunityService_GetForumByID_Handler,
		},
		{
			MethodName: "AddCommentForum",
			Handler:    _CommunityService_AddCommentForum_Handler,
		},
		{
			MethodName: "GetForumCommentByID",
			Handler:    _CommunityService_GetForumCommentByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/community.proto",
}
