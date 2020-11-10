// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gkeystore.proto

package gkeystoreproto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetDatabaseRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDatabaseRequest) Reset()         { *m = GetDatabaseRequest{} }
func (m *GetDatabaseRequest) String() string { return proto.CompactTextString(m) }
func (*GetDatabaseRequest) ProtoMessage()    {}
func (*GetDatabaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afa8d7e5e06303c6, []int{0}
}

func (m *GetDatabaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDatabaseRequest.Unmarshal(m, b)
}
func (m *GetDatabaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDatabaseRequest.Marshal(b, m, deterministic)
}
func (m *GetDatabaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDatabaseRequest.Merge(m, src)
}
func (m *GetDatabaseRequest) XXX_Size() int {
	return xxx_messageInfo_GetDatabaseRequest.Size(m)
}
func (m *GetDatabaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDatabaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDatabaseRequest proto.InternalMessageInfo

func (m *GetDatabaseRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetDatabaseRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type GetDatabaseResponse struct {
	DbServer             uint32   `protobuf:"varint,1,opt,name=dbServer,proto3" json:"dbServer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDatabaseResponse) Reset()         { *m = GetDatabaseResponse{} }
func (m *GetDatabaseResponse) String() string { return proto.CompactTextString(m) }
func (*GetDatabaseResponse) ProtoMessage()    {}
func (*GetDatabaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_afa8d7e5e06303c6, []int{1}
}

func (m *GetDatabaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDatabaseResponse.Unmarshal(m, b)
}
func (m *GetDatabaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDatabaseResponse.Marshal(b, m, deterministic)
}
func (m *GetDatabaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDatabaseResponse.Merge(m, src)
}
func (m *GetDatabaseResponse) XXX_Size() int {
	return xxx_messageInfo_GetDatabaseResponse.Size(m)
}
func (m *GetDatabaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDatabaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDatabaseResponse proto.InternalMessageInfo

func (m *GetDatabaseResponse) GetDbServer() uint32 {
	if m != nil {
		return m.DbServer
	}
	return 0
}

func init() {
	proto.RegisterType((*GetDatabaseRequest)(nil), "gkeystoreproto.GetDatabaseRequest")
	proto.RegisterType((*GetDatabaseResponse)(nil), "gkeystoreproto.GetDatabaseResponse")
}

func init() { proto.RegisterFile("gkeystore.proto", fileDescriptor_afa8d7e5e06303c6) }

var fileDescriptor_afa8d7e5e06303c6 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xcf, 0x4e, 0xad,
	0x2c, 0x2e, 0xc9, 0x2f, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0x0b, 0x80,
	0xf9, 0x4a, 0x3e, 0x5c, 0x42, 0xee, 0xa9, 0x25, 0x2e, 0x89, 0x25, 0x89, 0x49, 0x89, 0xc5, 0xa9,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x52, 0x5c, 0x1c, 0xa5, 0xc5, 0xa9, 0x45, 0x79,
	0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0x3e, 0x48, 0xae, 0x20, 0xb1,
	0xb8, 0xb8, 0x3c, 0xbf, 0x28, 0x45, 0x82, 0x09, 0x22, 0x07, 0xe3, 0x2b, 0x19, 0x72, 0x09, 0xa3,
	0x98, 0x56, 0x5c, 0x90, 0x9f, 0x57, 0x0c, 0xd6, 0x92, 0x92, 0x14, 0x9c, 0x5a, 0x54, 0x96, 0x5a,
	0x04, 0x36, 0x8e, 0x37, 0x08, 0xce, 0x37, 0x4a, 0xe2, 0xe2, 0xf0, 0x86, 0xba, 0x48, 0x28, 0x8c,
	0x8b, 0x1b, 0x49, 0xbb, 0x90, 0x92, 0x1e, 0xaa, 0x63, 0xf5, 0x30, 0x5d, 0x2a, 0xa5, 0x8c, 0x57,
	0x0d, 0xc4, 0xfe, 0x24, 0x36, 0xb0, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x96, 0x74,
	0x7f, 0x0e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeystoreClient is the client API for Keystore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeystoreClient interface {
	GetDatabase(ctx context.Context, in *GetDatabaseRequest, opts ...grpc.CallOption) (*GetDatabaseResponse, error)
}

type keystoreClient struct {
	cc grpc.ClientConnInterface
}

func NewKeystoreClient(cc grpc.ClientConnInterface) KeystoreClient {
	return &keystoreClient{cc}
}

func (c *keystoreClient) GetDatabase(ctx context.Context, in *GetDatabaseRequest, opts ...grpc.CallOption) (*GetDatabaseResponse, error) {
	out := new(GetDatabaseResponse)
	err := c.cc.Invoke(ctx, "/gkeystoreproto.Keystore/GetDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeystoreServer is the server API for Keystore service.
type KeystoreServer interface {
	GetDatabase(context.Context, *GetDatabaseRequest) (*GetDatabaseResponse, error)
}

// UnimplementedKeystoreServer can be embedded to have forward compatible implementations.
type UnimplementedKeystoreServer struct {
}

func (*UnimplementedKeystoreServer) GetDatabase(ctx context.Context, req *GetDatabaseRequest) (*GetDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatabase not implemented")
}

func RegisterKeystoreServer(s *grpc.Server, srv KeystoreServer) {
	s.RegisterService(&_Keystore_serviceDesc, srv)
}

func _Keystore_GetDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeystoreServer).GetDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gkeystoreproto.Keystore/GetDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeystoreServer).GetDatabase(ctx, req.(*GetDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Keystore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gkeystoreproto.Keystore",
	HandlerType: (*KeystoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDatabase",
			Handler:    _Keystore_GetDatabase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gkeystore.proto",
}
