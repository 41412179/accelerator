package service

import (
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"
	"accelerator/mysql"
	"accelerator/util"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserService 管理用户登录的服务
type UserService struct {
	Email        string `form:"email" json:"email" binding:"required,email"`
	ChannelId    int64  `form:"channel_id" json:"channel_id"`
	Source       string `form:"source" json:"source" binding:"required"`
	InviterId    int64  `form:"inviter_id" json:"inviter_id"`
	orderService *OrderService
	token        string
}

func NewUserService() *UserService {
	return &UserService{
		orderService: NewOrderService(),
	}
}

// Login 用户登录函数
func (u *UserService) Login(c *gin.Context) response.Response {

	// 设置session
	// service.setSession(c, user)

	localUser, err := mysql.GetUserByEmail(u.Email)

	// 判断用户是否存在
	if err == gorm.ErrRecordNotFound {
		user := u.createNewUser()
		_, err := mysql.InsertUser(user)
		if err != nil {
			util.Log().Error("insert user err: %v", err)
		}
		user, err = mysql.GetUserByEmail(u.Email)
		if err != nil {
			util.Log().Error("get user err: %v", err)
		}
		if err := u.createToken(user.ID); err != nil {
			util.Log().Error("create token err: %v", err)
		}
		localUser = user
	} else if err != nil {
		util.Log().Error("get user by email err: %v", err)
		return errcode.NewErr(errcode.CodeDBError, err)
	}

	// 如果存在，则查询剩余时间
	remainingTime, err := u.orderService.GetRemainingTimeByUserId(localUser.ID)
	if err != nil {
		util.Log().Error("get remaining time by user id err: %v", err)
		return errcode.NewErr(errcode.CodeDBError, err)
	}

	// 查询token
	u.getTokenByUserID(localUser.ID)
	return u.setRsponse(localUser, remainingTime)
}

// getTokenByUserID 根据用户id获取token
func (u *UserService) getTokenByUserID(id int64) error {
	token, err := mysql.GetTokenByUserID(id)

	// 如果不存在就创建
	if err == gorm.ErrRecordNotFound {
		if err := u.createToken(id); err != nil {
			util.Log().Error("create token err: %v", err)
			return err
		}
	}
	// 如果异常
	if err != nil {
		util.Log().Error("get token by user id err: %v", err)
		return err
	}
	// 如果过期就续期
	if token.ExpireDate.Before(time.Now()) {
		token.ExpireDate = time.Now().AddDate(1, 0, 0)
		if err := mysql.UpdateToken(token); err != nil {
			util.Log().Error("update token err: %v", err)
			return err
		}
	}
	u.token = token.Token
	return nil

}

// setRsponse 设置返回值
func (u *UserService) setRsponse(user *table.User, remainingTime int64) response.Response {
	return response.Response{
		Code: errcode.CodeSuccess,
		Msg:  errcode.Text(errcode.CodeSuccess),
		Data: response.UserServiceRsp{
			ID:            user.ID,
			Email:         user.Email,
			Token:         u.token,
			RemainingTime: remainingTime,
			ExpireDate:    time.Now().Unix() + remainingTime,
		},
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

// createToken 创建token
func (u *UserService) createToken(id int64) error {
	token := new(table.Token)
	token.UserId = id
	token.Token = util.TokenByMD5(u.Email, u.Email, 5)
	u.token = token.Token
	token.ExpireDate = time.Now().AddDate(1, 0, 0)
	return mysql.InsertToken(token)
}
