package service

import (
	// "accelerator/model"

	"accelerator/mysql"
	"accelerator/serializer"
	"accelerator/util"

	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	Email     string `form:"user_name" json:"user_name" binding:"required"`
	ChannelId int64  `form:"channel_id" json:"channel_id" binding:"required"`
	Source    string `form:"source" json:"source" binding:"required"`
}

// setSession 设置session
// func (service *UserLoginService) setSession(c *gin.Context, user table.Login) {
// 	s := sessions.Default(c)
// 	s.Clear()
// 	s.Set("user_id", user.Id)
// 	s.Save()
// }

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {

	// 设置session
	// service.setSession(c, user)

	user, err := mysql.GetUserByEmail(service.Email)
	if err != nil {
		util.Log().Error("服务器错误: %s", err)
		return serializer.NewErr(serializer.CodeDBError, err)
	}

	return serializer.BuildUserResponse(*user)
}
