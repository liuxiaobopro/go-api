package middleware

import (
	"net/http"

	"github.com/liuxiaobopro/go-api/app/admin/dao"
	"github.com/liuxiaobopro/go-api/app/admin/model"

	"github.com/gin-gonic/gin"
	"github.com/liuxiaobopro/go-lib/ecode"
	"github.com/liuxiaobopro/go-lib/response"
	stringl "github.com/liuxiaobopro/go-lib/utils/string"
	timel "github.com/liuxiaobopro/go-lib/utils/time"
)

type JwtMiddlewareType struct {
	// 过期时间
	Expire int
}

var JwtMiddleware = &JwtMiddlewareType{
	Expire: 7 * 86400, // 7天
}

// GenerateToken 生成token
func (*JwtMiddlewareType) GenerateToken(uid int) (string, ecode.BizErr) {
	userToken, err := dao.TokenDao.GetByUserId(uid)
	if err == ecode.ERROR_RESOURCE_DONT_EXISTS {
		// 不存在token, 生成记录
		token := stringl.RandString(24)
		expire := int(timel.GetNowTimeUnix()) + JwtMiddleware.Expire
		_, err := dao.TokenDao.Add(&model.Token{
			Uid:            uid,
			Token:          token,
			ExpirationTime: expire,
		})
		if err != ecode.SUCCSESS {
			return "", err
		}
		return token, ecode.SUCCSESS
	} else {
		// 存在token, 校验token是否过期
		t := timel.GetNowTimeUnix()
		if int64(userToken.ExpirationTime) <= t {
			// 已过期, 生成新token
			token := stringl.RandString(24)
			// 更新token
			err := dao.TokenDao.UpdateById(userToken.Id, &model.Token{
				Token:          token,
				ExpirationTime: int(t) + JwtMiddleware.Expire,
			})
			if err.Code != ecode.SUCCSESS.Code {
				return "", err
			}
			return token, ecode.SUCCSESS
		} else {
			// 未过期, 返回token
			return userToken.Token, ecode.SUCCSESS
		}
	}
}

// CheckToken 校验token
func (*JwtMiddlewareType) CheckToken(token string) bool {
	// 根据token获取用户信息
	userToken, err := dao.TokenDao.GetByToken(token)
	if err.Code != ecode.SUCCSESS.Code {
		return false
	}
	// 校验token是否过期
	return int64(userToken.ExpirationTime) > timel.GetNowTimeUnix()
}

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取token
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, response.GetErrRes(ecode.ERROR_TOKEN_INVALID))
			c.Abort()
			return
		}
		if !JwtMiddleware.CheckToken(token) {
			c.JSON(http.StatusOK, response.GetErrRes(ecode.ERROR_TOKEN_INVALID))
			c.Abort()
			return
		}
		c.Next()
	}
}
