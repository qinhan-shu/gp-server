// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/problem_manage.proto

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

// 获取全部题目信息（只下发 id & title & difficulty & pass_rate）
type GetProblemsReq struct {
	Tag                  string   `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	PageIndex            int64    `protobuf:"varint,3,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageNum              int64    `protobuf:"varint,4,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProblemsReq) Reset()         { *m = GetProblemsReq{} }
func (m *GetProblemsReq) String() string { return proto.CompactTextString(m) }
func (*GetProblemsReq) ProtoMessage()    {}
func (*GetProblemsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e4af9cb4edd521, []int{0}
}

func (m *GetProblemsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProblemsReq.Unmarshal(m, b)
}
func (m *GetProblemsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProblemsReq.Marshal(b, m, deterministic)
}
func (m *GetProblemsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProblemsReq.Merge(m, src)
}
func (m *GetProblemsReq) XXX_Size() int {
	return xxx_messageInfo_GetProblemsReq.Size(m)
}
func (m *GetProblemsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProblemsReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetProblemsReq proto.InternalMessageInfo

func (m *GetProblemsReq) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *GetProblemsReq) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *GetProblemsReq) GetPageNum() int64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

type GetProblemsResp struct {
	Status               *Status    `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Problems             []*Problem `protobuf:"bytes,2,rep,name=problems,proto3" json:"problems,omitempty"`
	PageIndex            int64      `protobuf:"varint,3,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageNum              int64      `protobuf:"varint,4,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetProblemsResp) Reset()         { *m = GetProblemsResp{} }
func (m *GetProblemsResp) String() string { return proto.CompactTextString(m) }
func (*GetProblemsResp) ProtoMessage()    {}
func (*GetProblemsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e4af9cb4edd521, []int{1}
}

func (m *GetProblemsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProblemsResp.Unmarshal(m, b)
}
func (m *GetProblemsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProblemsResp.Marshal(b, m, deterministic)
}
func (m *GetProblemsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProblemsResp.Merge(m, src)
}
func (m *GetProblemsResp) XXX_Size() int {
	return xxx_messageInfo_GetProblemsResp.Size(m)
}
func (m *GetProblemsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProblemsResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetProblemsResp proto.InternalMessageInfo

func (m *GetProblemsResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetProblemsResp) GetProblems() []*Problem {
	if m != nil {
		return m.Problems
	}
	return nil
}

func (m *GetProblemsResp) GetPageIndex() int64 {
	if m != nil {
		return m.PageIndex
	}
	return 0
}

func (m *GetProblemsResp) GetPageNum() int64 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

// 根据ID获得题目具体信息
type GetProblemByIDReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProblemByIDReq) Reset()         { *m = GetProblemByIDReq{} }
func (m *GetProblemByIDReq) String() string { return proto.CompactTextString(m) }
func (*GetProblemByIDReq) ProtoMessage()    {}
func (*GetProblemByIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e4af9cb4edd521, []int{2}
}

