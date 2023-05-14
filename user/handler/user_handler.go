package handler

import (
	"context"
	"errors"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/util/log"
	"github.com/qyh794/go-paas/common"
	"github.com/qyh794/go-paas/user/domain/model"
	"github.com/qyh794/go-paas/user/domain/service"
	"github.com/qyh794/go-paas/user/proto/user"
)

type UserHandler struct {
	UserDataService service.IUserDataService
}

func (u *UserHandler) Login(ctx context.Context, requestLogin *user.RequestLogin, responseLogin *user.ResponseLogin) error {
	log.Info("接收到登录请求")
	userObj := &model.User{}
	if err := common.SwapTo(requestLogin, userObj); err != nil {
		logger.Error(err)
		return err
	}
	token, err := u.UserDataService.Login(userObj)
	if err != nil {
		return ResponseError(responseLogin, err)
	}
	return ResponseLogin(responseLogin, err, token)
}

func (u *UserHandler) SignUp(ctx context.Context, requestSignUp *user.RequestSignUp, responseSignUp *user.ResponseSignUp) error {
	log.Info("接收到注册请求")
	if requestSignUp.Password != requestSignUp.RePassword {
		return errors.New("两次密码不一致")
	}
	userObj := &model.User{}
	if err := common.SwapTo(requestSignUp, userObj); err != nil {
		return err
	}
	if err := u.UserDataService.SignUp(userObj); err != nil {
		return err
	}
	return ResponseSignUp(responseSignUp)
}
