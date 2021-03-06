// Code generated by protoc-gen-go. DO NOT EDIT.
// source: publisher.proto

package publisher

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

// 业务测试结构体
type BusinessTest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Secret               string   `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	TestType             string   `protobuf:"bytes,3,opt,name=testType,proto3" json:"testType,omitempty"`
	TestOption           int64    `protobuf:"varint,4,opt,name=testOption,proto3" json:"testOption,omitempty"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BusinessTest) Reset()         { *m = BusinessTest{} }
func (m *BusinessTest) String() string { return proto.CompactTextString(m) }
func (*BusinessTest) ProtoMessage()    {}
func (*BusinessTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_41489454d08668ce, []int{0}
}

func (m *BusinessTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BusinessTest.Unmarshal(m, b)
}
func (m *BusinessTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BusinessTest.Marshal(b, m, deterministic)
}
func (m *BusinessTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BusinessTest.Merge(m, src)
}
func (m *BusinessTest) XXX_Size() int {
	return xxx_messageInfo_BusinessTest.Size(m)
}
func (m *BusinessTest) XXX_DiscardUnknown() {
	xxx_messageInfo_BusinessTest.DiscardUnknown(m)
}

var xxx_messageInfo_BusinessTest proto.InternalMessageInfo

func (m *BusinessTest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *BusinessTest) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *BusinessTest) GetTestType() string {
	if m != nil {
		return m.TestType
	}
	return ""
}

func (m *BusinessTest) GetTestOption() int64 {
	if m != nil {
		return m.TestOption
	}
	return 0
}

func (m *BusinessTest) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

// 业务结构体
type Business struct {
	Secret               string          `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	SecretValue          string          `protobuf:"bytes,2,opt,name=secretValue,proto3" json:"secretValue,omitempty"`
	ApkName              string          `protobuf:"bytes,3,opt,name=apkName,proto3" json:"apkName,omitempty"`
	PullFrequency        int64           `protobuf:"varint,4,opt,name=pullFrequency,proto3" json:"pullFrequency,omitempty"`
	CurrentEdition       float32         `protobuf:"fixed32,5,opt,name=currentEdition,proto3" json:"currentEdition,omitempty"`
	FileUrl              string          `protobuf:"bytes,6,opt,name=fileUrl,proto3" json:"fileUrl,omitempty"`
	IsStop               bool            `protobuf:"varint,7,opt,name=isStop,proto3" json:"isStop,omitempty"`
	UpdateTime           string          `protobuf:"bytes,8,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	BusinessTests        []*BusinessTest `protobuf:"bytes,9,rep,name=businessTests,proto3" json:"businessTests,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Business) Reset()         { *m = Business{} }
func (m *Business) String() string { return proto.CompactTextString(m) }
func (*Business) ProtoMessage()    {}
func (*Business) Descriptor() ([]byte, []int) {
	return fileDescriptor_41489454d08668ce, []int{1}
}

func (m *Business) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Business.Unmarshal(m, b)
}
func (m *Business) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Business.Marshal(b, m, deterministic)
}
func (m *Business) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Business.Merge(m, src)
}
func (m *Business) XXX_Size() int {
	return xxx_messageInfo_Business.Size(m)
}
func (m *Business) XXX_DiscardUnknown() {
	xxx_messageInfo_Business.DiscardUnknown(m)
}

var xxx_messageInfo_Business proto.InternalMessageInfo

func (m *Business) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Business) GetSecretValue() string {
	if m != nil {
		return m.SecretValue
	}
	return ""
}

func (m *Business) GetApkName() string {
	if m != nil {
		return m.ApkName
	}
	return ""
}

func (m *Business) GetPullFrequency() int64 {
	if m != nil {
		return m.PullFrequency
	}
	return 0
}

func (m *Business) GetCurrentEdition() float32 {
	if m != nil {
		return m.CurrentEdition
	}
	return 0
}

