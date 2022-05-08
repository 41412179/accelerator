package service

import (
	// "accelerator/model"

	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/mysql"
	"accelerator/util"
	"time"

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
func (u *UserLoginService) Login(c *gin.Context) response.Response {

	// 设置session
	// service.setSession(c, user)

	user, err := mysql.GetUserByEmail(u.Email)
	if err != nil {
		util.Log().Error("get user by email err: %s", err)
		return errcode.NewErr(errcode.CodeDBError, err)
	}
	// 判断用户是否存在
	if user.ID == 0 {
		user := u.createNewUser()
		id, err := mysql.InsertUser(user)
		if err != nil {
			util.Log().Error("insert user err: %v", err)
			return errcode.NewErr(errcode.CodeDBError, err)
		}
		if err := u.createToken(id); err != nil {
			util.Log().Error("create token err: %v", err)
			return errcode.NewErr(errcode.CodeDBError, err)
		}
	}
	// 如果存在，则查询需要的其他信息

	return response.BuildUserResponse(*user)
}

// createNewUser 创建新用户
func (u *UserLoginService) createNewUser() *table.User {
	user := new(table.User)
	user.Email = u.Email
	user.ChannelId = u.ChannelId
	user.Source = u.Source
	return user
}

func (u *UserLoginService) createToken(id int64) error {
	token := new(table.Token)
	token.UserId = id
	token.Token = util.RandStringRunes(int(id))
	token.ExpireDate = time.Now().AddDate(1, 0, 0)
	return mysql.InsertToken(token)

}
