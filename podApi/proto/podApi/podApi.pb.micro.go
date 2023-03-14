// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/podApi/podApi.proto

package podApi

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

// Api Endpoints for PodApi service

func NewPodApiEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PodApi service

type PodApiService interface {
	QueryPodByID(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	AddPod(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	DeletePodByID(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	UpdatePod(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type podApiService struct {
	c    client.Client
	name string
}

func NewPodApiService(name string, c client.Client) PodApiService {
	return &podApiService{
		c:    c,
		name: name,
	}
}

func (c *podApiService) QueryPodByID(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PodApi.QueryPodByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podApiService) AddPod(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PodApi.AddPod", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podApiService) DeletePodByID(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PodApi.DeletePodByID", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podApiService) UpdatePod(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PodApi.UpdatePod", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *podApiService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PodApi.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PodApi service

type PodApiHandler interface {
	QueryPodByID(context.Context, *Request, *Response) error
	AddPod(context.Context, *Request, *Response) error
	DeletePodByID(context.Context, *Request, *Response) error
	UpdatePod(context.Context, *Request, *Response) error
	Call(context.Context, *Request, *Response) error
}

func RegisterPodApiHandler(s server.Server, hdlr PodApiHandler, opts ...server.HandlerOption) error {
	type podApi interface {
		QueryPodByID(ctx context.Context, in *Request, out *Response) error
		AddPod(ctx context.Context, in *Request, out *Response) error
		DeletePodByID(ctx context.Context, in *Request, out *Response) error
		UpdatePod(ctx context.Context, in *Request, out *Response) error
		Call(ctx context.Context, in *Request, out *Response) error
	}
	type PodApi struct {
		podApi
	}
	h := &podApiHandler{hdlr}
	return s.Handle(s.NewHandler(&PodApi{h}, opts...))
}

type podApiHandler struct {
	PodApiHandler
}

func (h *podApiHandler) QueryPodByID(ctx context.Context, in *Request, out *Response) error {
	return h.PodApiHandler.QueryPodByID(ctx, in, out)
}

func (h *podApiHandler) AddPod(ctx context.Context, in *Request, out *Response) error {
	return h.PodApiHandler.AddPod(ctx, in, out)
}

func (h *podApiHandler) DeletePodByID(ctx context.Context, in *Request, out *Response) error {
	return h.PodApiHandler.DeletePodByID(ctx, in, out)
}

func (h *podApiHandler) UpdatePod(ctx context.Context, in *Request, out *Response) error {
	return h.PodApiHandler.UpdatePod(ctx, in, out)
}

func (h *podApiHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.PodApiHandler.Call(ctx, in, out)
}
