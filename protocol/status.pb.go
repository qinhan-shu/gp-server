// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/status.proto

package protocol

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

type Code int32

const (
	Code_OK                Code = 0
	Code_INTERNAL          Code = 1
	Code_DATA_LOSE         Code = 2
	Code_NO_TOKEN          Code = 3
	Code_UNAUTHORIZATED    Code = 4
	Code_PERMISSION_DENIED Code = 5
)

var Code_name = map[int32]string{
	0: "OK",
	1: "INTERNAL",
	2: "DATA_LOSE",
	3: "NO_TOKEN",
	4: "UNAUTHORIZATED",
	5: "PERMISSION_DENIED",
}

var Code_value = map[string]int32{
	"OK":                0,
	"INTERNAL":          1,
	"DATA_LOSE":         2,
	"NO_TOKEN":          3,
	"UNAUTHORIZATED":    4,
	"PERMISSION_DENIED": 5,
}

func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}

func (Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_17c96f8603cffea9, []int{0}
}

type Status struct {
	Code                 Code     `protobuf:"varint,1,opt,name=code,proto3,enum=protocol.Code" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_17c96f8603cffea9, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() Code {
	if m != nil {
		return m.Code
	}
	return Code_OK
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("protocol.Code", Code_name, Code_value)
	proto.RegisterType((*Status)(nil), "protocol.Status")
}

func init() { proto.RegisterFile("proto/status.proto", fileDescriptor_17c96f8603cffea9) }

var fileDescriptor_17c96f8603cffea9 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0x87, 0x4d, 0xad, 0x75, 0x7b, 0x68, 0xc9, 0x9e, 0x08, 0x3d, 0x8e, 0x9d, 0x86, 0x87, 0x0a,
	0x7a, 0xf4, 0x14, 0x4d, 0xc4, 0xb0, 0x99, 0x48, 0x92, 0x5d, 0x76, 0x29, 0xb5, 0x0d, 0x5e, 0x94,
	0x88, 0xa9, 0xff, 0xff, 0x68, 0x20, 0xa7, 0xf7, 0xfb, 0xf8, 0xbe, 0xc3, 0x03, 0xfc, 0xfd, 0x0b,
	0x53, 0xb8, 0x8f, 0x53, 0x3f, 0xfd, 0xc7, 0x36, 0x01, 0x2e, 0xd2, 0x19, 0xc2, 0xf7, 0xe6, 0x15,
	0x2a, 0x9b, 0x0c, 0x6e, 0xa0, 0x1c, 0xc2, 0xe8, 0x1b, 0xb2, 0x26, 0xdb, 0xfa, 0xa1, 0x6e, 0x73,
	0xd2, 0xbe, 0x84, 0xd1, 0x9b, 0xe4, 0xb0, 0x81, 0xcb, 0x1f, 0x1f, 0x63, 0xff, 0xe5, 0x9b, 0x62,
	0x4d, 0xb6, 0x4b, 0x93, 0xf1, 0x6e, 0x84, 0x72, 0xee, 0xb0, 0x82, 0x42, 0xef, 0xe8, 0x19, 0x5e,
	0xc1, 0x42, 0x2a, 0x27, 0x8c, 0x62, 0x7b, 0x4a, 0xf0, 0x1a, 0x96, 0x9c, 0x39, 0xd6, 0xed, 0xb5,
	0x15, 0xb4, 0x98, 0xa5, 0xd2, 0x9d, 0xd3, 0x3b, 0xa1, 0xe8, 0x39, 0x22, 0xd4, 0x07, 0xc5, 0x0e,
	0xee, 0x4d, 0x1b, 0x79, 0x64, 0x4e, 0x70, 0x5a, 0xe2, 0x2d, 0xac, 0x3e, 0x84, 0x79, 0x97, 0xd6,
	0x4a, 0xad, 0x3a, 0x2e, 0x94, 0x14, 0x9c, 0x5e, 0x3c, 0xdf, 0x1c, 0x57, 0xf9, 0xad, 0xa7, 0x3c,
	0x3e, 0xab, 0xb4, 0x1e, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x32, 0x88, 0x3c, 0xe9, 0x00,
	0x00, 0x00,
}
