package main

import (
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/config"
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/log"
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/server"
	"go.uber.org/zap"
	)

func main() {
	logConfig := config.GetAppConfig().Application.Logging
	var logger *zap.Logger
	if logConfig.Env == "production" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync() // flushes buffer, if any
	log.ConfigureLogger(logger)
	err := server.StartServer()
	if err != nil {
		panic(err)
	}
}




