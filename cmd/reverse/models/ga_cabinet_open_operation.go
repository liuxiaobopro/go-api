package models

import (
	"time"
)

type GaCabinetOpenOperation struct {
	Id         string    `xorm:"not null pk autoincr UNSIGNED INT(11)"`
	Command    int       `xorm:"not null comment('开柜命令') INT(11)"`
	Uid        int       `xorm:"not null comment('操作人') INT(11)"`
	Cid        int       `xorm:"not null comment('快递柜id') INT(11)"`
	Status     int       `xorm:"not null comment('状态(1待执行  2正在执行  3执行成功  4执行失败)') index TINYINT(1)"`
	CreateTime time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"comment('修改时间') DATETIME"`
	DeleteTime time.Time `xorm:"comment('删除时间') DATETIME"`
}

func (m *GaCabinetOpenOperation) TableName() string {
	return "ga_cabinet_open_operation"
}

func (m *GaCabinetOpenOperation) TableComment() string {
	return "ga_cabinet_open_operation"
}
