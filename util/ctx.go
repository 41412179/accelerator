package util

import (
	"accelerator/entity/table"
	// "accelerator/util"

	"github.com/gin-gonic/gin"
)

// GetUserByCtx 获取用户
func GetUserByCtx(c *gin.Context) *table.User {
	if user, ok := c.Get("user"); ok {
		Log().Info("get user by ctx success, user: %+v", user)
		if u, ok := user.(*table.User); ok {
			return u
		}
	}
	return nil
}
