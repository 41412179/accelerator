package api

import (
	"accelerator/conf"
	"accelerator/entity/errcode"
	"accelerator/entity/response"
	"accelerator/entity/table"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v8"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, response.Response{
		Code: 0,
		Msg:  "Pong",
	})
}

func GetRoomToken(c *gin.Context) {
	// userId := c.get
	// qbox.Mac
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *table.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*table.User); ok {
			return u
		}
	}
	return nil
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) response.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return errcode.ParamErr(
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return errcode.ParamErr("JSON类型不匹配", err)
	}

	return errcode.ParamErr("参数错误", err)
}
