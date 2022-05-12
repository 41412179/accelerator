package table

import (
	"time"
)

const (
	// 增加佣金
	AddCommissionType = 0
	// 减少佣金
	SubCommissionType = 1
)

type Commission struct {
	Id        int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	OrderId   int64     `gorm:"column:order_id;type:bigint(20);NOT NULL" json:"order_id"`
	UserId    int64     `gorm:"column:user_id;type:bigint(20);NOT NULL" json:"user_id"`     // 佣金所有人id
	Change    float64   `gorm:"column:change;type:double;default:0;NOT NULL" json:"change"` // 费用变化
	Type      int       `gorm:"column:type;type:int(11);default:0;NOT NULL" json:"type"`    // 类型：0：分享人下单，1:体现
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
}

func (m *Commission) TableName() string {
	return "commission"
}
