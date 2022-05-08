package response

// User 用户序列化器
type UserServiceRsp struct {
	ID uint `json:"id"`
}

// BuildUser 序列化用户
// func BuildUser(user table.User) User {
// 	return User{
// 		ID: uint(user.ID),
// 	}
// }
