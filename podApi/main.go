package main

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/opentracing/opentracing-go"
	"github.com/qyh794/go-paas/pod/common"
	hystrix2 "github.com/qyh794/go-paas/pod/plugin/hystrix"
	"github.com/qyh794/go-paas/pod/proto/pod"
	"github.com/qyh794/go-paas/podApi/handler"
	"github.com/qyh794/go-paas/podApi/proto/podApi"
	"net"
	"net/http"
	"strconv"
)

var (
	//服务地址
	hostIp = "192.168.0.105"
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

const QPS = 1000

func main() {
	// 注册中心
	newRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":" + strconv.FormatInt(consulPort, 10),
		}
	})
	// 链路追踪
	tracer, io, err := common.NewTracer("go.micro.api.podApi", tracerHost+":"+strconv.Itoa(tracerPort))
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
		err := http.ListenAndServe(net.JoinHostPort("0.0.0.0", strconv.Itoa(hystrixPort)), streamHandler)
		if err != nil {
			logger.Error(err)
		}
	}()
	// 添加监控采集地址
	common.PrometheusBoot(prometheusPort)
	// 创建服务
	service := micro.NewService(
		micro.Server(server.NewServer(func(options *server.Options) {
			options.Advertise = serviceHost + ":" + servicePort
		})),
		micro.Name("go.micro.api.podApi"),
		micro.Version("latest"),
		micro.Address(":"+servicePort),
		// 注册中心
		micro.Registry(newRegistry),
		// 链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		// 作为客户端启动熔断
		micro.WrapClient(hystrix2.NewClientHystrixWrapper()),
		// 限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)
	service.Init()
	// 作为客户端调用pod服务
	podService := pod.NewPodService("go.micro.service.pod", service.Client())
	err = podApi.RegisterPodApiHandler(service.Server(), &handler.PodApi{PodService: podService})
	if err != nil {
		logger.Error(err)
	}
	if err := service.Run(); err != nil {
		logger.Error(err)
	}

}
