// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb_gen

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

// GraduateDesignApiClient is the client API for GraduateDesignApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GraduateDesignApiClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	CreateDir(ctx context.Context, in *CreateDirRequest, opts ...grpc.CallOption) (*CreateDirResponse, error)
	UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileReponse, error)
	DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (*DownloadFileResponse, error)
	GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error)
	SearchFile(ctx context.Context, in *SearchFileRequest, opts ...grpc.CallOption) (*SearchFileResponse, error)
}

type graduateDesignApiClient struct {
	cc grpc.ClientConnInterface
}

func NewGraduateDesignApiClient(cc grpc.ClientConnInterface) GraduateDesignApiClient {
	return &graduateDesignApiClient{cc}
}

func (c *graduateDesignApiClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/graduate_design.GraduateDesignApi/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graduateDesignApiClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, "/graduate_design.GraduateDesignApi/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graduateDesignApiClient) CreateDir(ctx context.Context, in *CreateDirRequest, opts ...grpc.CallOption) (*CreateDirResponse, error) {
	out := new(CreateDirResponse)
	err := c.cc.Invoke(ctx, "/graduate_design.GraduateDesignApi/CreateDir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graduateDesignApiClient) UploadFile(ctx context.Context, in *UploadFileRequest, opts ...grpc.CallOption) (*UploadFileReponse, error) {
	out := new(UploadFileReponse)
	err := c.cc.Invoke(ctx, "/graduate_design.GraduateDesignApi/UploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graduateDesignApiClient) DownloadFile(ctx context.Context, in *DownloadFileRequest, opts ...grpc.CallOption) (*DownloadFileResponse, error) {
	out := new(DownloadFileResponse)
	err := c.cc.Invoke(ctx, "/graduate_design.GraduateDesignApi/DownloadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graduateDesignApiClient) GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error) {
	out := new(GetNodeResponse)
	err := c.cc.Invoke(ctx, "/graduate_design.GraduateDesignApi/GetNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *graduateDesignApiClient) SearchFile(ctx context.Context, in *SearchFileRequest, opts ...grpc.CallOption) (*SearchFileResponse, error) {
	out := new(SearchFileResponse)
	err := c.cc.Invoke(ctx, "/graduate_design.GraduateDesignApi/SearchFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GraduateDesignApiServer is the server API for GraduateDesignApi service.
// All implementations must embed UnimplementedGraduateDesignApiServer
// for forward compatibility
type GraduateDesignApiServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	CreateDir(context.Context, *CreateDirRequest) (*CreateDirResponse, error)
	UploadFile(context.Context, *UploadFileRequest) (*UploadFileReponse, error)
	DownloadFile(context.Context, *DownloadFileRequest) (*DownloadFileResponse, error)
	GetNode(context.Context, *GetNodeRequest) (*GetNodeResponse, error)
	SearchFile(context.Context, *SearchFileRequest) (*SearchFileResponse, error)
	mustEmbedUnimplementedGraduateDesignApiServer()
}

// UnimplementedGraduateDesignApiServer must be embedded to have forward compatible implementations.
type UnimplementedGraduateDesignApiServer struct {
}

func (UnimplementedGraduateDesignApiServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedGraduateDesignApiServer) UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedGraduateDesignApiServer) CreateDir(context.Context, *CreateDirRequest) (*CreateDirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDir not implemented")
}
func (UnimplementedGraduateDesignApiServer) UploadFile(context.Context, *UploadFileRequest) (*UploadFileReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedGraduateDesignApiServer) DownloadFile(context.Context, *DownloadFileRequest) (*DownloadFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedGraduateDesignApiServer) GetNode(context.Context, *GetNodeRequest) (*GetNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNode not implemented")
}
func (UnimplementedGraduateDesignApiServer) SearchFile(context.Context, *SearchFileRequest) (*SearchFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchFile not implemented")
}
func (UnimplementedGraduateDesignApiServer) mustEmbedUnimplementedGraduateDesignApiServer() {}

// UnsafeGraduateDesignApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GraduateDesignApiServer will
// result in compilation errors.
type UnsafeGraduateDesignApiServer interface {
	mustEmbedUnimplementedGraduateDesignApiServer()
}

func RegisterGraduateDesignApiServer(s grpc.ServiceRegistrar, srv GraduateDesignApiServer) {
	s.RegisterService(&GraduateDesignApi_ServiceDesc, srv)
}

func _GraduateDesignApi_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraduateDesignApiServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graduate_design.GraduateDesignApi/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraduateDesignApiServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GraduateDesignApi_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraduateDesignApiServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graduate_design.GraduateDesignApi/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraduateDesignApiServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GraduateDesignApi_CreateDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraduateDesignApiServer).CreateDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graduate_design.GraduateDesignApi/CreateDir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraduateDesignApiServer).CreateDir(ctx, req.(*CreateDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GraduateDesignApi_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraduateDesignApiServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graduate_design.GraduateDesignApi/UploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraduateDesignApiServer).UploadFile(ctx, req.(*UploadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GraduateDesignApi_DownloadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraduateDesignApiServer).DownloadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graduate_design.GraduateDesignApi/DownloadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraduateDesignApiServer).DownloadFile(ctx, req.(*DownloadFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GraduateDesignApi_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraduateDesignApiServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graduate_design.GraduateDesignApi/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraduateDesignApiServer).GetNode(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GraduateDesignApi_SearchFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GraduateDesignApiServer).SearchFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/graduate_design.GraduateDesignApi/SearchFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GraduateDesignApiServer).SearchFile(ctx, req.(*SearchFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GraduateDesignApi_ServiceDesc is the grpc.ServiceDesc for GraduateDesignApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GraduateDesignApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "graduate_design.GraduateDesignApi",
	HandlerType: (*GraduateDesignApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _GraduateDesignApi_Ping_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _GraduateDesignApi_UserLogin_Handler,
		},
		{
			MethodName: "CreateDir",
			Handler:    _GraduateDesignApi_CreateDir_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _GraduateDesignApi_UploadFile_Handler,
		},
		{
			MethodName: "DownloadFile",
			Handler:    _GraduateDesignApi_DownloadFile_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _GraduateDesignApi_GetNode_Handler,
		},
		{
			MethodName: "SearchFile",
			Handler:    _GraduateDesignApi_SearchFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
