package main

import (
	"fmt"
	"github.com/asim/go-micro/v3/logger"
	"github.com/qyh794/go-paas/common"
	"github.com/qyh794/go-paas/route/Init/consulInit"
	"github.com/qyh794/go-paas/route/Init/k8s"
	"github.com/qyh794/go-paas/route/Init/mysql"
	"github.com/qyh794/go-paas/route/Init/serviceInit"
	"github.com/qyh794/go-paas/route/domain/repository"
	service2 "github.com/qyh794/go-paas/route/domain/service"
	"github.com/qyh794/go-paas/route/handler"
	"github.com/qyh794/go-paas/route/proto/route"
	"github.com/qyh794/go-paas/route/settings"
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Printf("init setting failed, err:%v", err)
		return
	}
	newRegistry := consulInit.Init(settings.Conf.Consul.Host, settings.Conf.Consul.Port)

	consulConfig, err := common.GetConsulConfig(settings.Conf.Consul.Host, settings.Conf.Consul.Port, "/micro/config")
	if err != nil {
		logger.Error(err)
	}
	mysqlConfig := common.GetMysqlConfigFromConsul(consulConfig, "mysql")
	if err = mysql.Init("mysql", mysqlConfig); err != nil {
		logger.Error(err)
	}
	defer mysql.Close()

	clientset := k8s.Init()
	service := serviceInit.Init(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, newRegistry)
	service.Init()
	err = repository.NewRouteRepository(mysql.DB).InitTable()
	if err != nil {
		logger.Error(err)
	}
	routeDataService := service2.NewRouteDataService(repository.NewRouteRepository(mysql.DB), clientset)
	err = route.RegisterRouteHandler(service.Server(), &handler.RouteHandler{RouteDateService: routeDataService})
	if err != nil {
		logger.Error(err)
	}
	err = service.Run()
	if err != nil {
		logger.Error(err)
	}
}
