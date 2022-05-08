package table

// User 用户模型
type User struct {
	Id        int64  `gorm:"column:id;" json:"id"`
	Email     string `gorm:"column:email;" json:"email"`
	UserId    string `gorm:"column:user_id;" json:"userId"`
	ChannelId int64  `gorm:"column:channel_id;" json:"channelId"`
	Source    string `gorm:"column:source;" json:"source"`
	CreatedAt string `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt string `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updatedAt"`
	DeletedAt string `gorm:"column:deleted_at;" json:"deletedAt"`
}
