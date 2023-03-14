package main

import (
	"flag"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3"
	"github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/jinzhu/gorm"
	opentracing2 "github.com/opentracing/opentracing-go"
	"github.com/qyh794/go-paas/pod/common"
	"github.com/qyh794/go-paas/svc/domain/repository"
	"github.com/qyh794/go-paas/svc/domain/service"
	"github.com/qyh794/go-paas/svc/handler"
	"github.com/qyh794/go-paas/svc/proto/svc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"net"
	"net/http"
	"path/filepath"
	"strconv"
)

var (
	//服务地址
	hostIp = "192.168.0.105"
	//服务地址
	serviceHost = hostIp
	//服务端口
	servicePort = "8083"

	//注册中心配置
	consulHost       = hostIp
	consulPort int64 = 8500
	//链路追踪
	tracerHost = hostIp
	tracerPort = 6831
	//熔断端口，每个服务不能重复
	hystrixPort = 9093
	//监控端口，每个服务不能重复
	prometheusPort = 9193
	QPS            = 1000
)

func main() {
	// 注册中心
	newRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":" + strconv.FormatInt(consulPort, 10),
		}
	})
	// 配置中心
	consulConfig, err := common.GetConsulConfig(consulHost, consulPort, "/micro/config")
	if err != nil {
		logger.Error(err)
	}
	// 从配置中心获取MySQL配置
	mysqlConfig := common.GetMysqlConfigFromConsul(consulConfig, "mysql")
	// 初始化数据库链接
	db, err := gorm.Open("mysql", mysqlConfig.User+":"+mysqlConfig.Password+"@("+mysqlConfig.Host+":3306)/"+mysqlConfig.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logger.Error(err)
	}
	defer func() {
		_ = db.Close()
	}()
	// 添加链路追踪
	tracer, io, err := common.NewTracer("go.micro.service.svc", tracerHost+":"+strconv.Itoa(tracerPort))
	if err != nil {
		logger.Error(err)
	}
	defer func() {
		_ = io.Close()
	}()
	opentracing2.SetGlobalTracer(tracer)
	// 添加熔断
	streamHandler := hystrix.NewStreamHandler()
	// 启动熔断
	streamHandler.Start()
	// 监听程序
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", strconv.Itoa(hystrixPort)), streamHandler)
		if err != nil {
			return
		}
	}()
	// 添加监控
	common.PrometheusBoot(prometheusPort)
	// 创建k8s连接
	var kubeconfig *string
	// 返回当前用户的主目录
	if dir := homedir.HomeDir(); dir != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(dir, ".kube", "config"), "")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		logger.Error(err)
	}
	// 创建k8s客户端,为给定的配置创建一个新的ClientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Error(err)
	}
	// 创建服务
	newService := micro.NewService(
		micro.Server(server.NewServer(func(options *server.Options) {
			options.Advertise = serviceHost + ":" + servicePort
		})),
		micro.Name("go.micro.service.svc"),
		micro.Version("latest"),
		micro.Address(":"+servicePort),
		micro.Registry(newRegistry),
		micro.WrapHandler(opentracing.NewHandlerWrapper(opentracing2.GlobalTracer())),
		micro.WrapClient(opentracing.NewClientWrapper(opentracing2.GlobalTracer())),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)
	// 建表
	svcRepository := repository.NewSvcRepository(db)
	err = svcRepository.InitTable()
	if err != nil {
		logger.Error(err)
	}
	SvcDateService := service.NewServiceDateService(svcRepository, clientSet)
	// 注册句柄
	_ = svc.RegisterSvcHandler(newService.Server(), &handler.SvcHandler{SvcDataService: SvcDateService})
	// 启动服务
	err = newService.Run()
	if err != nil {
		logger.Error(err)
	}
}
