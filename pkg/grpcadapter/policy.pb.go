// Code generated by protoc-gen-go. DO NOT EDIT.
// source: policy.proto

package grpcadapter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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

type Policy struct {
	Ptype                string   `protobuf:"bytes,1,opt,name=ptype,proto3" json:"ptype,omitempty"`
	Rule                 []string `protobuf:"bytes,2,rep,name=rule,proto3" json:"rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_edd0b7aa982d151a, []int{0}
}
func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (dst *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(dst, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetPtype() string {
	if m != nil {
		return m.Ptype
	}
	return ""
}

func (m *Policy) GetRule() []string {
	if m != nil {
		return m.Rule
	}
	return nil
}

type Policies struct {
	Policies             []*Policy `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Policies) Reset()         { *m = Policies{} }
func (m *Policies) String() string { return proto.CompactTextString(m) }
func (*Policies) ProtoMessage()    {}
func (*Policies) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_edd0b7aa982d151a, []int{1}
}
func (m *Policies) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policies.Unmarshal(m, b)
}
func (m *Policies) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policies.Marshal(b, m, deterministic)
}
func (dst *Policies) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policies.Merge(dst, src)
}
func (m *Policies) XXX_Size() int {
	return xxx_messageInfo_Policies.Size(m)
}
func (m *Policies) XXX_DiscardUnknown() {
	xxx_messageInfo_Policies.DiscardUnknown(m)
}

var xxx_messageInfo_Policies proto.InternalMessageInfo

