package serializer

import "accelerator/model"

// User 用户序列化器
type User struct {
	ID uint `json:"id"`
}

// BuildUser 序列化用户
func BuildUser(user model.Login) User {
	return User{
		ID: user.ID,
	}
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.Login) Response {
	return Response{
		Data: BuildUser(user),
	}
}
