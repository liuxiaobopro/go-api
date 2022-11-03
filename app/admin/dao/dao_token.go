package dao

import (
	"github.com/liuxiaobopro/go-api/app/admin/model"
	"github.com/liuxiaobopro/go-api/global"

	"github.com/liuxiaobopro/go-lib/console"
	"github.com/liuxiaobopro/go-lib/ecode"
	"xorm.io/xorm"
)

type TokenDaoType struct{}

var TokenDao = new(TokenDaoType)

func sessToken() *xorm.Session {
	return global.Db.Table(global.Conf.Database.Mdb.Prefix + "token")
}

// Add 添加token
func (th *TokenDaoType) Add(token *model.Token) (int, ecode.BizErr) {
	insert, err := sessToken().Insert(&token)
	if err != nil {
		console.Console.Error("添加token失败", "")
		return 0, ecode.ERROR_SERVER
	}
	return int(insert), ecode.SUCCSESS
}

// UpdateById 通过id修改token
func (th *TokenDaoType) UpdateById(id int, token *model.Token) ecode.BizErr {
	_, err := sessToken().Where("id = ?", id).Update(token)
	if err != nil {
		console.Console.Error("通过id修改token失败", "")
		return ecode.ERROR_SERVER
	}
	return ecode.SUCCSESS
}

// UpdateByUid 根据uid修改token
func (th *TokenDaoType) UpdateByUid(uid int, token *model.Token) ecode.BizErr {
	_, err := sessToken().Where("uid = ?", uid).Update(token)
	if err != nil {
		console.Console.Error("根据uid修改token失败", "")
		return ecode.ERROR_SERVER
	}
	return ecode.SUCCSESS
}

// GetByToken 通过token获取用户信息
func (th *TokenDaoType) GetByToken(token string) (*model.Token, ecode.BizErr) {
	var info model.Token
	_, err := sessToken().Where("token = ?", token).Get(&info)
	if err != nil {
		console.Console.Error("通过token获取用户信息失败", "")
		return nil, ecode.ERROR_SERVER
	}
	return &info, ecode.SUCCSESS
}

// GetByUserId 通过用户id获取token
func (th *TokenDaoType) GetByUserId(userId int) (*model.Token, ecode.BizErr) {
	var info model.Token
	has, err := sessToken().Where("uid = ?", userId).Get(&info)
	if err != nil {
		console.Console.Error("通过用户id获取token失败", "")
		return nil, ecode.ERROR_SERVER
	}
	if !has {
		return nil, ecode.ERROR_RESOURCE_DONT_EXISTS
	}
	return &info, ecode.SUCCSESS
}

// UpdateByUidField 通过uid修改指定字段
func (th *TokenDaoType) UpdateByUidField(uid int, field string, value interface{}) ecode.BizErr {
	_, err := sessToken().Where("uid = ?", uid).Cols(field).Update(map[string]interface{}{field: value})
	if err != nil {
		console.Console.Error("通过uid修改指定字段失败", "")
		return ecode.ERROR_SERVER
	}
	return ecode.SUCCSESS
}
