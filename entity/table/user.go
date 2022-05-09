package table

import "time"

// User 用户模型
type User struct {
	ID        int64     `gorm:"column:id" db:"id" json:"id" form:"id"`                                 //  每个用户唯一的userid
	Email     string    `gorm:"column:email" db:"email" json:"email" form:"email"`                     //  邮箱信息，非空且唯一
	UserId    string    `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`             //  用户id-备用的唯一id
	ChannelId int64     `gorm:"column:channel_id" db:"channel_id" json:"channel_id" form:"channel_id"` //  渠道id，默认为0
	Source    string    `gorm:"column:source" db:"source" json:"source" form:"source"`                 //  来源：android
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
}

// TableName 会将 User 的表名重写为 `user`
func (User) TableName() string {
	return "user"
}
