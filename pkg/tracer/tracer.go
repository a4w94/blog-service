package tracer

import (
	"io"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	//jaeger client 設定項目
	cfg := &config.Configuration{
		ServiceName: serviceName,
		//固定取樣 對所有資料都進行取樣
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,            //是否啟用LoggingReporter
			BufferFlushInterval: 1 * time.Second, //更新緩衝區的頻率
			LocalAgentHostPort:  agentHostPort,   //上報的Agent位址
		},
	}

	//初始化Tracer物件
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}

	//設定全域的Tracer物件
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
