package gin

import (
	"fmt"
	"os"

	"github.com/liuxiaobopro/go-api/core/ws"
	"github.com/liuxiaobopro/go-api/global"
	"github.com/liuxiaobopro/go-api/router"

	"github.com/gin-gonic/gin"
	"github.com/liuxiaobopro/go-lib/console"
	"github.com/liuxiaobopro/go-lib/gin/middleware/cors"
	"github.com/liuxiaobopro/go-lib/gin/middleware/logger"
)

func InitGin() {
	r := gin.Default()

	logger.Configs.Path = global.Conf.App.LoggerFolderPath
	logger.Configs.IsRetuenErr = true // Panic返回错误信息

	r.Use(logger.Log)
	r.Use(cors.Cors())

	if global.Conf.Ws.IsOpen {
		// 初始化websocket
		ws.InitWs(r)
	}

	// 初始化路由
	router.InitRouter(r)

	ip := fmt.Sprintf("%s:%d", global.Conf.App.Host, global.Conf.App.Port)

	console.Console.Info(fmt.Sprintf("服务启动成功，地址：http://%s", ip))
	console.Console.Info(fmt.Sprintf("Runmode: %s", global.Conf.App.Runmode))

	console.Console.Info(fmt.Sprintf("是否首次安装1: %v", !global.Conf.App.IsInstall))
	generateLockFile()
	console.Console.Info(fmt.Sprintf("是否首次安装2: %v", !global.Conf.App.IsInstall))

	if err := r.Run(ip); err != nil {
		console.Console.Error("Gin启动失败", err.Error())
	}
}

// generateLockFile 生成锁文件
func generateLockFile() {
	if !global.Conf.App.IsInstall {
		// 创建锁文件
		_, err := os.Create(global.Conf.App.InstallFileName)
		if err != nil {
			console.Console.Error("创建锁文件失败", err.Error())
		}
		global.Conf.App.IsInstall = true
	}
}
