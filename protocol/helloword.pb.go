// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol/helloword.proto

package protocol

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_daa63b1c226d28e9, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_daa63b1c226d28e9, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Buy request
type BuyRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuyRequest) Reset()         { *m = BuyRequest{} }
func (m *BuyRequest) String() string { return proto.CompactTextString(m) }
func (*BuyRequest) ProtoMessage()    {}
func (*BuyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_daa63b1c226d28e9, []int{2}
}

func (m *BuyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyRequest.Unmarshal(m, b)
}
func (m *BuyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyRequest.Marshal(b, m, deterministic)
}
func (m *BuyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyRequest.Merge(m, src)
}
func (m *BuyRequest) XXX_Size() int {
	return xxx_messageInfo_BuyRequest.Size(m)
}
func (m *BuyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuyRequest proto.InternalMessageInfo

func (m *BuyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Buy response
type BuyReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Price                int32    `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuyReply) Reset()         { *m = BuyReply{} }
func (m *BuyReply) String() string { return proto.CompactTextString(m) }
func (*BuyReply) ProtoMessage()    {}
func (*BuyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_daa63b1c226d28e9, []int{3}
}

func (m *BuyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuyReply.Unmarshal(m, b)
}
func (m *BuyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuyReply.Marshal(b, m, deterministic)
}
func (m *BuyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuyReply.Merge(m, src)
}
func (m *BuyReply) XXX_Size() int {
	return xxx_messageInfo_BuyReply.Size(m)
}
func (m *BuyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BuyReply.DiscardUnknown(m)
}

var xxx_messageInfo_BuyReply proto.InternalMessageInfo

func (m *BuyReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *BuyReply) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "protocol.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "protocol.HelloReply")
	proto.RegisterType((*BuyRequest)(nil), "protocol.BuyRequest")
	proto.RegisterType((*BuyReply)(nil), "protocol.BuyReply")
}

func init() { proto.RegisterFile("protocol/helloword.proto", fileDescriptor_daa63b1c226d28e9) }

var fileDescriptor_daa63b1c226d28e9 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0x2f, 0xcf, 0x2f, 0x4a, 0xd1, 0x03,
	0x0b, 0x09, 0x71, 0xc0, 0x64, 0x94, 0x94, 0xb8, 0x78, 0x3c, 0x40, 0x92, 0x41, 0xa9, 0x85, 0xa5,
	0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x60, 0xb6, 0x92, 0x1a, 0x17, 0x17, 0x54, 0x4d, 0x41, 0x4e, 0xa5, 0x90, 0x04, 0x17,
	0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x3a, 0x4c, 0x11, 0x8c, 0xab, 0xa4, 0xc0, 0xc5, 0xe5, 0x54,
	0x5a, 0x89, 0xcf, 0x24, 0x2b, 0x2e, 0x0e, 0xb0, 0x0a, 0xbc, 0xe6, 0x08, 0x89, 0x70, 0xb1, 0x16,
	0x14, 0x65, 0x26, 0xa7, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x38, 0x46, 0xae, 0x5c,
	0xec, 0xee, 0x45, 0xa9, 0xa9, 0x25, 0xa9, 0x45, 0x42, 0x56, 0x5c, 0x1c, 0xc1, 0x89, 0x95, 0x60,
	0x37, 0x09, 0x89, 0xe9, 0xc1, 0xfc, 0xa2, 0x87, 0xec, 0x11, 0x29, 0x11, 0x0c, 0xf1, 0x82, 0x9c,
	0x4a, 0x25, 0x06, 0x23, 0x47, 0x2e, 0x4e, 0xa7, 0xd2, 0xca, 0xe0, 0xd4, 0xa2, 0xb2, 0xd4, 0x22,
	0x21, 0x13, 0x2e, 0x36, 0xa7, 0xd2, 0xca, 0x80, 0xa2, 0x7c, 0x21, 0x24, 0xe5, 0x08, 0x3f, 0x48,
	0x09, 0xa1, 0x89, 0x82, 0x8d, 0x70, 0x62, 0x5b, 0xc4, 0xc4, 0xec, 0xe1, 0x13, 0x9e, 0xc4, 0x06,
	0x96, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x95, 0x0a, 0xe0, 0x8e, 0x68, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/protocol.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol/helloword.proto",
}

// BuyServerClient is the client API for BuyServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuyServerClient interface {
	BuyPro(ctx context.Context, in *BuyRequest, opts ...grpc.CallOption) (*BuyReply, error)
}

type buyServerClient struct {
	cc *grpc.ClientConn
}

func NewBuyServerClient(cc *grpc.ClientConn) BuyServerClient {
	return &buyServerClient{cc}
}

func (c *buyServerClient) BuyPro(ctx context.Context, in *BuyRequest, opts ...grpc.CallOption) (*BuyReply, error) {
	out := new(BuyReply)
	err := c.cc.Invoke(ctx, "/protocol.BuyServer/BuyPro", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuyServerServer is the server API for BuyServer service.
type BuyServerServer interface {
	BuyPro(context.Context, *BuyRequest) (*BuyReply, error)
}

func RegisterBuyServerServer(s *grpc.Server, srv BuyServerServer) {
	s.RegisterService(&_BuyServer_serviceDesc, srv)
}

func _BuyServer_BuyPro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuyServerServer).BuyPro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.BuyServer/BuyPro",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuyServerServer).BuyPro(ctx, req.(*BuyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BuyServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.BuyServer",
	HandlerType: (*BuyServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuyPro",
			Handler:    _BuyServer_BuyPro_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol/helloword.proto",
}
