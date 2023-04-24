package consul

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/registry"
	"strconv"
)

func Init(consulHost string, consulPort int) registry.Registry {
	NewConsul := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":" + strconv.Itoa(consulPort),
		}
	})
	return NewConsul
}
