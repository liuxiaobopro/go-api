package service

import (
	"github.com/liuxiaobopro/go-api/app/common/errcode"
	"github.com/liuxiaobopro/go-api/app/common/middleware"
	dao2 "github.com/liuxiaobopro/go-api/app/dao"
	"github.com/liuxiaobopro/go-api/app/model"
	"github.com/liuxiaobopro/go-api/global"

	"github.com/liuxiaobopro/go-lib/ecode"
	"github.com/liuxiaobopro/go-lib/utils/encryption"
)

type UserServiceType struct{}

var UserSrv = new(UserServiceType)

// AddUser 添加用户
func (*UserServiceType) AddUser(req *model.AddUser) (int, ecode.BizErr) {
	// 如果两次密码一致
	if req.Password != req.RePassword {
		return 0, errcode.TWO_PASSWORDS_ARE_INCONSISTENT
	}
	// 判断用户是否存在
	user, _ := dao2.UserDao.GetUserByUsername(req.Username)
	if user.Id > 0 {
		return 0, ecode.ERROR_RESOURCE_ALREADY_EXISTS
	}
	// 密码加密
	pass, _ := encryption.BcryptEncrypt(req.Password + global.Conf.App.Sale)

	var userAdd = &model.User{
		Username: req.Username,
		Nickname: req.Nickname,
		Password: pass,
		Phone:    req.Phone,
	}
	// 添加用户
	insert, err := dao2.UserDao.AddUser(userAdd)
	return insert, err
}

// Login 用户登录
func (*UserServiceType) Login(req *model.LoginUser) (map[string]interface{}, ecode.BizErr) {
	//#region 判断用户是否存在
	user, _ := dao2.UserDao.GetUserByUsername(req.Username)
	if user.Id == 0 {
		return nil, ecode.ERROR_RESOURCE_DONT_EXISTS
	}
	//#endregion

	//#region 密码校验
	if ok := encryption.BcryptCheck(user.Password, req.Password+global.Conf.App.Sale); !ok {
		return nil, errcode.ACCOUNT_PASSWORD_ERROR
	}
	//#endregion

	//#region 生成token
	token, err := middleware.JwtMiddleware.GenerateToken(user.Id)
	if err != ecode.SUCCSESS {
		return nil, ecode.ERROR_SERVER
	}
	//#endregion

	var detailUser = &model.DetailUser{
		Id:         user.Id,
		Username:   user.Username,
		Nickname:   user.Nickname,
		Phone:      user.Phone,
		CreateTime: user.CreateTime,
	}
	var d = make(map[string]interface{})
	d["token"] = token
	d["user"] = detailUser
	return d, nil
}

// Logout 用户退出
func (*UserServiceType) Logout(uid int) ecode.BizErr {
	// 将token置空
	dao2.TokenDao.UpdateByUidField(uid, "token", "")
	dao2.TokenDao.UpdateByUidField(uid, "expiration_time", 0)
	return ecode.SUCCSESS
}

// UpdateUserById 根据id修改用户
func (*UserServiceType) UpdateUserById(id int, req *model.UpdateUser) (int, ecode.BizErr) {
	// 判断用户是否存在
	user, _ := dao2.UserDao.GetUserById(id)
	if user.Id == 0 {
		return 0, ecode.ERROR_RESOURCE_DONT_EXISTS
	}
	var userUpdate = new(model.UpdateUser)
	// 判断req三个字段是否为空
	if req.Password != "" {
		// 密码加密
		pass, _ := encryption.BcryptEncrypt(req.Password + global.Conf.App.Sale)
		userUpdate.Password = pass
	}
	if req.Nickname != "" {
		// 判断昵称唯一
		user, _ := dao2.UserDao.GetUserByNickname(req.Nickname)
		if user.Id > 0 {
			return 0, errcode.NICKNAME_ALREADY_EXIST
		}
		userUpdate.Nickname = req.Nickname
	}
	if req.Phone != "" {
		// 判断手机号唯一
		user, _ := dao2.UserDao.GetUserByPhone(req.Phone)
		if user.Id > 0 {
			return 0, errcode.PHONE_ALREADY_EXIST
		}
		userUpdate.Phone = req.Phone
	}
	// 判断userUpdate是否为空
	if userUpdate.Nickname == "" && userUpdate.Phone == "" && req.Password == "" {
		return 0, ecode.ERROR_PARAMETER_EXCEPTION
	}
	// 修改用户
	update, err := dao2.UserDao.UpdateUserById(id, userUpdate)
	return update, err
}

// DeleteUserById 根据id删除用户
func (*UserServiceType) DeleteUserById(id int) (int, ecode.BizErr) {
	// 判断用户是否存在
	user, _ := dao2.UserDao.GetUserById(id)
	if user.Id == 0 {
		return 0, ecode.ERROR_RESOURCE_DONT_EXISTS
	}
	// 删除用户
	delete, err := dao2.UserDao.DeleteUserById(id)
	return delete, err
}

// GetUserList 获取用户列表
func (*UserServiceType) GetUserList(req *model.SearchUser) ([]*model.DetailUser, ecode.BizErr) {
	// 获取用户列表
	users, err := dao2.UserDao.GetUserList(req)
	if err != ecode.SUCCSESS {
		return nil, err
	}
	var detailUsers []*model.DetailUser
	for _, user := range users {
		var detailUser = &model.DetailUser{
			Id:         user.Id,
			Username:   user.Username,
			Nickname:   user.Nickname,
			Phone:      user.Phone,
			CreateTime: user.CreateTime,
		}
		detailUsers = append(detailUsers, detailUser)
	}
	return detailUsers, err
}

// 获取用户列表总数
func (*UserServiceType) GetUserListTotal(req *model.SearchUser) (int, ecode.BizErr) {
	// 获取用户列表总数
	total, err := dao2.UserDao.GetUserListTotal(req)
	return total, err
}
