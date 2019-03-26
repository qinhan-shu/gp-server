// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common.proto

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

// Role : 用户角色（学生/老师...）
type Role int32

const (
	Role_STUDENT Role = 0
	Role_TEACHER Role = 1
	Role_MANAGER Role = 2
)

var Role_name = map[int32]string{
	0: "STUDENT",
	1: "TEACHER",
	2: "MANAGER",
}

var Role_value = map[string]int32{
	"STUDENT": 0,
	"TEACHER": 1,
	"MANAGER": 2,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0}
}

// ProblemDifficluty : 题目难度
type ProblemDifficluty int32

const (
	ProblemDifficluty_EASY   ProblemDifficluty = 0
	ProblemDifficluty_MEDIUM ProblemDifficluty = 1
	ProblemDifficluty_HARD   ProblemDifficluty = 2
)

var ProblemDifficluty_name = map[int32]string{
	0: "EASY",
	1: "MEDIUM",
	2: "HARD",
}

var ProblemDifficluty_value = map[string]int32{
	"EASY":   0,
	"MEDIUM": 1,
	"HARD":   2,
}

func (x ProblemDifficluty) String() string {
	return proto.EnumName(ProblemDifficluty_name, int32(x))
}

func (ProblemDifficluty) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{1}
}

// UserInfo : 用户基本信息
type UserInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Role                 Role     `protobuf:"varint,4,opt,name=role,proto3,enum=protocol.Role" json:"role,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Sex                  bool     `protobuf:"varint,6,opt,name=sex,proto3" json:"sex,omitempty"`
	Phone                string   `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                string   `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	School               string   `protobuf:"bytes,9,opt,name=school,proto3" json:"school,omitempty"`
	LastLogin            int64    `protobuf:"varint,10,opt,name=last_login,json=lastLogin,proto3" json:"last_login,omitempty"`
	Create               int64    `protobuf:"varint,11,opt,name=create,proto3" json:"create,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfo) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_STUDENT
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetSex() bool {
	if m != nil {
		return m.Sex
	}
	return false
}

func (m *UserInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserInfo) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func (m *UserInfo) GetLastLogin() int64 {
	if m != nil {
		return m.LastLogin
	}
	return 0
}

func (m *UserInfo) GetCreate() int64 {
	if m != nil {
		return m.Create
	}
	return 0
}

