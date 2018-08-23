// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/uaa.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type OAuthType int32

const (
	OAuthType_GOOGLE   OAuthType = 0
	OAuthType_WECHAT   OAuthType = 1
	OAuthType_QQ       OAuthType = 2
	OAuthType_SINA     OAuthType = 3
	OAuthType_FACEBOOK OAuthType = 4
	OAuthType_TWITTER  OAuthType = 5
)

var OAuthType_name = map[int32]string{
	0: "GOOGLE",
	1: "WECHAT",
	2: "QQ",
	3: "SINA",
	4: "FACEBOOK",
	5: "TWITTER",
}
var OAuthType_value = map[string]int32{
	"GOOGLE":   0,
	"WECHAT":   1,
	"QQ":       2,
	"SINA":     3,
	"FACEBOOK": 4,
	"TWITTER":  5,
}

func (x OAuthType) String() string {
	return proto.EnumName(OAuthType_name, int32(x))
}
func (OAuthType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_uaa_b1b23e9c9eb173e5, []int{0}
}

type Account struct {
	Uid                  string               `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             []byte               `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	AccountExpired       bool                 `protobuf:"varint,5,opt,name=accountExpired,proto3" json:"accountExpired,omitempty"`
	AccountLocked        bool                 `protobuf:"varint,6,opt,name=accountLocked,proto3" json:"accountLocked,omitempty"`
	CredentialsExpired   bool                 `protobuf:"varint,7,opt,name=credentialsExpired,proto3" json:"credentialsExpired,omitempty"`
	Roles                []string             `protobuf:"bytes,8,rep,name=roles,proto3" json:"roles,omitempty"`
	OauthUserIds         map[uint32]string    `protobuf:"bytes,9,rep,name=oauthUserIds,proto3" json:"oauthUserIds,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreateDate           *timestamp.Timestamp `protobuf:"bytes,10,opt,name=createDate,proto3" json:"createDate,omitempty"`
	UpdateDate           *timestamp.Timestamp `protobuf:"bytes,11,opt,name=updateDate,proto3" json:"updateDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_uaa_b1b23e9c9eb173e5, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (dst *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(dst, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Account) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Account) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Account) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *Account) GetAccountExpired() bool {
	if m != nil {
		return m.AccountExpired
	}
	return false
}

func (m *Account) GetAccountLocked() bool {
	if m != nil {
		return m.AccountLocked
	}
	return false
}

func (m *Account) GetCredentialsExpired() bool {
	if m != nil {
		return m.CredentialsExpired
	}
	return false
}

func (m *Account) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Account) GetOauthUserIds() map[uint32]string {
	if m != nil {
		return m.OauthUserIds
	}
	return nil
}

func (m *Account) GetCreateDate() *timestamp.Timestamp {
	if m != nil {
		return m.CreateDate
	}
	return nil
}

func (m *Account) GetUpdateDate() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateDate
	}
	return nil
}

type GetAllResp struct {
	Accounts             []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetAllResp) Reset()         { *m = GetAllResp{} }
