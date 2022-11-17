package model

import jsonl "github.com/liuxiaobopro/go-lib/json"

type User struct {
	Id           int         `json:"id" xorm:"pk autoincr unsigned not null int(11) 'id'"`
	Nickname     string      `json:"nickname" xorm:"default '' not null comment('昵称') varchar(50) 'nickname'"`
	Username     string      `json:"username" xorm:"default '' not null comment('账号') varchar(50) 'username'"`
	Password     string      `json:"password" xorm:"default '' not null comment('密码') varchar(100) 'password'"`
	Phone        string      `json:"phone" xorm:"default '' comment('电话号') char(11) 'phone'"`
	CreateTime   jsonl.Time  `json:"create_time" xorm:"default '0000-00-00 00:00:00' comment('创建时间') datetime 'create_time' created"`
	UpdateTime   jsonl.Time  `json:"-" xorm:"default '0000-00-00 00:00:00' comment('修改时间') datetime 'update_time' updated"`
	DeleteTime   *jsonl.Time `json:"-" xorm:"default '0000-00-00 00:00:00' comment('删除时间') datetime 'delete_time' deleted"`
	LastUpdateId string      `json:"-" xorm:"default '' comment('最后修改用户标识') varchar(50) 'last_update_id'"`
}

// TableComment 设置表注释
func (th *User) TableComment() string {
	return "用户表"
}

// AddUser 添加用户
type AddUser struct {
	Nickname   string `json:"nickname" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
}

// DetailUser 用户信息
type DetailUser struct {
	Id         int        `json:"id"`
	Nickname   string     `json:"nickname"`
	Username   string     `json:"username"`
	Phone      string     `json:"phone"`
	CreateTime jsonl.Time `json:"create_time"`
}

// LoginUser 用户登录
type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UpdateUser 修改用户
type UpdateUser struct {
	Nickname   string     `json:"nickname"`
	Password   string     `json:"password"`
	Phone      string     `json:"phone"`
	UpdateTime jsonl.Time `json:"-" xorm:"default '0000-00-00 00:00:00' comment('修改时间') datetime 'update_time' updated"`
}

// UserList 用户列表
type SearchUser struct {
	Page            int        `json:"page" binding:"required"`
	Limit           int        `json:"limit" binding:"required"`
	CreateTimeStart jsonl.Time `json:"create_time_start"`
	CreateTimeEnd   jsonl.Time `json:"create_time_end"`
}
