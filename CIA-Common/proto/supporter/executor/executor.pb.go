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

// 任务结构体（Redis）
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
	PhoneType            string   `protobuf:"bytes,5,opt,name=phoneType,proto3" json:"phoneType,omitempty"`
	PhoneEdition         string   `protobuf:"bytes,6,opt,name=phoneEdition,proto3" json:"phoneEdition,omitempty"`
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

func (m *Phone) GetPhoneType() string {
	if m != nil {
		return m.PhoneType
	}
	return ""
}

func (m *Phone) GetPhoneEdition() string {
	if m != nil {
		return m.PhoneEdition
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
	Secret               string   `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
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

func (m *FindTaskKeyReq) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

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
	return fileDescriptor_12d1cdcda51e000f, []int{9}
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
	return fileDescriptor_12d1cdcda51e000f, []int{10}
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
	return fileDescriptor_12d1cdcda51e000f, []int{11}
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
	return fileDescriptor_12d1cdcda51e000f, []int{12}
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

// 更新任务开始信息
type UpdateTaskStartByIdReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Client               string   `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTaskStartByIdReq) Reset()         { *m = UpdateTaskStartByIdReq{} }
func (m *UpdateTaskStartByIdReq) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskStartByIdReq) ProtoMessage()    {}
func (*UpdateTaskStartByIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{13}
}

func (m *UpdateTaskStartByIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskStartByIdReq.Unmarshal(m, b)
}
func (m *UpdateTaskStartByIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskStartByIdReq.Marshal(b, m, deterministic)
}
func (m *UpdateTaskStartByIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskStartByIdReq.Merge(m, src)
}
func (m *UpdateTaskStartByIdReq) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskStartByIdReq.Size(m)
}
func (m *UpdateTaskStartByIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskStartByIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskStartByIdReq proto.InternalMessageInfo

func (m *UpdateTaskStartByIdReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateTaskStartByIdReq) GetClient() string {
	if m != nil {
		return m.Client
	}
	return ""
}

type UpdateTaskStartByIdResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTaskStartByIdResp) Reset()         { *m = UpdateTaskStartByIdResp{} }
func (m *UpdateTaskStartByIdResp) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskStartByIdResp) ProtoMessage()    {}
func (*UpdateTaskStartByIdResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{14}
}

func (m *UpdateTaskStartByIdResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskStartByIdResp.Unmarshal(m, b)
}
func (m *UpdateTaskStartByIdResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskStartByIdResp.Marshal(b, m, deterministic)
}
func (m *UpdateTaskStartByIdResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskStartByIdResp.Merge(m, src)
}
func (m *UpdateTaskStartByIdResp) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskStartByIdResp.Size(m)
}
func (m *UpdateTaskStartByIdResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskStartByIdResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskStartByIdResp proto.InternalMessageInfo

// 更新任务结束信息
type UpdateTaskEndByIdReq struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsSuccess            bool     `protobuf:"varint,2,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
	ResultDesc           string   `protobuf:"bytes,3,opt,name=resultDesc,proto3" json:"resultDesc,omitempty"`
	ResultLocation       string   `protobuf:"bytes,4,opt,name=resultLocation,proto3" json:"resultLocation,omitempty"`
	ResultImageLocation  string   `protobuf:"bytes,5,opt,name=resultImageLocation,proto3" json:"resultImageLocation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTaskEndByIdReq) Reset()         { *m = UpdateTaskEndByIdReq{} }
func (m *UpdateTaskEndByIdReq) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskEndByIdReq) ProtoMessage()    {}
func (*UpdateTaskEndByIdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{15}
}

func (m *UpdateTaskEndByIdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskEndByIdReq.Unmarshal(m, b)
}
func (m *UpdateTaskEndByIdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskEndByIdReq.Marshal(b, m, deterministic)
}
func (m *UpdateTaskEndByIdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskEndByIdReq.Merge(m, src)
}
func (m *UpdateTaskEndByIdReq) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskEndByIdReq.Size(m)
}
func (m *UpdateTaskEndByIdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskEndByIdReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskEndByIdReq proto.InternalMessageInfo

func (m *UpdateTaskEndByIdReq) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateTaskEndByIdReq) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func (m *UpdateTaskEndByIdReq) GetResultDesc() string {
	if m != nil {
		return m.ResultDesc
	}
	return ""
}

func (m *UpdateTaskEndByIdReq) GetResultLocation() string {
	if m != nil {
		return m.ResultLocation
	}
	return ""
}

func (m *UpdateTaskEndByIdReq) GetResultImageLocation() string {
	if m != nil {
		return m.ResultImageLocation
	}
	return ""
}

type UpdateTaskEndByIdResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTaskEndByIdResp) Reset()         { *m = UpdateTaskEndByIdResp{} }
func (m *UpdateTaskEndByIdResp) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskEndByIdResp) ProtoMessage()    {}
func (*UpdateTaskEndByIdResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d1cdcda51e000f, []int{16}
}

func (m *UpdateTaskEndByIdResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskEndByIdResp.Unmarshal(m, b)
}
func (m *UpdateTaskEndByIdResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskEndByIdResp.Marshal(b, m, deterministic)
}
func (m *UpdateTaskEndByIdResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskEndByIdResp.Merge(m, src)
}
func (m *UpdateTaskEndByIdResp) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskEndByIdResp.Size(m)
}
func (m *UpdateTaskEndByIdResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskEndByIdResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskEndByIdResp proto.InternalMessageInfo

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
	proto.RegisterType((*InsertPhoneReq)(nil), "executor.InsertPhoneReq")
	proto.RegisterType((*InsertPhoneResp)(nil), "executor.InsertPhoneResp")
	proto.RegisterType((*FindClientByIpReq)(nil), "executor.FindClientByIpReq")
	proto.RegisterType((*FindClientByIpResp)(nil), "executor.FindClientByIpResp")
	proto.RegisterType((*UpdateTaskStartByIdReq)(nil), "executor.UpdateTaskStartByIdReq")
	proto.RegisterType((*UpdateTaskStartByIdResp)(nil), "executor.UpdateTaskStartByIdResp")
	proto.RegisterType((*UpdateTaskEndByIdReq)(nil), "executor.UpdateTaskEndByIdReq")
	proto.RegisterType((*UpdateTaskEndByIdResp)(nil), "executor.UpdateTaskEndByIdResp")
}

func init() { proto.RegisterFile("executor.proto", fileDescriptor_12d1cdcda51e000f) }

