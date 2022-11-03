package errcode

import (
	"github.com/liuxiaobopro/go-lib/ecode"
)

var (
	USER_NOT_EXIST                 = &ecode.Res{Code: 10001, Desc: "用户不存在"}
	TWO_PASSWORDS_ARE_INCONSISTENT = &ecode.Res{Code: 10002, Desc: "两次密码不一致"}
	USER_ALREADY_EXIST             = &ecode.Res{Code: 10003, Desc: "用户已存在"}
	ACCOUNT_PASSWORD_ERROR         = &ecode.Res{Code: 10004, Desc: "账号密码错误"}
	ERROR_TOKEN                    = &ecode.Res{Code: 20001, Desc: "token错误"}
	NICKNAME_ALREADY_EXIST         = &ecode.Res{Code: 10005, Desc: "用户昵称已存在"}
	PHONE_ALREADY_EXIST            = &ecode.Res{Code: 10006, Desc: "用户手机号已存在"}
	DEPARTMENT_NOT_EXIST           = &ecode.Res{Code: 10007, Desc: "不存在该部门"}
	CABINET_NOT_EXIST              = &ecode.Res{Code: 10008, Desc: "不存在这个快递柜"}
	ERROR_RESOURCE_ALREADY_EXISTS  = &ecode.Res{Code: 10009, Desc: "已存在该运单号"}
)
