// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/test.proto

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

type TestRequest struct {
	A                    string   `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	B                    string   `protobuf:"bytes,2,opt,name=B,proto3" json:"B,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc70169f84046e97, []int{0}
}

func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRequest.Unmarshal(m, b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRequest.Marshal(b, m, deterministic)
}
func (m *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(m, src)
}
func (m *TestRequest) XXX_Size() int {
	return xxx_messageInfo_TestRequest.Size(m)
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

func (m *TestRequest) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *TestRequest) GetB() string {
	if m != nil {
		return m.B
	}
	return ""
}

type TestResponse struct {
	A                    string   `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	B                    string   `protobuf:"bytes,2,opt,name=B,proto3" json:"B,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestResponse) Reset()         { *m = TestResponse{} }
func (m *TestResponse) String() string { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()    {}
func (*TestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc70169f84046e97, []int{1}
}

func (m *TestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResponse.Unmarshal(m, b)
}
func (m *TestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResponse.Marshal(b, m, deterministic)
}
func (m *TestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResponse.Merge(m, src)
}
func (m *TestResponse) XXX_Size() int {
	return xxx_messageInfo_TestResponse.Size(m)
}
func (m *TestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestResponse proto.InternalMessageInfo

func (m *TestResponse) GetA() string {
	if m != nil {
		return m.A
	}
	return ""
}

func (m *TestResponse) GetB() string {
	if m != nil {
		return m.B
	}
	return ""
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "protocol.TestRequest")
	proto.RegisterType((*TestResponse)(nil), "protocol.TestResponse")
}

func init() { proto.RegisterFile("proto/test.proto", fileDescriptor_fc70169f84046e97) }

var fileDescriptor_fc70169f84046e97 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x03, 0x33, 0x85, 0x38, 0xc0, 0x54, 0x72, 0x7e, 0x8e,
	0x92, 0x26, 0x17, 0x77, 0x48, 0x6a, 0x71, 0x49, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x10,
	0x0f, 0x17, 0xa3, 0xa3, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0xa3, 0x23, 0x88, 0xe7, 0x24,
	0xc1, 0x04, 0xe1, 0x39, 0x29, 0x69, 0x71, 0xf1, 0x40, 0x94, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0xe2, 0x53, 0xeb, 0xc4, 0x1b, 0xc5, 0x9d, 0x9e, 0x6f, 0x0d, 0xb3, 0x25, 0x89, 0x0d, 0xcc, 0x32,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x22, 0x6e, 0x38, 0x8a, 0x00, 0x00, 0x00,
}
