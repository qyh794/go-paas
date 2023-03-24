package common

import (
	"github.com/asim/go-micro/plugins/config/source/consul/v3"
	"github.com/asim/go-micro/v3/config"
	"strconv"
)

//func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
//	consulSource := consul.NewSource(
//		// 配置中心地址
//		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
//		// 设置前缀
//		consul.WithPrefix(prefix),
//		// 配置项中删除前缀，还是保留前缀
//		consul.StripPrefix(true),
//	)
//	conf, err := config.NewConfig()
//	if err != nil {
//		return conf, err
//	}
//	err = conf.Load(consulSource)
//	return conf, err
//}

func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		consul.WithPrefix(prefix),
		consul.StripPrefix(true),
	)
	cf, err := config.NewConfig()
	if err != nil {
		return cf, err
	}
	//err = config.Load(consulSource)  <-----原先错误的写法
	err = cf.Load(consulSource)
	return cf, err
}
