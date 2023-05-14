package handler

import (
	"context"
	"encoding/json"
	"github.com/qyh794/go-paas/user/proto/user"
	"github.com/qyh794/go-paas/userApi/proto/userApi"
)

type UserApi struct {
	UserService user.UserService
}

func (u *UserApi) SignUp(ctx context.Context, request *userApi.Request, response *userApi.Response) error {
	userInfo := &user.RequestSignUp{}
	err := json.Unmarshal([]byte(request.Body), userInfo)
	if err != nil {
		return ResponseBadRequest(ctx, err, response)
	}

	responseSignUp, err := u.UserService.SignUp(ctx, userInfo)
	if err != nil {
		return ResponseServerError(ctx, err, response)
	}
	return ResponseSucceed(ctx, responseSignUp, response)
}

func (u *UserApi) Login(ctx context.Context, request *userApi.Request, response *userApi.Response) error {
	userInfo := &user.RequestLogin{}
	err := json.Unmarshal([]byte(request.Body), userInfo)
	if err != nil {
		return ResponseBadRequest(ctx, err, response)
	}
	responseLogin, err := u.UserService.Login(ctx, userInfo)
	if err != nil {
		return ResponseServerError(ctx, err, response)

	}
	return ResponseSucceed(ctx, responseLogin, response)
}
