package app

import (
	controller2 "github.com/liuxiaobopro/go-api/app/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.RouterGroup) {
	// demo
	g.GET("/demo/conf", controller2.DemoController.Conf)
	g.POST("/demo/add", controller2.DemoController.Add)

	// 文件相关
	g.POST("/file/uploadqiniu", controller2.FileController.UploadQiniu) // 上传文件到七牛云
	g.POST("/file/uploadlocal", controller2.FileController.UploadLocal) // 上传文件到本地

	// 用户相关
	g.POST("/user/logout", controller2.UserController.Logout) // 退出登录
	g.PUT("/user/:id", controller2.UserController.Update)     // 修改用户
	g.DELETE("/user/:id", controller2.UserController.Delete)  // 删除用户
	g.GET("/user", controller2.UserController.List)           // 用户列表
}
