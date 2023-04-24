package main

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/qyh794/go-paas/pod/proto/pod"
	"github.com/qyh794/go-paas/podApi/Init/consulInit"
	"github.com/qyh794/go-paas/podApi/Init/service"
	"github.com/qyh794/go-paas/podApi/handler"
	"github.com/qyh794/go-paas/podApi/proto/podApi"
	"github.com/qyh794/go-paas/podApi/settings"
)

func main() {
	if err := settings.Init(); err != nil {
		logger.Error(err)
	}

	// 注册中心
	newRegistry := consulInit.Init(settings.Conf.Consul.Host, settings.Conf.Consul.Port)

	// 创建服务
	newService := service.Init(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, newRegistry)
	newService.Init()
	// 作为客户端调用pod服务
	podService := pod.NewPodService("go.micro.service.pod", newService.Client())
	err := podApi.RegisterPodApiHandler(newService.Server(), &handler.PodApi{PodService: podService})
	if err != nil {
		logger.Error(err)
	}
	if err = newService.Run(); err != nil {
		logger.Error(err)
	}
}
