package main

import (
	"fmt"
	"github.com/asim/go-micro/v3/logger"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/qyh794/go-paas/common"
	"github.com/qyh794/go-paas/pod/domain/model"
	"github.com/qyh794/go-paas/pod/domain/repository"
	"github.com/qyh794/go-paas/pod/domain/service"
	"github.com/qyh794/go-paas/pod/handler"
	"github.com/qyh794/go-paas/pod/init/consulInit"
	"github.com/qyh794/go-paas/pod/init/k8s"
	"github.com/qyh794/go-paas/pod/init/mysql"
	"github.com/qyh794/go-paas/pod/init/serviceInit"
	"github.com/qyh794/go-paas/pod/proto/pod"
	"github.com/qyh794/go-paas/pod/settings"
)

func main() {
	// 0. 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init setting failed, err:%v", err)
		return
	}
	// 1.注册中心
	newRegistry := consulInit.Init(settings.Conf.Consul.Host, settings.Conf.Consul.Port)
	// 2.配置中心
	consulConfig, err := common.GetConsulConfig(settings.Conf.Consul.Host, settings.Conf.Consul.Port, "/micro/config")
	if err != nil {
		logger.Error(err)
	}
	// 3.获取配置中心数据
	mysqlConfig := common.GetMysqlConfigFromConsul(consulConfig, "mysql")
	//fmt.Printf("user:%s\n password:%s\n database:%s\n port:%s\n host:%s\n", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Database, mysqlConfig.Port, mysqlConfig.Host)

	// 4.初始化数据库连接
	if err = mysql.Init("mysql", mysqlConfig); err != nil {
		logger.Error(err)
	}
	defer mysql.Close()

	clientset := k8s.Init()
	// 13.创建服务实例
	newService := serviceInit.Init(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, newRegistry)
	// 14.初始化服务
	newService.Init()
	if err = mysql.DB.AutoMigrate(&model.Pod{}, &model.PodPort{}, &model.PodEnv{}).Error; err != nil {
		logger.Error(err)
	}

	//err = repository.NewPodRepository(mysql.DB).InitTable()
	//if err != nil {
	//	logger.Fatal(err)
	//}
	// 15.注册句柄
	podDateService := service.NewPodDateService(repository.NewPodRepository(mysql.DB), clientset)
	_ = pod.RegisterPodHandler(newService.Server(), &handler.PodHandler{PodDataService: podDateService})
	// 16.启动服务
	if err = newService.Run(); err != nil {
		logger.Fatal(err)
	}
}
