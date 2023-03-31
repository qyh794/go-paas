package common

import (
	consul2 "github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/registry"
)

func InitConsul(consulHost, consulPort string) registry.Registry {
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":" + consulPort,
		}
	})
	return consul
}
