// Code generated by protoc-gen-go. DO NOT EDIT.
// source: executor.proto

package executor

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

// 任务结构体
type Task struct {
	TaskId               int64    `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	TaskType             string   `protobuf:"bytes,2,opt,name=TaskType,proto3" json:"TaskType,omitempty"`
	TaskFile             string   `protobuf:"bytes,3,opt,name=TaskFile,proto3" json:"TaskFile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetTaskId() int64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *Task) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *Task) GetTaskFile() string {
	if m != nil {
		return m.TaskFile
	}
	return ""
}

// 手机结构体
type Phone struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Client               string   `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	TestType             string   `protobuf:"bytes,3,opt,name=testType,proto3" json:"testType,omitempty"`
	Secret               string   `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
	UpdateTime           string   `protobuf:"bytes,5,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Phone) Reset()         { *m = Phone{} }
func (m *Phone) String() string { return proto.CompactTextString(m) }
func (*Phone) ProtoMessage()    {}
func (*Phone) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{1}
}

func (m *Phone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Phone.Unmarshal(m, b)
}
func (m *Phone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Phone.Marshal(b, m, deterministic)
}
func (m *Phone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Phone.Merge(m, src)
}
func (m *Phone) XXX_Size() int {
	return xxx_messageInfo_Phone.Size(m)
}
func (m *Phone) XXX_DiscardUnknown() {
	xxx_messageInfo_Phone.DiscardUnknown(m)
}

var xxx_messageInfo_Phone proto.InternalMessageInfo

func (m *Phone) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Phone) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

func (m *Phone) GetTestType() string {
	if m != nil {
		return m.TestType
	}
	return ""
}

func (m *Phone) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Phone) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

// 客户端结构体
type Client struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	TestType             string   `protobuf:"bytes,2,opt,name=testType,proto3" json:"testType,omitempty"`
	Secret               string   `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	UpdateTime           string   `protobuf:"bytes,4,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{2}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Client) GetTestType() string {
	if m != nil {
		return m.TestType
	}
	return ""
}

func (m *Client) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Client) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

// 获取任务 key
type FindTaskKeyReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindTaskKeyReq) Reset()         { *m = FindTaskKeyReq{} }
func (m *FindTaskKeyReq) String() string { return proto.CompactTextString(m) }
func (*FindTaskKeyReq) ProtoMessage()    {}
func (*FindTaskKeyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{3}
}

func (m *FindTaskKeyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindTaskKeyReq.Unmarshal(m, b)
}
func (m *FindTaskKeyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindTaskKeyReq.Marshal(b, m, deterministic)
}
func (m *FindTaskKeyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindTaskKeyReq.Merge(m, src)
}
func (m *FindTaskKeyReq) XXX_Size() int {
	return xxx_messageInfo_FindTaskKeyReq.Size(m)
}
func (m *FindTaskKeyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FindTaskKeyReq.DiscardUnknown(m)
}

var xxx_messageInfo_FindTaskKeyReq proto.InternalMessageInfo

type FindTaskKeyResp struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindTaskKeyResp) Reset()         { *m = FindTaskKeyResp{} }
func (m *FindTaskKeyResp) String() string { return proto.CompactTextString(m) }
func (*FindTaskKeyResp) ProtoMessage()    {}
func (*FindTaskKeyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{4}
}

func (m *FindTaskKeyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindTaskKeyResp.Unmarshal(m, b)
}
func (m *FindTaskKeyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindTaskKeyResp.Marshal(b, m, deterministic)
}
func (m *FindTaskKeyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindTaskKeyResp.Merge(m, src)
}
func (m *FindTaskKeyResp) XXX_Size() int {
	return xxx_messageInfo_FindTaskKeyResp.Size(m)
}
func (m *FindTaskKeyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FindTaskKeyResp.DiscardUnknown(m)
}

var xxx_messageInfo_FindTaskKeyResp proto.InternalMessageInfo

func (m *FindTaskKeyResp) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// 获取任务
type FindTaskByKeyReq struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindTaskByKeyReq) Reset()         { *m = FindTaskByKeyReq{} }
func (m *FindTaskByKeyReq) String() string { return proto.CompactTextString(m) }
func (*FindTaskByKeyReq) ProtoMessage()    {}
func (*FindTaskByKeyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{5}
}

func (m *FindTaskByKeyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindTaskByKeyReq.Unmarshal(m, b)
}
func (m *FindTaskByKeyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindTaskByKeyReq.Marshal(b, m, deterministic)
}
func (m *FindTaskByKeyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindTaskByKeyReq.Merge(m, src)
}
func (m *FindTaskByKeyReq) XXX_Size() int {
	return xxx_messageInfo_FindTaskByKeyReq.Size(m)
}
func (m *FindTaskByKeyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FindTaskByKeyReq.DiscardUnknown(m)
}

var xxx_messageInfo_FindTaskByKeyReq proto.InternalMessageInfo

func (m *FindTaskByKeyReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type FindTaskByKeyResp struct {
	Task                 *Task    `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindTaskByKeyResp) Reset()         { *m = FindTaskByKeyResp{} }
