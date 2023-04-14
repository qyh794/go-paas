package main

import (
	"fmt"
	"github.com/asim/go-micro/v3/logger"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/opentracing/opentracing-go"
	"github.com/qyh794/go-paas/common"
	"github.com/qyh794/go-paas/pod/domain/repository"
	"github.com/qyh794/go-paas/pod/domain/service"
	"github.com/qyh794/go-paas/pod/handler"
	"github.com/qyh794/go-paas/pod/init/consulInit"
	"github.com/qyh794/go-paas/pod/init/k8s"
	"github.com/qyh794/go-paas/pod/init/mysql"
	"github.com/qyh794/go-paas/pod/init/serviceInit"
	"github.com/qyh794/go-paas/pod/proto/pod"
	"github.com/qyh794/go-paas/pod/settings"
)

func main() {
	// 0. 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init setting failed, err:%v", err)
		return
	}
	// 1.注册中心
	newRegistry := consulInit.Init(settings.Conf.Consul.Host, settings.Conf.Consul.Port)
	// 2.配置中心
	consulConfig, err := common.GetConsulConfig(settings.Conf.Consul.Host, settings.Conf.Consul.Port, "/micro/config")
	if err != nil {
		logger.Error(err)
	}
	// 3.获取配置中心数据
	mysqlConfig := common.GetMysqlConfigFromConsul(consulConfig, "mysql")
	//fmt.Printf("user:%s\n password:%s\n database:%s\n port:%s\n host:%s\n", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Database, mysqlConfig.Port, mysqlConfig.Host)

	// 4.初始化数据库连接
	if err = mysql.Init("mysql", mysqlConfig); err != nil {
		logger.Error(err)
	}
	defer mysql.Close()

	// 5.添加链路追踪
	tracer, closer, err := common.NewTracer(settings.Conf.Name, settings.Conf.Tracer.Host, settings.Conf.Tracer.Port)
	if err != nil {
		logger.Error(err)
	}
	defer func() {
		_ = closer.Close()
	}()
	opentracing.SetGlobalTracer(tracer)

	//// 6.添加熔断器作为客户端才使用,返回能够通过HTTP公开仪表板指标的服务器
	//streamHandler := hystrix.NewStreamHandler()
	//// Start()开始观察内存断路器的指标
	//streamHandler.Start()

	// 8.添加日志中心
	//1）需要程序日志打入到日志文件中
	//2）在程序中添加filebeat.yml 文件
	//3) 启动filebeat，启动命令 ./filebeat -e -c filebeat.yml
	// 9.添加监控
	common.PrometheusBoot(settings.Conf.Prometheus.Port)
	clientset := k8s.Init()
	// 13.创建服务实例
	newService := serviceInit.Init(settings.Conf.ServiceHost, settings.Conf.ServicePort, settings.Conf.Name, settings.Conf.Version, newRegistry)
	// 14.初始化服务
	newService.Init()
	//err = repository.NewPodRepository(db).InitTable()
	//if err != nil {
	//	logger.Fatal(err)
	//}
	// 15.注册句柄
	podDateService := service.NewPodDateService(repository.NewPodRepository(mysql.DB), clientset)
	_ = pod.RegisterPodHandler(newService.Server(), &handler.PodHandler{PodDataService: podDateService})
	// 16.启动服务
	if err = newService.Run(); err != nil {
		logger.Fatal(err)
	}
}
