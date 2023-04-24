package main

import (
	"appStoreApi/Init/consul"
	"appStoreApi/Init/service"
	"appStoreApi/handler"
	"appStoreApi/proto/appStoreApi"
	"appStoreApi/settings"
	"github.com/asim/go-micro/v3/logger"
	"github.com/qyh794/go-paas/appStore/proto/appStore"
)

func main() {
	if err := settings.Init(); err != nil {
		logger.Error(err)
	}
	consulRegistry := consul.Init(settings.Conf.Consul.Host, settings.Conf.Consul.Port)

	microService := service.Init(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, consulRegistry)
	appStoreService := appStore.NewAppStoreService("go.micro.service.appStore", microService.Client())

	if err := appStoreApi.RegisterAppStoreApiHandler(microService.Server(), &handler.AppStoreApi{AppStoreService: appStoreService}); err != nil {
		logger.Error(err)
	}

	if err := microService.Run(); err != nil {
		logger.Error(err)
	}
}
