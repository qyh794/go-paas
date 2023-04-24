package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(ServiceConfig)

type ServiceConfig struct {
	Name        string `mapstructure:"name"`
	ServiceHost string `mapstructure:"service_Host"`
	ServicePort string `mapstructure:"service_Port"`
	Version     string `mapstructure:"version"`
	*Consul     `mapstructure:"consul"`
	*Tracer     `mapstructure:"tracer"`
	*Prometheus `mapstructure:"prometheus"`
	*Hystrix    `mapstructure:"hystrix"`
}

type Hystrix struct {
	Port int `mapstructure:"port"`
}

type Consul struct {
	Host   string `mapstructure:"host"`
	Prefix string `mapstructure:"prefix"`
	Port   int    `mapstructure:"port"`
}

type Tracer struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type Prometheus struct {
	Port int `mapstructure:"port"`
}

// 从配置文件里读取配置
func Init() (err error) {
	viper.SetConfigFile("./conf/config.yaml")
	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {            // 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}

	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}
	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})
	return
}
