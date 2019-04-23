package server

import (
	"fmt"
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/config"
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/route"
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/log"
	"github.com/gin-gonic/gin"
)

func StartServer() error {
	r := gin.Default()
	route.ConfigureRoutes(r)

	serverBindAddr := fmt.Sprintf(":%d", config.GetAppConfig().Server.Port)
	log.Infof("Run server[service-name=%s] at %s", config.GetAppConfig().Application.Service.Name, serverBindAddr)
	return r.Run(serverBindAddr) // listen and serve
}
