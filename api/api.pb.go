// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Api_Request struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Api_Request) Reset()         { *m = Api_Request{} }
func (m *Api_Request) String() string { return proto.CompactTextString(m) }
func (*Api_Request) ProtoMessage()    {}
func (*Api_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{0}
}

func (m *Api_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Api_Request.Unmarshal(m, b)
}
func (m *Api_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Api_Request.Marshal(b, m, deterministic)
}
func (m *Api_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Api_Request.Merge(m, src)
}
func (m *Api_Request) XXX_Size() int {
	return xxx_messageInfo_Api_Request.Size(m)
}
func (m *Api_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Api_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Api_Request proto.InternalMessageInfo

func (m *Api_Request) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Api_Response struct {
	Code                 int32    `protobuf:"zigzag32,1,opt,name=code,proto3" json:"code,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Api_Response) Reset()         { *m = Api_Response{} }
func (m *Api_Response) String() string { return proto.CompactTextString(m) }
func (*Api_Response) ProtoMessage()    {}
func (*Api_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b40cafcd4234784, []int{1}
}

func (m *Api_Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Api_Response.Unmarshal(m, b)
}
func (m *Api_Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Api_Response.Marshal(b, m, deterministic)
}
func (m *Api_Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Api_Response.Merge(m, src)
}
func (m *Api_Response) XXX_Size() int {
	return xxx_messageInfo_Api_Response.Size(m)
}
func (m *Api_Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Api_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Api_Response proto.InternalMessageInfo

func (m *Api_Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Api_Response) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Api_Request)(nil), "api.Api_Request")
	proto.RegisterType((*Api_Response)(nil), "api.Api_Response")
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor_1b40cafcd4234784) }

var fileDescriptor_1b40cafcd4234784 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2c, 0xc8, 0xd4,
	0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x52,
	0xe4, 0xe2, 0x76, 0x2c, 0xc8, 0x8c, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2,
	0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0xcc,
	0xb8, 0x78, 0x20, 0x4a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x41, 0x6a, 0x92, 0xf3, 0x53, 0x52,
	0xc1, 0x6a, 0x04, 0x83, 0xc0, 0x6c, 0xb8, 0x3e, 0x26, 0x84, 0x3e, 0xa3, 0x9f, 0x8c, 0x5c, 0xec,
	0x21, 0xa9, 0xc5, 0x25, 0x8e, 0x05, 0x99, 0x42, 0x7a, 0x5c, 0xac, 0xa1, 0x79, 0x89, 0x45, 0x95,
	0x42, 0x02, 0x7a, 0x20, 0x07, 0x20, 0x59, 0x29, 0x25, 0x88, 0x24, 0x02, 0xb1, 0x41, 0x89, 0x41,
	0xc8, 0x8e, 0x4b, 0x38, 0x38, 0xb5, 0xa8, 0x2c, 0xb5, 0x28, 0x38, 0x33, 0x25, 0x35, 0xb8, 0xa4,
	0x28, 0x35, 0x31, 0x37, 0x33, 0x2f, 0x9d, 0x48, 0xdd, 0x06, 0x8c, 0x20, 0xfd, 0xce, 0x39, 0x99,
	0xa9, 0x79, 0x25, 0xe4, 0xe8, 0xd7, 0x00, 0xe9, 0x17, 0x74, 0xca, 0x4c, 0xc9, 0x24, 0x4f, 0xb7,
	0x01, 0xa3, 0x13, 0x6b, 0x14, 0x28, 0x74, 0x93, 0xd8, 0xc0, 0x21, 0x6d, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xb7, 0xbe, 0x43, 0x8f, 0x7a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestApiClient is the client API for TestApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestApiClient interface {
	Unary(ctx context.Context, in *Api_Request, opts ...grpc.CallOption) (*Api_Response, error)
	ServerSideStreaming(ctx context.Context, in *Api_Request, opts ...grpc.CallOption) (TestApi_ServerSideStreamingClient, error)
	ClientSideStreaming(ctx context.Context, opts ...grpc.CallOption) (TestApi_ClientSideStreamingClient, error)
	BidiSideStreaming(ctx context.Context, opts ...grpc.CallOption) (TestApi_BidiSideStreamingClient, error)
}

type testApiClient struct {
	cc *grpc.ClientConn
}

func NewTestApiClient(cc *grpc.ClientConn) TestApiClient {
	return &testApiClient{cc}
}

func (c *testApiClient) Unary(ctx context.Context, in *Api_Request, opts ...grpc.CallOption) (*Api_Response, error) {
	out := new(Api_Response)
	err := c.cc.Invoke(ctx, "/api.TestApi/Unary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testApiClient) ServerSideStreaming(ctx context.Context, in *Api_Request, opts ...grpc.CallOption) (TestApi_ServerSideStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestApi_serviceDesc.Streams[0], "/api.TestApi/ServerSideStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &testApiServerSideStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestApi_ServerSideStreamingClient interface {
	Recv() (*Api_Response, error)
	grpc.ClientStream
}

type testApiServerSideStreamingClient struct {
	grpc.ClientStream
}

func (x *testApiServerSideStreamingClient) Recv() (*Api_Response, error) {
	m := new(Api_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testApiClient) ClientSideStreaming(ctx context.Context, opts ...grpc.CallOption) (TestApi_ClientSideStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestApi_serviceDesc.Streams[1], "/api.TestApi/ClientSideStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &testApiClientSideStreamingClient{stream}
	return x, nil
}

type TestApi_ClientSideStreamingClient interface {
	Send(*Api_Request) error
	CloseAndRecv() (*Api_Response, error)
	grpc.ClientStream
}

type testApiClientSideStreamingClient struct {
	grpc.ClientStream
}

func (x *testApiClientSideStreamingClient) Send(m *Api_Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testApiClientSideStreamingClient) CloseAndRecv() (*Api_Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Api_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testApiClient) BidiSideStreaming(ctx context.Context, opts ...grpc.CallOption) (TestApi_BidiSideStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestApi_serviceDesc.Streams[2], "/api.TestApi/BidiSideStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &testApiBidiSideStreamingClient{stream}
	return x, nil
}

type TestApi_BidiSideStreamingClient interface {
	Send(*Api_Request) error
	Recv() (*Api_Response, error)
	grpc.ClientStream
}

type testApiBidiSideStreamingClient struct {
	grpc.ClientStream
}

func (x *testApiBidiSideStreamingClient) Send(m *Api_Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testApiBidiSideStreamingClient) Recv() (*Api_Response, error) {
	m := new(Api_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestApiServer is the server API for TestApi service.
type TestApiServer interface {
	Unary(context.Context, *Api_Request) (*Api_Response, error)
	ServerSideStreaming(*Api_Request, TestApi_ServerSideStreamingServer) error
	ClientSideStreaming(TestApi_ClientSideStreamingServer) error
	BidiSideStreaming(TestApi_BidiSideStreamingServer) error
}

// UnimplementedTestApiServer can be embedded to have forward compatible implementations.
type UnimplementedTestApiServer struct {
}

func (*UnimplementedTestApiServer) Unary(ctx context.Context, req *Api_Request) (*Api_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary not implemented")
}
func (*UnimplementedTestApiServer) ServerSideStreaming(req *Api_Request, srv TestApi_ServerSideStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerSideStreaming not implemented")
}
func (*UnimplementedTestApiServer) ClientSideStreaming(srv TestApi_ClientSideStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientSideStreaming not implemented")
}
func (*UnimplementedTestApiServer) BidiSideStreaming(srv TestApi_BidiSideStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method BidiSideStreaming not implemented")
}

func RegisterTestApiServer(s *grpc.Server, srv TestApiServer) {
	s.RegisterService(&_TestApi_serviceDesc, srv)
}

func _TestApi_Unary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Api_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestApiServer).Unary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TestApi/Unary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestApiServer).Unary(ctx, req.(*Api_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestApi_ServerSideStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Api_Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestApiServer).ServerSideStreaming(m, &testApiServerSideStreamingServer{stream})
}

type TestApi_ServerSideStreamingServer interface {
	Send(*Api_Response) error
	grpc.ServerStream
}

type testApiServerSideStreamingServer struct {
	grpc.ServerStream
}

func (x *testApiServerSideStreamingServer) Send(m *Api_Response) error {
	return x.ServerStream.SendMsg(m)
}

func _TestApi_ClientSideStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestApiServer).ClientSideStreaming(&testApiClientSideStreamingServer{stream})
}

type TestApi_ClientSideStreamingServer interface {
	SendAndClose(*Api_Response) error
	Recv() (*Api_Request, error)
	grpc.ServerStream
}

type testApiClientSideStreamingServer struct {
	grpc.ServerStream
}

func (x *testApiClientSideStreamingServer) SendAndClose(m *Api_Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testApiClientSideStreamingServer) Recv() (*Api_Request, error) {
	m := new(Api_Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestApi_BidiSideStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestApiServer).BidiSideStreaming(&testApiBidiSideStreamingServer{stream})
}

type TestApi_BidiSideStreamingServer interface {
	Send(*Api_Response) error
	Recv() (*Api_Request, error)
	grpc.ServerStream
}

type testApiBidiSideStreamingServer struct {
	grpc.ServerStream
}

func (x *testApiBidiSideStreamingServer) Send(m *Api_Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testApiBidiSideStreamingServer) Recv() (*Api_Request, error) {
	m := new(Api_Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.TestApi",
	HandlerType: (*TestApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unary",
			Handler:    _TestApi_Unary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerSideStreaming",
			Handler:       _TestApi_ServerSideStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientSideStreaming",
			Handler:       _TestApi_ClientSideStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidiSideStreaming",
			Handler:       _TestApi_BidiSideStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/api.proto",
}
