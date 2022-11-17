package controller

import (
	"github.com/liuxiaobopro/go-api/app/model"
	"github.com/liuxiaobopro/go-api/app/service"
	"github.com/liuxiaobopro/go-api/global"

	"github.com/gin-gonic/gin"
	"github.com/liuxiaobopro/go-lib/ecode"
	ginl "github.com/liuxiaobopro/go-lib/gin"
)

type DemoControllerType struct {
	ginl.Handler
}

var DemoController = new(DemoControllerType)

func (th *DemoControllerType) Conf(c *gin.Context) {
	th.SendSucc(c, global.Conf, nil)
}

func (th *DemoControllerType) Add(c *gin.Context) {
	var demo = new(model.Demo)
	if err := c.ShouldBindJSON(&demo); err != nil {
		th.SendError(c, err, ecode.ERROR_PARAMETER_EXCEPTION)
		return
	}
	res, _ := service.DemoSrv.Add(demo)
	th.SendSucc(c, res, nil)
}
