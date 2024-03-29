// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/user.proto

package user

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

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	Login(ctx context.Context, in *RequestLogin, opts ...client.CallOption) (*ResponseLogin, error)
	SignUp(ctx context.Context, in *RequestSignUp, opts ...client.CallOption) (*ResponseSignUp, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Login(ctx context.Context, in *RequestLogin, opts ...client.CallOption) (*ResponseLogin, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(ResponseLogin)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SignUp(ctx context.Context, in *RequestSignUp, opts ...client.CallOption) (*ResponseSignUp, error) {
	req := c.c.NewRequest(c.name, "User.SignUp", in)
	out := new(ResponseSignUp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Login(context.Context, *RequestLogin, *ResponseLogin) error
	SignUp(context.Context, *RequestSignUp, *ResponseSignUp) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Login(ctx context.Context, in *RequestLogin, out *ResponseLogin) error
		SignUp(ctx context.Context, in *RequestSignUp, out *ResponseSignUp) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Login(ctx context.Context, in *RequestLogin, out *ResponseLogin) error {
	return h.UserHandler.Login(ctx, in, out)
}

func (h *userHandler) SignUp(ctx context.Context, in *RequestSignUp, out *ResponseSignUp) error {
	return h.UserHandler.SignUp(ctx, in, out)
}