func (m *FindTaskByKeyResp) String() string { return proto.CompactTextString(m) }
func (*FindTaskByKeyResp) ProtoMessage()    {}
func (*FindTaskByKeyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{6}
}

func (m *FindTaskByKeyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindTaskByKeyResp.Unmarshal(m, b)
}
func (m *FindTaskByKeyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindTaskByKeyResp.Marshal(b, m, deterministic)
}
func (m *FindTaskByKeyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindTaskByKeyResp.Merge(m, src)
}
func (m *FindTaskByKeyResp) XXX_Size() int {
	return xxx_messageInfo_FindTaskByKeyResp.Size(m)
}
func (m *FindTaskByKeyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FindTaskByKeyResp.DiscardUnknown(m)
}

var xxx_messageInfo_FindTaskByKeyResp proto.InternalMessageInfo

func (m *FindTaskByKeyResp) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

// 删除任务
type DeleteTaskByKeyReq struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskByKeyReq) Reset()         { *m = DeleteTaskByKeyReq{} }
func (m *DeleteTaskByKeyReq) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskByKeyReq) ProtoMessage()    {}
func (*DeleteTaskByKeyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{7}
}

func (m *DeleteTaskByKeyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskByKeyReq.Unmarshal(m, b)
}
func (m *DeleteTaskByKeyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskByKeyReq.Marshal(b, m, deterministic)
}
func (m *DeleteTaskByKeyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskByKeyReq.Merge(m, src)
}
func (m *DeleteTaskByKeyReq) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskByKeyReq.Size(m)
}
func (m *DeleteTaskByKeyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskByKeyReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskByKeyReq proto.InternalMessageInfo

func (m *DeleteTaskByKeyReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DeleteTaskByKeyResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskByKeyResp) Reset()         { *m = DeleteTaskByKeyResp{} }
func (m *DeleteTaskByKeyResp) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskByKeyResp) ProtoMessage()    {}
func (*DeleteTaskByKeyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{8}
}

func (m *DeleteTaskByKeyResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskByKeyResp.Unmarshal(m, b)
}
func (m *DeleteTaskByKeyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskByKeyResp.Marshal(b, m, deterministic)
}
func (m *DeleteTaskByKeyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskByKeyResp.Merge(m, src)
}
func (m *DeleteTaskByKeyResp) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskByKeyResp.Size(m)
}
func (m *DeleteTaskByKeyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskByKeyResp.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskByKeyResp proto.InternalMessageInfo

// 推送任务
type PushTestTaskReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushTestTaskReq) Reset()         { *m = PushTestTaskReq{} }
func (m *PushTestTaskReq) String() string { return proto.CompactTextString(m) }
func (*PushTestTaskReq) ProtoMessage()    {}
func (*PushTestTaskReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{9}
}

func (m *PushTestTaskReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushTestTaskReq.Unmarshal(m, b)
}
func (m *PushTestTaskReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushTestTaskReq.Marshal(b, m, deterministic)
}
func (m *PushTestTaskReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushTestTaskReq.Merge(m, src)
}
func (m *PushTestTaskReq) XXX_Size() int {
	return xxx_messageInfo_PushTestTaskReq.Size(m)
}
func (m *PushTestTaskReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PushTestTaskReq.DiscardUnknown(m)
}

var xxx_messageInfo_PushTestTaskReq proto.InternalMessageInfo

type PushTestTaskResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushTestTaskResp) Reset()         { *m = PushTestTaskResp{} }
func (m *PushTestTaskResp) String() string { return proto.CompactTextString(m) }
func (*PushTestTaskResp) ProtoMessage()    {}
func (*PushTestTaskResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{10}
}

