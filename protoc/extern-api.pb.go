// Code generated by protoc-gen-go. DO NOT EDIT.
// source: extern-api.proto

package extern_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type Timepoint struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Timepoint) Reset()         { *m = Timepoint{} }
func (m *Timepoint) String() string { return proto.CompactTextString(m) }
func (*Timepoint) ProtoMessage()    {}
func (*Timepoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_extern_api_f7a471be37fb4c9e, []int{0}
}
func (m *Timepoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timepoint.Unmarshal(m, b)
}
func (m *Timepoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timepoint.Marshal(b, m, deterministic)
}
func (dst *Timepoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timepoint.Merge(dst, src)
}
func (m *Timepoint) XXX_Size() int {
	return xxx_messageInfo_Timepoint.Size(m)
}
func (m *Timepoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Timepoint.DiscardUnknown(m)
}

var xxx_messageInfo_Timepoint proto.InternalMessageInfo

func (m *Timepoint) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type Endpoints struct {
	Endpoint             []string             `protobuf:"bytes,1,rep,name=endpoint" json:"endpoint,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Endpoints) Reset()         { *m = Endpoints{} }
func (m *Endpoints) String() string { return proto.CompactTextString(m) }
func (*Endpoints) ProtoMessage()    {}
func (*Endpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_extern_api_f7a471be37fb4c9e, []int{1}
}
func (m *Endpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoints.Unmarshal(m, b)
}
func (m *Endpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoints.Marshal(b, m, deterministic)
}
func (dst *Endpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoints.Merge(dst, src)
}
func (m *Endpoints) XXX_Size() int {
	return xxx_messageInfo_Endpoints.Size(m)
}
func (m *Endpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoints.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoints proto.InternalMessageInfo

func (m *Endpoints) GetEndpoint() []string {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *Endpoints) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type Master struct {
	Master               string               `protobuf:"bytes,1,opt,name=master" json:"master,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Master) Reset()         { *m = Master{} }
func (m *Master) String() string { return proto.CompactTextString(m) }
func (*Master) ProtoMessage()    {}
func (*Master) Descriptor() ([]byte, []int) {
	return fileDescriptor_extern_api_f7a471be37fb4c9e, []int{2}
}
func (m *Master) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Master.Unmarshal(m, b)
}
func (m *Master) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Master.Marshal(b, m, deterministic)
}
func (dst *Master) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Master.Merge(dst, src)
}
func (m *Master) XXX_Size() int {
	return xxx_messageInfo_Master.Size(m)
}
func (m *Master) XXX_DiscardUnknown() {
	xxx_messageInfo_Master.DiscardUnknown(m)
}

var xxx_messageInfo_Master proto.InternalMessageInfo

