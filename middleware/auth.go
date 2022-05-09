package middleware

import (
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/mysql"

	"accelerator/entity/errcode"

	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, ok := c.Get("token")
		var uid int64
		if ok {
			t, err := mysql.GetToken(token.(string))
			if err != nil {
				c.JSON(200, errcode.Err(errcode.CodeDBError, errcode.Text(errcode.CodeDBError), err))
				c.Abort()
				return
			}
			uid = t.UserId
		}

		user, err := mysql.GetUserByID(uid)
		if err == nil {
			c.Set("user", &user)
		}

		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*table.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, CheckLogin())
		c.Abort()
	}
}

// CheckLogin 检查登录
func CheckLogin() response.Response {
	return response.Response{
		Code: errcode.CodeCheckLogin,
		Msg:  errcode.Text(errcode.CodeCheckLogin),
	}
}
