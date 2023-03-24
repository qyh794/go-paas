package main

import (
	"flag"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/jinzhu/gorm"
	"github.com/micro/micro/v3/service/logger"
	"github.com/opentracing/opentracing-go"
	"github.com/qyh794/go-paas/pod/common"
	"github.com/qyh794/go-paas/volume/domain/repository"
	service2 "github.com/qyh794/go-paas/volume/domain/service"
	"github.com/qyh794/go-paas/volume/handler"
	"github.com/qyh794/go-paas/volume/proto/volume"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"strconv"
)

var (
	//服务地址
	hostIp = "0.0.0.0"
	//服务地址
	serviceHost = hostIp
	//服务端口
	servicePort = "8081"

	//注册中心配置
	consulHost       = hostIp
	consulPort int64 = 8500
	//链路追踪
	tracerHost = hostIp
	tracerPort = 6831
	//熔断端口，每个服务不能重复
	//hystrixPort = 9092
	//监控端口，每个服务不能重复
	prometheusPort = 9192
)

const QPS = 1000

func main() {
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":" + strconv.FormatInt(consulPort, 10),
		}
	})
	consulConfig, err := common.GetConsulConfig(consulHost, consulPort, "/micro/config")
	if err != nil {
		logger.Error(err)
	}
	mysqlConfig := common.GetMysqlConfigFromConsul(consulConfig, "mysql")
	db, err := gorm.Open("mysql", mysqlConfig.User+":"+mysqlConfig.Password+"@("+mysqlConfig.Host+":3306)/"+mysqlConfig.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logger.Error(err)
	}
	defer func() {
		_ = db.Close()
	}()
	db.SingularTable(true)
	tracer, io, err := common.NewTracer("go.micro.service.volume", tracerHost+":"+strconv.Itoa(tracerPort))
	if err != nil {
		logger.Error(err)
	}
	defer func() {
		_ = io.Close()
	}()
	opentracing.SetGlobalTracer(tracer)
	common.PrometheusBoot(prometheusPort)
	var kubeConfig *string
	if dir := homedir.HomeDir(); dir != "" {
		kubeConfig = flag.String("kubeconfig", filepath.Join(dir, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		logger.Error(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Error(err)
	}
	service := micro.NewService(
		micro.Server(server.NewServer(func(options *server.Options) {
			options.Advertise = serviceHost + ":" + servicePort
		})),
		micro.Name("go.micro.service.volume"),
		micro.Version("latest"),
		micro.Address(":"+servicePort),
		micro.Registry(consulRegistry),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(1000)),
	)
	service.Init()
	volumeDataService := service2.NewVolumeDataService(clientset, repository.NewVolumeRepository(db))
	err = volume.RegisterVolumeHandler(service.Server(), &handler.VolumeHandler{VolumeDataService: volumeDataService})
	if err != nil {
		logger.Error(err)
	}
	if err = service.Run(); err != nil {
		logger.Error(err)
	}
}
