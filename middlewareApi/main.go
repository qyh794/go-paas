package main

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/qyh794/go-paas/middleware/proto/middleware"
	"github.com/qyh794/go-paas/middlewareApi/Init/consulInit"
	"github.com/qyh794/go-paas/middlewareApi/Init/service"
	"github.com/qyh794/go-paas/middlewareApi/handler"
	"github.com/qyh794/go-paas/middlewareApi/proto/middlewareApi"
	"github.com/qyh794/go-paas/middlewareApi/settings"
)

func main() {
	if err := settings.Init(); err != nil {
		return
	}
	//需要本地启动，mysql，consul中间件服务
	//1.注册中心
	newRegistry := consulInit.Init(settings.Conf.Consul.Host, settings.Conf.Consul.Port)

	//5.创建服务
	newService := service.Init(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, newRegistry)
	newService.Init()
	middlewareService := middleware.NewMiddlewareService("go.micro.service.middleware", newService.Client())
	// 注册控制器
	err := middlewareApi.RegisterMiddlewareApiHandler(newService.Server(), &handler.MiddlewareApi{MiddlewareService: middlewareService})
	if err != nil {
		logger.Error(err)
	}

	// 启动服务
	if err = newService.Run(); err != nil {
		//输出启动失败信息
		logger.Fatal(err)
	}
}
