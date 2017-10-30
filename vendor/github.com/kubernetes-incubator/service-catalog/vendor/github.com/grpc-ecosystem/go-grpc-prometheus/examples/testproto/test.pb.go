// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package mwitkow_testproto is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Empty
	PingRequest
	PingResponse
*/
package mwitkow_testproto

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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PingRequest struct {
	Value             string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	SleepTimeMs       int32  `protobuf:"varint,2,opt,name=sleep_time_ms,json=sleepTimeMs" json:"sleep_time_ms,omitempty"`
	ErrorCodeReturned uint32 `protobuf:"varint,3,opt,name=error_code_returned,json=errorCodeReturned" json:"error_code_returned,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PingRequest) GetSleepTimeMs() int32 {
	if m != nil {
		return m.SleepTimeMs
	}
	return 0
}

func (m *PingRequest) GetErrorCodeReturned() uint32 {
	if m != nil {
		return m.ErrorCodeReturned
	}
	return 0
}

type PingResponse struct {
	Value   string `protobuf:"bytes,1,opt,name=Value,json=value" json:"Value,omitempty"`
	Counter int32  `protobuf:"varint,2,opt,name=counter" json:"counter,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PingResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PingResponse) GetCounter() int32 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "mwitkow.testproto.Empty")
	proto.RegisterType((*PingRequest)(nil), "mwitkow.testproto.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "mwitkow.testproto.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	PingEmpty(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingResponse, error)
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	PingError(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*Empty, error)
	PingList(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (TestService_PingListClient, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) PingEmpty(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/mwitkow.testproto.TestService/PingEmpty", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/mwitkow.testproto.TestService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PingError(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/mwitkow.testproto.TestService/PingError", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) PingList(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (TestService_PingListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[0], c.cc, "/mwitkow.testproto.TestService/PingList", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServicePingListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_PingListClient interface {
	Recv() (*PingResponse, error)
	grpc.ClientStream
}

type testServicePingListClient struct {
	grpc.ClientStream
}

func (x *testServicePingListClient) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TestService service

type TestServiceServer interface {
	PingEmpty(context.Context, *Empty) (*PingResponse, error)
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	PingError(context.Context, *PingRequest) (*Empty, error)
	PingList(*PingRequest, TestService_PingListServer) error
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_PingEmpty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingEmpty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mwitkow.testproto.TestService/PingEmpty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingEmpty(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mwitkow.testproto.TestService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PingError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).PingError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mwitkow.testproto.TestService/PingError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).PingError(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_PingList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PingRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).PingList(m, &testServicePingListServer{stream})
}

type TestService_PingListServer interface {
	Send(*PingResponse) error
	grpc.ServerStream
}

type testServicePingListServer struct {
	grpc.ServerStream
}

func (x *testServicePingListServer) Send(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mwitkow.testproto.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingEmpty",
			Handler:    _TestService_PingEmpty_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _TestService_Ping_Handler,
		},
		{
			MethodName: "PingError",
			Handler:    _TestService_PingError_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingList",
			Handler:       _TestService_PingList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xd7, 0x69, 0x9d, 0x7b, 0x75, 0x87, 0x45, 0x0f, 0xc5, 0x83, 0x96, 0x9c, 0x7a, 0x0a,
	0xa2, 0x77, 0x2f, 0x22, 0x2a, 0x28, 0x4a, 0x1c, 0x5e, 0x8b, 0xb6, 0x0f, 0x09, 0x2e, 0x4d, 0x4d,
	0x5e, 0x57, 0xfc, 0xdf, 0xfc, 0xe3, 0x24, 0x59, 0x05, 0x61, 0x0e, 0x3d, 0xec, 0x98, 0xcf, 0xf7,
	0xf1, 0xfd, 0x11, 0x00, 0x42, 0x47, 0xa2, 0xb1, 0x86, 0x0c, 0x9b, 0xea, 0x4e, 0xd1, 0x9b, 0xe9,
	0x84, 0x67, 0x01, 0xf1, 0x11, 0xc4, 0x97, 0xba, 0xa1, 0x0f, 0xde, 0x41, 0xf2, 0xa0, 0xea, 0x57,
	0x89, 0xef, 0x2d, 0x3a, 0x62, 0x07, 0x10, 0x2f, 0x9e, 0xe7, 0x2d, 0xa6, 0x51, 0x16, 0xe5, 0x63,
	0xb9, 0x7c, 0x30, 0x0e, 0x13, 0x37, 0x47, 0x6c, 0x0a, 0x52, 0x1a, 0x0b, 0xed, 0xd2, 0x61, 0x16,
	0xe5, 0xb1, 0x4c, 0x02, 0x9c, 0x29, 0x8d, 0x77, 0x8e, 0x09, 0xd8, 0x47, 0x6b, 0x8d, 0x2d, 0x4a,
	0x53, 0x61, 0x61, 0x91, 0x5a, 0x5b, 0x63, 0x95, 0x6e, 0x65, 0x51, 0x3e, 0x91, 0xd3, 0x20, 0x5d,
	0x98, 0x0a, 0x65, 0x2f, 0xf0, 0x73, 0xd8, 0x5b, 0x06, 0xbb, 0xc6, 0xd4, 0x0e, 0x7d, 0xf2, 0xd3,
	0x6a, 0x72, 0x0a, 0xa3, 0xd2, 0xb4, 0x35, 0xa1, 0xed, 0x33, 0xbf, 0x9f, 0xa7, 0x9f, 0x43, 0x48,
	0x66, 0xe8, 0xe8, 0x11, 0xed, 0x42, 0x95, 0xc8, 0xae, 0x61, 0xec, 0xfd, 0xc2, 0x2a, 0x96, 0x8a,
	0x95, 0xc9, 0x22, 0x28, 0x87, 0xc7, 0xbf, 0x28, 0x3f, 0x7b, 0xf0, 0x01, 0xbb, 0x81, 0x6d, 0x4f,
	0xd8, 0xd1, 0xda, 0xd3, 0xf0, 0x57, 0xff, 0xb1, 0xba, 0xea, 0x4b, 0xf9, 0xf5, 0x7f, 0xfa, 0xad,
	0x2d, 0xcd, 0x07, 0xec, 0x1e, 0x76, 0xfd, 0xe9, 0xad, 0x72, 0xb4, 0x81, 0x5e, 0x27, 0xd1, 0xcb,
	0x4e, 0xe0, 0x67, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x38, 0x3e, 0x02, 0xe9, 0x28, 0x02, 0x00,
	0x00,
}
