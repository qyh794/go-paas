package common

import (
	consul2 "github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/registry"
	"strconv"
)

func InitConsul(consulHost string, consulPort int64) registry.Registry {
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":" + strconv.FormatInt(consulPort, 10),
		}
	})
	return consul
}
