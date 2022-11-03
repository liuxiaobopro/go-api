package models

import (
	"time"
)

type GaToken struct {
	Id             string    `xorm:"not null pk autoincr UNSIGNED INT(11)"`
	Uid            int       `xorm:"not null default 0 INT(11)"`
	Token          string    `xorm:"not null unique VARCHAR(50)"`
	ExpirationTime int       `xorm:"not null comment('过期时间') INT(11)"`
	CreateTime     time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateTime     time.Time `xorm:"comment('修改时间') DATETIME"`
	DeleteTime     time.Time `xorm:"comment('删除时间') DATETIME"`
	LastUpdateId   string    `xorm:"default '' comment('最后修改用户标识') VARCHAR(50)"`
}

func (m *GaToken) TableName() string {
	return "ga_token"
}

func (m *GaToken) TableComment() string {
	return "ga_token"
}
