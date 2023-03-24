package main

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/opentracing/opentracing-go"
	"github.com/qyh794/go-paas/pod/common"
	"github.com/qyh794/go-paas/volume/proto/volume"
	"github.com/qyh794/go-paas/volumeApi/handler"
	hystrix2 "github.com/qyh794/go-paas/volumeApi/plugin/hystrix"
	"github.com/qyh794/go-paas/volumeApi/proto/volumeApi"
	"net"
	"net/http"
	"strconv"
)

var (
	//服务地址
	hostIp = "192.168.0.108"
	//服务地址
	serviceHost = hostIp
	//服务端口
	servicePort = "8082"
	//注册中心配置
	consulHost       = hostIp
	consulPort int64 = 8500
	//链路追踪
	tracerHost = hostIp
	tracerPort = 6831
	//熔断端口，每个服务不能重复
	hystrixPort = 9092
	//监控端口，每个服务不能重复
	prometheusPort = 9192
)

func main() {
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":" + strconv.FormatInt(consulPort, 10),
		}
	})
	tracer, io, err := common.NewTracer("go.micro.api.volumeApi", tracerHost+":"+strconv.Itoa(tracerPort))
	if err != nil {
		logger.Error(err)
	}
	defer func() {
		_ = io.Close()
	}()
	opentracing.SetGlobalTracer(tracer)
	newStreamHandler := hystrix.NewStreamHandler()
	newStreamHandler.Start()
	go func() {
		err := http.ListenAndServe(net.JoinHostPort("0.0.0.0", strconv.Itoa(hystrixPort)), newStreamHandler)
		if err != nil {
			logger.Error(err)
		}
	}()
	common.PrometheusBoot(prometheusPort)
	service := micro.NewService(
		micro.Server(server.NewServer(func(options *server.Options) {
			options.Advertise = serviceHost + ":" + servicePort
		})),
		micro.Name("go.micro.api.volumeApi"),
		micro.Version("latest"),
		micro.Address(":"+servicePort),
		micro.Registry(consulRegistry),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(hystrix2.NewClientHystrixWrapper()),
		micro.WrapClient(),
	)
	service.Init()
	volumeService := volume.NewVolumeService("go.micro.service.volume", service.Client())
	if err = volumeApi.RegisterVolumeApiHandler(service.Server(), &handler.VolumeApi{VolumeService: volumeService}); err != nil {
		logger.Error(err)
	}
	if err = service.Run(); err != nil {
		logger.Error(err)
	}
}
