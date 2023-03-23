package main

import (
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3"
	"github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/micro/micro/v3/service/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/qyh794/go-paas/pod/common"
	"github.com/qyh794/go-paas/route/proto/route"
	"github.com/qyh794/go-paas/routeApi/handler"
	hystrix2 "github.com/qyh794/go-paas/routeApi/plugin/hystrix"
	"github.com/qyh794/go-paas/routeApi/proto/routeApi"
	"net"
	"net/http"
	"strconv"
)

var (
	//服务地址
	hostIp = "0.0.0.0"
	//服务地址
	serviceHost = hostIp
	//服务端口
	servicePort = "8088"
	//注册中心配置
	consulHost       = hostIp
	consulPort int64 = 8500
	//链路追踪
	tracerHost = hostIp
	tracerPort = 6831
	//熔断端口，每个服务不能重复
	hystrixPort = 9098
	//监控端口，每个服务不能重复
	prometheusPort = 9198
)

const QPS = 1000

func main() {
	newRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":" + strconv.FormatInt(consulPort, 10),
		}
	})
	tracer, io, err := common.NewTracer("go.mirco.api.routeApi", tracerHost+":"+strconv.Itoa(tracerPort))
	if err != nil {
		logger.Error(err)
	}
	defer func() {
		_ = io.Close()
	}()
	opentracing.SetGlobalTracer(tracer)
	streamHandler := hystrix.NewStreamHandler()
	streamHandler.Start()
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", strconv.Itoa(hystrixPort)), streamHandler)
		if err != nil {
			logger.Error(err)
		}
	}()
	common.PrometheusBoot(prometheusPort)
	service := micro.NewService(
		micro.Server(server.NewServer(func(options *server.Options) {
			options.Advertise = serviceHost + ":" + servicePort
		})),
		micro.Name("go.micro.api.routeApi"),
		micro.Version("latest"),
		micro.Address(":"+servicePort),
		micro.Registry(newRegistry),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(hystrix2.NewClientHystrixWrapper()),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)
	service.Init()
	routeService := route.NewRouteService("go.micro.service.route", service.Client())
	if err = routeApi.RegisterRouteApiHandler(service.Server(), &handler.RouteApi{RouteService: routeService}); err != nil {
		logger.Error(err)
	}
	if err = service.Run(); err != nil {
		logger.Error(err)
	}
}