func (m *UserInfo) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *UserInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// ProblemExample : 题目输入输出样例
type ProblemExample struct {
	Input                string   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Output               string   `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProblemExample) Reset()         { *m = ProblemExample{} }
func (m *ProblemExample) String() string { return proto.CompactTextString(m) }
func (*ProblemExample) ProtoMessage()    {}
func (*ProblemExample) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{1}
}

func (m *ProblemExample) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProblemExample.Unmarshal(m, b)
}
func (m *ProblemExample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProblemExample.Marshal(b, m, deterministic)
}
func (m *ProblemExample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProblemExample.Merge(m, src)
}
func (m *ProblemExample) XXX_Size() int {
	return xxx_messageInfo_ProblemExample.Size(m)
}
func (m *ProblemExample) XXX_DiscardUnknown() {
	xxx_messageInfo_ProblemExample.DiscardUnknown(m)
}

var xxx_messageInfo_ProblemExample proto.InternalMessageInfo

func (m *ProblemExample) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *ProblemExample) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

// Problem : 题目
type Problem struct {
	Id                   int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string            `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	In                   string            `protobuf:"bytes,4,opt,name=in,proto3" json:"in,omitempty"`
	Out                  string            `protobuf:"bytes,5,opt,name=out,proto3" json:"out,omitempty"`
	Hint                 string            `protobuf:"bytes,6,opt,name=hint,proto3" json:"hint,omitempty"`
	InOutExamples        []*ProblemExample `protobuf:"bytes,7,rep,name=in_out_examples,json=inOutExamples,proto3" json:"in_out_examples,omitempty"`
	JudgeLimitTime       int64             `protobuf:"varint,8,opt,name=judge_limit_time,json=judgeLimitTime,proto3" json:"judge_limit_time,omitempty"`
	JudgeLimitMem        int64             `protobuf:"varint,9,opt,name=judge_limit_mem,json=judgeLimitMem,proto3" json:"judge_limit_mem,omitempty"`
	Tags                 []int64           `protobuf:"varint,10,rep,packed,name=tags,proto3" json:"tags,omitempty"`
	Difficluty           ProblemDifficluty `protobuf:"varint,11,opt,name=difficluty,proto3,enum=protocol.ProblemDifficluty" json:"difficluty,omitempty"`
	SubmitTime           int64             `protobuf:"varint,12,opt,name=submit_time,json=submitTime,proto3" json:"submit_time,omitempty"`
	AcceptTime           int64             `protobuf:"varint,13,opt,name=accept_time,json=acceptTime,proto3" json:"accept_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Problem) Reset()         { *m = Problem{} }
func (m *Problem) String() string { return proto.CompactTextString(m) }
func (*Problem) ProtoMessage()    {}
func (*Problem) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{2}
}

func (m *Problem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Problem.Unmarshal(m, b)
}
func (m *Problem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Problem.Marshal(b, m, deterministic)
}
func (m *Problem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Problem.Merge(m, src)
}
func (m *Problem) XXX_Size() int {
	return xxx_messageInfo_Problem.Size(m)
}
func (m *Problem) XXX_DiscardUnknown() {
	xxx_messageInfo_Problem.DiscardUnknown(m)
}

var xxx_messageInfo_Problem proto.InternalMessageInfo

func (m *Problem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Problem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Problem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Problem) GetIn() string {
	if m != nil {
		return m.In
	}
	return ""
}

func (m *Problem) GetOut() string {
	if m != nil {
		return m.Out
	}
	return ""
}

func (m *Problem) GetHint() string {
	if m != nil {
		return m.Hint
	}
	return ""
}

func (m *Problem) GetInOutExamples() []*ProblemExample {
	if m != nil {
		return m.InOutExamples
	}
	return nil
}

func (m *Problem) GetJudgeLimitTime() int64 {
	if m != nil {
		return m.JudgeLimitTime
	}
	return 0
}

func (m *Problem) GetJudgeLimitMem() int64 {
	if m != nil {
		return m.JudgeLimitMem
	}
	return 0
}

func (m *Problem) GetTags() []int64 {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Problem) GetDifficluty() ProblemDifficluty {
	if m != nil {
		return m.Difficluty
	}
	return ProblemDifficluty_EASY
}

func (m *Problem) GetSubmitTime() int64 {
	if m != nil {
		return m.SubmitTime
	}
	return 0
}

func (m *Problem) GetAcceptTime() int64 {
	if m != nil {
		return m.AcceptTime
	}
	return 0
}

// SubmitRecord : 提交情况
type SubmitRecord struct {
	Problem              *Problem `protobuf:"bytes,1,opt,name=problem,proto3" json:"problem,omitempty"`
	SubmitTime           int64    `protobuf:"varint,2,opt,name=submit_time,json=submitTime,proto3" json:"submit_time,omitempty"`
	IsPass               bool     `protobuf:"varint,3,opt,name=is_pass,json=isPass,proto3" json:"is_pass,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitRecord) Reset()         { *m = SubmitRecord{} }
func (m *SubmitRecord) String() string { return proto.CompactTextString(m) }
func (*SubmitRecord) ProtoMessage()    {}
func (*SubmitRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{3}
}

func (m *SubmitRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitRecord.Unmarshal(m, b)
}
func (m *SubmitRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitRecord.Marshal(b, m, deterministic)
}
func (m *SubmitRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitRecord.Merge(m, src)
}
func (m *SubmitRecord) XXX_Size() int {
	return xxx_messageInfo_SubmitRecord.Size(m)
}
func (m *SubmitRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitRecord.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitRecord proto.InternalMessageInfo

func (m *SubmitRecord) GetProblem() *Problem {
	if m != nil {
		return m.Problem
	}
	return nil
}

func (m *SubmitRecord) GetSubmitTime() int64 {
	if m != nil {
		return m.SubmitTime
	}
	return 0
}

func (m *SubmitRecord) GetIsPass() bool {
	if m != nil {
		return m.IsPass
	}
	return false
}

// Announcement : 公告，包括班级公告和全局公告
type Announcement struct {
	Publisher            string   `protobuf:"bytes,1,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	CreateTime           int64    `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	DisableTime          int64    `protobuf:"varint,4,opt,name=disable_time,json=disableTime,proto3" json:"disable_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Announcement) Reset()         { *m = Announcement{} }
func (m *Announcement) String() string { return proto.CompactTextString(m) }
func (*Announcement) ProtoMessage()    {}
func (*Announcement) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{4}
}

func (m *Announcement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Announcement.Unmarshal(m, b)
}
func (m *Announcement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Announcement.Marshal(b, m, deterministic)
}
func (m *Announcement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Announcement.Merge(m, src)
}
func (m *Announcement) XXX_Size() int {
	return xxx_messageInfo_Announcement.Size(m)
}
func (m *Announcement) XXX_DiscardUnknown() {
	xxx_messageInfo_Announcement.DiscardUnknown(m)
}

var xxx_messageInfo_Announcement proto.InternalMessageInfo

func (m *Announcement) GetPublisher() string {
	if m != nil {
		return m.Publisher
	}
	return ""
}

func (m *Announcement) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *Announcement) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Announcement) GetDisableTime() int64 {
	if m != nil {
		return m.DisableTime
	}
	return 0
}

// Class : 班级信息
type Class struct {
	Id                   int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Tutor                string          `protobuf:"bytes,2,opt,name=tutor,proto3" json:"tutor,omitempty"`
	Name                 string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Introduction         string          `protobuf:"bytes,4,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Number               int64           `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	IsCheck              bool            `protobuf:"varint,6,opt,name=is_check,json=isCheck,proto3" json:"is_check,omitempty"`
	CreateTime           int64           `protobuf:"varint,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Announcements        []*Announcement `protobuf:"bytes,8,rep,name=announcements,proto3" json:"announcements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Class) Reset()         { *m = Class{} }
func (m *Class) String() string { return proto.CompactTextString(m) }
func (*Class) ProtoMessage()    {}
func (*Class) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{5}
}

func (m *Class) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Class.Unmarshal(m, b)
}
func (m *Class) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Class.Marshal(b, m, deterministic)
}
func (m *Class) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Class.Merge(m, src)
}
func (m *Class) XXX_Size() int {
	return xxx_messageInfo_Class.Size(m)
}
func (m *Class) XXX_DiscardUnknown() {
	xxx_messageInfo_Class.DiscardUnknown(m)
}

