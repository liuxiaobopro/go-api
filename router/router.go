package router

import (
	"github.com/liuxiaobopro/go-api/app"
	"github.com/liuxiaobopro/go-api/app/common/middleware"
	"github.com/liuxiaobopro/go-api/app/controller"

	"github.com/gin-gonic/gin"
	"github.com/liuxiaobopro/go-lib/response"
)

func InitRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, response.GetSuccRes(nil))
	})

	r.POST("/admin/user/login", controller.UserController.Login)
	r.POST("/admin/user", controller.UserController.Add)

	g1 := r.Group("/admin")
	g1.Use(middleware.Jwt())
	app.InitRouter(g1)
}
