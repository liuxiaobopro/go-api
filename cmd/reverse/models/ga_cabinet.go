package models

import (
	"time"
)

type GaCabinet struct {
	Id           string    `xorm:"not null pk autoincr UNSIGNED INT(11)"`
	Title        string    `xorm:"not null default '' comment('名称') unique VARCHAR(50)"`
	Number       int       `xorm:"not null comment('编号') unique INT(11)"`
	RdId         int       `xorm:"not null comment('部门id') INT(11)"`
	MailNum      int       `xorm:"not null comment('当前柜内邮件数') INT(11)"`
	UpdateMid    int       `xorm:"default 0 comment('开柜后,最后更新的邮件id') INT(11)"`
	LastMid      int       `xorm:"default 0 comment('最后id') INT(11)"`
	OpenStatus   int       `xorm:"not null default 1 comment('是否开门(1关门  2程序开门  3手动开门)') index(open_status) TINYINT(1)"`
	OnlineStatus int       `xorm:"not null default 1 comment('是否在运行(1开机  2关机)') index(open_status) TINYINT(1)"`
	LastOpenTime time.Time `xorm:"comment('最后开锁时间') DATETIME"`
	IsSelected   int       `xorm:"not null default 1 comment('码枪是否被选中(1未选中  2选中)') TINYINT(1)"`
	UpdateTime   time.Time `xorm:"comment('修改时间') DATETIME"`
	DeleteTime   time.Time `xorm:"comment('删除时间') DATETIME"`
}

func (m *GaCabinet) TableName() string {
	return "ga_cabinet"
}

func (m *GaCabinet) TableComment() string {
	return "ga_cabinet"
}
