package handler

import (
	"context"
	"encoding/json"
	"github.com/asim/go-micro/v3/logger"
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
		logger.Error(err)
		response.Body = err.Error()
		return err
	}

	responseSignUp, err := u.UserService.SignUp(ctx, userInfo)
	if err != nil {
		logger.Error(err)
		response.Body = err.Error()
		return err
	}
	response.StatusCode = 200
	bytes, _ := json.Marshal(responseSignUp)
	response.Body = string(bytes)
	return nil
}

func (u *UserApi) Login(ctx context.Context, request *userApi.Request, response *userApi.Response) error {
	userInfo := &user.RequestLogin{}
	err := json.Unmarshal([]byte(request.Body), userInfo)
	if err != nil {
		response.StatusCode = 403
		response.Body = err.Error()
		return err
	}
	responseLogin, err := u.UserService.Login(ctx, userInfo)
	if err != nil {
		logger.Error(err)
		response.Body = err.Error()
		return err
	}
	response.StatusCode = 200
	bytes, _ := json.Marshal(responseLogin)
	response.Body = string(bytes)
	return nil
}
