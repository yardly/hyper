// Code generated by protoc-gen-go. DO NOT EDIT.
// source: distribution.proto

/*
Package sync is a generated protocol buffer package.

It is generated from these files:
	distribution.proto
	packet.proto

It has these top-level messages:
	Distribution
	Condition
	Packet
*/
package sync

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Distribution struct {
	Packet    *Packet    `protobuf:"bytes,40,opt,name=Packet" json:"Packet,omitempty"`
	Condition *Condition `protobuf:"bytes,50,opt,name=Condition" json:"Condition,omitempty"`
}

func (m *Distribution) Reset()                    { *m = Distribution{} }
func (m *Distribution) String() string            { return proto.CompactTextString(m) }
func (*Distribution) ProtoMessage()               {}
func (*Distribution) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Distribution) GetPacket() *Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *Distribution) GetCondition() *Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

type Condition struct {
	EqIDs  []int64  `protobuf:"varint,10,rep,packed,name=EqIDs" json:"EqIDs,omitempty"`
	NeIDs  []int64  `protobuf:"varint,11,rep,packed,name=NeIDs" json:"NeIDs,omitempty"`
	EqKeys []string `protobuf:"bytes,20,rep,name=EqKeys" json:"EqKeys,omitempty"`
	NeKeys []string `protobuf:"bytes,21,rep,name=NeKeys" json:"NeKeys,omitempty"`
}

func (m *Condition) Reset()                    { *m = Condition{} }
func (m *Condition) String() string            { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()               {}
func (*Condition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Condition) GetEqIDs() []int64 {
	if m != nil {
		return m.EqIDs
	}
	return nil
}

func (m *Condition) GetNeIDs() []int64 {
	if m != nil {
		return m.NeIDs
	}
	return nil
}

func (m *Condition) GetEqKeys() []string {
	if m != nil {
		return m.EqKeys
	}
	return nil
}

func (m *Condition) GetNeKeys() []string {
	if m != nil {
		return m.NeKeys
	}
	return nil
}

func init() {
	proto.RegisterType((*Distribution)(nil), "sync.Distribution")
	proto.RegisterType((*Condition)(nil), "sync.Condition")
}

func init() { proto.RegisterFile("distribution.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xc9, 0x2c, 0x2e,
	0x29, 0xca, 0x4c, 0x2a, 0x2d, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x29, 0xae, 0xcc, 0x4b, 0x96, 0xe2, 0x29, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0x81, 0x88, 0x29, 0x25,
	0x73, 0xf1, 0xb8, 0x20, 0xa9, 0x14, 0x52, 0xe1, 0x62, 0x0b, 0x00, 0xcb, 0x4b, 0x68, 0x28, 0x30,
	0x6a, 0x70, 0x1b, 0xf1, 0xe8, 0x81, 0x34, 0xe9, 0x41, 0xc4, 0x82, 0xa0, 0x72, 0x42, 0xba, 0x5c,
	0x9c, 0xce, 0xf9, 0x79, 0x29, 0x99, 0x20, 0x2d, 0x12, 0x46, 0x60, 0x85, 0xfc, 0x10, 0x85, 0x70,
	0xe1, 0x20, 0x84, 0x0a, 0xa5, 0x74, 0x24, 0xe5, 0x42, 0x22, 0x5c, 0xac, 0xae, 0x85, 0x9e, 0x2e,
	0xc5, 0x12, 0x5c, 0x0a, 0xcc, 0x1a, 0xcc, 0x41, 0x10, 0x0e, 0x48, 0xd4, 0x2f, 0x15, 0x24, 0xca,
	0x0d, 0x11, 0x05, 0x73, 0x84, 0xc4, 0xb8, 0xd8, 0x5c, 0x0b, 0xbd, 0x53, 0x2b, 0x8b, 0x25, 0x44,
	0x14, 0x98, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x90, 0xb8, 0x5f, 0x2a, 0x58, 0x5c, 0x14, 0x22, 0x0e,
	0xe1, 0x25, 0xb1, 0x81, 0x3d, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x33, 0x4a, 0x4a, 0xd5,
	0xfe, 0x00, 0x00, 0x00,
}
