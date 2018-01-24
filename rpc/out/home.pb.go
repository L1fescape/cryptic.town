// Code generated by protoc-gen-go. DO NOT EDIT.
// source: home.proto

/*
Package home is a generated protocol buffer package.

It is generated from these files:
	home.proto

It has these top-level messages:
	SetHomeRequest
	Home
*/
package home

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

type SetHomeRequest struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Body  string `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
}

func (m *SetHomeRequest) Reset()                    { *m = SetHomeRequest{} }
func (m *SetHomeRequest) String() string            { return proto.CompactTextString(m) }
func (*SetHomeRequest) ProtoMessage()               {}
func (*SetHomeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SetHomeRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SetHomeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SetHomeRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type Home struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Body string `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *Home) Reset()                    { *m = Home{} }
func (m *Home) String() string            { return proto.CompactTextString(m) }
func (*Home) ProtoMessage()               {}
func (*Home) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Home) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Home) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*SetHomeRequest)(nil), "SetHomeRequest")
	proto.RegisterType((*Home)(nil), "Home")
}

func init() { proto.RegisterFile("home.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xc8, 0xcf, 0x4d,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xf2, 0xe3, 0xe2, 0x0b, 0x4e, 0x2d, 0xf1, 0xc8, 0xcf,
	0x4d, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0xc9, 0xcf, 0x4e,
	0xcd, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12,
	0x73, 0x53, 0x25, 0x98, 0xc0, 0x82, 0x60, 0x36, 0x48, 0x2c, 0x29, 0x3f, 0xa5, 0x52, 0x82, 0x19,
	0x22, 0x06, 0x62, 0x2b, 0xe9, 0x71, 0xb1, 0x80, 0x0c, 0x83, 0xab, 0x67, 0xc4, 0xa2, 0x9e, 0x09,
	0xa1, 0xde, 0xc8, 0x80, 0x8b, 0xdb, 0xb9, 0xa8, 0xb2, 0xa0, 0x24, 0x33, 0x39, 0x24, 0xbf, 0x3c,
	0x4f, 0x48, 0x91, 0x8b, 0x1d, 0xea, 0x1c, 0x21, 0x7e, 0x3d, 0x54, 0x87, 0x49, 0xb1, 0xea, 0x81,
	0x78, 0x49, 0x6c, 0x60, 0x87, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x05, 0x3b, 0x83, 0xf3,
	0xc6, 0x00, 0x00, 0x00,
}
