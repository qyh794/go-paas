// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/svcApi/svcApi.proto

package svcApi

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for SvcApi service

func NewSvcApiEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SvcApi service

type SvcApiService interface {
	QuerySvcByID(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	AddSvc(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	DeleteSvcByID(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UpdateSvc(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type svcApiService struct {
	c    client.Client
	name string
}

func NewSvcApiService(name string, c client.Client) SvcApiService {
	return &svcApiService{
		c:    c,
		name: name,
	}
}

func (c *svcApiService) QuerySvcByID(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SvcApi.QuerySvcByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcApiService) AddSvc(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SvcApi.AddSvc", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcApiService) DeleteSvcByID(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SvcApi.DeleteSvcByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcApiService) UpdateSvc(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SvcApi.UpdateSvc", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcApiService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "SvcApi.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SvcApi service

type SvcApiHandler interface {
	QuerySvcByID(context.Context, *Request, *Response) error
	AddSvc(context.Context, *Request, *Response) error
	DeleteSvcByID(context.Context, *Request, *Response) error
	UpdateSvc(context.Context, *Request, *Response) error
	Call(context.Context, *Request, *Response) error
}

func RegisterSvcApiHandler(s server.Server, hdlr SvcApiHandler, opts ...server.HandlerOption) error {
	type svcApi interface {
		QuerySvcByID(ctx context.Context, in *Request, out *Response) error
		AddSvc(ctx context.Context, in *Request, out *Response) error
		DeleteSvcByID(ctx context.Context, in *Request, out *Response) error
		UpdateSvc(ctx context.Context, in *Request, out *Response) error
		Call(ctx context.Context, in *Request, out *Response) error
	}
	type SvcApi struct {
		svcApi
	}
	h := &svcApiHandler{hdlr}
	return s.Handle(s.NewHandler(&SvcApi{h}, opts...))
}

type svcApiHandler struct {
	SvcApiHandler
}

func (h *svcApiHandler) QuerySvcByID(ctx context.Context, in *Request, out *Response) error {
	return h.SvcApiHandler.QuerySvcByID(ctx, in, out)
}

func (h *svcApiHandler) AddSvc(ctx context.Context, in *Request, out *Response) error {
	return h.SvcApiHandler.AddSvc(ctx, in, out)
}

func (h *svcApiHandler) DeleteSvcByID(ctx context.Context, in *Request, out *Response) error {
	return h.SvcApiHandler.DeleteSvcByID(ctx, in, out)
}

func (h *svcApiHandler) UpdateSvc(ctx context.Context, in *Request, out *Response) error {
	return h.SvcApiHandler.UpdateSvc(ctx, in, out)
}

func (h *svcApiHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.SvcApiHandler.Call(ctx, in, out)
}
