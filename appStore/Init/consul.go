package Init

import (
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/registry"
	"strconv"
)

func SetupConsul(consulHost string, consulPort int) registry.Registry {
	newConsul := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":" + strconv.Itoa(consulPort),
		}
	})
	return newConsul
}