func (m *Policies) GetPolicies() []*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type AddPolicyReq struct {
	Sec                  string   `protobuf:"bytes,1,opt,name=sec,proto3" json:"sec,omitempty"`
	Ptype                string   `protobuf:"bytes,2,opt,name=ptype,proto3" json:"ptype,omitempty"`
	Rule                 []string `protobuf:"bytes,3,rep,name=rule,proto3" json:"rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPolicyReq) Reset()         { *m = AddPolicyReq{} }
func (m *AddPolicyReq) String() string { return proto.CompactTextString(m) }
func (*AddPolicyReq) ProtoMessage()    {}
func (*AddPolicyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_edd0b7aa982d151a, []int{2}
}
func (m *AddPolicyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPolicyReq.Unmarshal(m, b)
}
func (m *AddPolicyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPolicyReq.Marshal(b, m, deterministic)
}
func (dst *AddPolicyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPolicyReq.Merge(dst, src)
}
func (m *AddPolicyReq) XXX_Size() int {
	return xxx_messageInfo_AddPolicyReq.Size(m)
}
func (m *AddPolicyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPolicyReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddPolicyReq proto.InternalMessageInfo

func (m *AddPolicyReq) GetSec() string {
	if m != nil {
		return m.Sec
	}
	return ""
}

func (m *AddPolicyReq) GetPtype() string {
	if m != nil {
		return m.Ptype
	}
	return ""
}

func (m *AddPolicyReq) GetRule() []string {
	if m != nil {
		return m.Rule
	}
	return nil
}

type RemovePolicyReq struct {
	Sec                  string   `protobuf:"bytes,1,opt,name=sec,proto3" json:"sec,omitempty"`
	Ptype                string   `protobuf:"bytes,2,opt,name=ptype,proto3" json:"ptype,omitempty"`
	Rule                 []string `protobuf:"bytes,3,rep,name=rule,proto3" json:"rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemovePolicyReq) Reset()         { *m = RemovePolicyReq{} }
func (m *RemovePolicyReq) String() string { return proto.CompactTextString(m) }
func (*RemovePolicyReq) ProtoMessage()    {}
func (*RemovePolicyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_edd0b7aa982d151a, []int{3}
}
func (m *RemovePolicyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemovePolicyReq.Unmarshal(m, b)
}
func (m *RemovePolicyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemovePolicyReq.Marshal(b, m, deterministic)
}
func (dst *RemovePolicyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemovePolicyReq.Merge(dst, src)
}
func (m *RemovePolicyReq) XXX_Size() int {
	return xxx_messageInfo_RemovePolicyReq.Size(m)
}
func (m *RemovePolicyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RemovePolicyReq.DiscardUnknown(m)
}

var xxx_messageInfo_RemovePolicyReq proto.InternalMessageInfo

func (m *RemovePolicyReq) GetSec() string {
	if m != nil {
		return m.Sec
	}
	return ""
}

func (m *RemovePolicyReq) GetPtype() string {
	if m != nil {
		return m.Ptype
	}
	return ""
}

func (m *RemovePolicyReq) GetRule() []string {
	if m != nil {
		return m.Rule
	}
	return nil
}

type RemoveFilteredPolicyReq struct {
	Sec                  string   `protobuf:"bytes,1,opt,name=sec,proto3" json:"sec,omitempty"`
	Ptype                string   `protobuf:"bytes,2,opt,name=ptype,proto3" json:"ptype,omitempty"`
	FieldIndex           int64    `protobuf:"varint,3,opt,name=fieldIndex,proto3" json:"fieldIndex,omitempty"`
	FieldValues          []string `protobuf:"bytes,4,rep,name=fieldValues,proto3" json:"fieldValues,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveFilteredPolicyReq) Reset()         { *m = RemoveFilteredPolicyReq{} }
func (m *RemoveFilteredPolicyReq) String() string { return proto.CompactTextString(m) }
func (*RemoveFilteredPolicyReq) ProtoMessage()    {}
func (*RemoveFilteredPolicyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_policy_edd0b7aa982d151a, []int{4}
}
func (m *RemoveFilteredPolicyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveFilteredPolicyReq.Unmarshal(m, b)
}
func (m *RemoveFilteredPolicyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveFilteredPolicyReq.Marshal(b, m, deterministic)
}
func (dst *RemoveFilteredPolicyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveFilteredPolicyReq.Merge(dst, src)
}
func (m *RemoveFilteredPolicyReq) XXX_Size() int {
	return xxx_messageInfo_RemoveFilteredPolicyReq.Size(m)
}
func (m *RemoveFilteredPolicyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveFilteredPolicyReq.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveFilteredPolicyReq proto.InternalMessageInfo

func (m *RemoveFilteredPolicyReq) GetSec() string {
	if m != nil {
		return m.Sec
	}
	return ""
}

func (m *RemoveFilteredPolicyReq) GetPtype() string {
	if m != nil {
		return m.Ptype
	}
	return ""
}

func (m *RemoveFilteredPolicyReq) GetFieldIndex() int64 {
	if m != nil {
		return m.FieldIndex
	}
	return 0
}

func (m *RemoveFilteredPolicyReq) GetFieldValues() []string {
	if m != nil {
		return m.FieldValues
	}
	return nil
}

func init() {
	proto.RegisterType((*Policy)(nil), "grpcadapter.Policy")
	proto.RegisterType((*Policies)(nil), "grpcadapter.Policies")
	proto.RegisterType((*AddPolicyReq)(nil), "grpcadapter.AddPolicyReq")
	proto.RegisterType((*RemovePolicyReq)(nil), "grpcadapter.RemovePolicyReq")
	proto.RegisterType((*RemoveFilteredPolicyReq)(nil), "grpcadapter.RemoveFilteredPolicyReq")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PolicyAdapterClient is the client API for PolicyAdapter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PolicyAdapterClient interface {
	LoadPolicy(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Policies, error)
	SavePolicy(ctx context.Context, in *Policies, opts ...grpc.CallOption) (*empty.Empty, error)
	AddPolicy(ctx context.Context, in *AddPolicyReq, opts ...grpc.CallOption) (*empty.Empty, error)
	RemovePolicy(ctx context.Context, in *RemovePolicyReq, opts ...grpc.CallOption) (*empty.Empty, error)
	RemoveFilteredPolicy(ctx context.Context, in *RemoveFilteredPolicyReq, opts ...grpc.CallOption) (*empty.Empty, error)
}

type policyAdapterClient struct {
	cc *grpc.ClientConn
}

func NewPolicyAdapterClient(cc *grpc.ClientConn) PolicyAdapterClient {
	return &policyAdapterClient{cc}
}

func (c *policyAdapterClient) LoadPolicy(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Policies, error) {
	out := new(Policies)
	err := c.cc.Invoke(ctx, "/grpcadapter.PolicyAdapter/loadPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyAdapterClient) SavePolicy(ctx context.Context, in *Policies, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpcadapter.PolicyAdapter/savePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyAdapterClient) AddPolicy(ctx context.Context, in *AddPolicyReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpcadapter.PolicyAdapter/addPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyAdapterClient) RemovePolicy(ctx context.Context, in *RemovePolicyReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpcadapter.PolicyAdapter/removePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyAdapterClient) RemoveFilteredPolicy(ctx context.Context, in *RemoveFilteredPolicyReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpcadapter.PolicyAdapter/removeFilteredPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyAdapterServer is the server API for PolicyAdapter service.
type PolicyAdapterServer interface {
	LoadPolicy(context.Context, *empty.Empty) (*Policies, error)
	SavePolicy(context.Context, *Policies) (*empty.Empty, error)
	AddPolicy(context.Context, *AddPolicyReq) (*empty.Empty, error)
	RemovePolicy(context.Context, *RemovePolicyReq) (*empty.Empty, error)
	RemoveFilteredPolicy(context.Context, *RemoveFilteredPolicyReq) (*empty.Empty, error)
}

func RegisterPolicyAdapterServer(s *grpc.Server, srv PolicyAdapterServer) {
	s.RegisterService(&_PolicyAdapter_serviceDesc, srv)
}

func _PolicyAdapter_LoadPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyAdapterServer).LoadPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcadapter.PolicyAdapter/LoadPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyAdapterServer).LoadPolicy(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyAdapter_SavePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Policies)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyAdapterServer).SavePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcadapter.PolicyAdapter/SavePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyAdapterServer).SavePolicy(ctx, req.(*Policies))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyAdapter_AddPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPolicyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyAdapterServer).AddPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcadapter.PolicyAdapter/AddPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyAdapterServer).AddPolicy(ctx, req.(*AddPolicyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyAdapter_RemovePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePolicyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyAdapterServer).RemovePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcadapter.PolicyAdapter/RemovePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyAdapterServer).RemovePolicy(ctx, req.(*RemovePolicyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyAdapter_RemoveFilteredPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFilteredPolicyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyAdapterServer).RemoveFilteredPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcadapter.PolicyAdapter/RemoveFilteredPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyAdapterServer).RemoveFilteredPolicy(ctx, req.(*RemoveFilteredPolicyReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PolicyAdapter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcadapter.PolicyAdapter",
	HandlerType: (*PolicyAdapterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "loadPolicy",
			Handler:    _PolicyAdapter_LoadPolicy_Handler,
		},
		{
			MethodName: "savePolicy",
			Handler:    _PolicyAdapter_SavePolicy_Handler,
		},
		{
			MethodName: "addPolicy",
			Handler:    _PolicyAdapter_AddPolicy_Handler,
		},
		{
			MethodName: "removePolicy",
			Handler:    _PolicyAdapter_RemovePolicy_Handler,
		},
		{
			MethodName: "removeFilteredPolicy",
			Handler:    _PolicyAdapter_RemoveFilteredPolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "policy.proto",
}

func init() { proto.RegisterFile("policy.proto", fileDescriptor_policy_edd0b7aa982d151a) }

var fileDescriptor_policy_edd0b7aa982d151a = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x9b, 0xa6, 0x6f, 0x69, 0xa7, 0x7d, 0x79, 0x5f, 0xc6, 0xaa, 0xb1, 0x88, 0x84, 0xe0,
	0x21, 0xa7, 0x04, 0xda, 0xa3, 0x08, 0xf6, 0x60, 0x41, 0x41, 0x90, 0x1c, 0x7a, 0x4f, 0x93, 0x69,
	0x0c, 0x6c, 0xbb, 0x6b, 0xfe, 0x88, 0x3d, 0xfb, 0x81, 0xfd, 0x0a, 0x92, 0xdd, 0x54, 0x53, 0x9b,
	0xf6, 0x20, 0x5e, 0x96, 0xd9, 0x61, 0xe6, 0x37, 0x0f, 0xcf, 0x03, 0x7d, 0xc1, 0x59, 0x1c, 0xac,
	0x1d, 0x91, 0xf0, 0x8c, 0x63, 0x2f, 0x4a, 0x44, 0xe0, 0x87, 0xbe, 0xc8, 0x28, 0x19, 0x8e, 0xa3,
	0x38, 0x7b, 0xca, 0xe7, 0x4e, 0xc0, 0x97, 0x6e, 0xc4, 0x99, 0xbf, 0x8a, 0x5c, 0x39, 0x35, 0xcf,
	0x17, 0xae, 0xc8, 0xd6, 0x82, 0x52, 0x97, 0x96, 0x22, 0x5b, 0xab, 0x57, 0x11, 0xac, 0x11, 0xb4,
	0x1f, 0x25, 0x11, 0x07, 0xf0, 0x47, 0x4e, 0x19, 0x9a, 0xa9, 0xd9, 0x5d, 0x4f, 0x7d, 0x10, 0xa1,
	0x95, 0xe4, 0x8c, 0x8c, 0xa6, 0xa9, 0xdb, 0x5d, 0x4f, 0xd6, 0xd6, 0x15, 0x74, 0xe4, 0x4e, 0x4c,
	0x29, 0xba, 0xd0, 0x11, 0x65, 0x6d, 0x68, 0xa6, 0x6e, 0xf7, 0x46, 0x47, 0x4e, 0x45, 0x94, 0xa3,
	0xe0, 0xde, 0xe7, 0x90, 0x75, 0x0f, 0xfd, 0x49, 0x18, 0x96, 0x6d, 0x7a, 0xc6, 0xff, 0xa0, 0xa7,
	0x14, 0x94, 0x47, 0x8b, 0xf2, 0x4b, 0x48, 0xb3, 0x4e, 0x88, 0x5e, 0x11, 0xf2, 0x00, 0xff, 0x3c,
	0x5a, 0xf2, 0x17, 0xfa, 0x1d, 0xdc, 0x9b, 0x06, 0xa7, 0x8a, 0x37, 0x8d, 0x59, 0x46, 0x09, 0xfd,
	0x40, 0xe6, 0x05, 0xc0, 0x22, 0x26, 0x16, 0xde, 0xad, 0x42, 0x7a, 0x35, 0x74, 0x53, 0xb3, 0x75,
	0xaf, 0xd2, 0x41, 0x13, 0x7a, 0xf2, 0x37, 0xf3, 0x59, 0x4e, 0xa9, 0xd1, 0x92, 0xe7, 0xab, 0xad,
	0xd1, 0x7b, 0x13, 0xfe, 0xaa, 0xbb, 0x13, 0xe5, 0x21, 0x5e, 0x03, 0x30, 0xee, 0x97, 0x62, 0xf0,
	0xc4, 0x89, 0x38, 0x8f, 0x18, 0x39, 0x9b, 0x70, 0x9d, 0xdb, 0x22, 0xcf, 0xe1, 0xf1, 0xae, 0xef,
	0x85, 0xdf, 0x8d, 0x62, 0x3d, 0xf5, 0x37, 0x1e, 0x61, 0xfd, 0xd8, 0x70, 0x0f, 0xd5, 0x6a, 0xe0,
	0x0d, 0x74, 0xfd, 0x4d, 0x60, 0x78, 0xb6, 0xb5, 0x5d, 0x0d, 0xf2, 0x00, 0x61, 0x0a, 0xfd, 0xa4,
	0x12, 0x13, 0x9e, 0x6f, 0x41, 0xbe, 0x25, 0x78, 0x80, 0x33, 0x83, 0x41, 0x52, 0x13, 0x0f, 0x5e,
	0xd6, 0xf0, 0x76, 0x12, 0xdc, 0xcf, 0x9d, 0xb7, 0x65, 0x67, 0xfc, 0x11, 0x00, 0x00, 0xff, 0xff,
	0x60, 0x14, 0x50, 0x09, 0x5c, 0x03, 0x00, 0x00,
}