package main

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/qyh794/go-paas/volume/proto/volume"
	"github.com/qyh794/go-paas/volumeApi/Init/consulInit"
	"github.com/qyh794/go-paas/volumeApi/Init/service"
	"github.com/qyh794/go-paas/volumeApi/handler"
	"github.com/qyh794/go-paas/volumeApi/proto/volumeApi"
	"github.com/qyh794/go-paas/volumeApi/settings"
)

func main() {
	if err := settings.Init(); err != nil {
		logger.Error(err)
	}
	newRegistry := consulInit.Init(settings.Conf.Consul.Host, settings.Conf.Consul.Port)

	newService := service.Init(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, newRegistry)
	newService.Init()

	volumeService := volume.NewVolumeService("go.micro.service.volume", newService.Client())
	if err := volumeApi.RegisterVolumeApiHandler(newService.Server(), &handler.VolumeApi{VolumeService: volumeService}); err != nil {
		logger.Error(err)
	}
	if err := newService.Run(); err != nil {
		logger.Error(err)
	}
}
