package serializer

import (
	"accelerator/entity/table"
)

// User 用户序列化器
type User struct {
	ID uint `json:"id"`
}

// BuildUser 序列化用户
func BuildUser(user table.User) User {
	return User{
		ID: uint(user.Id),
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user table.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
