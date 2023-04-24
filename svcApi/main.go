package main

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/qyh794/go-paas/svc/proto/svc"
	consulInit "github.com/qyh794/go-paas/svcApi/Init/consul"
	"github.com/qyh794/go-paas/svcApi/Init/service"
	"github.com/qyh794/go-paas/svcApi/handler"
	"github.com/qyh794/go-paas/svcApi/settings"

	"github.com/qyh794/go-paas/svcApi/proto/svcApi"
)

func main() {
	if err := settings.Init(); err != nil {
		logger.Error(err)
	}

	// 注册中心
	newRegistry := consulInit.Init(settings.Conf.Consul.Host, settings.Conf.Consul.Port)

	// 创建服务
	srv := service.Init(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, newRegistry)
	srv.Init()
	svcService := svc.NewSvcService("go.micro.service.svc", srv.Client())
	err := svcApi.RegisterSvcApiHandler(srv.Server(), &handler.SvcApi{SvcService: svcService})
	if err != nil {
		logger.Error(err)
	}
	if err = srv.Run(); err != nil {
		logger.Error(err)
	}
}
