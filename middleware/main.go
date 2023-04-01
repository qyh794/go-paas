package main

import (
	"fmt"
	"github.com/asim/go-micro/v3/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/qyh794/go-paas/common"
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
	// 加载配置文件中的配置
	if err := settings.Init(); err != nil {
		return
	}

	// 初始化注册中心
	consul := common.InitConsul(settings.Conf.ServiceHost, settings.Conf.ServicePort)
	// 配置中心
	consulConfig, err := common.GetConsulConfig(settings.Conf.Consul.Host, settings.Conf.Consul.Port, settings.Conf.Consul.Prefix)
	if err != nil {
		logger.Error(err)
	}

	// 配置中心获取MySQL配置
	mysqlConfig := common.GetMysqlConfigFromConsul(consulConfig, settings.Conf.Consul.Path)
	fmt.Println("user: ", mysqlConfig.User)
	// 初始化MySQL
	db, err := common.InitMySQL("mysql", mysqlConfig)
	if err != nil {
		logger.Error(err)
	}
	defer func() {
		_ = db.Close()
	}()
	db.SingularTable(true)

	// 初始化链路追踪
	tracer, io, err := common.NewTracer(settings.Conf.Name, settings.Conf.Tracer.Host, settings.Conf.Tracer.Port)
	if err != nil {
		logger.Error(err)
	}
	defer func() {
		_ = io.Close()
	}()
	opentracing.SetGlobalTracer(tracer)

	// 添加监控
	common.PrometheusBoot(settings.Conf.Prometheus.Port)

	// 初始化k8s
	k8sClient := common.InitK8sClient()

	// 初始化服务
	service := common.InitService(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, 1000, consul)
	service.Init()
	middlewareDataService := service2.NewMiddlewareDataService(repository.NewMiddlewareRepository(db), k8sClient)
	middlewareTypeDataService := service2.NewMiddlewareTypeDataService(repository.NewMiddleTypeRepository(db))
	err = middleware.RegisterMiddlewareHandler(service.Server(), &handler.MiddlewareHandler{MiddlewareDataService: middlewareDataService, MiddlewareTypeDataService: middlewareTypeDataService})
	if err != nil {
		logger.Error(err)
	}
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
