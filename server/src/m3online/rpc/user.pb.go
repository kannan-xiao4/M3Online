// Code generated by protoc-gen-go. DO NOT EDIT.
// source: domain/user.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type UserId struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserId) Reset()         { *m = UserId{} }
func (m *UserId) String() string { return proto.CompactTextString(m) }
func (*UserId) ProtoMessage()    {}
func (*UserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bf984dc6decf814, []int{0}
}

func (m *UserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserId.Unmarshal(m, b)
}
func (m *UserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserId.Marshal(b, m, deterministic)
}
func (m *UserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserId.Merge(m, src)
}
func (m *UserId) XXX_Size() int {
	return xxx_messageInfo_UserId.Size(m)
}
func (m *UserId) XXX_DiscardUnknown() {
	xxx_messageInfo_UserId.DiscardUnknown(m)
}

var xxx_messageInfo_UserId proto.InternalMessageInfo

func (m *UserId) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type User struct {
	Id                   *UserId  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bf984dc6decf814, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() *UserId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *User) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterType((*UserId)(nil), "UserId")
	proto.RegisterType((*User)(nil), "User")
}

func init() { proto.RegisterFile("domain/user.proto", fileDescriptor_8bf984dc6decf814) }

var fileDescriptor_8bf984dc6decf814 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xc9, 0xcf, 0x4d,
	0xcc, 0xcc, 0xd3, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe4,
	0x62, 0x0b, 0x2d, 0x4e, 0x2d, 0xf2, 0x4c, 0x11, 0x12, 0xe7, 0x62, 0x07, 0x89, 0xc7, 0x67, 0xa6,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xb1, 0x95, 0x82, 0x25, 0x94, 0x6c, 0xb8, 0x58, 0x40,
	0x4a, 0x84, 0xc4, 0xb9, 0x98, 0xa0, 0x72, 0xdc, 0x46, 0xec, 0x7a, 0x10, 0x5d, 0x41, 0x4c, 0x99,
	0x29, 0x42, 0xd2, 0x5c, 0x9c, 0x60, 0x9d, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x60, 0xbd, 0x1c,
	0x20, 0x01, 0xbf, 0xc4, 0xdc, 0x54, 0x27, 0xd5, 0x28, 0x9e, 0x5c, 0xe3, 0xfc, 0xbc, 0x9c, 0xcc,
	0xbc, 0x54, 0xfd, 0xa2, 0x82, 0xe4, 0x55, 0x4c, 0x22, 0xbe, 0xc6, 0xfe, 0x60, 0xae, 0x9e, 0x0b,
	0xd8, 0x35, 0x60, 0x93, 0x92, 0xd8, 0xc0, 0xce, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x08,
	0x00, 0x0f, 0x67, 0xa3, 0x00, 0x00, 0x00,
}