func (m *GetAllResp) String() string { return proto.CompactTextString(m) }
func (*GetAllResp) ProtoMessage()    {}
func (*GetAllResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_uaa_b1b23e9c9eb173e5, []int{1}
}
func (m *GetAllResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllResp.Unmarshal(m, b)
}
func (m *GetAllResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllResp.Marshal(b, m, deterministic)
}
func (dst *GetAllResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllResp.Merge(dst, src)
}
func (m *GetAllResp) XXX_Size() int {
	return xxx_messageInfo_GetAllResp.Size(m)
}
func (m *GetAllResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllResp proto.InternalMessageInfo

func (m *GetAllResp) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type GetByUsernameReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByUsernameReq) Reset()         { *m = GetByUsernameReq{} }
func (m *GetByUsernameReq) String() string { return proto.CompactTextString(m) }
func (*GetByUsernameReq) ProtoMessage()    {}
func (*GetByUsernameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_uaa_b1b23e9c9eb173e5, []int{2}
}
func (m *GetByUsernameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByUsernameReq.Unmarshal(m, b)
}
func (m *GetByUsernameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByUsernameReq.Marshal(b, m, deterministic)
}
func (dst *GetByUsernameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByUsernameReq.Merge(dst, src)
}
func (m *GetByUsernameReq) XXX_Size() int {
	return xxx_messageInfo_GetByUsernameReq.Size(m)
}
func (m *GetByUsernameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByUsernameReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetByUsernameReq proto.InternalMessageInfo

func (m *GetByUsernameReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type DeleteByUsernameReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteByUsernameReq) Reset()         { *m = DeleteByUsernameReq{} }
func (m *DeleteByUsernameReq) String() string { return proto.CompactTextString(m) }
func (*DeleteByUsernameReq) ProtoMessage()    {}
func (*DeleteByUsernameReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_uaa_b1b23e9c9eb173e5, []int{3}
}
func (m *DeleteByUsernameReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteByUsernameReq.Unmarshal(m, b)
}
func (m *DeleteByUsernameReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteByUsernameReq.Marshal(b, m, deterministic)
}
func (dst *DeleteByUsernameReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteByUsernameReq.Merge(dst, src)
}
func (m *DeleteByUsernameReq) XXX_Size() int {
	return xxx_messageInfo_DeleteByUsernameReq.Size(m)
}
func (m *DeleteByUsernameReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteByUsernameReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteByUsernameReq proto.InternalMessageInfo

func (m *DeleteByUsernameReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type RegisterReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Roles                []string `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_uaa_b1b23e9c9eb173e5, []int{4}
}
func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (dst *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(dst, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterReq) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

type VerifyPasswordReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyPasswordReq) Reset()         { *m = VerifyPasswordReq{} }
func (m *VerifyPasswordReq) String() string { return proto.CompactTextString(m) }
func (*VerifyPasswordReq) ProtoMessage()    {}
func (*VerifyPasswordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_uaa_b1b23e9c9eb173e5, []int{5}
}
func (m *VerifyPasswordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyPasswordReq.Unmarshal(m, b)
}
func (m *VerifyPasswordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyPasswordReq.Marshal(b, m, deterministic)
}
func (dst *VerifyPasswordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyPasswordReq.Merge(dst, src)
}
func (m *VerifyPasswordReq) XXX_Size() int {
	return xxx_messageInfo_VerifyPasswordReq.Size(m)
}
func (m *VerifyPasswordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyPasswordReq.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyPasswordReq proto.InternalMessageInfo

func (m *VerifyPasswordReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *VerifyPasswordReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ChangePasswordReq struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	OldPassword          string   `protobuf:"bytes,2,opt,name=oldPassword,proto3" json:"oldPassword,omitempty"`
	NewPassword          string   `protobuf:"bytes,3,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangePasswordReq) Reset()         { *m = ChangePasswordReq{} }
func (m *ChangePasswordReq) String() string { return proto.CompactTextString(m) }
func (*ChangePasswordReq) ProtoMessage()    {}
func (*ChangePasswordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_uaa_b1b23e9c9eb173e5, []int{6}
}
func (m *ChangePasswordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangePasswordReq.Unmarshal(m, b)
}
func (m *ChangePasswordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangePasswordReq.Marshal(b, m, deterministic)
}
func (dst *ChangePasswordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangePasswordReq.Merge(dst, src)
}
func (m *ChangePasswordReq) XXX_Size() int {
	return xxx_messageInfo_ChangePasswordReq.Size(m)
}
func (m *ChangePasswordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangePasswordReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChangePasswordReq proto.InternalMessageInfo

func (m *ChangePasswordReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ChangePasswordReq) GetOldPassword() string {
	if m != nil {
		return m.OldPassword
	}
	return ""
}

func (m *ChangePasswordReq) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "com.teddy.srv.uaa.Account")
	proto.RegisterMapType((map[uint32]string)(nil), "com.teddy.srv.uaa.Account.OauthUserIdsEntry")
	proto.RegisterType((*GetAllResp)(nil), "com.teddy.srv.uaa.GetAllResp")
	proto.RegisterType((*GetByUsernameReq)(nil), "com.teddy.srv.uaa.GetByUsernameReq")
	proto.RegisterType((*DeleteByUsernameReq)(nil), "com.teddy.srv.uaa.DeleteByUsernameReq")
	proto.RegisterType((*RegisterReq)(nil), "com.teddy.srv.uaa.RegisterReq")
	proto.RegisterType((*VerifyPasswordReq)(nil), "com.teddy.srv.uaa.VerifyPasswordReq")
	proto.RegisterType((*ChangePasswordReq)(nil), "com.teddy.srv.uaa.ChangePasswordReq")
	proto.RegisterEnum("com.teddy.srv.uaa.OAuthType", OAuthType_name, OAuthType_value)
}

func init() { proto.RegisterFile("proto/uaa.proto", fileDescriptor_uaa_b1b23e9c9eb173e5) }

var fileDescriptor_uaa_b1b23e9c9eb173e5 = []byte{
	// 689 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x94, 0xdd, 0x4e, 0xdb, 0x4c,
	0x10, 0x86, 0x49, 0x9c, 0x1f, 0x33, 0x01, 0x3e, 0x67, 0xbf, 0xaa, 0xb2, 0x2c, 0xb5, 0x42, 0x2e,
	0x42, 0x55, 0xd5, 0x3a, 0x2a, 0x48, 0x55, 0x45, 0x0f, 0x2a, 0x03, 0x2e, 0x20, 0x50, 0x13, 0x5c,
	0x53, 0xa4, 0xf6, 0xc8, 0xc4, 0x8b, 0xb1, 0x70, 0xb2, 0xae, 0x77, 0x0d, 0xcd, 0x1d, 0xf4, 0x42,
	0x7b, 0x21, 0xdd, 0xb5, 0x9d, 0x1f, 0x13, 0xa7, 0x41, 0xea, 0x01, 0x64, 0x76, 0xf6, 0x9d, 0xd7,
	0xb3, 0x3b, 0x8f, 0x0d, 0xff, 0x45, 0x31, 0x61, 0xa4, 0x93, 0xb8, 0xae, 0x91, 0x46, 0xa8, 0xdd,
	0x27, 0x03, 0x83, 0x61, 0xcf, 0x1b, 0x19, 0x34, 0xbe, 0x33, 0xf8, 0x86, 0xb6, 0xeb, 0x07, 0xec,
	0x26, 0xb9, 0x32, 0xf8, 0x4e, 0xc7, 0x27, 0xa1, 0x3b, 0xf4, 0x3b, 0xa9, 0xf6, 0x2a, 0xb9, 0xee,
	0x44, 0x6c, 0x14, 0x61, 0xda, 0xc1, 0x03, 0x1e, 0x64, 0xff, 0x33, 0x1f, 0xed, 0xc3, 0xf2, 0x22,
	0x16, 0x0c, 0x30, 0x65, 0xee, 0x20, 0x9a, 0x46, 0x79, 0xf1, 0x9b, 0x42, 0xb1, 0x4f, 0xa6, 0xa5,
	0x62, 0x95, 0xb5, 0x2c, 0xa2, 0x4c, 0xae, 0xff, 0xaa, 0x41, 0xd3, 0xec, 0xf7, 0x49, 0x32, 0x64,
	0x48, 0x01, 0x29, 0x09, 0x3c, 0xb5, 0xb2, 0x59, 0x79, 0xb9, 0x6a, 0x8b, 0x10, 0x69, 0x20, 0x27,
	0x14, 0xc7, 0x43, 0x77, 0x80, 0xd5, 0x6a, 0x9a, 0x9e, 0xac, 0xd1, 0x13, 0xa8, 0xe3, 0x81, 0x1b,
	0x84, 0xaa, 0x94, 0x6e, 0x64, 0x0b, 0x51, 0x11, 0xb9, 0x94, 0xde, 0x93, 0xd8, 0x53, 0x6b, 0x7c,
	0x63, 0xcd, 0x9e, 0xac, 0xd1, 0x36, 0x6c, 0xb8, 0xd9, 0xa3, 0xac, 0x9f, 0x51, 0x10, 0x63, 0x4f,
	0xad, 0x73, 0x85, 0x6c, 0x3f, 0xc8, 0xa2, 0x2d, 0x58, 0xcf, 0x33, 0x67, 0xa4, 0x7f, 0xcb, 0x65,
	0x8d, 0x54, 0x56, 0x4c, 0x22, 0x03, 0x50, 0x9f, 0xab, 0xf1, 0x90, 0x05, 0x6e, 0x48, 0xc7, 0x8e,
	0xcd, 0x54, 0x5a, 0xb2, 0x23, 0xfa, 0x8d, 0x49, 0x88, 0xa9, 0x2a, 0x6f, 0x4a, 0xa2, 0xdf, 0x74,
	0x81, 0x7a, 0xb0, 0x46, 0xdc, 0x84, 0xdd, 0x5c, 0xf0, 0x63, 0x9d, 0x78, 0x54, 0x5d, 0xe5, 0x9b,
	0xad, 0x9d, 0xd7, 0xc6, 0xdc, 0x28, 0x8d, 0xfc, 0x96, 0x8c, 0xee, 0x8c, 0xdc, 0x1a, 0xb2, 0x78,
	0x64, 0x17, 0x1c, 0xd0, 0x1e, 0x00, 0x7f, 0xba, 0xcb, 0xf0, 0x21, 0xff, 0x53, 0x81, 0xf7, 0xd3,
	0xda, 0xd1, 0x0c, 0x9f, 0x10, 0x3f, 0xc4, 0xc6, 0x78, 0x18, 0x86, 0x33, 0x1e, 0x9b, 0x3d, 0xa3,
	0x16, 0xb5, 0x49, 0xe4, 0x8d, 0x6b, 0x5b, 0xcb, 0x6b, 0xa7, 0x6a, 0xed, 0x23, 0xb4, 0xe7, 0x5a,
	0x13, 0x23, 0xbd, 0xc5, 0xa3, 0x74, 0xa4, 0xeb, 0xb6, 0x08, 0xc5, 0x35, 0xdc, 0xb9, 0x61, 0x32,
	0x9e, 0x67, 0xb6, 0xd8, 0xab, 0xbe, 0xaf, 0xe8, 0x87, 0x00, 0x47, 0x98, 0x99, 0x61, 0x68, 0x63,
	0x1a, 0xa1, 0x77, 0x20, 0xe7, 0xf7, 0x4d, 0x79, 0xb9, 0x94, 0x36, 0xb2, 0xf0, 0x52, 0xec, 0x89,
	0x56, 0x37, 0x40, 0xe1, 0x2e, 0xfb, 0xa3, 0x8b, 0x9c, 0x13, 0x1b, 0xff, 0x28, 0x60, 0x54, 0x29,
	0x62, 0xa4, 0xbf, 0x85, 0xff, 0x0f, 0x71, 0x88, 0x19, 0x7e, 0x7c, 0xc9, 0x77, 0x68, 0xd9, 0xd8,
	0x0f, 0x28, 0xc3, 0xf1, 0x12, 0x69, 0x01, 0xc7, 0x1c, 0xe0, 0x09, 0x8e, 0x13, 0x20, 0xa4, 0x19,
	0x20, 0xf4, 0x53, 0x68, 0x7f, 0xc5, 0x71, 0x70, 0x3d, 0xea, 0xe5, 0xba, 0x7f, 0x78, 0x84, 0x4e,
	0xa1, 0x7d, 0x70, 0xc3, 0xdf, 0x5f, 0xfc, 0x58, 0xb3, 0x4d, 0x68, 0x91, 0xd0, 0xeb, 0x15, 0xfd,
	0x66, 0x53, 0x42, 0x31, 0xc4, 0xf7, 0x13, 0x45, 0xf6, 0xf2, 0xcd, 0xa6, 0x5e, 0xf5, 0x60, 0xb5,
	0x6b, 0x72, 0x10, 0x1c, 0xfe, 0xa5, 0x40, 0x00, 0x8d, 0xa3, 0x6e, 0xf7, 0xe8, 0xcc, 0x52, 0x56,
	0x44, 0x7c, 0x69, 0x1d, 0x1c, 0x9b, 0x8e, 0x52, 0x41, 0x0d, 0xa8, 0x9e, 0x9f, 0x2b, 0x55, 0x24,
	0x43, 0xed, 0xcb, 0xc9, 0x67, 0x53, 0x91, 0xd0, 0x1a, 0xc8, 0x9f, 0xcc, 0x03, 0x6b, 0xbf, 0xdb,
	0x3d, 0x55, 0x6a, 0xa8, 0x05, 0x4d, 0xe7, 0xf2, 0xc4, 0x71, 0x2c, 0x5b, 0xa9, 0xef, 0xfc, 0x96,
	0x40, 0xba, 0x30, 0x4d, 0x64, 0x72, 0xb3, 0x94, 0x10, 0xf4, 0x74, 0x0e, 0x4a, 0x4b, 0x7c, 0xc0,
	0xb4, 0x67, 0x25, 0x8c, 0x4c, 0xa1, 0xd2, 0x57, 0x90, 0x0d, 0xeb, 0x05, 0x3c, 0xd0, 0x8b, 0xf2,
	0x8a, 0x02, 0x0d, 0xda, 0x5f, 0xd0, 0xe3, 0x9e, 0x0e, 0x28, 0x0f, 0x11, 0x42, 0xdb, 0x25, 0x15,
	0x25, 0x9c, 0x69, 0x0b, 0x0e, 0xc2, 0x5d, 0x8f, 0x41, 0x1e, 0x53, 0x86, 0x9e, 0x97, 0xb8, 0xcd,
	0x20, 0xb8, 0xb4, 0xbf, 0x8d, 0x22, 0x52, 0x68, 0xab, 0x44, 0x3f, 0x47, 0xdd, 0x12, 0xd7, 0x1e,
	0x6c, 0x14, 0xd9, 0x2a, 0x75, 0x9d, 0xc3, 0x6f, 0xf1, 0x89, 0xf7, 0x9b, 0xdf, 0xea, 0x59, 0xae,
	0x91, 0xfe, 0xec, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x86, 0xe6, 0xad, 0x2b, 0xe2, 0x06, 0x00,
	0x00,
}