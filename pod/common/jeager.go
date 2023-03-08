package common

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

// go get github.com/asim/go-mirco/plugins/wrapper/trace/opentracing/v3
// 失败, err:Failed connect to github.com:443; Connection refused
// 原因: 下载地址有误

func NewTracer(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	// 配置并创建Jaeger Tracer
	cfg := &jaegerCfg.Configuration{
		// 服务名称
		ServiceName: serviceName,
		// 允许初始化非默认采样器
		Sampler: &jaegerCfg.SamplerConfig{
			// 指定采样器的类型
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		// 配置报告程序
		Reporter: &jaegerCfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  addr,
		},
	}
	return cfg.NewTracer()
}
