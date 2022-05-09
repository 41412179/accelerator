package table

import "time"

type Token struct {
	ID         int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	UserId     int64     `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`                 //  user id
	Token      string    `gorm:"column:token" db:"token" json:"token" form:"token"`                         //  token
	ExpireDate time.Time `gorm:"column:expire_date" db:"expire_date" json:"expire_date" form:"expire_date"` //  过期时间，开始时间+1年
	CreatedAt  time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`     //  记录创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`     //  记录更新时间
}

// TableName 会将 Token 的表名重写为 `token`
func (Token) TableName() string {
	return "token"
}