func (m *GetProblemByIDReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProblemByIDReq.Unmarshal(m, b)
}
func (m *GetProblemByIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProblemByIDReq.Marshal(b, m, deterministic)
}
func (m *GetProblemByIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProblemByIDReq.Merge(m, src)
}
func (m *GetProblemByIDReq) XXX_Size() int {
	return xxx_messageInfo_GetProblemByIDReq.Size(m)
}
func (m *GetProblemByIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProblemByIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetProblemByIDReq proto.InternalMessageInfo

func (m *GetProblemByIDReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetProblemByIDResp struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Problem              *Problem `protobuf:"bytes,2,opt,name=problem,proto3" json:"problem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProblemByIDResp) Reset()         { *m = GetProblemByIDResp{} }
func (m *GetProblemByIDResp) String() string { return proto.CompactTextString(m) }
func (*GetProblemByIDResp) ProtoMessage()    {}
func (*GetProblemByIDResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e4af9cb4edd521, []int{3}
}

func (m *GetProblemByIDResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProblemByIDResp.Unmarshal(m, b)
}
func (m *GetProblemByIDResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProblemByIDResp.Marshal(b, m, deterministic)
}
func (m *GetProblemByIDResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProblemByIDResp.Merge(m, src)
}
func (m *GetProblemByIDResp) XXX_Size() int {
	return xxx_messageInfo_GetProblemByIDResp.Size(m)
}
func (m *GetProblemByIDResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProblemByIDResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetProblemByIDResp proto.InternalMessageInfo

func (m *GetProblemByIDResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetProblemByIDResp) GetProblem() *Problem {
	if m != nil {
		return m.Problem
	}
	return nil
}

// 新增题目
type AddProblemReq struct {
	Problem              *Problem `protobuf:"bytes,1,opt,name=problem,proto3" json:"problem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddProblemReq) Reset()         { *m = AddProblemReq{} }
func (m *AddProblemReq) String() string { return proto.CompactTextString(m) }
func (*AddProblemReq) ProtoMessage()    {}
func (*AddProblemReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e4af9cb4edd521, []int{4}
}

func (m *AddProblemReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddProblemReq.Unmarshal(m, b)
}
func (m *AddProblemReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddProblemReq.Marshal(b, m, deterministic)
}
func (m *AddProblemReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProblemReq.Merge(m, src)
}
func (m *AddProblemReq) XXX_Size() int {
	return xxx_messageInfo_AddProblemReq.Size(m)
}
func (m *AddProblemReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProblemReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddProblemReq proto.InternalMessageInfo

func (m *AddProblemReq) GetProblem() *Problem {
	if m != nil {
		return m.Problem
	}
	return nil
}

type AddProblemResp struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IsSuccess            bool     `protobuf:"varint,2,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddProblemResp) Reset()         { *m = AddProblemResp{} }
func (m *AddProblemResp) String() string { return proto.CompactTextString(m) }
func (*AddProblemResp) ProtoMessage()    {}
func (*AddProblemResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e4af9cb4edd521, []int{5}
}

func (m *AddProblemResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddProblemResp.Unmarshal(m, b)
}
func (m *AddProblemResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddProblemResp.Marshal(b, m, deterministic)
}
func (m *AddProblemResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProblemResp.Merge(m, src)
}
func (m *AddProblemResp) XXX_Size() int {
	return xxx_messageInfo_AddProblemResp.Size(m)
}
func (m *AddProblemResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProblemResp.DiscardUnknown(m)
}

var xxx_messageInfo_AddProblemResp proto.InternalMessageInfo

func (m *AddProblemResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AddProblemResp) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

// 编辑题目
type EditProblemReq struct {
	Problem              *Problem `protobuf:"bytes,1,opt,name=problem,proto3" json:"problem,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditProblemReq) Reset()         { *m = EditProblemReq{} }
func (m *EditProblemReq) String() string { return proto.CompactTextString(m) }
func (*EditProblemReq) ProtoMessage()    {}
func (*EditProblemReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e4af9cb4edd521, []int{6}
}

func (m *EditProblemReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditProblemReq.Unmarshal(m, b)
}
func (m *EditProblemReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditProblemReq.Marshal(b, m, deterministic)
}
func (m *EditProblemReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditProblemReq.Merge(m, src)
}
func (m *EditProblemReq) XXX_Size() int {
	return xxx_messageInfo_EditProblemReq.Size(m)
}
func (m *EditProblemReq) XXX_DiscardUnknown() {
	xxx_messageInfo_EditProblemReq.DiscardUnknown(m)
}

var xxx_messageInfo_EditProblemReq proto.InternalMessageInfo

func (m *EditProblemReq) GetProblem() *Problem {
	if m != nil {
		return m.Problem
	}
	return nil
}

type EditProblemResp struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IsSuccess            bool     `protobuf:"varint,2,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditProblemResp) Reset()         { *m = EditProblemResp{} }
func (m *EditProblemResp) String() string { return proto.CompactTextString(m) }
func (*EditProblemResp) ProtoMessage()    {}
func (*EditProblemResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e4af9cb4edd521, []int{7}
}

func (m *EditProblemResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditProblemResp.Unmarshal(m, b)
}
func (m *EditProblemResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditProblemResp.Marshal(b, m, deterministic)
}
func (m *EditProblemResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditProblemResp.Merge(m, src)
}
func (m *EditProblemResp) XXX_Size() int {
	return xxx_messageInfo_EditProblemResp.Size(m)
}
func (m *EditProblemResp) XXX_DiscardUnknown() {
	xxx_messageInfo_EditProblemResp.DiscardUnknown(m)
}

var xxx_messageInfo_EditProblemResp proto.InternalMessageInfo

func (m *EditProblemResp) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *EditProblemResp) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func init() {
	proto.RegisterType((*GetProblemsReq)(nil), "protocol.GetProblemsReq")
	proto.RegisterType((*GetProblemsResp)(nil), "protocol.GetProblemsResp")
	proto.RegisterType((*GetProblemByIDReq)(nil), "protocol.GetProblemByIDReq")
	proto.RegisterType((*GetProblemByIDResp)(nil), "protocol.GetProblemByIDResp")
	proto.RegisterType((*AddProblemReq)(nil), "protocol.AddProblemReq")
	proto.RegisterType((*AddProblemResp)(nil), "protocol.AddProblemResp")
	proto.RegisterType((*EditProblemReq)(nil), "protocol.EditProblemReq")
	proto.RegisterType((*EditProblemResp)(nil), "protocol.EditProblemResp")
}

func init() { proto.RegisterFile("proto/problem_manage.proto", fileDescriptor_a2e4af9cb4edd521) }

var fileDescriptor_a2e4af9cb4edd521 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x69, 0x2b, 0x5b, 0xf7, 0xc6, 0xba, 0x2d, 0xa7, 0x3a, 0x18, 0x8c, 0x7a, 0x29, 0x88,
	0x13, 0xe6, 0x51, 0x3d, 0x38, 0x14, 0xd9, 0x45, 0x24, 0x3b, 0x39, 0x84, 0xd2, 0x35, 0xa1, 0x04,
	0x97, 0xa6, 0x5b, 0x52, 0xd0, 0xff, 0xc7, 0x3f, 0x54, 0x9a, 0xb4, 0x6e, 0x13, 0x41, 0xa6, 0x9e,
	0xfa, 0xfa, 0xfd, 0xf1, 0xfa, 0x79, 0x85, 0x41, 0xbe, 0x11, 0x4a, 0x9c, 0xe7, 0x1b, 0xb1, 0x5c,
	0x51, 0x1e, 0xf1, 0x38, 0x8b, 0x53, 0x3a, 0xd6, 0x22, 0x72, 0xf5, 0x23, 0x11, 0xab, 0x01, 0x32,
	0xa9, 0x44, 0x70, 0x2e, 0x32, 0xe3, 0xd6, 0x9a, 0x54, 0xb1, 0x2a, 0xa4, 0xd1, 0x82, 0x67, 0xf0,
	0xee, 0xa9, 0x7a, 0x34, 0xcb, 0x24, 0xa6, 0x6b, 0xd4, 0x03, 0x47, 0xc5, 0xa9, 0x6f, 0x8d, 0xac,
	0xb0, 0x85, 0xcb, 0x11, 0x0d, 0x01, 0xf2, 0x38, 0xa5, 0x11, 0xcb, 0x08, 0x7d, 0xf5, 0x9d, 0x91,
	0x15, 0x3a, 0xb8, 0x55, 0x2a, 0xb3, 0x52, 0x40, 0xc7, 0xe0, 0x6a, 0x3b, 0x2b, 0xb8, 0x7f, 0xa4,
	0xcd, 0x66, 0xf9, 0xfe, 0x50, 0xf0, 0xe0, 0xdd, 0x82, 0xee, 0xde, 0x7a, 0x99, 0xa3, 0x10, 0x1a,
	0x86, 0x40, 0x7f, 0xa2, 0x3d, 0xe9, 0x8d, 0x6b, 0xe8, 0xf1, 0x5c, 0xeb, 0xb8, 0xf2, 0xd1, 0x19,
	0xb8, 0xd5, 0x95, 0xd2, 0xb7, 0x47, 0x4e, 0xd8, 0x9e, 0xf4, 0xb7, 0xd9, 0x6a, 0x27, 0xfe, 0x8c,
	0xfc, 0x01, 0xf3, 0x04, 0xfa, 0x5b, 0xca, 0xe9, 0xdb, 0xec, 0xb6, 0xfc, 0x0f, 0x1e, 0xd8, 0x8c,
	0x68, 0x46, 0x07, 0xdb, 0x8c, 0x04, 0x2f, 0x80, 0xbe, 0x86, 0x0e, 0xba, 0xe6, 0x14, 0x9a, 0x15,
	0xaa, 0x6f, 0xeb, 0xe8, 0x37, 0xc7, 0xd4, 0x89, 0xe0, 0x0a, 0x3a, 0x37, 0x84, 0xd4, 0x32, 0x5d,
	0xef, 0xb6, 0xad, 0x1f, 0xdb, 0x4f, 0xe0, 0xed, 0xb6, 0x0f, 0xc2, 0x1c, 0x02, 0x30, 0x19, 0xc9,
	0x22, 0x49, 0xa8, 0x94, 0x9a, 0xd4, 0xc5, 0x2d, 0x26, 0xe7, 0x46, 0x08, 0xae, 0xc1, 0xbb, 0x23,
	0x4c, 0xfd, 0x96, 0x6c, 0x01, 0xdd, 0xbd, 0xfa, 0x3f, 0xa2, 0x4d, 0x3b, 0x8b, 0x76, 0x2a, 0x2e,
	0xeb, 0xf2, 0xb2, 0xa1, 0xa7, 0x8b, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x42, 0x57, 0xa6,
	0x30, 0x03, 0x00, 0x00,
}
