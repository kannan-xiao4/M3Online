// Code generated by protoc-gen-go. DO NOT EDIT.
// source: domain/enemy.proto

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

type Enemy struct {
	InstantId            uint32   `protobuf:"varint,1,opt,name=instant_id,json=instantId,proto3" json:"instant_id,omitempty"`
	MasterId             uint32   `protobuf:"varint,2,opt,name=master_id,json=masterId,proto3" json:"master_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Hp                   int32    `protobuf:"varint,4,opt,name=hp,proto3" json:"hp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Enemy) Reset()         { *m = Enemy{} }
func (m *Enemy) String() string { return proto.CompactTextString(m) }
func (*Enemy) ProtoMessage()    {}
func (*Enemy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dc6b9c0185ae71e, []int{0}
}

func (m *Enemy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Enemy.Unmarshal(m, b)
}
func (m *Enemy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Enemy.Marshal(b, m, deterministic)
}
func (m *Enemy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Enemy.Merge(m, src)
}
func (m *Enemy) XXX_Size() int {
	return xxx_messageInfo_Enemy.Size(m)
}
func (m *Enemy) XXX_DiscardUnknown() {
	xxx_messageInfo_Enemy.DiscardUnknown(m)
}

var xxx_messageInfo_Enemy proto.InternalMessageInfo

func (m *Enemy) GetInstantId() uint32 {
	if m != nil {
		return m.InstantId
	}
	return 0
}

func (m *Enemy) GetMasterId() uint32 {
	if m != nil {
		return m.MasterId
	}
	return 0
}

func (m *Enemy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Enemy) GetHp() int32 {
	if m != nil {
		return m.Hp
	}
	return 0
}

func init() {
	proto.RegisterType((*Enemy)(nil), "Enemy")
}

func init() { proto.RegisterFile("domain/enemy.proto", fileDescriptor_0dc6b9c0185ae71e) }

var fileDescriptor_0dc6b9c0185ae71e = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xc9, 0xcf, 0x4d,
	0xcc, 0xcc, 0xd3, 0x4f, 0xcd, 0x4b, 0xcd, 0xad, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x4a,
	0xe7, 0x62, 0x75, 0x05, 0x71, 0x85, 0x64, 0xb9, 0xb8, 0x32, 0xf3, 0x8a, 0x4b, 0x12, 0xf3, 0x4a,
	0xe2, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x83, 0x38, 0xa1, 0x22, 0x9e, 0x29, 0x42,
	0xd2, 0x5c, 0x9c, 0xb9, 0x89, 0xc5, 0x25, 0xa9, 0x45, 0x20, 0x59, 0x26, 0xb0, 0x2c, 0x07, 0x44,
	0xc0, 0x33, 0x45, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x59, 0x81, 0x51, 0x83,
	0x33, 0x08, 0xcc, 0x16, 0xe2, 0xe3, 0x62, 0xca, 0x28, 0x90, 0x60, 0x51, 0x60, 0xd4, 0x60, 0x0d,
	0x62, 0xca, 0x28, 0x70, 0x52, 0x8b, 0xe2, 0xc9, 0x35, 0xce, 0xcf, 0xcb, 0xc9, 0xcc, 0x4b, 0xd5,
	0x2f, 0x2a, 0x48, 0x5e, 0xc5, 0x24, 0xea, 0x6b, 0xec, 0x0f, 0xe6, 0xea, 0xb9, 0x80, 0x9d, 0xa5,
	0x07, 0x76, 0x47, 0x12, 0x1b, 0xd8, 0x5d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0x56,
	0xc8, 0x6c, 0xad, 0x00, 0x00, 0x00,
}
