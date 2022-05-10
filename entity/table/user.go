package table

import "time"

// User 用户
type User struct {
	Id        int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`         // 每个用户唯一的userid
	Email     string    `gorm:"column:email;type:varchar(45);NOT NULL" json:"email"`                    // 邮箱信息，非空且唯一
	UserId    string    `gorm:"column:user_id;type:varchar(45);NOT NULL" json:"user_id"`                // 用户id-备用的唯一id
	ChannelId int64     `gorm:"column:channel_id;type:bigint(20);NOT NULL" json:"channel_id"`           // 渠道id，默认为0
	InviterId int64     `gorm:"column:inviter_id;type:bigint(20);default:0;NOT NULL" json:"inviter_id"` // 邀请人id
	Source    string    `gorm:"column:source;type:varchar(45);NOT NULL" json:"source"`                  // 来源：Android
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
}

// TableName 会将 User 的表名重写为 `user`
func (User) TableName() string {
	return "user"
}
