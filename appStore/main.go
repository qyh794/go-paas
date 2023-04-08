package main

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/qyh794/go-paas/appStore/Init"
	"github.com/qyh794/go-paas/appStore/domain/repository"
	service2 "github.com/qyh794/go-paas/appStore/domain/service"
	"github.com/qyh794/go-paas/appStore/handler"
	"github.com/qyh794/go-paas/appStore/proto/appStore"
	"github.com/qyh794/go-paas/appStore/settings"
	"github.com/qyh794/go-paas/common"
)

func main() {
	if err := settings.Init(); err != nil {
		return
	}

	consul := Init.SetupConsul(settings.Conf.Consul.Host, settings.Conf.Consul.Port)
	consulConfig, err := common.GetConsulConfig(settings.Conf.Consul.Host, settings.Conf.Consul.Port, settings.Conf.Consul.Prefix)
	if err != nil {
		logger.Error(err)
	}

	mysqlConfig := common.GetMysqlConfigFromConsul(consulConfig, settings.Conf.Consul.Path)
	if err = Init.SetupMysql("mysql", mysqlConfig); err != nil {
		logger.Error(err)
	}
	Init.Close()

	tracer, io, err := common.NewTracer(settings.Conf.Name, settings.Conf.Tracer.Host, settings.Conf.Tracer.Port)
	if err != nil {
		logger.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(tracer)

	common.PrometheusBoot(settings.Conf.Prometheus.Port)

	service := Init.SetupService(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, consul)
	service.Init()
	appStoreDataService := service2.NewAppStoreDataService(repository.NewAppStoreRepository(Init.DB))
	_ = appStore.RegisterAppStoreHandler(service.Server(), &handler.AppStoreHandler{AppStoreDataService: appStoreDataService})

	if err = service.Run(); err != nil {
		logger.Error(err)
	}
}
