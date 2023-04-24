package main

import (
	"fmt"
	"github.com/asim/go-micro/v3/logger"

	"github.com/qyh794/go-paas/common"

	"github.com/qyh794/go-paas/svc/Init/consulInit"
	"github.com/qyh794/go-paas/svc/Init/k8s"
	"github.com/qyh794/go-paas/svc/Init/mysql"
	"github.com/qyh794/go-paas/svc/Init/serviceInit"
	"github.com/qyh794/go-paas/svc/domain/repository"
	"github.com/qyh794/go-paas/svc/domain/service"
	"github.com/qyh794/go-paas/svc/handler"
	"github.com/qyh794/go-paas/svc/proto/svc"
	"github.com/qyh794/go-paas/svc/settings"
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Printf("init setting failed, err:%v", err)
		return
	}
	// 注册中心
	newRegistry := consulInit.Init(settings.Conf.Consul.Host, settings.Conf.Consul.Port)
	// 配置中心
	consulConfig, err := common.GetConsulConfig(settings.Conf.Consul.Host, settings.Conf.Consul.Port, "/micro/config")
	if err != nil {
		logger.Error(err)
	}
	// 从配置中心获取MySQL配置
	mysqlConfig := common.GetMysqlConfigFromConsul(consulConfig, "mysql")
	// 初始化数据库链接
	if err = mysql.Init("mysql", mysqlConfig); err != nil {
		logger.Error(err)
	}
	defer mysql.Close()

	// 创建k8s连接
	clientset := k8s.Init()
	// 创建服务
	newService := serviceInit.Init(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, newRegistry)
	// 建表
	svcRepository := repository.NewSvcRepository(mysql.DB)
	err = svcRepository.InitTable()
	if err != nil {
		logger.Error(err)
	}
	SvcDateService := service.NewServiceDateService(svcRepository, clientset)
	// 注册句柄
	_ = svc.RegisterSvcHandler(newService.Server(), &handler.SvcHandler{SvcDataService: SvcDateService})
	// 启动服务
	err = newService.Run()
	if err != nil {
		logger.Error(err)
	}
}
