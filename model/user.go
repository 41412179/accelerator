package model

import (
	"gorm.io/gorm"
)

// Login 用户模型
type Login struct {
	gorm.Model
	Email     string
	UserId    string
	ChannelId string
	Source    string
}

// GetUser 用ID获取用户
func GetUser(ID interface{}) (Login, error) {
	var user Login
	result := DB.First(&user, ID)
	return user, result.Error
}
