package response

// User 用户序列化器
type UserServiceRsp struct {
	ID            int64  `json:"id"`
	Email         string `json:"email"`
	Token         string `json:"token"`
	RemainingTime int64  `json:"remaining_time"`
}
