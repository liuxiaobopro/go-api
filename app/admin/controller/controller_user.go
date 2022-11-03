package controller

import (
	"strconv"

	"github.com/liuxiaobopro/go-api/app/admin/model"
	"github.com/liuxiaobopro/go-api/app/admin/service"

	"github.com/gin-gonic/gin"
	"github.com/liuxiaobopro/go-lib/ecode"
	ginl "github.com/liuxiaobopro/go-lib/gin"
	jsonl "github.com/liuxiaobopro/go-lib/json"
	timel "github.com/liuxiaobopro/go-lib/utils/time"
)

type UserControllerType struct {
	ginl.Handler
}

var UserController = new(UserControllerType)

// Add 添加用户
func (th *UserControllerType) Add(c *gin.Context) {
	var AddUser = new(model.AddUser)
	if err := c.ShouldBindJSON(AddUser); err != nil {
		th.SendErr(c, nil, ecode.ERROR_PARAMETER_EXCEPTION)
		return
	}
	res, err := service.UserSrv.AddUser(AddUser)
	if err != nil {
		th.SendErr(c, nil, err)
		return
	}
	th.SendSucc(c, res, nil)
}

// Login 用户登录
func (th *UserControllerType) Login(c *gin.Context) {
	var LoginUser = new(model.LoginUser)
	if err := c.ShouldBindJSON(LoginUser); err != nil {
		th.SendErr(c, nil, ecode.ERROR_PARAMETER_EXCEPTION)
		return
	}
	res, err := service.UserSrv.Login(LoginUser)
	if err != nil {
		th.SendErr(c, nil, err)
		return
	}
	th.SendSucc(c, res, nil)
}

// Logout 用户登出
func (th *UserControllerType) Logout(c *gin.Context) {
	type reqType struct {
		Uid int `json:"uid" binding:"required"`
	}
	var req = new(reqType)
	if err := c.ShouldBindJSON(req); err != nil {
		th.SendErr(c, nil, ecode.ERROR_PARAMETER_EXCEPTION)
		return
	}
	_ = service.UserSrv.Logout(req.Uid)
	th.SendSucc(c, nil, nil)
}

// Update 修改用户
func (th *UserControllerType) Update(c *gin.Context) {
	id, _ := strconv.Atoi(th.GetParam(c, "id"))
	var UpdateUser = new(model.UpdateUser)
	if err := c.ShouldBindJSON(UpdateUser); err != nil {
		th.SendErr(c, nil, ecode.ERROR_PARAMETER_EXCEPTION)
		return
	}
	res, err := service.UserSrv.UpdateUserById(id, UpdateUser)
	th.Send(c, res, err)
}

// Delete 删除用户
func (th *UserControllerType) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(th.GetParam(c, "id"))
	res, err := service.UserSrv.DeleteUserById(id)
	th.Send(c, res, err)
}

// List 用户列表
func (th *UserControllerType) List(c *gin.Context) {
	var SearchUser = new(model.SearchUser)
	SearchUser.Page, _ = strconv.Atoi(th.GetQueryDefault(c, "page", "1"))
	SearchUser.Limit, _ = strconv.Atoi(th.GetQueryDefault(c, "limit", "10"))
	if cts := th.GetQueryDefault(c, "create_time_start", ""); cts != "" {
		SearchUser.CreateTimeStart = jsonl.Time(timel.StringToTime(cts))
	}
	if cte := th.GetQueryDefault(c, "create_time_end", ""); cte != "" {
		SearchUser.CreateTimeEnd = jsonl.Time(timel.StringToTime(cte))
	}
	res, err := service.UserSrv.GetUserList(SearchUser)
	if err != ecode.SUCCSESS {
		th.SendErr(c, nil, err)
		return
	}
	total, err := service.UserSrv.GetUserListTotal(SearchUser)
	if err != ecode.SUCCSESS {
		th.SendErr(c, nil, err)
		return
	}
	d := map[string]interface{}{
		"list":  res,
		"total": total,
	}
	th.Send(c, d, ecode.SUCCSESS)
}
