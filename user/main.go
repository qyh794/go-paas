package main

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/qyh794/go-paas/common"
	"github.com/qyh794/go-paas/user/Init/consulInit"
	"github.com/qyh794/go-paas/user/Init/mysql"
	"github.com/qyh794/go-paas/user/Init/serviceInit"
	"github.com/qyh794/go-paas/user/domain/repository"
	service2 "github.com/qyh794/go-paas/user/domain/service"
	"github.com/qyh794/go-paas/user/handler"
	"github.com/qyh794/go-paas/user/proto/user"
	"github.com/qyh794/go-paas/user/settings"
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
	userDataService := service2.NewUserDataService(repository.NewUserRepository(mysql.DB))
	_ = user.RegisterUserHandler(service.Server(), &handler.UserHandler{UserDataService: userDataService})
	if err = service.Run(); err != nil {
		logger.Error(err)
	}
}
