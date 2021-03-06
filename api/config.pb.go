// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package config

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

type PutConfigRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutConfigRequest) Reset()         { *m = PutConfigRequest{} }
func (m *PutConfigRequest) String() string { return proto.CompactTextString(m) }
func (*PutConfigRequest) ProtoMessage()    {}
func (*PutConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}

func (m *PutConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutConfigRequest.Unmarshal(m, b)
}
func (m *PutConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutConfigRequest.Marshal(b, m, deterministic)
}
func (m *PutConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutConfigRequest.Merge(m, src)
}
func (m *PutConfigRequest) XXX_Size() int {
	return xxx_messageInfo_PutConfigRequest.Size(m)
}
func (m *PutConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutConfigRequest proto.InternalMessageInfo

func (m *PutConfigRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutConfigRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetConfigRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConfigRequest) Reset()         { *m = GetConfigRequest{} }
func (m *GetConfigRequest) String() string { return proto.CompactTextString(m) }
func (*GetConfigRequest) ProtoMessage()    {}
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{1}
}

func (m *GetConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConfigRequest.Unmarshal(m, b)
}
func (m *GetConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConfigRequest.Marshal(b, m, deterministic)
}
func (m *GetConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfigRequest.Merge(m, src)
}
func (m *GetConfigRequest) XXX_Size() int {
	return xxx_messageInfo_GetConfigRequest.Size(m)
}
func (m *GetConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfigRequest proto.InternalMessageInfo

func (m *GetConfigRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Config struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{2}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Config) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*PutConfigRequest)(nil), "config.PutConfigRequest")
	proto.RegisterType((*GetConfigRequest)(nil), "config.GetConfigRequest")
	proto.RegisterType((*Config)(nil), "config.Config")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4) }

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0xac, 0xb8, 0x04,
	0x02, 0x4a, 0x4b, 0x9c, 0xc1, 0x9c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x01, 0x2e,
	0xe6, 0xec, 0xd4, 0x4a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x48, 0x84, 0x8b,
	0xb5, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x09, 0x2c, 0x06, 0xe1, 0x28, 0xa9, 0x70, 0x09, 0xb8,
	0xa7, 0x12, 0xd2, 0xab, 0x64, 0xc0, 0xc5, 0x06, 0x51, 0x42, 0xac, 0xb9, 0x46, 0x85, 0x5c, 0xbc,
	0x10, 0x1d, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xfa, 0x5c, 0xcc, 0x01, 0xa5, 0x25,
	0x42, 0x12, 0x7a, 0x50, 0x2f, 0xa0, 0xbb, 0x58, 0x8a, 0x0f, 0x26, 0x03, 0xb5, 0x49, 0x9f, 0x8b,
	0xd9, 0x3d, 0x15, 0x49, 0x03, 0xba, 0x33, 0xd1, 0x35, 0x24, 0xb1, 0x81, 0x43, 0xc5, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x9f, 0x60, 0xc3, 0x19, 0x25, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigServiceClient is the client API for ConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigServiceClient interface {
	Put(ctx context.Context, in *PutConfigRequest, opts ...grpc.CallOption) (*Config, error)
	Get(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*Config, error)
}

type configServiceClient struct {
	cc *grpc.ClientConn
}

func NewConfigServiceClient(cc *grpc.ClientConn) ConfigServiceClient {
	return &configServiceClient{cc}
}

func (c *configServiceClient) Put(ctx context.Context, in *PutConfigRequest, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, "/config.ConfigService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) Get(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, "/config.ConfigService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServiceServer is the server API for ConfigService service.
type ConfigServiceServer interface {
	Put(context.Context, *PutConfigRequest) (*Config, error)
	Get(context.Context, *GetConfigRequest) (*Config, error)
}

// UnimplementedConfigServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConfigServiceServer struct {
}

func (*UnimplementedConfigServiceServer) Put(ctx context.Context, req *PutConfigRequest) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (*UnimplementedConfigServiceServer) Get(ctx context.Context, req *GetConfigRequest) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterConfigServiceServer(s *grpc.Server, srv ConfigServiceServer) {
	s.RegisterService(&_ConfigService_serviceDesc, srv)
}

func _ConfigService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).Put(ctx, req.(*PutConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.ConfigService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).Get(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "config.ConfigService",
	HandlerType: (*ConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _ConfigService_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ConfigService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config.proto",
}
