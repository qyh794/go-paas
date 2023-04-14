package service

import (
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3"
	"github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v3"
	"github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	opentracing2 "github.com/opentracing/opentracing-go"
	"github.com/qyh794/go-paas/podApi/plugin/hystrix"
)

const QPS = 1000

func Init(serverHost, serverPort, serverName, serverVersion string, consul registry.Registry) micro.Service {
	newService := micro.NewService(
		micro.Server(server.NewServer(func(options *server.Options) {
			options.Advertise = serverHost + ":" + serverPort
		})),
		micro.Name(serverName),
		micro.Version(serverVersion),
		micro.Address(":"+serverPort),
		micro.Registry(consul),
		micro.WrapHandler(opentracing.NewHandlerWrapper(opentracing2.GlobalTracer())),
		micro.WrapClient(opentracing.NewClientWrapper(opentracing2.GlobalTracer())),
		micro.WrapClient(hystrix.NewClientHystrixWrapper()),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)
	newService.Init()
	return newService
}
