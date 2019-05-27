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

type JudgeResults struct {
	Status               *Status          `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	JudgeResults         map[int64]string `protobuf:"bytes,2,rep,name=judge_results,json=judgeResults,proto3" json:"judge_results,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *JudgeResults) Reset()         { *m = JudgeResults{} }
func (m *JudgeResults) String() string { return proto.CompactTextString(m) }
func (*JudgeResults) ProtoMessage()    {}
func (*JudgeResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_69971211ee56777f, []int{3}
}

func (m *JudgeResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JudgeResults.Unmarshal(m, b)
}
func (m *JudgeResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JudgeResults.Marshal(b, m, deterministic)
}
func (m *JudgeResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JudgeResults.Merge(m, src)
}
func (m *JudgeResults) XXX_Size() int {
	return xxx_messageInfo_JudgeResults.Size(m)
}
func (m *JudgeResults) XXX_DiscardUnknown() {
	xxx_messageInfo_JudgeResults.DiscardUnknown(m)
}

var xxx_messageInfo_JudgeResults proto.InternalMessageInfo

func (m *JudgeResults) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *JudgeResults) GetJudgeResults() map[int64]string {
	if m != nil {
		return m.JudgeResults
	}
	return nil
}

// PaperCompose 组卷算法
type PaperComposeAlgorithm struct {
	Status               *Status          `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Algorithm            map[int64]string `protobuf:"bytes,2,rep,name=algorithm,proto3" json:"algorithm,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PaperComposeAlgorithm) Reset()         { *m = PaperComposeAlgorithm{} }
func (m *PaperComposeAlgorithm) String() string { return proto.CompactTextString(m) }
func (*PaperComposeAlgorithm) ProtoMessage()    {}
func (*PaperComposeAlgorithm) Descriptor() ([]byte, []int) {
	return fileDescriptor_69971211ee56777f, []int{4}
}

func (m *PaperComposeAlgorithm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaperComposeAlgorithm.Unmarshal(m, b)
}
func (m *PaperComposeAlgorithm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaperComposeAlgorithm.Marshal(b, m, deterministic)
}
func (m *PaperComposeAlgorithm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaperComposeAlgorithm.Merge(m, src)
}
func (m *PaperComposeAlgorithm) XXX_Size() int {
	return xxx_messageInfo_PaperComposeAlgorithm.Size(m)
}
func (m *PaperComposeAlgorithm) XXX_DiscardUnknown() {
	xxx_messageInfo_PaperComposeAlgorithm.DiscardUnknown(m)
}

var xxx_messageInfo_PaperComposeAlgorithm proto.InternalMessageInfo

func (m *PaperComposeAlgorithm) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PaperComposeAlgorithm) GetAlgorithm() map[int64]string {
	if m != nil {
		return m.Algorithm
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
	proto.RegisterType((*JudgeResults)(nil), "protocol.JudgeResults")
	proto.RegisterMapType((map[int64]string)(nil), "protocol.JudgeResults.JudgeResultsEntry")
	proto.RegisterType((*PaperComposeAlgorithm)(nil), "protocol.PaperComposeAlgorithm")
	proto.RegisterMapType((map[int64]string)(nil), "protocol.PaperComposeAlgorithm.AlgorithmEntry")
}

func init() { proto.RegisterFile("proto/conf.proto", fileDescriptor_69971211ee56777f) }

var fileDescriptor_69971211ee56777f = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x50, 0xcf, 0x0a, 0xd3, 0x30,
	0x18, 0x27, 0xed, 0x1c, 0xdb, 0xb7, 0x4d, 0xb7, 0xaa, 0x50, 0x8a, 0x87, 0x32, 0x10, 0x7a, 0x8a,
	0x32, 0x0f, 0x8a, 0x53, 0x74, 0x4e, 0x2f, 0x32, 0x41, 0xaa, 0x5e, 0xbc, 0x48, 0xdc, 0xd2, 0xda,
	0x99, 0x35, 0xa3, 0x49, 0x85, 0xbd, 0x8b, 0xe0, 0x73, 0xf8, 0x06, 0x5e, 0x7c, 0x27, 0x69, 0x92,
	0xb6, 0xeb, 0xe6, 0x25, 0x97, 0xf6, 0x4b, 0x7e, 0x7f, 0xf2, 0xfb, 0x7d, 0x30, 0x3d, 0x16, 0x5c,
	0xf2, 0x07, 0x5b, 0x9e, 0x27, 0x58, 0x8d, 0xde, 0x40, 0xfd, 0xb6, 0x9c, 0x05, 0x9e, 0xc6, 0x84,
	0x24, 0xb2, 0x14, 0x1a, 0x9d, 0xff, 0x74, 0xa0, 0xbf, 0xe6, 0x79, 0x92, 0xa5, 0x5e, 0x04, 0x7d,
	0x0d, 0xf9, 0x28, 0x44, 0xd1, 0x68, 0x31, 0xc5, 0xb5, 0x12, 0x7f, 0x50, 0xf7, 0xb1, 0xc1, 0x3d,
	0x0c, 0x3d, 0x49, 0x52, 0xe1, 0x3b, 0xa1, 0x1b, 0x8d, 0x16, 0x41, 0xcb, 0xd3, 0x4e, 0xf8, 0x23,
	0x49, 0xc5, 0x9b, 0x5c, 0x16, 0xa7, 0x58, 0xf1, 0xbc, 0x97, 0x00, 0xbb, 0x2c, 0x49, 0xb2, 0x6d,
	0xc9, 0xe4, 0xc9, 0x77, 0x95, 0x2a, 0xbc, 0x52, 0xbd, 0x6e, 0x28, 0x5a, 0x7b, 0xa6, 0x09, 0x1e,
	0xc3, 0xb0, 0x31, 0xf5, 0xa6, 0xe0, 0x7e, 0xa7, 0x27, 0x95, 0xd2, 0x8d, 0xab, 0xd1, 0xbb, 0x03,
	0x37, 0x7e, 0x10, 0x56, 0x52, 0xdf, 0x09, 0x51, 0x34, 0x8c, 0xf5, 0xe1, 0xa9, 0xf3, 0x04, 0x05,
	0xcf, 0xe1, 0xd6, 0x85, 0xaf, 0x8d, 0x7c, 0xfe, 0x0b, 0xc1, 0xe0, 0x93, 0xa0, 0x45, 0xcc, 0x19,
	0xb5, 0x58, 0xd0, 0x43, 0xe8, 0x15, 0x9c, 0x51, 0xb3, 0xa0, 0x7b, 0x2d, 0xaf, 0xf6, 0xc2, 0xd5,
	0xc7, 0xac, 0xa8, 0x62, 0x56, 0x05, 0x9b, 0x2b, 0xab, 0x84, 0xbf, 0x11, 0x4c, 0xde, 0x96, 0xbb,
	0x94, 0x6e, 0x48, 0x9e, 0x96, 0x24, 0xb5, 0x89, 0xb9, 0x82, 0x01, 0x33, 0x2a, 0x13, 0xf5, 0x7e,
	0xcb, 0xed, 0x98, 0xe2, 0x7a, 0xd0, 0x99, 0x1b, 0x59, 0xb0, 0x84, 0x49, 0x07, 0xb2, 0xca, 0xfe,
	0x07, 0xc1, 0x58, 0x3d, 0x13, 0x53, 0x51, 0x32, 0x29, 0x2c, 0xa2, 0xbf, 0x83, 0xc9, 0xbe, 0x52,
	0x7e, 0x29, 0xb4, 0xd4, 0xe4, 0x8f, 0x2e, 0xf2, 0x1b, 0xe3, 0xce, 0x41, 0x57, 0x18, 0xef, 0xcf,
	0xae, 0x82, 0x17, 0x30, 0xbb, 0xa2, 0x58, 0x55, 0xf9, 0x8b, 0xe0, 0xee, 0x7b, 0x72, 0xa4, 0xc5,
	0x9a, 0x1f, 0x8e, 0x5c, 0xd0, 0x15, 0x4b, 0x79, 0x91, 0xc9, 0x6f, 0x07, 0x8b, 0x4e, 0x1b, 0x18,
	0x92, 0x5a, 0x66, 0xfa, 0xe0, 0x96, 0xfc, 0x5f, 0x77, 0xdc, 0x4c, 0xba, 0x55, 0x6b, 0x10, 0x3c,
	0x83, 0x9b, 0x5d, 0xd0, 0xa6, 0xcf, 0xab, 0xdb, 0x9f, 0x67, 0xf5, 0xcb, 0xcb, 0x7a, 0xf8, 0xda,
	0x57, 0xd3, 0xa3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x85, 0x4a, 0x03, 0x65, 0x04, 0x00,
	0x00,
}
