package util

import (
	"accelerator/entity/table"

	"github.com/gin-gonic/gin"
)

// GetUserByCtx 获取用户
func GetUserByCtx(c *gin.Context) *table.User {
	if user, ok := c.Get("user"); ok {
		if u, ok := user.(*table.User); ok {
			return u
		}
	}
	return nil
}
