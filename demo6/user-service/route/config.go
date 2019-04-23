package route

import (
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/controller"
	"git.code.oa.com/tcnp/arch-revolution/demo6/user-service/tracer"
	"github.com/gin-gonic/gin"
	"net/http"

)

func ConfigureRoutes(r *gin.Engine){
	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	root := r.Group("/api/v1")

	userController := controller.NewUserController()

	root.GET("/users", tracer.ExtractSpanContext, userController.GetUsers)
	root.POST("/users", tracer.ExtractSpanContext, userController.AddUser)
	root.GET("/users/:id", tracer.ExtractSpanContext, userController.GetUser)
	root.PUT("/users/:id", tracer.ExtractSpanContext, userController.UpdateUser)
	root.DELETE("/users/:id", tracer.ExtractSpanContext, userController.DeleteUser)

}