func (m *Business) GetFileUrl() string {
	if m != nil {
		return m.FileUrl
	}
	return ""
}

func (m *Business) GetIsStop() bool {
	if m != nil {
		return m.IsStop
	}
	return false
}

func (m *Business) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *Business) GetBusinessTests() []*BusinessTest {
	if m != nil {
		return m.BusinessTests
	}
	return nil
}

// 获取所有业务
type FindAllBusinessReq struct {
	All                  bool     `protobuf:"varint,1,opt,name=all,proto3" json:"all,omitempty"`
	IsStop               bool     `protobuf:"varint,2,opt,name=isStop,proto3" json:"isStop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindAllBusinessReq) Reset()         { *m = FindAllBusinessReq{} }
func (m *FindAllBusinessReq) String() string { return proto.CompactTextString(m) }
func (*FindAllBusinessReq) ProtoMessage()    {}
func (*FindAllBusinessReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_41489454d08668ce, []int{2}
}

func (m *FindAllBusinessReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllBusinessReq.Unmarshal(m, b)
}
func (m *FindAllBusinessReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllBusinessReq.Marshal(b, m, deterministic)
}
func (m *FindAllBusinessReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllBusinessReq.Merge(m, src)
}
func (m *FindAllBusinessReq) XXX_Size() int {
	return xxx_messageInfo_FindAllBusinessReq.Size(m)
}
func (m *FindAllBusinessReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllBusinessReq.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllBusinessReq proto.InternalMessageInfo

func (m *FindAllBusinessReq) GetAll() bool {
	if m != nil {
		return m.All
	}
	return false
}

func (m *FindAllBusinessReq) GetIsStop() bool {
	if m != nil {
		return m.IsStop
	}
	return false
}

type FindAllBusinessResp struct {
	Businesses           []*Business `protobuf:"bytes,1,rep,name=businesses,proto3" json:"businesses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FindAllBusinessResp) Reset()         { *m = FindAllBusinessResp{} }
func (m *FindAllBusinessResp) String() string { return proto.CompactTextString(m) }
func (*FindAllBusinessResp) ProtoMessage()    {}
func (*FindAllBusinessResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_41489454d08668ce, []int{3}
}

func (m *FindAllBusinessResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindAllBusinessResp.Unmarshal(m, b)
}
func (m *FindAllBusinessResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindAllBusinessResp.Marshal(b, m, deterministic)
}
func (m *FindAllBusinessResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindAllBusinessResp.Merge(m, src)
}
func (m *FindAllBusinessResp) XXX_Size() int {
	return xxx_messageInfo_FindAllBusinessResp.Size(m)
}
func (m *FindAllBusinessResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FindAllBusinessResp.DiscardUnknown(m)
}

var xxx_messageInfo_FindAllBusinessResp proto.InternalMessageInfo

func (m *FindAllBusinessResp) GetBusinesses() []*Business {
	if m != nil {
		return m.Businesses
	}
	return nil
}

// 更新业务版本
type UpdateBusinessEditionReq struct {
	Edition              float32  `protobuf:"fixed32,1,opt,name=edition,proto3" json:"edition,omitempty"`
	Secret               string   `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBusinessEditionReq) Reset()         { *m = UpdateBusinessEditionReq{} }
func (m *UpdateBusinessEditionReq) String() string { return proto.CompactTextString(m) }
func (*UpdateBusinessEditionReq) ProtoMessage()    {}
func (*UpdateBusinessEditionReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_41489454d08668ce, []int{4}
}

func (m *UpdateBusinessEditionReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBusinessEditionReq.Unmarshal(m, b)
}
func (m *UpdateBusinessEditionReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBusinessEditionReq.Marshal(b, m, deterministic)
}
func (m *UpdateBusinessEditionReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBusinessEditionReq.Merge(m, src)
}
func (m *UpdateBusinessEditionReq) XXX_Size() int {
	return xxx_messageInfo_UpdateBusinessEditionReq.Size(m)
}
func (m *UpdateBusinessEditionReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBusinessEditionReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBusinessEditionReq proto.InternalMessageInfo

func (m *UpdateBusinessEditionReq) GetEdition() float32 {
	if m != nil {
		return m.Edition
	}
	return 0
}

func (m *UpdateBusinessEditionReq) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

// 更新业务版本
type UpdateBusinessEditionResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBusinessEditionResp) Reset()         { *m = UpdateBusinessEditionResp{} }
func (m *UpdateBusinessEditionResp) String() string { return proto.CompactTextString(m) }
func (*UpdateBusinessEditionResp) ProtoMessage()    {}
func (*UpdateBusinessEditionResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_41489454d08668ce, []int{5}
}

func (m *UpdateBusinessEditionResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBusinessEditionResp.Unmarshal(m, b)
}
func (m *UpdateBusinessEditionResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBusinessEditionResp.Marshal(b, m, deterministic)
}
func (m *UpdateBusinessEditionResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBusinessEditionResp.Merge(m, src)
}
func (m *UpdateBusinessEditionResp) XXX_Size() int {
	return xxx_messageInfo_UpdateBusinessEditionResp.Size(m)
}
func (m *UpdateBusinessEditionResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBusinessEditionResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBusinessEditionResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BusinessTest)(nil), "publisher.BusinessTest")
	proto.RegisterType((*Business)(nil), "publisher.Business")
	proto.RegisterType((*FindAllBusinessReq)(nil), "publisher.FindAllBusinessReq")
	proto.RegisterType((*FindAllBusinessResp)(nil), "publisher.FindAllBusinessResp")
	proto.RegisterType((*UpdateBusinessEditionReq)(nil), "publisher.UpdateBusinessEditionReq")
	proto.RegisterType((*UpdateBusinessEditionResp)(nil), "publisher.UpdateBusinessEditionResp")
}

func init() { proto.RegisterFile("publisher.proto", fileDescriptor_41489454d08668ce) }

var fileDescriptor_41489454d08668ce = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0x71, 0x16, 0x76, 0xb3, 0x53, 0xda, 0xa2, 0xa9, 0x00, 0xb3, 0x88, 0x2a, 0x0a, 0x15,
	0xda, 0x53, 0x0f, 0xed, 0x19, 0x24, 0x90, 0xe8, 0x01, 0x21, 0x40, 0xee, 0x96, 0x7b, 0x76, 0x33,
	0x08, 0x0b, 0x37, 0x71, 0x6d, 0xe7, 0xd0, 0xd7, 0xe0, 0x7d, 0x78, 0x05, 0x9e, 0x09, 0xd9, 0x8d,
	0x8b, 0x37, 0xdd, 0x85, 0x9b, 0xe7, 0x9f, 0xb1, 0xe7, 0x9b, 0x7f, 0x12, 0xd8, 0xd7, 0xdd, 0x52,
	0x49, 0xfb, 0x9d, 0xcc, 0xb1, 0x36, 0xad, 0x6b, 0x71, 0x7a, 0x2b, 0x94, 0x3f, 0x19, 0x3c, 0x7c,
	0xd7, 0x59, 0xd9, 0x90, 0xb5, 0x0b, 0xb2, 0x0e, 0xf7, 0x20, 0x93, 0x35, 0x67, 0x05, 0x9b, 0x8f,
	0x44, 0x26, 0x6b, 0x7c, 0x02, 0x63, 0x4b, 0x2b, 0x43, 0x8e, 0x67, 0x05, 0x9b, 0x4f, 0x45, 0x1f,
	0xe1, 0x0c, 0x72, 0x47, 0xd6, 0x2d, 0xae, 0x35, 0xf1, 0x51, 0xc8, 0xdc, 0xc6, 0x78, 0x08, 0xe0,
	0xcf, 0x9f, 0xb5, 0x93, 0x6d, 0xc3, 0xef, 0x87, 0xb7, 0x12, 0xc5, 0xe7, 0x3b, 0x5d, 0x57, 0x8e,
	0x16, 0xf2, 0x92, 0xf8, 0x83, 0x70, 0x3b, 0x51, 0xca, 0x5f, 0x19, 0xe4, 0x11, 0x2a, 0x01, 0x60,
	0x6b, 0x00, 0x05, 0xec, 0xdc, 0x9c, 0xbe, 0x56, 0xaa, 0xa3, 0x9e, 0x2e, 0x95, 0x90, 0xc3, 0xa4,
	0xd2, 0x3f, 0x3e, 0x55, 0x97, 0x91, 0x30, 0x86, 0x78, 0x04, 0xbb, 0xba, 0x53, 0xea, 0xcc, 0xd0,
	0x55, 0x47, 0xcd, 0xea, 0xba, 0x67, 0x5c, 0x17, 0xf1, 0x15, 0xec, 0xad, 0x3a, 0x63, 0xa8, 0x71,
	0xef, 0x6b, 0x19, 0x46, 0xf1, 0xa8, 0x99, 0x18, 0xa8, 0xbe, 0xcf, 0x37, 0xa9, 0xe8, 0xc2, 0x28,
	0x3e, 0xbe, 0xe9, 0xd3, 0x87, 0x9e, 0x5d, 0xda, 0x73, 0xd7, 0x6a, 0x3e, 0x29, 0xd8, 0x3c, 0x17,
	0x7d, 0x34, 0x30, 0x20, 0x1f, 0x1a, 0x80, 0xaf, 0x61, 0x77, 0x99, 0x2c, 0xc5, 0xf2, 0x69, 0x31,
	0x9a, 0xef, 0x9c, 0x3c, 0x3d, 0xfe, 0xbb, 0xc9, 0x74, 0x69, 0x62, 0xbd, 0xba, 0x7c, 0x03, 0x78,
	0x26, 0x9b, 0xfa, 0xad, 0x52, 0xb1, 0x4a, 0xd0, 0x15, 0x3e, 0x82, 0x51, 0xa5, 0x54, 0x70, 0x31,
	0x17, 0xfe, 0x98, 0xe0, 0x65, 0x29, 0x5e, 0xf9, 0x01, 0x0e, 0xee, 0xdc, 0xb7, 0x1a, 0x4f, 0x01,
	0x62, 0x1f, 0xb2, 0x9c, 0x05, 0xa4, 0x83, 0x0d, 0x48, 0x22, 0x29, 0x2b, 0x3f, 0x02, 0xbf, 0x08,
	0x83, 0xc5, 0x6c, 0xef, 0x9a, 0x27, 0xe2, 0x30, 0xa1, 0xde, 0x59, 0x16, 0x9c, 0x8d, 0xe1, 0xb6,
	0xaf, 0xae, 0x7c, 0x0e, 0xcf, 0xb6, 0xbc, 0x66, 0xf5, 0xc9, 0x6f, 0x06, 0x78, 0xde, 0x69, 0xdd,
	0x1a, 0x47, 0xe6, 0x4b, 0xc4, 0x42, 0x01, 0xfb, 0x83, 0x69, 0xf0, 0x45, 0x42, 0x7d, 0xd7, 0xa9,
	0xd9, 0xe1, 0xbf, 0xd2, 0x56, 0x97, 0xf7, 0xb0, 0x86, 0xc7, 0x1b, 0x39, 0xf0, 0x65, 0x72, 0x75,
	0xdb, 0xdc, 0xb3, 0xa3, 0xff, 0x17, 0xf9, 0x2e, 0xcb, 0x71, 0xf8, 0x5d, 0x4f, 0xff, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xcc, 0x55, 0x39, 0x1c, 0xc1, 0x03, 0x00, 0x00,
}