var xxx_messageInfo_Class proto.InternalMessageInfo

func (m *Class) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Class) GetTutor() string {
	if m != nil {
		return m.Tutor
	}
	return ""
}

func (m *Class) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Class) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

func (m *Class) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Class) GetIsCheck() bool {
	if m != nil {
		return m.IsCheck
	}
	return false
}

func (m *Class) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Class) GetAnnouncements() []*Announcement {
	if m != nil {
		return m.Announcements
	}
	return nil
}

func init() {
	proto.RegisterEnum("protocol.Role", Role_name, Role_value)
	proto.RegisterEnum("protocol.ProblemDifficluty", ProblemDifficluty_name, ProblemDifficluty_value)
	proto.RegisterType((*UserInfo)(nil), "protocol.UserInfo")
	proto.RegisterType((*ProblemExample)(nil), "protocol.ProblemExample")
	proto.RegisterType((*Problem)(nil), "protocol.Problem")
	proto.RegisterType((*SubmitRecord)(nil), "protocol.SubmitRecord")
	proto.RegisterType((*Announcement)(nil), "protocol.Announcement")
	proto.RegisterType((*Class)(nil), "protocol.Class")
}

func init() { proto.RegisterFile("proto/common.proto", fileDescriptor_1747d3070a2311a0) }

