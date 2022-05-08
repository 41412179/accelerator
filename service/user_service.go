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

// UserService 管理用户登录的服务
type UserService struct {
	Email        string `form:"user_name" json:"user_name" binding:"required"`
	ChannelId    int64  `form:"channel_id" json:"channel_id" binding:"required"`
	Source       string `form:"source" json:"source" binding:"required"`
	orderService *OrderService
}

func NewUserService() *UserService {
	return &UserService{
		orderService: NewOrderService(),
	}
}

// setSession 设置session
// func (service *UserLoginService) setSession(c *gin.Context, user table.Login) {
// 	s := sessions.Default(c)
// 	s.Clear()
// 	s.Set("user_id", user.Id)
// 	s.Save()
// }

// Login 用户登录函数
func (u *UserService) Login(c *gin.Context) response.Response {

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
	remainingTime, err := u.orderService.GetRemainingTimeByUserId(user.ID)
	if err != nil {
		util.Log().Error("get remaining time by user id err: %v", err)
		return errcode.NewErr(errcode.CodeDBError, err)
	}
	// return response.BuildUserResponse(*user)
	return u.setRsponse(user, remainingTime)
}

func (u *UserService) setRsponse(user *table.User, remainingTime int64) response.Response {
	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
		Data: nil,
	}
}

// createNewUser 创建新用户
func (u *UserService) createNewUser() *table.User {
	user := new(table.User)
	user.Email = u.Email
	user.ChannelId = u.ChannelId
	user.Source = u.Source
	return user
}

func (u *UserService) createToken(id int64) error {
	token := new(table.Token)
	token.UserId = id
	token.Token = util.RandStringRunes(int(id))
	token.ExpireDate = time.Now().AddDate(1, 0, 0)
	return mysql.InsertToken(token)

}