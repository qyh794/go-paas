package main

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/qyh794/go-paas/user/proto/user"
	"github.com/qyh794/go-paas/userApi/Init/consul"
	"github.com/qyh794/go-paas/userApi/Init/service"
	"github.com/qyh794/go-paas/userApi/handler"
	"github.com/qyh794/go-paas/userApi/proto/userApi"
	"github.com/qyh794/go-paas/userApi/settings"
)

func main() {
	if err := settings.Init(); err != nil {
		logger.Error(err)
	}
	consulRegistry := consul.Init(settings.Conf.Consul.Host, settings.Conf.Consul.Port)

	microService := service.Init(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, consulRegistry)
	userService := user.NewUserService("go.micro.service.user", microService.Client())
	if err := userApi.RegisterUserApiHandler(microService.Server(), &handler.UserApi{UserService: userService}); err != nil {
		logger.Error(err)
	}
	if err := microService.Run(); err != nil {
		logger.Error(err)
	}
}
