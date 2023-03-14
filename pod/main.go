package main

import (
	"flag"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	ratelimit "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v3"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/opentracing/opentracing-go"
	"github.com/qyh794/go-paas/pod/common"
	"github.com/qyh794/go-paas/pod/domain/repository"
	"github.com/qyh794/go-paas/pod/domain/service"
	"github.com/qyh794/go-paas/pod/handler"
	hystrix2 "github.com/qyh794/go-paas/pod/plugin/hystrix"
	"github.com/qyh794/go-paas/pod/proto/pod"
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
	//熔断器
	hystrixPort = 9091
	//监控端口
	prometheusPort = 9191

	QPS = 1000
)

func main() {
	// Delete pod directory
	// 1.注册中心
	consul := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":" + strconv.FormatInt(consulPort, 10),
		}
	})
	// 2.配置中心
	consulConfig, err := common.GetConsulConfig(consulHost, consulPort, "/micro/config")
	if err != nil {
		logger.Error(err)
	}
	// 3.获取配置中心数据
	mysqlConfig := common.GetMysqlConfigFromConsul(consulConfig, "mysql")
	//fmt.Printf("user:%s\n password:%s\n database:%s\n port:%s\n host:%s\n", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Database, mysqlConfig.Port, mysqlConfig.Host)

	// 4.初始化数据库连接
	db, err := gorm.Open("mysql", mysqlConfig.User+":"+mysqlConfig.Password+"@("+
		mysqlConfig.Host+":"+mysqlConfig.Port+")/"+mysqlConfig.Database+
		"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logger.Error(err)
	}
	defer func() {
		_ = db.Close()
	}()
	db.SingularTable(true)
	// 5.添加链路追踪
	tracer, closer, err := common.NewTracer("go.micro.service.pod", tracerHost+":"+strconv.Itoa(tracerPort))
	if err != nil {
		logger.Error(err)
	}
	defer func() {
		_ = closer.Close()
	}()
	opentracing.SetGlobalTracer(tracer)
	// 6.添加熔断器,返回能够通过HTTP公开仪表板指标的服务器
	streamHandler := hystrix.NewStreamHandler()
	// Start()开始观察内存断路器的指标
	streamHandler.Start()
	// 7.添加监听程序
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", strconv.Itoa(hystrixPort)), streamHandler)
		if err != nil {
			logger.Error(err)
		}
	}()
	// 8.添加日志中心
	//1）需要程序日志打入到日志文件中
	//2）在程序中添加filebeat.yml 文件
	//3) 启动filebeat，启动命令 ./filebeat -e -c filebeat.yml
	// 9.添加监控
	common.PrometheusBoot(prometheusPort)
	// 10.创建k8s连接
	var kubeConfig *string
	if dir := homedir.HomeDir(); dir != "" {
		kubeConfig = flag.String("kubeconfig", filepath.Join(dir, ".kube", "config"), "kubeconfig file 在当前系统中的地址")
	} else {
		kubeConfig = flag.String("kubeconfig", "", "kubeconfig file 在当前系统中的地址")
	}
	flag.Parse()
	// 11.创建config
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		logger.Fatal(err.Error())
	}
	// 12.创建集群外可操作的k8s客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Error(err)
	}
	// 13.创建服务实例
	newService := micro.NewService(
		// 自定义服务地址
		micro.Server(server.NewServer(func(options *server.Options) {
			options.Advertise = serviceHost + ":" + servicePort
		})),
		micro.Name("go.micro.service.pod"),
		micro.Version("latest"),
		micro.Address(":"+servicePort),
		micro.Registry(consul),
		// 添加链路追踪
		// 服务端添加,将处理程序Wrapper添加到传递到服务器的选项列表中
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		// 客户端添加
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		// 作为客户端使用添加熔断
		micro.WrapClient(hystrix2.NewClientHystrixWrapper()),
		// 添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)
	// 14.初始化服务
	newService.Init()
	//err = repository.NewPodRepository(db).InitTable()
	//if err != nil {
	//	logger.Fatal(err)
	//}
	// 15.注册句柄
	podDateService := service.NewPodDateService(repository.NewPodRepository(db), clientset)
	_ = pod.RegisterPodHandler(newService.Server(), &handler.PodHandler{PodDataService: podDateService})
	// 16.启动服务
	if err := newService.Run(); err != nil {
		logger.Fatal(err)
	}
}
