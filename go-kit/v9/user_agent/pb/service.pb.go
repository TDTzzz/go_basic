// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	RpcUserLogin(ctx context.Context, in *Login, opts ...grpc.CallOption) (*LoginAck, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) RpcUserLogin(ctx context.Context, in *Login, opts ...grpc.CallOption) (*LoginAck, error) {
	out := new(LoginAck)
	err := c.cc.Invoke(ctx, "/pb.User/RpcUserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	RpcUserLogin(context.Context, *Login) (*LoginAck, error)
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_RpcUserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Login)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RpcUserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.User/RpcUserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RpcUserLogin(ctx, req.(*Login))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RpcUserLogin",
			Handler:    _User_RpcUserLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_196f7be73e208444) }

var fileDescriptor_service_196f7be73e208444 = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2, 0x2a,
	0x2d, 0x4e, 0x2d, 0x82, 0xf0, 0x8d, 0x0c, 0xb9, 0x58, 0x42, 0x8b, 0x53, 0x8b, 0x84, 0x34, 0xb9,
	0x78, 0x82, 0x0a, 0x92, 0x41, 0x4c, 0x9f, 0xfc, 0xf4, 0xcc, 0x3c, 0x21, 0x4e, 0xbd, 0x82, 0x24,
	0x3d, 0x30, 0x53, 0x8a, 0x07, 0xce, 0x74, 0x4c, 0xce, 0x56, 0x62, 0x48, 0x62, 0x03, 0xeb, 0x34,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xc1, 0xaa, 0xb9, 0x5a, 0x00, 0x00, 0x00,
}
