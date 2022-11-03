package models

import (
	"time"
)

type GaMail struct {
	Id           string    `xorm:"not null pk autoincr index(id) UNSIGNED INT(11)"`
	WaybillNo    string    `xorm:"not null comment('运单号') unique VARCHAR(100)"`
	CId          int       `xorm:"not null comment('快递柜id') INT(11)"`
	RdId         int       `xorm:"not null comment('部门id') INT(11)"`
	RegisterTime time.Time `xorm:"comment('登记时间') DATETIME"`
	PickupTime   time.Time `xorm:"comment('取件时间') DATETIME"`
	FormType     int       `xorm:"comment('来源(1投递  2扫码枪  3手动录入)') index(id) TINYINT(1)"`
	Remark       string    `xorm:"comment('备注') TEXT"`
	CreateTime   time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateTime   time.Time `xorm:"comment('修改时间') DATETIME"`
	DeleteTime   time.Time `xorm:"comment('删除时间') DATETIME"`
	LastUpdateId string    `xorm:"default '' comment('最后修改用户标识') VARCHAR(50)"`
}

func (m *GaMail) TableName() string {
	return "ga_mail"
}

func (m *GaMail) TableComment() string {
	return "ga_mail"
}
