package dao

import (
	"github.com/liuxiaobopro/go-api/app/model"
	"github.com/liuxiaobopro/go-api/global"

	"github.com/liuxiaobopro/go-lib/ecode"
	"github.com/liuxiaobopro/go-lib/response"
	"xorm.io/xorm"
)

type DemoDaoType struct{}

var DemoDao = new(DemoDaoType)

func sessDemo() *xorm.Session {
	return global.Db.Table(global.Conf.Database.Mdb.Prefix + "demo")
}

// List 商品列表
func (th *DemoDaoType) List(req *model.DemoSearchModel) *response.Result {
	return response.GetSuccRes(req)
}

// Add 添加商品
func (th *DemoDaoType) Add(req *model.Demo) (int, ecode.BizErr) {
	insert, err := sessDemo().Insert(req)
	if err != nil {
		return 0, ecode.ERROR_SERVER
	}
	return int(insert), ecode.SUCCSESS
}
