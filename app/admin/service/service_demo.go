package service

import (
	"github.com/liuxiaobopro/go-api/app/admin/dao"
	"github.com/liuxiaobopro/go-api/app/admin/model"

	"github.com/liuxiaobopro/go-lib/ecode"
	"github.com/liuxiaobopro/go-lib/response"
)

type DemoServiceType struct{}

var DemoSrv = new(DemoServiceType)

func (*DemoServiceType) GetList(req *model.DemoSearchModel) *response.Result {
	return dao.DemoDao.List(req)
}

func (*DemoServiceType) Add(req *model.Demo) (int, ecode.BizErr) {
	return dao.DemoDao.Add(req)
}
