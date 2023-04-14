package main

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/v3/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/qyh794/go-paas/common"
	"github.com/qyh794/go-paas/pod/proto/pod"
	"github.com/qyh794/go-paas/podApi/Init/consulInit"
	"github.com/qyh794/go-paas/podApi/Init/service"
	"github.com/qyh794/go-paas/podApi/handler"
	"github.com/qyh794/go-paas/podApi/proto/podApi"
	"github.com/qyh794/go-paas/podApi/settings"
	"net"
	"net/http"
	"strconv"
)

func main() {
	if err := settings.Init(); err != nil {
		logger.Error(err)
	}

	// 注册中心
	newRegistry := consulInit.Init(settings.Conf.Consul.Host, settings.Conf.Consul.Port)
	// 链路追踪
	tracer, io, err := common.NewTracer(settings.Conf.Name, settings.Conf.Tracer.Host, settings.Conf.Tracer.Port)
	if err != nil {
		logger.Error(err)
	}
	defer func() {
		_ = io.Close()
	}()
	opentracing.SetGlobalTracer(tracer)
	// 熔断器
	streamHandler := hystrix.NewStreamHandler()
	streamHandler.Start()
	// 熔断器监听程序
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", strconv.Itoa(settings.Conf.Hystrix.Port)), streamHandler)
		if err != nil {
			logger.Error(err)
		}
	}()
	// 添加监控采集地址
	common.PrometheusBoot(settings.Conf.Prometheus.Port)
	// 创建服务
	service := service.Init(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, newRegistry)
	service.Init()
	// 作为客户端调用pod服务
	podService := pod.NewPodService("go.micro.service.pod", service.Client())
	err = podApi.RegisterPodApiHandler(service.Server(), &handler.PodApi{PodService: podService})
	if err != nil {
		logger.Error(err)
	}
	if err = service.Run(); err != nil {
		logger.Error(err)
	}
}
