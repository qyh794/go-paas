package main

import (
	"fmt"
	"github.com/micro/micro/v3/service/logger"
	"github.com/qyh794/go-paas/common"
	"github.com/qyh794/go-paas/volume/Init/consulInit"
	"github.com/qyh794/go-paas/volume/Init/k8s"
	"github.com/qyh794/go-paas/volume/Init/mysql"
	"github.com/qyh794/go-paas/volume/Init/serviceInit"
	"github.com/qyh794/go-paas/volume/domain/repository"
	service2 "github.com/qyh794/go-paas/volume/domain/service"
	"github.com/qyh794/go-paas/volume/handler"
	"github.com/qyh794/go-paas/volume/proto/volume"
	"github.com/qyh794/go-paas/volume/settings"
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

	newService := serviceInit.Init(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, newRegistry)

	newService.Init()

	volumeDataService := service2.NewVolumeDataService(clientset, repository.NewVolumeRepository(mysql.DB))
	err = volume.RegisterVolumeHandler(newService.Server(), &handler.VolumeHandler{VolumeDataService: volumeDataService})
	if err != nil {
		logger.Error(err)
	}
	if err = newService.Run(); err != nil {
		logger.Error(err)
	}
}
