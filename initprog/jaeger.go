package initprog

import (
	"fmt"
	"github.com/liuhangkaixcode/websocket/global"

	"github.com/opentracing/opentracing-go"

	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"

)

func initJaeger()  {
	service:="websocketService"
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			LocalAgentHostPort:global.Global_Config_Manger.Jaeger.Host,
		},
	}
	tracer, _, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)

	global.Global_Jaeger=tracer
}
