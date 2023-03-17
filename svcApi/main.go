package main

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3"
	"github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/opentracing/opentracing-go"
	"github.com/qyh794/go-paas/pod/common"
	"github.com/qyh794/go-paas/svc/proto/svc"
	"net"
	"net/http"
	"strconv"
	"svcApi/handler"
	hystrix2 "svcApi/plugin/hystrix"
	"svcApi/proto/svcApi"
)

var (
	//服务地址
	hostIp = "0.0.0.0"
	//服务地址
	serviceHost = hostIp
	//服务端口
	servicePort = "8084"
	//注册中心配置
	consulHost       = hostIp
	consulPort int64 = 8500
	//链路追踪
	tracerHost = hostIp
	tracerPort = 6831
	//熔断端口，每个服务不能重复
	hystrixPort = 9094
	//监控端口，每个服务不能重复
	prometheusPort = 9194
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
	tracer, io, err := common.NewTracer("go.micro.api.svcApi", tracerHost+":"+strconv.Itoa(tracerPort))
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
	// 启动监听程序
	go func() {
		err := http.ListenAndServe(net.JoinHostPort("0.0.0.0", strconv.Itoa(hystrixPort)), streamHandler)
		if err != nil {
			logger.Error(err)
		}
	}()
	// 添加监控
	common.PrometheusBoot(prometheusPort)
	// 创建服务
	srv := micro.NewService(
		micro.Server(server.NewServer(func(options *server.Options) {
			options.Advertise = serviceHost + ":" + servicePort
		})),
		micro.Name("go.micro.api.svcApi"),
		micro.Version("latest"),
		micro.Address(":"+servicePort),
		micro.Registry(newRegistry),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(hystrix2.NewClientHystrixWrapper()),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)
	srv.Init()
	svcService := svc.NewSvcService("go.micro.service.pod", srv.Client())
	err = svcApi.RegisterSvcApiHandler(srv.Server(), &handler.SvcApi{SvcService: svcService})
	if err != nil {
		logger.Error(err)
	}
	if err = srv.Run(); err != nil {
		logger.Error(err)
	}
}
