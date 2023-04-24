package main

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/qyh794/go-paas/appStore/Init/consulInit"
	"github.com/qyh794/go-paas/appStore/Init/mysql"
	"github.com/qyh794/go-paas/appStore/Init/serviceInit"
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

	consul := consulInit.Init(settings.Conf.Consul.Host, settings.Conf.Consul.Port)
	consulConfig, err := common.GetConsulConfig(settings.Conf.Consul.Host, settings.Conf.Consul.Port, settings.Conf.Consul.Prefix)
	if err != nil {
		logger.Error(err)
	}

	mysqlConfig := common.GetMysqlConfigFromConsul(consulConfig, settings.Conf.Consul.Path)
	if err = mysql.Init("mysql", mysqlConfig); err != nil {
		logger.Error(err)
	}
	defer mysql.Close()

	service := serviceInit.Init(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, consul)
	service.Init()
	appStoreDataService := service2.NewAppStoreDataService(repository.NewAppStoreRepository(mysql.DB))
	_ = appStore.RegisterAppStoreHandler(service.Server(), &handler.AppStoreHandler{AppStoreDataService: appStoreDataService})

	if err = service.Run(); err != nil {
		logger.Error(err)
	}
}