func (m *PushTestTaskResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushTestTaskResp.Unmarshal(m, b)
}
func (m *PushTestTaskResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushTestTaskResp.Marshal(b, m, deterministic)
}
func (m *PushTestTaskResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushTestTaskResp.Merge(m, src)
}
func (m *PushTestTaskResp) XXX_Size() int {
	return xxx_messageInfo_PushTestTaskResp.Size(m)
}
func (m *PushTestTaskResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PushTestTaskResp.DiscardUnknown(m)
}

var xxx_messageInfo_PushTestTaskResp proto.InternalMessageInfo

// 存入手机信息
type InsertPhoneReq struct {
	Phones               []*Phone `protobuf:"bytes,1,rep,name=phones,proto3" json:"phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsertPhoneReq) Reset()         { *m = InsertPhoneReq{} }
func (m *InsertPhoneReq) String() string { return proto.CompactTextString(m) }
func (*InsertPhoneReq) ProtoMessage()    {}
func (*InsertPhoneReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{11}
}

func (m *InsertPhoneReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertPhoneReq.Unmarshal(m, b)
}
func (m *InsertPhoneReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertPhoneReq.Marshal(b, m, deterministic)
}
func (m *InsertPhoneReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertPhoneReq.Merge(m, src)
}
func (m *InsertPhoneReq) XXX_Size() int {
	return xxx_messageInfo_InsertPhoneReq.Size(m)
}
func (m *InsertPhoneReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertPhoneReq.DiscardUnknown(m)
}

var xxx_messageInfo_InsertPhoneReq proto.InternalMessageInfo

func (m *InsertPhoneReq) GetPhones() []*Phone {
	if m != nil {
		return m.Phones
	}
	return nil
}

type InsertPhoneResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsertPhoneResp) Reset()         { *m = InsertPhoneResp{} }
func (m *InsertPhoneResp) String() string { return proto.CompactTextString(m) }
func (*InsertPhoneResp) ProtoMessage()    {}
func (*InsertPhoneResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{12}
}

func (m *InsertPhoneResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertPhoneResp.Unmarshal(m, b)
}
func (m *InsertPhoneResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertPhoneResp.Marshal(b, m, deterministic)
}
func (m *InsertPhoneResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertPhoneResp.Merge(m, src)
}
func (m *InsertPhoneResp) XXX_Size() int {
	return xxx_messageInfo_InsertPhoneResp.Size(m)
}
func (m *InsertPhoneResp) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertPhoneResp.DiscardUnknown(m)
}

var xxx_messageInfo_InsertPhoneResp proto.InternalMessageInfo

// 根据 ip 获取客户端
type FindClientByIpReq struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindClientByIpReq) Reset()         { *m = FindClientByIpReq{} }
func (m *FindClientByIpReq) String() string { return proto.CompactTextString(m) }
func (*FindClientByIpReq) ProtoMessage()    {}
func (*FindClientByIpReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{13}
}

func (m *FindClientByIpReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindClientByIpReq.Unmarshal(m, b)
}
func (m *FindClientByIpReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindClientByIpReq.Marshal(b, m, deterministic)
}
func (m *FindClientByIpReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindClientByIpReq.Merge(m, src)
}
func (m *FindClientByIpReq) XXX_Size() int {
	return xxx_messageInfo_FindClientByIpReq.Size(m)
}
func (m *FindClientByIpReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FindClientByIpReq.DiscardUnknown(m)
}

var xxx_messageInfo_FindClientByIpReq proto.InternalMessageInfo

func (m *FindClientByIpReq) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type FindClientByIpResp struct {
	Client               *Client  `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindClientByIpResp) Reset()         { *m = FindClientByIpResp{} }
func (m *FindClientByIpResp) String() string { return proto.CompactTextString(m) }
func (*FindClientByIpResp) ProtoMessage()    {}
func (*FindClientByIpResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{14}
}

func (m *FindClientByIpResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindClientByIpResp.Unmarshal(m, b)
}
func (m *FindClientByIpResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindClientByIpResp.Marshal(b, m, deterministic)
}
func (m *FindClientByIpResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindClientByIpResp.Merge(m, src)
}
func (m *FindClientByIpResp) XXX_Size() int {
	return xxx_messageInfo_FindClientByIpResp.Size(m)
}
func (m *FindClientByIpResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FindClientByIpResp.DiscardUnknown(m)
}

