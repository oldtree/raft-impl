// Code generated by protoc-gen-go. DO NOT EDIT.
// source: extern-api.proto

package protobuffer

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
	return fileDescriptor_extern_api_838574f0eb510ab3, []int{0}
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
	return fileDescriptor_extern_api_838574f0eb510ab3, []int{1}
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
	return fileDescriptor_extern_api_838574f0eb510ab3, []int{2}
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

type Member struct {
	Member               string               `protobuf:"bytes,1,opt,name=member" json:"member,omitempty"`
	Terms                int64                `protobuf:"varint,2,opt,name=terms" json:"terms,omitempty"`
	Address              string               `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_extern_api_838574f0eb510ab3, []int{3}
}
func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (dst *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(dst, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetMember() string {
	if m != nil {
		return m.Member
	}
	return ""
}

func (m *Member) GetTerms() int64 {
	if m != nil {
		return m.Terms
	}
	return 0
}

func (m *Member) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Member) GetTimestamp() *timestamp.Timestamp {
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
	return fileDescriptor_extern_api_838574f0eb510ab3, []int{4}
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

type ClusterSnapshot struct {
	Master               *Master  `protobuf:"bytes,1,opt,name=master" json:"master,omitempty"`
	Member               *Member  `protobuf:"bytes,2,opt,name=member" json:"member,omitempty"`
	Terms                *Terms   `protobuf:"bytes,3,opt,name=terms" json:"terms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClusterSnapshot) Reset()         { *m = ClusterSnapshot{} }
func (m *ClusterSnapshot) String() string { return proto.CompactTextString(m) }
func (*ClusterSnapshot) ProtoMessage()    {}
func (*ClusterSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_extern_api_838574f0eb510ab3, []int{5}
}
func (m *ClusterSnapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterSnapshot.Unmarshal(m, b)
}
func (m *ClusterSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterSnapshot.Marshal(b, m, deterministic)
}
func (dst *ClusterSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterSnapshot.Merge(dst, src)
}
func (m *ClusterSnapshot) XXX_Size() int {
	return xxx_messageInfo_ClusterSnapshot.Size(m)
}
func (m *ClusterSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterSnapshot proto.InternalMessageInfo

func (m *ClusterSnapshot) GetMaster() *Master {
	if m != nil {
		return m.Master
	}
	return nil
}

func (m *ClusterSnapshot) GetMember() *Member {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *ClusterSnapshot) GetTerms() *Terms {
	if m != nil {
		return m.Terms
	}
	return nil
}

func init() {
	proto.RegisterType((*Timepoint)(nil), "protobuffer.Timepoint")
	proto.RegisterType((*Endpoints)(nil), "protobuffer.Endpoints")
	proto.RegisterType((*Master)(nil), "protobuffer.Master")
	proto.RegisterType((*Member)(nil), "protobuffer.Member")
	proto.RegisterType((*Terms)(nil), "protobuffer.Terms")
	proto.RegisterType((*ClusterSnapshot)(nil), "protobuffer.ClusterSnapshot")
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
	GetClusterSnapshot(ctx context.Context, in *Timepoint, opts ...grpc.CallOption) (*ClusterSnapshot, error)
}

type externApiClient struct {
	cc *grpc.ClientConn
}

func NewExternApiClient(cc *grpc.ClientConn) ExternApiClient {
	return &externApiClient{cc}
}

func (c *externApiClient) GetEndpoints(ctx context.Context, in *Timepoint, opts ...grpc.CallOption) (*Endpoints, error) {
	out := new(Endpoints)
	err := c.cc.Invoke(ctx, "/protobuffer.ExternApi/GetEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externApiClient) GetMaster(ctx context.Context, in *Timepoint, opts ...grpc.CallOption) (*Master, error) {
	out := new(Master)
	err := c.cc.Invoke(ctx, "/protobuffer.ExternApi/GetMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externApiClient) GetCurrentTerms(ctx context.Context, in *Timepoint, opts ...grpc.CallOption) (*Terms, error) {
	out := new(Terms)
	err := c.cc.Invoke(ctx, "/protobuffer.ExternApi/GetCurrentTerms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externApiClient) GetClusterSnapshot(ctx context.Context, in *Timepoint, opts ...grpc.CallOption) (*ClusterSnapshot, error) {
	out := new(ClusterSnapshot)
	err := c.cc.Invoke(ctx, "/protobuffer.ExternApi/GetClusterSnapshot", in, out, opts...)
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
	GetClusterSnapshot(context.Context, *Timepoint) (*ClusterSnapshot, error)
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
		FullMethod: "/protobuffer.ExternApi/GetEndpoints",
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
		FullMethod: "/protobuffer.ExternApi/GetMaster",
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
		FullMethod: "/protobuffer.ExternApi/GetCurrentTerms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternApiServer).GetCurrentTerms(ctx, req.(*Timepoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternApi_GetClusterSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Timepoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternApiServer).GetClusterSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuffer.ExternApi/GetClusterSnapshot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternApiServer).GetClusterSnapshot(ctx, req.(*Timepoint))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExternApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuffer.ExternApi",
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
		{
			MethodName: "GetClusterSnapshot",
			Handler:    _ExternApi_GetClusterSnapshot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "extern-api.proto",
}

func init() { proto.RegisterFile("extern-api.proto", fileDescriptor_extern_api_838574f0eb510ab3) }

var fileDescriptor_extern_api_838574f0eb510ab3 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x41, 0x6b, 0xf2, 0x40,
	0x10, 0x25, 0xe6, 0x33, 0x9f, 0x19, 0x0b, 0x96, 0x6d, 0x91, 0x10, 0x0a, 0x15, 0x4f, 0x82, 0x34,
	0x82, 0xbd, 0x78, 0x69, 0xa1, 0x88, 0x48, 0x0b, 0x5e, 0x52, 0x4f, 0xbd, 0xc5, 0x66, 0xb4, 0x01,
	0x37, 0x09, 0xbb, 0x2b, 0xf4, 0x37, 0xf4, 0xd4, 0x4b, 0xff, 0x6f, 0x71, 0x37, 0xbb, 0xba, 0xa2,
	0x07, 0xe9, 0x29, 0x99, 0x99, 0x37, 0xef, 0xcd, 0xbc, 0x1d, 0xb8, 0xc4, 0x4f, 0x81, 0x2c, 0xbf,
	0x4b, 0xca, 0x2c, 0x2a, 0x59, 0x21, 0x0a, 0xd2, 0x94, 0x9f, 0xc5, 0x66, 0xb9, 0x44, 0x16, 0xde,
	0xae, 0x8a, 0x62, 0xb5, 0xc6, 0x81, 0xce, 0x0d, 0x44, 0x46, 0x91, 0x8b, 0x84, 0x96, 0x0a, 0xdd,
	0x9d, 0x80, 0x3f, 0xcf, 0x28, 0x96, 0x45, 0x96, 0x0b, 0x32, 0x02, 0xdf, 0xd4, 0x03, 0xa7, 0xe3,
	0xf4, 0x9a, 0xc3, 0x30, 0x52, 0x0c, 0x91, 0x66, 0x88, 0xe6, 0x1a, 0x11, 0xef, 0xc0, 0xdd, 0x04,
	0xfc, 0x49, 0x9e, 0x4a, 0x16, 0x4e, 0x42, 0x68, 0x60, 0x15, 0x04, 0x4e, 0xc7, 0xed, 0xf9, 0xb1,
	0x89, 0x6d, 0x89, 0xda, 0x39, 0x12, 0x6f, 0xe0, 0xcd, 0x12, 0x2e, 0x90, 0x91, 0x36, 0x78, 0x54,
	0xfe, 0xc9, 0x19, 0xfd, 0xb8, 0x8a, 0xfe, 0xc0, 0xfd, 0xe5, 0x80, 0x37, 0x43, 0xba, 0xa8, 0xc8,
	0xe5, 0x9f, 0x21, 0x57, 0xf9, 0x6b, 0xa8, 0x0b, 0x64, 0x94, 0x4b, 0x62, 0x37, 0x56, 0x01, 0x09,
	0xe0, 0x7f, 0x92, 0xa6, 0x0c, 0x39, 0x0f, 0x5c, 0x09, 0xd7, 0xa1, 0x3d, 0xcc, 0xbf, 0x73, 0x86,
	0xe1, 0x50, 0x9f, 0x4b, 0x72, 0x23, 0xe9, 0xec, 0x4b, 0x86, 0xd0, 0x78, 0x2f, 0x28, 0xcd, 0xc4,
	0x73, 0x5a, 0xcd, 0x62, 0x62, 0x5b, 0xd4, 0x3d, 0x47, 0xf4, 0xc7, 0x81, 0xd6, 0x78, 0xbd, 0xd9,
	0xfa, 0xf8, 0x9a, 0x27, 0x25, 0xff, 0x28, 0x04, 0xe9, 0x5b, 0x3e, 0x37, 0x87, 0x57, 0xd1, 0xde,
	0x69, 0x45, 0xea, 0x31, 0x8c, 0xf9, 0x7d, 0xe3, 0x5b, 0xed, 0x18, 0x58, 0x96, 0x8c, 0x99, 0x3d,
	0xbd, 0x99, 0x9a, 0x91, 0x58, 0x58, 0xb9, 0x7c, 0xb5, 0xed, 0xf0, 0xbb, 0x06, 0xfe, 0x44, 0x9e,
	0xf8, 0x53, 0x99, 0x91, 0x47, 0xb8, 0x98, 0xa2, 0xd8, 0x5d, 0x5a, 0xdb, 0x6e, 0xd4, 0x87, 0x1c,
	0xda, 0xf9, 0x1d, 0x7e, 0x04, 0xfe, 0x14, 0x85, 0x3e, 0xa3, 0x13, 0xcd, 0xc7, 0xd6, 0x24, 0x0f,
	0xd0, 0x9a, 0xa2, 0x18, 0x6f, 0x18, 0xc3, 0x5c, 0xa8, 0xe7, 0x39, 0xd5, 0x7f, 0x64, 0x1b, 0xf2,
	0x02, 0x64, 0xdb, 0x7e, 0x60, 0xf0, 0x29, 0x86, 0x1b, 0x2b, 0x7f, 0xd0, 0xb5, 0xf0, 0x64, 0xf1,
	0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x83, 0x0b, 0xa8, 0x9d, 0xfb, 0x03, 0x00, 0x00,
}