package main

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/v3/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/qyh794/go-paas/common"
	"github.com/qyh794/go-paas/middleware/proto/middleware"
	"github.com/qyh794/go-paas/middlewareApi/Init"
	"github.com/qyh794/go-paas/middlewareApi/handler"
	"github.com/qyh794/go-paas/middlewareApi/proto/middlewareApi"
	"github.com/qyh794/go-paas/middlewareApi/settings"
	"net"
	"net/http"
	"strconv"
)

func main() {
	if err := settings.Init(); err != nil {
		return
	}
	//需要本地启动，mysql，consul中间件服务
	//1.注册中心
	newRegistry := Init.InitConsul(settings.Conf.Consul.Host, settings.Conf.Consul.Port)

	//2.添加链路追踪
	t, io, err := common.NewTracer("go.micro.api.middlewareApi", settings.Conf.Tracer.Host, settings.Conf.Tracer.Port)
	if err != nil {
		logger.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)
	//3.添加熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	//启动监听程序
	go func() {
		//http://192.168.0.112:9092/turbine/turbine.stream
		//看板访问地址 http://127.0.0.1:9002/hystrix，url后面一定要带 /hystrix
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", strconv.Itoa(settings.Conf.Hystrix.Port)), hystrixStreamHandler)
		if err != nil {
			logger.Error(err)
		}
	}()

	//4.添加监控
	common.PrometheusBoot(settings.Conf.Prometheus.Port)
	//5.创建服务
	service := Init.InitServer(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, newRegistry)
	service.Init()
	middlewareService := middleware.NewMiddlewareService("go.micro.service.middleware", service.Client())
	// 注册控制器
	err = middlewareApi.RegisterMiddlewareApiHandler(service.Server(), &handler.MiddlewareApi{MiddlewareService: middlewareService})
	if err != nil {
		logger.Error(err)
	}

	// 启动服务
	if err = service.Run(); err != nil {
		//输出启动失败信息
		logger.Fatal(err)
	}
}
