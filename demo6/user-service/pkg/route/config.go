package route

import (
	_ "git.code.oa.com/tcnp/arch-revolution/demo6/user-service/docs"
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/pkg/controller"
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/pkg/tracer"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger user-service
// @version 1.0
// @description This is a user-service api server.
// @BasePath /api/v1
func ConfigureRoutes(r *gin.Engine){
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	root := r.Group("/api/v1")

	root.Use(tracer.ExtractSpanContextMiddleware)

	userController := controller.NewUserController()

	root.GET("/users", userController.GetUsers)
	root.POST("/users", userController.AddUser)
	root.GET("/users/:id", userController.GetUser)
	root.PUT("/users/:id", userController.UpdateUser)
	root.DELETE("/users/:id", userController.DeleteUser)

}
