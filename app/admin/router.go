package admin

import (
	"github.com/liuxiaobopro/go-api/app/admin/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.RouterGroup) {
	// demo
	g.GET("/demo/conf", controller.DemoController.Conf)
	g.POST("/demo/add", controller.DemoController.Add)

	// 文件相关
	g.POST("/file/uploadqiniu", controller.FileController.UploadQiniu) // 上传文件到七牛云
	g.POST("/file/uploadlocal", controller.FileController.UploadLocal) // 上传文件到本地

	// 用户相关
	g.POST("/user/logout", controller.UserController.Logout) // 退出登录
	g.PUT("/user/:id", controller.UserController.Update)     // 修改用户
	g.DELETE("/user/:id", controller.UserController.Delete)  // 删除用户
	g.GET("/user", controller.UserController.List)           // 用户列表
}