var fileDescriptor_12d1cdcda51e000f = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0x4f, 0x6f, 0xda, 0x4e,
	0x10, 0x8d, 0x81, 0xa0, 0x30, 0xfc, 0x7e, 0x26, 0x6c, 0x9a, 0xc4, 0x71, 0x69, 0x4a, 0x37, 0x55,
	0xca, 0x29, 0xaa, 0xe8, 0xa1, 0xea, 0xa5, 0xaa, 0x52, 0x88, 0x44, 0xff, 0x48, 0x91, 0xa1, 0x39,
	0xf5, 0xe2, 0xda, 0xa3, 0xd6, 0x82, 0xd8, 0x5b, 0xef, 0x22, 0x95, 0x0f, 0xd4, 0xaf, 0xd1, 0x43,
	0x3f, 0x59, 0xb5, 0xbb, 0x36, 0xfe, 0x83, 0x21, 0x37, 0x66, 0xde, 0x9b, 0x37, 0x33, 0xbb, 0x6f,
	0x31, 0x98, 0xf8, 0x0b, 0xbd, 0xa5, 0x88, 0xe2, 0x2b, 0x16, 0x47, 0x22, 0x22, 0x07, 0x69, 0x4c,
	0xef, 0xa0, 0x31, 0x73, 0xf9, 0x9c, 0x9c, 0x40, 0x53, 0xb8, 0x7c, 0x3e, 0xf1, 0x2d, 0xa3, 0x6f,
	0x0c, 0xea, 0x4e, 0x12, 0x11, 0x1b, 0x0e, 0x24, 0x3e, 0x5b, 0x31, 0xb4, 0x6a, 0x7d, 0x63, 0xd0,
	0x72, 0xd6, 0x71, 0x8a, 0xdd, 0x04, 0x0b, 0xb4, 0xea, 0x19, 0x26, 0x63, 0xfa, 0xdb, 0x80, 0xfd,
	0xdb, 0x1f, 0x51, 0x88, 0xc4, 0x84, 0x5a, 0xa0, 0x55, 0x5b, 0x4e, 0x2d, 0xf0, 0x65, 0x27, 0x6f,
	0x11, 0x60, 0x28, 0x12, 0xbd, 0x24, 0x92, 0x6a, 0x02, 0xb9, 0x50, 0x9d, 0x12, 0xb5, 0x34, 0x96,
	0x35, 0x1c, 0xbd, 0x18, 0x85, 0xd5, 0xd0, 0x35, 0x3a, 0x22, 0x3d, 0x68, 0x31, 0xd9, 0x44, 0x15,
	0xed, 0x2b, 0x28, 0x4b, 0x10, 0x0a, 0xff, 0xa9, 0x60, 0xec, 0x07, 0x22, 0x88, 0x42, 0xab, 0xa9,
	0x08, 0x85, 0x1c, 0x5d, 0x40, 0xf3, 0xbd, 0xee, 0x2f, 0xe7, 0x64, 0xeb, 0x39, 0x59, 0x61, 0x9e,
	0xda, 0xd6, 0x79, 0xea, 0x85, 0x79, 0xce, 0x01, 0x96, 0xcc, 0x77, 0x05, 0xce, 0x82, 0x7b, 0x4c,
	0x66, 0xcd, 0x65, 0xe8, 0x00, 0xcc, 0x9b, 0x20, 0xf4, 0xe5, 0x29, 0x7d, 0xc4, 0x95, 0x83, 0x3f,
	0x73, 0x4a, 0x46, 0x5e, 0x89, 0x5e, 0x40, 0xa7, 0xc0, 0xe4, 0x8c, 0x1c, 0x42, 0x7d, 0x8e, 0xab,
	0x84, 0x27, 0x7f, 0xd2, 0xe7, 0x70, 0x98, 0x92, 0xae, 0x57, 0x89, 0xe0, 0x26, 0xeb, 0x35, 0x74,
	0x4b, 0x2c, 0xce, 0x08, 0x85, 0x86, 0xbc, 0x61, 0xc5, 0x6b, 0x0f, 0xcd, 0xab, 0xb5, 0x41, 0x24,
	0xcd, 0x51, 0x18, 0xbd, 0x04, 0x32, 0xc2, 0x05, 0x0a, 0x7c, 0xa0, 0xc1, 0x31, 0x1c, 0x6d, 0xf0,
	0x38, 0xa3, 0x6f, 0xc0, 0x9c, 0x84, 0x1c, 0x63, 0xa1, 0x7c, 0x20, 0x4b, 0x5f, 0x40, 0x53, 0x1d,
	0x3e, 0xb7, 0x8c, 0x7e, 0x7d, 0xd0, 0x1e, 0x76, 0xb2, 0xb6, 0x9a, 0x93, 0xc0, 0xb4, 0x0b, 0x9d,
	0x42, 0x29, 0x67, 0xf4, 0x42, 0x6f, 0xa1, 0x2f, 0xeb, 0x7a, 0x35, 0x61, 0x52, 0xb0, 0x74, 0x67,
	0xf4, 0x2d, 0x90, 0x32, 0x89, 0x33, 0x32, 0x58, 0x3b, 0x4e, 0x6f, 0x7b, 0x98, 0xb5, 0xd5, 0xcc,
	0xd4, 0x83, 0xf4, 0x1d, 0x9c, 0x7c, 0xd1, 0xb7, 0xe5, 0xf2, 0xf9, 0x54, 0xb8, 0xb1, 0x14, 0xf1,
	0xd3, 0x4e, 0xe9, 0xdb, 0xd8, 0xe1, 0x62, 0x7a, 0x06, 0xa7, 0x95, 0x0a, 0x9c, 0xd1, 0x3f, 0x06,
	0x3c, 0xca, 0xb0, 0x71, 0xe8, 0x6f, 0xd3, 0xee, 0x41, 0x2b, 0xe0, 0xd3, 0xa5, 0xe7, 0x21, 0xe7,
	0x4a, 0xfe, 0xc0, 0xc9, 0x12, 0xd2, 0x63, 0x31, 0xf2, 0xe5, 0x42, 0x8c, 0x90, 0x7b, 0x89, 0xff,
	0x72, 0x19, 0x72, 0x09, 0xa6, 0x8e, 0x3e, 0x45, 0x9e, 0xab, 0x7c, 0xaf, 0x7d, 0x58, 0xca, 0x92,
	0x97, 0x70, 0xa4, 0x33, 0x93, 0x7b, 0xf7, 0x3b, 0xae, 0xc9, 0xfa, 0x15, 0x55, 0x41, 0xf4, 0x14,
	0x8e, 0x2b, 0xe6, 0xe7, 0x6c, 0xf8, 0xb7, 0x01, 0xdd, 0xe9, 0x92, 0xb1, 0x28, 0x16, 0x18, 0x8f,
	0x93, 0xb3, 0x25, 0x23, 0x68, 0xe7, 0x2c, 0x4c, 0xac, 0xec, 0xd4, 0x8b, 0x6f, 0xc0, 0x3e, 0xdb,
	0x82, 0x70, 0x46, 0xf7, 0xc8, 0x07, 0xf8, 0xbf, 0xe0, 0x5e, 0x62, 0x6f, 0xb2, 0x53, 0x6f, 0xda,
	0x8f, 0xb7, 0x62, 0x4a, 0xeb, 0x16, 0x3a, 0x25, 0xa3, 0x92, 0x5e, 0x56, 0xb1, 0xe9, 0x75, 0xfb,
	0xc9, 0x0e, 0x54, 0x29, 0x8e, 0xa0, 0x9d, 0x33, 0x6a, 0x7e, 0xc7, 0xa2, 0xf5, 0xf3, 0x3b, 0x96,
	0x9d, 0xbd, 0x47, 0x3e, 0xeb, 0xbf, 0x85, 0xcc, 0xb6, 0xa4, 0xb4, 0x48, 0xc1, 0xf5, 0x76, 0x6f,
	0x3b, 0xa8, 0xe4, 0xbe, 0xc2, 0x51, 0x85, 0x07, 0x49, 0x3f, 0x2b, 0xab, 0x36, 0xb9, 0xfd, 0xec,
	0x01, 0x86, 0x52, 0xbf, 0x83, 0xee, 0x86, 0x0b, 0xc8, 0x79, 0x55, 0x65, 0x66, 0x71, 0xfb, 0xe9,
	0x4e, 0x5c, 0xea, 0x7e, 0x6b, 0xaa, 0x4f, 0xd3, 0xab, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8e,
	0xf8, 0xae, 0x5f, 0xac, 0x06, 0x00, 0x00,
}