func (m *Master) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *Master) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type Terms struct {
	Terms                int64                `protobuf:"varint,1,opt,name=terms" json:"terms,omitempty"`
	CommitId             int64                `protobuf:"varint,2,opt,name=commitId" json:"commitId,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Terms) Reset()         { *m = Terms{} }
func (m *Terms) String() string { return proto.CompactTextString(m) }
func (*Terms) ProtoMessage()    {}
func (*Terms) Descriptor() ([]byte, []int) {
	return fileDescriptor_extern_api_f7a471be37fb4c9e, []int{3}
}
func (m *Terms) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Terms.Unmarshal(m, b)
}
func (m *Terms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Terms.Marshal(b, m, deterministic)
}
func (dst *Terms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Terms.Merge(dst, src)
}
func (m *Terms) XXX_Size() int {
	return xxx_messageInfo_Terms.Size(m)
}
func (m *Terms) XXX_DiscardUnknown() {
	xxx_messageInfo_Terms.DiscardUnknown(m)
}

var xxx_messageInfo_Terms proto.InternalMessageInfo

func (m *Terms) GetTerms() int64 {
	if m != nil {
		return m.Terms
	}
	return 0
}

func (m *Terms) GetCommitId() int64 {
	if m != nil {
		return m.CommitId
	}
	return 0
}

func (m *Terms) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*Timepoint)(nil), "Timepoint")
	proto.RegisterType((*Endpoints)(nil), "Endpoints")
	proto.RegisterType((*Master)(nil), "Master")
	proto.RegisterType((*Terms)(nil), "Terms")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExternApiClient is the client API for ExternApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExternApiClient interface {
	GetEndpoints(ctx context.Context, in *Timepoint, opts ...grpc.CallOption) (*Endpoints, error)
	GetMaster(ctx context.Context, in *Timepoint, opts ...grpc.CallOption) (*Master, error)
	GetCurrentTerms(ctx context.Context, in *Timepoint, opts ...grpc.CallOption) (*Terms, error)
}

type externApiClient struct {
	cc *grpc.ClientConn
}

func NewExternApiClient(cc *grpc.ClientConn) ExternApiClient {
	return &externApiClient{cc}
}

func (c *externApiClient) GetEndpoints(ctx context.Context, in *Timepoint, opts ...grpc.CallOption) (*Endpoints, error) {
	out := new(Endpoints)
	err := c.cc.Invoke(ctx, "/ExternApi/GetEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externApiClient) GetMaster(ctx context.Context, in *Timepoint, opts ...grpc.CallOption) (*Master, error) {
	out := new(Master)
	err := c.cc.Invoke(ctx, "/ExternApi/GetMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externApiClient) GetCurrentTerms(ctx context.Context, in *Timepoint, opts ...grpc.CallOption) (*Terms, error) {
	out := new(Terms)
	err := c.cc.Invoke(ctx, "/ExternApi/GetCurrentTerms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternApiServer is the server API for ExternApi service.
type ExternApiServer interface {
	GetEndpoints(context.Context, *Timepoint) (*Endpoints, error)
	GetMaster(context.Context, *Timepoint) (*Master, error)
	GetCurrentTerms(context.Context, *Timepoint) (*Terms, error)
}

func RegisterExternApiServer(s *grpc.Server, srv ExternApiServer) {
	s.RegisterService(&_ExternApi_serviceDesc, srv)
}

func _ExternApi_GetEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Timepoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternApiServer).GetEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternApi/GetEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternApiServer).GetEndpoints(ctx, req.(*Timepoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternApi_GetMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Timepoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternApiServer).GetMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternApi/GetMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternApiServer).GetMaster(ctx, req.(*Timepoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternApi_GetCurrentTerms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Timepoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternApiServer).GetCurrentTerms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ExternApi/GetCurrentTerms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternApiServer).GetCurrentTerms(ctx, req.(*Timepoint))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExternApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ExternApi",
	HandlerType: (*ExternApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEndpoints",
			Handler:    _ExternApi_GetEndpoints_Handler,
		},
		{
			MethodName: "GetMaster",
			Handler:    _ExternApi_GetMaster_Handler,
		},
		{
			MethodName: "GetCurrentTerms",
			Handler:    _ExternApi_GetCurrentTerms_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "extern-api.proto",
}

func init() { proto.RegisterFile("extern-api.proto", fileDescriptor_extern_api_f7a471be37fb4c9e) }

var fileDescriptor_extern_api_f7a471be37fb4c9e = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xcf, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xc9, 0x37, 0x34, 0x5f, 0x77, 0x14, 0x94, 0x45, 0xa4, 0xe4, 0x62, 0x28, 0x28, 0xbd,
	0xb8, 0x85, 0x7a, 0xf1, 0x2a, 0x12, 0x82, 0x07, 0x2f, 0x4b, 0x4e, 0xde, 0x52, 0x3b, 0x96, 0x05,
	0x37, 0xbb, 0xec, 0x4e, 0xc1, 0x83, 0x7f, 0xbc, 0x74, 0x92, 0x26, 0xed, 0xb1, 0x78, 0xdb, 0x37,
	0xfb, 0xf8, 0xbc, 0xf9, 0x01, 0x57, 0xf8, 0x4d, 0x18, 0xda, 0x87, 0xc6, 0x1b, 0xe5, 0x83, 0x23,
	0x97, 0xdf, 0x6e, 0x9c, 0xdb, 0x7c, 0xe1, 0x82, 0xd5, 0x6a, 0xfb, 0xb9, 0x20, 0x63, 0x31, 0x52,
	0x63, 0x7d, 0x67, 0x98, 0x95, 0x20, 0x6a, 0x63, 0xd1, 0x3b, 0xd3, 0x92, 0x7c, 0x02, 0x31, 0xfc,
	0x4f, 0x93, 0x22, 0x99, 0x9f, 0x2f, 0x73, 0xd5, 0x11, 0xd4, 0x9e, 0xa0, 0xea, 0xbd, 0x43, 0x8f,
	0xe6, 0x59, 0x03, 0xa2, 0x6c, 0xd7, 0x4c, 0x89, 0x32, 0x87, 0x33, 0xec, 0xc5, 0x34, 0x29, 0xd2,
	0xb9, 0xd0, 0x83, 0x3e, 0x8e, 0xf8, 0x77, 0x4a, 0xc4, 0x3b, 0x64, 0x6f, 0x4d, 0x24, 0x0c, 0xf2,
	0x06, 0x32, 0xcb, 0x2f, 0xee, 0x51, 0xe8, 0x5e, 0xfd, 0x81, 0x1d, 0x61, 0x52, 0x63, 0xb0, 0x51,
	0x5e, 0xc3, 0x84, 0x76, 0x0f, 0x26, 0xa7, 0xba, 0x13, 0xbb, 0x81, 0x3e, 0x9c, 0xb5, 0x86, 0x5e,
	0xd7, 0xcc, 0x4d, 0xf5, 0xa0, 0x8f, 0x43, 0xd3, 0x13, 0x42, 0x97, 0x3f, 0x20, 0x4a, 0xbe, 0xd7,
	0xb3, 0x37, 0xf2, 0x1e, 0x2e, 0x2a, 0xa4, 0x71, 0x87, 0xa0, 0x86, 0xb3, 0xe4, 0xa0, 0xc6, 0x7a,
	0x01, 0xa2, 0x42, 0xea, 0x17, 0x71, 0x68, 0xfa, 0xaf, 0xfa, 0xe2, 0x1d, 0x5c, 0x56, 0x48, 0x2f,
	0xdb, 0x10, 0xb0, 0xa5, 0x6e, 0xaa, 0x43, 0x5f, 0xa6, 0xb8, 0xb6, 0xca, 0xb8, 0xb9, 0xc7, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x85, 0x7d, 0x11, 0x34, 0x02, 0x00, 0x00,
}