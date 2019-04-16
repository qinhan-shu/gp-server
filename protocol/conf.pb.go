// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/conf.proto

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

type Config struct {
	Status               *Status          `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Tags                 map[int64]string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Difficulty           map[int64]string `protobuf:"bytes,3,rep,name=difficulty,proto3" json:"difficulty,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_69971211ee56777f, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Config) GetTags() map[int64]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Config) GetDifficulty() map[int64]string {
	if m != nil {
		return m.Difficulty
	}
	return nil
}

type UserRole struct {
	Status               *Status          `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Role                 map[int64]string `protobuf:"bytes,2,rep,name=role,proto3" json:"role,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UserRole) Reset()         { *m = UserRole{} }
func (m *UserRole) String() string { return proto.CompactTextString(m) }
func (*UserRole) ProtoMessage()    {}
func (*UserRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_69971211ee56777f, []int{1}
}

func (m *UserRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRole.Unmarshal(m, b)
}
func (m *UserRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRole.Marshal(b, m, deterministic)
}
func (m *UserRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRole.Merge(m, src)
}
func (m *UserRole) XXX_Size() int {
	return xxx_messageInfo_UserRole.Size(m)
}
func (m *UserRole) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRole.DiscardUnknown(m)
}

var xxx_messageInfo_UserRole proto.InternalMessageInfo

func (m *UserRole) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *UserRole) GetRole() map[int64]string {
	if m != nil {
		return m.Role
	}
	return nil
}

type JudgeLanguage struct {
	Status               *Status          `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Language             map[int64]string `protobuf:"bytes,2,rep,name=language,proto3" json:"language,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *JudgeLanguage) Reset()         { *m = JudgeLanguage{} }
func (m *JudgeLanguage) String() string { return proto.CompactTextString(m) }
func (*JudgeLanguage) ProtoMessage()    {}
func (*JudgeLanguage) Descriptor() ([]byte, []int) {
	return fileDescriptor_69971211ee56777f, []int{2}
}

func (m *JudgeLanguage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JudgeLanguage.Unmarshal(m, b)
}
func (m *JudgeLanguage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JudgeLanguage.Marshal(b, m, deterministic)
}
func (m *JudgeLanguage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JudgeLanguage.Merge(m, src)
}
func (m *JudgeLanguage) XXX_Size() int {
	return xxx_messageInfo_JudgeLanguage.Size(m)
}
func (m *JudgeLanguage) XXX_DiscardUnknown() {
	xxx_messageInfo_JudgeLanguage.DiscardUnknown(m)
}

var xxx_messageInfo_JudgeLanguage proto.InternalMessageInfo

func (m *JudgeLanguage) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *JudgeLanguage) GetLanguage() map[int64]string {
	if m != nil {
		return m.Language
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "protocol.Config")
	proto.RegisterMapType((map[int64]string)(nil), "protocol.Config.DifficultyEntry")
	proto.RegisterMapType((map[int64]string)(nil), "protocol.Config.TagsEntry")
	proto.RegisterType((*UserRole)(nil), "protocol.UserRole")
	proto.RegisterMapType((map[int64]string)(nil), "protocol.UserRole.RoleEntry")
	proto.RegisterType((*JudgeLanguage)(nil), "protocol.JudgeLanguage")
	proto.RegisterMapType((map[int64]string)(nil), "protocol.JudgeLanguage.LanguageEntry")
}

func init() { proto.RegisterFile("proto/conf.proto", fileDescriptor_69971211ee56777f) }

var fileDescriptor_69971211ee56777f = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x50, 0xcb, 0x4a, 0xc4, 0x30,
	0x14, 0x25, 0xed, 0x58, 0x3a, 0x77, 0x18, 0xac, 0xd1, 0x45, 0x29, 0x2e, 0xca, 0x80, 0xd0, 0x55,
	0x94, 0x71, 0xa1, 0x38, 0x08, 0x3e, 0x37, 0xe2, 0xaa, 0xea, 0xc6, 0x5d, 0xec, 0xa4, 0xa1, 0x18,
	0x1a, 0x69, 0x53, 0xa1, 0xff, 0x22, 0xf8, 0x1d, 0xfe, 0x9d, 0x34, 0xe9, 0xc3, 0x71, 0x56, 0xdd,
	0xb4, 0x27, 0xf7, 0x3c, 0x72, 0x72, 0xc1, 0xfb, 0x28, 0xa4, 0x92, 0xc7, 0x89, 0xcc, 0x53, 0xa2,
	0x21, 0x76, 0xf5, 0x2f, 0x91, 0x22, 0xc0, 0x86, 0x2b, 0x15, 0x55, 0x55, 0x69, 0xd8, 0xc5, 0x97,
	0x05, 0xce, 0xad, 0xcc, 0xd3, 0x8c, 0xe3, 0x08, 0x1c, 0x43, 0xf9, 0x28, 0x44, 0xd1, 0x6c, 0xe9,
	0x91, 0xce, 0x49, 0x9e, 0xf4, 0x3c, 0x6e, 0x79, 0x4c, 0x60, 0xa2, 0x28, 0x2f, 0x7d, 0x2b, 0xb4,
	0xa3, 0xd9, 0x32, 0x18, 0x74, 0x26, 0x89, 0x3c, 0x53, 0x5e, 0xde, 0xe7, 0xaa, 0xa8, 0x63, 0xad,
	0xc3, 0x57, 0x00, 0xeb, 0x2c, 0x4d, 0xb3, 0xa4, 0x12, 0xaa, 0xf6, 0x6d, 0xed, 0x0a, 0xb7, 0x5c,
	0x77, 0xbd, 0xc4, 0x78, 0xff, 0x78, 0x82, 0x33, 0x98, 0xf6, 0xa1, 0xd8, 0x03, 0xfb, 0x9d, 0xd5,
	0xba, 0xa5, 0x1d, 0x37, 0x10, 0x1f, 0xc0, 0xce, 0x27, 0x15, 0x15, 0xf3, 0xad, 0x10, 0x45, 0xd3,
	0xd8, 0x1c, 0x2e, 0xac, 0x73, 0x14, 0x5c, 0xc2, 0xee, 0xbf, 0xdc, 0x31, 0xf6, 0xc5, 0x37, 0x02,
	0xf7, 0xa5, 0x64, 0x45, 0x2c, 0x05, 0x1b, 0xb1, 0xa0, 0x13, 0x98, 0x14, 0x52, 0xb0, 0x76, 0x41,
	0x87, 0x83, 0xae, 0xcb, 0x22, 0xcd, 0xa7, 0x5d, 0x51, 0xa3, 0x6c, 0x1e, 0xd8, 0x8f, 0x46, 0x35,
	0xfc, 0x41, 0x30, 0x7f, 0xa8, 0xd6, 0x9c, 0x3d, 0xd2, 0x9c, 0x57, 0x94, 0x8f, 0xa9, 0x79, 0x0d,
	0xae, 0x68, 0x5d, 0x6d, 0xd5, 0xa3, 0x41, 0xbb, 0x11, 0x4a, 0x3a, 0x60, 0x3a, 0xf7, 0xb6, 0x60,
	0x05, 0xf3, 0x0d, 0x6a, 0x4c, 0xf7, 0x9b, 0xfd, 0xd7, 0xbd, 0xee, 0xba, 0x55, 0x07, 0xde, 0x1c,
	0x8d, 0x4e, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x65, 0x6a, 0xb8, 0x9a, 0xca, 0x02, 0x00, 0x00,
}
