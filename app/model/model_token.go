package model

import jsonl "github.com/liuxiaobopro/go-lib/json"

type Token struct {
	Id             int         `json:"id" xorm:"pk autoincr unsigned not null int(11) 'id'"`
	Uid            int         `json:"uid" xorm:"default 0 not null int(11) 'uid'"`
	Token          string      `json:"token" xorm:"unique not null varchar(50) 'token'"`
	ExpirationTime int         `json:"expiration_time" xorm:"not null comment('过期时间') int(11) 'expiration_time'"`
	CreateTime     jsonl.Time  `json:"create_time" xorm:"default '0000-00-00 00:00:00' comment('创建时间') datetime 'create_time' created"`
	UpdateTime     jsonl.Time  `json:"-" xorm:"default '0000-00-00 00:00:00' comment('修改时间') datetime 'update_time' updated"`
	DeleteTime     *jsonl.Time `json:"-" xorm:"default '0000-00-00 00:00:00' comment('删除时间') datetime 'delete_time' deleted"`
	LastUpdateId   string      `json:"-" xorm:"default '' comment('最后修改用户标识') varchar(50) 'last_update_id'"`
}

// TableComment 设置表注释
func (th *Token) TableComment() string {
	return "token表"
}
