package models

import (
	"time"
)

type GaUser struct {
	Id           string    `xorm:"not null pk autoincr UNSIGNED INT(11)"`
	Nickname     string    `xorm:"not null default '' comment('昵称') VARCHAR(50)"`
	Username     string    `xorm:"not null default '' comment('账号') VARCHAR(50)"`
	Password     string    `xorm:"not null comment('密码') TEXT"`
	Phone        string    `xorm:"default '' comment('电话号') CHAR(11)"`
	CreateTime   time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateTime   time.Time `xorm:"comment('修改时间') DATETIME"`
	DeleteTime   time.Time `xorm:"comment('删除时间') DATETIME"`
	LastUpdateId string    `xorm:"default '' comment('最后修改用户标识') VARCHAR(50)"`
}

func (m *GaUser) TableName() string {
	return "ga_user"
}

func (m *GaUser) TableComment() string {
	return "ga_user"
}
