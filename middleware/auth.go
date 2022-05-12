package middleware

import (
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/mysql"
	"accelerator/util"

	"accelerator/entity/errcode"

	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := c.GetQuery("token")
		var uid int64
		if ok {
			t, err := mysql.GetToken(token)
			if err != nil {
				c.JSON(200, errcode.NewErr(errcode.CodeDBError, err))
				c.Abort()
				return
			}
			uid = t.UserId
		}

		user, err := mysql.GetUserByID(uid)
		if err == nil {
			c.Set("user", user)
		}

		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, ok := c.Get("user"); ok {
			if _, ok := user.(*table.User); ok {
				util.Log().Debug("user has logined")
				// c.Next()
				return
			}
		} else {
			util.Log().Debug("user not logined, user: %v", user)
			c.JSON(200, CheckLogin())
			c.Abort()
		}
	}
}

// CheckLogin 检查登录
func CheckLogin() response.Response {
	return response.Response{
		Code: errcode.CodeCheckLogin,
		Msg:  errcode.Text(errcode.CodeCheckLogin),
	}
}
