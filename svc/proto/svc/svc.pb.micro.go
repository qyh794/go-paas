// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/svc/svc.proto

package svc

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Svc service

func NewSvcEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Svc service

type SvcService interface {
	AddSvc(ctx context.Context, in *RSvcInfo, opts ...client.CallOption) (*Response, error)
	DeleteSvcByID(ctx context.Context, in *RequestSvcID, opts ...client.CallOption) (*Response, error)
	UpdateSvc(ctx context.Context, in *RSvcInfo, opts ...client.CallOption) (*Response, error)
	QuerySvcByID(ctx context.Context, in *RequestSvcID, opts ...client.CallOption) (*RSvcInfo, error)
	QueryAll(ctx context.Context, in *RequestQueryAll, opts ...client.CallOption) (*ResponseAllSvc, error)
}

type svcService struct {
	c    client.Client
	name string
}

func NewSvcService(name string, c client.Client) SvcService {
	return &svcService{
		c:    c,
		name: name,
	}
}

func (c *svcService) AddSvc(ctx context.Context, in *RSvcInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Svc.AddSvc", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcService) DeleteSvcByID(ctx context.Context, in *RequestSvcID, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Svc.DeleteSvcByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcService) UpdateSvc(ctx context.Context, in *RSvcInfo, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Svc.UpdateSvc", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcService) QuerySvcByID(ctx context.Context, in *RequestSvcID, opts ...client.CallOption) (*RSvcInfo, error) {
	req := c.c.NewRequest(c.name, "Svc.QuerySvcByID", in)
	out := new(RSvcInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *svcService) QueryAll(ctx context.Context, in *RequestQueryAll, opts ...client.CallOption) (*ResponseAllSvc, error) {
	req := c.c.NewRequest(c.name, "Svc.QueryAll", in)
	out := new(ResponseAllSvc)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Svc service

type SvcHandler interface {
	AddSvc(context.Context, *RSvcInfo, *Response) error
	DeleteSvcByID(context.Context, *RequestSvcID, *Response) error
	UpdateSvc(context.Context, *RSvcInfo, *Response) error
	QuerySvcByID(context.Context, *RequestSvcID, *RSvcInfo) error
	QueryAll(context.Context, *RequestQueryAll, *ResponseAllSvc) error
}

func RegisterSvcHandler(s server.Server, hdlr SvcHandler, opts ...server.HandlerOption) error {
	type svc interface {
		AddSvc(ctx context.Context, in *RSvcInfo, out *Response) error
		DeleteSvcByID(ctx context.Context, in *RequestSvcID, out *Response) error
		UpdateSvc(ctx context.Context, in *RSvcInfo, out *Response) error
		QuerySvcByID(ctx context.Context, in *RequestSvcID, out *RSvcInfo) error
		QueryAll(ctx context.Context, in *RequestQueryAll, out *ResponseAllSvc) error
	}
	type Svc struct {
		svc
	}
	h := &svcHandler{hdlr}
	return s.Handle(s.NewHandler(&Svc{h}, opts...))
}

type svcHandler struct {
	SvcHandler
}

func (h *svcHandler) AddSvc(ctx context.Context, in *RSvcInfo, out *Response) error {
	return h.SvcHandler.AddSvc(ctx, in, out)
}

func (h *svcHandler) DeleteSvcByID(ctx context.Context, in *RequestSvcID, out *Response) error {
	return h.SvcHandler.DeleteSvcByID(ctx, in, out)
}

func (h *svcHandler) UpdateSvc(ctx context.Context, in *RSvcInfo, out *Response) error {
	return h.SvcHandler.UpdateSvc(ctx, in, out)
}

func (h *svcHandler) QuerySvcByID(ctx context.Context, in *RequestSvcID, out *RSvcInfo) error {
	return h.SvcHandler.QuerySvcByID(ctx, in, out)
}

func (h *svcHandler) QueryAll(ctx context.Context, in *RequestQueryAll, out *ResponseAllSvc) error {
	return h.SvcHandler.QueryAll(ctx, in, out)
}
