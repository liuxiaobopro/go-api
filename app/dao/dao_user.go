package dao

import (
	"github.com/liuxiaobopro/go-api/app/model"
	"github.com/liuxiaobopro/go-api/global"

	"github.com/liuxiaobopro/go-lib/console"
	"github.com/liuxiaobopro/go-lib/ecode"
	jsonl "github.com/liuxiaobopro/go-lib/json"
	timel "github.com/liuxiaobopro/go-lib/utils/time"
	"xorm.io/xorm"
)

type UserDaoType struct{}

var UserDao = new(UserDaoType)

func sessUser() *xorm.Session {
	return global.Db.Table(global.Conf.Database.Mdb.Prefix + "user")
}

// AddUser 添加用户
func (th *UserDaoType) AddUser(req *model.User) (int, ecode.BizErr) {
	insert, err := sessUser().Insert(req)
	if err != nil {
		console.Console.Error("添加用户失败", "")
		return 0, ecode.ERROR_SERVER
	}
	return int(insert), ecode.SUCCSESS
}

// GetUserByUsername 根据用户名获取用户
func (th *UserDaoType) GetUserByUsername(userName string) (model.User, ecode.BizErr) {
	var user model.User
	_, err := sessUser().Where("username = ?", userName).Get(&user)
	if err != nil {
		console.Console.Error("根据用户名获取用户失败", "")
		return user, ecode.ERROR_SERVER
	}
	return user, nil
}

// GetUserById 根据id获取用户
func (th *UserDaoType) GetUserById(id int) (model.User, ecode.BizErr) {
	var user model.User
	_, err := sessUser().Where("id = ?", id).Get(&user)
	if err != nil {
		console.Console.Error("根据id获取用户失败", "")
		return user, ecode.ERROR_SERVER
	}
	return user, nil
}

// UpdateUserById 根据id修改用户
func (th *UserDaoType) UpdateUserById(id int, req *model.UpdateUser) (int, ecode.BizErr) {
	update, err := sessUser().Where("id = ?", id).Update(req)
	if err != nil {
		console.Console.Error("根据id修改用户失败", "")
		return 0, ecode.ERROR_SERVER
	}
	return int(update), ecode.SUCCSESS
}

// GetUserByNickname 根据nickname获取用户
func (th *UserDaoType) GetUserByNickname(nickname string) (model.User, ecode.BizErr) {
	var user model.User
	_, err := sessUser().Where("nickname = ?", nickname).Get(&user)
	if err != nil {
		console.Console.Error("根据nickname获取用户失败", "")
		return user, ecode.ERROR_SERVER
	}
	return user, nil
}

// GetUserByPhone 根据phone获取用户
func (th *UserDaoType) GetUserByPhone(phone string) (model.User, ecode.BizErr) {
	var user model.User
	_, err := sessUser().Where("phone = ?", phone).Get(&user)
	if err != nil {
		console.Console.Error("根据phone获取用户失败", "")
		return user, ecode.ERROR_SERVER
	}
	return user, nil
}

// DeleteUserById 根据id删除用户
func (th *UserDaoType) DeleteUserById(id int) (int, ecode.BizErr) {
	delete, err := sessUser().Where("id = ?", id).Delete(&model.User{})
	if err != nil {
		console.Console.Error("根据id删除用户失败", "")
		return 0, ecode.ERROR_SERVER
	}
	return int(delete), ecode.SUCCSESS
}

// GetUserList 获取用户列表
func (th *UserDaoType) GetUserList(req *model.SearchUser) ([]model.User, ecode.BizErr) {
	var (
		userList []model.User
		err      error
	)
	sess := sessUser()
	if req.CreateTimeStart != jsonl.NilTime && req.CreateTimeEnd != jsonl.NilTime {
		sess = sess.Where("create_time BETWEEN ? AND ?", timel.JsonlTimeToString(req.CreateTimeStart), timel.JsonlTimeToString(req.CreateTimeEnd))
	}
	sess.OrderBy("create_time desc")
	err = sess.Limit(req.Limit, (req.Page-1)*req.Limit).Find(&userList)
	if err != nil {
		console.Console.Error("获取用户列表失败", "")
		return userList, ecode.ERROR_SERVER
	}
	return userList, ecode.SUCCSESS
}

// GetUserListTotal 获取用户列表总数
func (th *UserDaoType) GetUserListTotal(req *model.SearchUser) (int, ecode.BizErr) {
	var (
		total int64
		err   error
	)
	sess := sessUser()
	if req.CreateTimeStart != jsonl.NilTime && req.CreateTimeEnd != jsonl.NilTime {
		sess = sess.Where("create_time BETWEEN ? AND ?", timel.JsonlTimeToString(req.CreateTimeStart), timel.JsonlTimeToString(req.CreateTimeEnd))
	}
	total, err = sess.Count(new(model.User))
	if err != nil {
		console.Console.Error("获取用户列表总数失败", "")
		return 0, ecode.ERROR_SERVER
	}
	return int(total), ecode.SUCCSESS
}