var xxx_messageInfo_FindClientByIpResp proto.InternalMessageInfo

func (m *FindClientByIpResp) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func init() {
	proto.RegisterType((*Task)(nil), "executor.Task")
	proto.RegisterType((*Phone)(nil), "executor.Phone")
	proto.RegisterType((*Client)(nil), "executor.Client")
	proto.RegisterType((*FindTaskKeyReq)(nil), "executor.FindTaskKeyReq")
	proto.RegisterType((*FindTaskKeyResp)(nil), "executor.FindTaskKeyResp")
	proto.RegisterType((*FindTaskByKeyReq)(nil), "executor.FindTaskByKeyReq")
	proto.RegisterType((*FindTaskByKeyResp)(nil), "executor.FindTaskByKeyResp")
	proto.RegisterType((*DeleteTaskByKeyReq)(nil), "executor.DeleteTaskByKeyReq")
	proto.RegisterType((*DeleteTaskByKeyResp)(nil), "executor.DeleteTaskByKeyResp")
	proto.RegisterType((*PushTestTaskReq)(nil), "executor.PushTestTaskReq")
	proto.RegisterType((*PushTestTaskResp)(nil), "executor.PushTestTaskResp")
	proto.RegisterType((*InsertPhoneReq)(nil), "executor.InsertPhoneReq")
	proto.RegisterType((*InsertPhoneResp)(nil), "executor.InsertPhoneResp")
	proto.RegisterType((*FindClientByIpReq)(nil), "executor.FindClientByIpReq")
	proto.RegisterType((*FindClientByIpResp)(nil), "executor.FindClientByIpResp")
}

func init() { proto.RegisterFile("executor.proto", fileDescriptor_12d1cdcda51e000f) }

