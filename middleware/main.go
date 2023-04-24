package main

import (
	"github.com/asim/go-micro/v3/logger"
	"github.com/qyh794/go-paas/common"
	"github.com/qyh794/go-paas/middleware/Init/consulInit"
	"github.com/qyh794/go-paas/middleware/Init/k8s"
	"github.com/qyh794/go-paas/middleware/Init/mysql"
	"github.com/qyh794/go-paas/middleware/Init/serviceInit"
	"github.com/qyh794/go-paas/middleware/domain/repository"
	service2 "github.com/qyh794/go-paas/middleware/domain/service"
	"github.com/qyh794/go-paas/middleware/handler"
	"github.com/qyh794/go-paas/middleware/proto/middleware"
	"github.com/qyh794/go-paas/middleware/settings"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//需要本地启动，mysql，consul中间件服务
	if err := settings.Init(); err != nil {
		return
	}
	//1.注册中心
	consulRegistry := consulInit.Init(settings.Conf.Consul.Host, settings.Conf.Consul.Port)
	//2.配置中心，存放经常变动的变量
	consulConfig, err := common.GetConsulConfig(settings.Conf.Consul.Host, settings.Conf.Consul.Port, "/micro/config")
	if err != nil {
		logger.Error(err)
	}
	//3.使用配置中心连接 mysql
	//common.GetMysqlConfigFromConsul()
	mysqlConfig := common.GetMysqlConfigFromConsul(consulConfig, "mysql")
	//初始化数据库
	if err = mysql.Init("mysql", mysqlConfig); err != nil {
		logger.Error(err)
	}
	defer mysql.Close()

	//6.创建k8s连接
	//在集群外面连接
	clientset := k8s.Init()

	//7.创建服务
	service := serviceInit.Init(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, consulRegistry)
	service.Init()

	// 注册句柄，可以快速操作已开发的服务
	middlewareDataService := service2.NewMiddlewareDataService(repository.NewMiddlewareRepository(mysql.DB), clientset)
	middleTypeDataService := service2.NewMiddlewareTypeDataService(repository.NewMiddleTypeRepository(mysql.DB))
	err = middleware.RegisterMiddlewareHandler(service.Server(), &handler.MiddlewareHandler{MiddlewareDataService: middlewareDataService, MiddlewareTypeDataService: middleTypeDataService})
	if err != nil {
		logger.Error(err)
	}
	// 启动服务
	go func() {
		if err = service.Run(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}

	}()
	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	logger.Info("Shutdown Server ...")
	logger.Info("Server exiting")
}
