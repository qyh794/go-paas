package main

import (
	"github.com/micro/micro/v3/service/logger"
	"github.com/qyh794/go-paas/route/proto/route"
	"github.com/qyh794/go-paas/routeApi/Init/consulInit"
	"github.com/qyh794/go-paas/routeApi/Init/service"
	"github.com/qyh794/go-paas/routeApi/handler"
	"github.com/qyh794/go-paas/routeApi/proto/routeApi"
	"github.com/qyh794/go-paas/routeApi/settings"
)

func main() {
	if err := settings.Init(); err != nil {
		logger.Error(err)
	}
	newRegistry := consulInit.Init(settings.Conf.Consul.Host, settings.Conf.Consul.Port)

	newService := service.Init(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, newRegistry)
	newService.Init()

	routeService := route.NewRouteService("go.micro.service.route", newService.Client())
	if err := routeApi.RegisterRouteApiHandler(newService.Server(), &handler.RouteApi{RouteService: routeService}); err != nil {
		logger.Error(err)
	}
	if err := newService.Run(); err != nil {
		logger.Error(err)
	}
}
