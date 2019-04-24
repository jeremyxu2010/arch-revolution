package main

import (
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/pkg/config"
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/pkg/log"
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/pkg/server"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	jaegerconfig "github.com/uber/jaeger-client-go/config"
	)

func main() {
	// Configure logging
	logConfig := config.GetAppConfig().Application.Logging
	var logger *zap.Logger
	if logConfig.Env == "production" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync() // flushes buffer, if any
	log.ConfigureLogger(logger)

	// Configure global tracer to help span context information propagation
	tracer, closer, err := jaegerconfig.Configuration{
		ServiceName: config.GetAppConfig().Application.Service.Name,
	}.NewTracer()
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	// Restful api server startup
	err = server.StartServer()
	if err != nil {
		panic(err)
	}
}




