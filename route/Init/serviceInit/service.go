package serviceInit

import (
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
)

const QPS = 1000

func Init(serverHost, serverPort, serverName, serverVersion string, consul registry.Registry) micro.Service {
	return micro.NewService(
		micro.Server(server.NewServer(func(options *server.Options) {
			options.Advertise = serverHost + ":" + serverPort
		})),
		micro.Name(serverName),
		micro.Version(serverVersion),
		micro.Address(":"+serverPort),
		micro.Registry(consul),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)
}