var fileDescriptor_1747d3070a2311a0 = []byte{
	// 757 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xcb, 0x8e, 0xdb, 0x36,
	0x14, 0x8d, 0x24, 0xdb, 0x92, 0xaf, 0x6c, 0x47, 0x21, 0x82, 0x94, 0x7d, 0x21, 0xae, 0x17, 0x85,
	0x91, 0xa2, 0x53, 0xc0, 0x59, 0xa6, 0x28, 0xea, 0x8e, 0x8d, 0x26, 0x40, 0x9c, 0x06, 0x1c, 0xcf,
	0xa2, 0xdd, 0x18, 0xb2, 0xc4, 0xb1, 0xd9, 0x4a, 0xa4, 0x20, 0x52, 0x68, 0xfa, 0x07, 0xfd, 0x81,
	0xfe, 0x46, 0x7f, 0xaf, 0xdb, 0x82, 0x97, 0xf2, 0x63, 0xc6, 0x59, 0xf9, 0x9e, 0xa3, 0x43, 0xf2,
	0xf0, 0xde, 0x43, 0x03, 0xa9, 0x6a, 0x65, 0xd4, 0x77, 0x99, 0x2a, 0x4b, 0x25, 0xaf, 0x10, 0x90,
	0x08, 0x7f, 0x32, 0x55, 0x4c, 0xfe, 0xf1, 0x21, 0xba, 0xd5, 0xbc, 0x7e, 0x23, 0xef, 0x14, 0x19,
	0x81, 0x2f, 0x72, 0xea, 0x8d, 0xbd, 0x69, 0xc0, 0x7c, 0x91, 0x93, 0x09, 0x74, 0x6a, 0x55, 0x70,
	0xda, 0x19, 0x7b, 0xd3, 0xd1, 0x6c, 0x74, 0x75, 0x58, 0x75, 0xc5, 0x54, 0xc1, 0x19, 0x7e, 0x23,
	0x04, 0x3a, 0x32, 0x2d, 0x39, 0xed, 0x8e, 0xbd, 0x69, 0x9f, 0x61, 0x4d, 0x12, 0x08, 0x34, 0xff,
	0x40, 0x7b, 0x63, 0x6f, 0x1a, 0x31, 0x5b, 0x92, 0xa7, 0xd0, 0xad, 0xf6, 0x4a, 0x72, 0x1a, 0xa2,
	0xcc, 0x01, 0xcb, 0xf2, 0x32, 0x15, 0x05, 0x8d, 0x1c, 0x8b, 0x80, 0x3c, 0x83, 0x9e, 0xce, 0xf6,
	0x4a, 0x15, 0xb4, 0x8f, 0x74, 0x8b, 0xc8, 0x97, 0x00, 0x45, 0xaa, 0xcd, 0xa6, 0x50, 0x3b, 0x21,
	0x29, 0xa0, 0xcb, 0xbe, 0x65, 0xde, 0x5a, 0xc2, 0x2e, 0xcb, 0x6a, 0x9e, 0x1a, 0x4e, 0x63, 0xfc,
	0xd4, 0x22, 0x42, 0x21, 0x4c, 0xb3, 0x4c, 0x35, 0xd2, 0x50, 0x1f, 0xf7, 0x3b, 0x40, 0xf2, 0x19,
	0x44, 0x55, 0xaa, 0xf5, 0x9f, 0xaa, 0xce, 0x69, 0x80, 0x9f, 0x8e, 0x78, 0xf2, 0x03, 0x8c, 0xde,
	0xd7, 0x6a, 0x5b, 0xf0, 0x72, 0xf9, 0x21, 0x2d, 0xab, 0x02, 0xcd, 0x0a, 0x59, 0x35, 0x06, 0xfb,
	0xd3, 0x67, 0x0e, 0xd8, 0x53, 0x55, 0x63, 0x2c, 0xed, 0x36, 0x6f, 0xd1, 0xe4, 0xdf, 0x00, 0xc2,
	0x76, 0x83, 0x8b, 0xb6, 0x3e, 0x85, 0xae, 0x11, 0xa6, 0xe0, 0xed, 0x12, 0x07, 0xc8, 0x18, 0xe2,
	0x9c, 0xeb, 0xac, 0x16, 0x95, 0x11, 0x4a, 0xb6, 0x86, 0xce, 0x29, 0xdc, 0x47, 0xe2, 0x30, 0xfa,
	0xcc, 0x17, 0xd2, 0xb6, 0x59, 0x35, 0xa6, 0xed, 0xbc, 0x2d, 0xed, 0x30, 0xf6, 0x42, 0x1a, 0xec,
	0x7c, 0x9f, 0x61, 0x4d, 0x7e, 0x84, 0xc7, 0x42, 0x6e, 0x54, 0x63, 0x36, 0xdc, 0xdd, 0x44, 0xd3,
	0x70, 0x1c, 0x4c, 0xe3, 0x19, 0x3d, 0xcd, 0xf3, 0xfe, 0x55, 0xd9, 0x50, 0xc8, 0x5f, 0x1a, 0xd3,
	0x22, 0x4d, 0xa6, 0x90, 0xfc, 0xde, 0xe4, 0x3b, 0xbe, 0x29, 0x44, 0x29, 0xcc, 0xc6, 0x88, 0x92,
	0xe3, 0xc4, 0x02, 0x36, 0x42, 0xfe, 0xad, 0xa5, 0xd7, 0xa2, 0xe4, 0xe4, 0x6b, 0x78, 0x7c, 0xae,
	0x2c, 0x79, 0x89, 0x33, 0x0c, 0xd8, 0xf0, 0x24, 0x5c, 0xf1, 0xd2, 0xfa, 0x34, 0xe9, 0x4e, 0x53,
	0x18, 0x07, 0xd3, 0x80, 0x61, 0x4d, 0x5e, 0x01, 0xe4, 0xe2, 0xee, 0x4e, 0x64, 0x45, 0x63, 0xfe,
	0xc2, 0x19, 0x8e, 0x66, 0x9f, 0x5f, 0x58, 0x5c, 0x1c, 0x25, 0xec, 0x4c, 0x4e, 0x9e, 0x43, 0xac,
	0x9b, 0xed, 0xd1, 0xdd, 0x00, 0x0f, 0x05, 0x47, 0xa1, 0xb3, 0xe7, 0x10, 0xa7, 0x59, 0xc6, 0xab,
	0x56, 0x30, 0x74, 0x02, 0x47, 0x59, 0xc1, 0xa4, 0x81, 0xc1, 0x0d, 0xca, 0x19, 0xcf, 0x54, 0x9d,
	0x93, 0x6f, 0x20, 0xac, 0xdc, 0x91, 0x38, 0xb9, 0x78, 0xf6, 0xe4, 0xc2, 0x0b, 0x3b, 0x28, 0x1e,
	0x1e, 0xef, 0x5f, 0x1c, 0xff, 0x09, 0x84, 0x42, 0x6f, 0x6c, 0xba, 0x70, 0xb0, 0x11, 0xeb, 0x09,
	0xfd, 0x3e, 0xd5, 0x7a, 0xf2, 0xb7, 0x07, 0x83, 0xb9, 0x94, 0xaa, 0x91, 0x19, 0x2f, 0xb9, 0x34,
	0xe4, 0x0b, 0xe8, 0x57, 0xcd, 0xb6, 0x10, 0x7a, 0xcf, 0xeb, 0x36, 0x6a, 0x27, 0xc2, 0xc6, 0x2d,
	0xe7, 0xc6, 0x3e, 0x99, 0x36, 0x6e, 0x0e, 0x59, 0x03, 0x2e, 0xee, 0xce, 0x40, 0xe0, 0x0c, 0x38,
	0x0a, 0x0d, 0x7c, 0x05, 0x83, 0x5c, 0xe8, 0x74, 0x5b, 0xb4, 0x8a, 0x0e, 0x2a, 0xe2, 0x96, 0xc3,
	0x0e, 0xfc, 0xe7, 0x41, 0xf7, 0xba, 0x48, 0xb5, 0xfe, 0x68, 0x60, 0x1b, 0xa3, 0xea, 0x63, 0x60,
	0x2d, 0x38, 0xbe, 0xfc, 0xe0, 0xec, 0xe5, 0x4f, 0x60, 0x20, 0xa4, 0xa9, 0x55, 0xde, 0x64, 0x98,
	0x62, 0x17, 0xd6, 0x7b, 0x9c, 0xbd, 0x83, 0x6c, 0xca, 0x2d, 0xaf, 0x31, 0xb9, 0x01, 0x6b, 0x11,
	0xf9, 0x14, 0x22, 0xa1, 0x37, 0xd9, 0x9e, 0x67, 0x7f, 0xb4, 0x7f, 0x1d, 0xa1, 0xd0, 0xd7, 0x16,
	0x3e, 0xbc, 0x5e, 0x78, 0x71, 0xbd, 0xef, 0x61, 0x98, 0x9e, 0x75, 0x51, 0xd3, 0x08, 0x23, 0xfe,
	0xec, 0x34, 0xb3, 0xf3, 0x26, 0xb3, 0xfb, 0xe2, 0x17, 0xdf, 0x42, 0xc7, 0xfe, 0xa3, 0x91, 0x18,
	0xc2, 0x9b, 0xf5, 0xed, 0x62, 0xf9, 0x6e, 0x9d, 0x3c, 0xb2, 0x60, 0xbd, 0x9c, 0x5f, 0xbf, 0x5e,
	0xb2, 0xc4, 0xb3, 0x60, 0x35, 0x7f, 0x37, 0xff, 0x79, 0xc9, 0x12, 0xff, 0xc5, 0x4b, 0x78, 0x72,
	0x91, 0x46, 0x12, 0x41, 0x67, 0x39, 0xbf, 0xf9, 0x35, 0x79, 0x44, 0x00, 0x7a, 0xab, 0xe5, 0xe2,
	0xcd, 0xed, 0x2a, 0xf1, 0x2c, 0xfb, 0x7a, 0xce, 0x16, 0x89, 0xff, 0xd3, 0xf0, 0xb7, 0x78, 0xa7,
	0x5e, 0x1d, 0xec, 0x6c, 0x7b, 0x58, 0xbd, 0xfc, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x87, 0x45, 0x48,
	0x33, 0x9e, 0x05, 0x00, 0x00,
}
