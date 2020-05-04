// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: manager.proto

package manager

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Manager service

type ManagerService interface {
	CreateTask(ctx context.Context, in *CreateTaskReq, opts ...client.CallOption) (*CreateTaskResp, error)
}

type managerService struct {
	c    client.Client
	name string
}

func NewManagerService(name string, c client.Client) ManagerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "manager"
	}
	return &managerService{
		c:    c,
		name: name,
	}
}

func (c *managerService) CreateTask(ctx context.Context, in *CreateTaskReq, opts ...client.CallOption) (*CreateTaskResp, error) {
	req := c.c.NewRequest(c.name, "Manager.CreateTask", in)
	out := new(CreateTaskResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Manager service

type ManagerHandler interface {
	CreateTask(context.Context, *CreateTaskReq, *CreateTaskResp) error
}

func RegisterManagerHandler(s server.Server, hdlr ManagerHandler, opts ...server.HandlerOption) error {
	type manager interface {
		CreateTask(ctx context.Context, in *CreateTaskReq, out *CreateTaskResp) error
	}
	type Manager struct {
		manager
	}
	h := &managerHandler{hdlr}
	return s.Handle(s.NewHandler(&Manager{h}, opts...))
}

type managerHandler struct {
	ManagerHandler
}

func (h *managerHandler) CreateTask(ctx context.Context, in *CreateTaskReq, out *CreateTaskResp) error {
	return h.ManagerHandler.CreateTask(ctx, in, out)
}