var fileDescriptor_12d1cdcda51e000f = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xe3, 0x34, 0x6a, 0x27, 0xe0, 0x38, 0x83, 0x40, 0xc6, 0x14, 0x14, 0x6d, 0x11, 0xe4,
	0xd4, 0x43, 0x38, 0x20, 0x2e, 0x1c, 0x4a, 0x28, 0x0a, 0x08, 0x29, 0x32, 0x11, 0xf7, 0x90, 0x8c,
	0x54, 0x2b, 0x21, 0xde, 0x7a, 0x37, 0x12, 0x3e, 0xf3, 0x1b, 0xf9, 0x3f, 0x68, 0x76, 0xed, 0xf8,
	0x23, 0xb8, 0xdc, 0x3c, 0xfb, 0xde, 0xbc, 0xb7, 0x3b, 0x6f, 0x12, 0xf0, 0xe8, 0x17, 0xad, 0xf6,
	0x3a, 0x49, 0xaf, 0x64, 0x9a, 0xe8, 0x04, 0xcf, 0x8a, 0x5a, 0x7c, 0x87, 0xee, 0x62, 0xa9, 0x36,
	0xf8, 0x04, 0x7a, 0x7a, 0xa9, 0x36, 0xb3, 0x75, 0xe0, 0x8c, 0x9c, 0xb1, 0x1b, 0xe5, 0x15, 0x86,
	0x70, 0xc6, 0xf8, 0x22, 0x93, 0x14, 0x74, 0x46, 0xce, 0xf8, 0x3c, 0x3a, 0xd4, 0x05, 0x76, 0x13,
	0x6f, 0x29, 0x70, 0x4b, 0x8c, 0x6b, 0xf1, 0xdb, 0x81, 0xd3, 0xf9, 0x6d, 0xb2, 0x23, 0xf4, 0xa0,
	0x13, 0x5b, 0xd5, 0xf3, 0xa8, 0x13, 0xaf, 0xd9, 0x69, 0xb5, 0x8d, 0x69, 0xa7, 0x73, 0xbd, 0xbc,
	0x62, 0x35, 0x4d, 0x4a, 0x1b, 0xa7, 0x5c, 0xad, 0xa8, 0xb9, 0x47, 0xd1, 0x2a, 0x25, 0x1d, 0x74,
	0x6d, 0x8f, 0xad, 0xf0, 0x05, 0xc0, 0x5e, 0xae, 0x97, 0x9a, 0x16, 0xf1, 0x4f, 0x0a, 0x4e, 0x0d,
	0x56, 0x39, 0x11, 0x5b, 0xe8, 0x7d, 0xb0, 0xea, 0x7c, 0x0b, 0x79, 0xb8, 0x85, 0xac, 0xb9, 0x75,
	0x5a, 0xdd, 0xdc, 0x7b, 0xdc, 0xba, 0x47, 0x6e, 0x3e, 0x78, 0x37, 0xf1, 0x6e, 0xcd, 0x33, 0xf8,
	0x42, 0x59, 0x44, 0x77, 0xe2, 0x12, 0x06, 0xb5, 0x13, 0x25, 0xd1, 0x07, 0x77, 0x43, 0x59, 0x7e,
	0x13, 0xfe, 0x14, 0x2f, 0xc1, 0x2f, 0x48, 0xd7, 0x99, 0x6d, 0xfc, 0x07, 0xeb, 0x2d, 0x0c, 0x1b,
	0x2c, 0x25, 0x51, 0x40, 0x97, 0x73, 0x32, 0xbc, 0xfe, 0xc4, 0xbb, 0x3a, 0xc4, 0xcc, 0xb4, 0xc8,
	0x60, 0xe2, 0x15, 0xe0, 0x94, 0xb6, 0xa4, 0xe9, 0x3f, 0x06, 0x8f, 0xe1, 0xd1, 0x11, 0x4f, 0x49,
	0x31, 0x84, 0xc1, 0x7c, 0xaf, 0x6e, 0x17, 0x3c, 0x1c, 0x16, 0xa5, 0x3b, 0x81, 0xe0, 0xd7, 0x8f,
	0x94, 0x14, 0xef, 0xc0, 0x9b, 0xed, 0x14, 0xa5, 0xda, 0x84, 0xce, 0x0e, 0xaf, 0xa1, 0x27, 0xf9,
	0x5b, 0x05, 0xce, 0xc8, 0x1d, 0xf7, 0x27, 0x83, 0xf2, 0x76, 0x96, 0x93, 0xc3, 0xec, 0x50, 0x6b,
	0x55, 0x52, 0x5c, 0xda, 0xc7, 0xda, 0xec, 0xae, 0xb3, 0x99, 0x64, 0xc1, 0x46, 0x84, 0xe2, 0x3d,
	0x60, 0x93, 0xa4, 0x24, 0x8e, 0x0f, 0xeb, 0x65, 0x87, 0xe2, 0x97, 0xb6, 0x96, 0x59, 0x2c, 0xdc,
	0xe4, 0x8f, 0x0b, 0xc3, 0x6f, 0x7b, 0x29, 0x93, 0x54, 0x53, 0xfa, 0x31, 0x27, 0xe1, 0x14, 0xfa,
	0x95, 0xc8, 0x30, 0x28, 0xdb, 0xeb, 0xd9, 0x86, 0x4f, 0x5b, 0x10, 0x25, 0xc5, 0x09, 0x7e, 0x86,
	0x87, 0xb5, 0xb4, 0x30, 0x3c, 0x66, 0x17, 0x59, 0x84, 0xcf, 0x5a, 0x31, 0xa3, 0x35, 0x87, 0x41,
	0x23, 0x18, 0xbc, 0x28, 0x3b, 0x8e, 0xb3, 0x0d, 0x9f, 0xdf, 0x83, 0x1a, 0xc5, 0x4f, 0xf0, 0xa0,
	0x1a, 0x20, 0x56, 0x9e, 0xd2, 0xc8, 0x3a, 0x0c, 0xdb, 0x20, 0x23, 0x34, 0x85, 0x7e, 0x25, 0xba,
	0xea, 0xb0, 0xea, 0xcb, 0x50, 0x1d, 0x56, 0x33, 0xeb, 0x13, 0xfc, 0x6a, 0x7f, 0x37, 0x65, 0x90,
	0xd8, 0x98, 0x48, 0x6d, 0x0f, 0xc2, 0x8b, 0x76, 0x90, 0xe5, 0x7e, 0xf4, 0xcc, 0x7f, 0xdc, 0x9b,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x43, 0x8f, 0x99, 0x2d, 0xf5, 0x04, 0x00, 0x00,
}
