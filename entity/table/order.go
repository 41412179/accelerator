package table

import "time"

// Order 订单
type Order struct {
	ID        int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	UserId    int64     `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`             //  用户id
	Type      int64     `gorm:"column:type" db:"type" json:"type" form:"type"`                         //  订单类型：0.5美刀/1天，,2美刀/7天，6美刀/30天，12美刀/90天，20美刀/180天，30美刀/360天\r
	Status    int64     `gorm:"column:status" db:"status" json:"status" form:"status"`                 //  订单状态：0：未支付，1:已支付
	StartTime time.Time `gorm:"column:start_time" db:"start_time" json:"start_time" form:"start_time"` //  订单开始时间
	EndTime   time.Time `gorm:"column:end_time" db:"end_time" json:"end_time" form:"end_time"`         //  订单结束时间
	CreatedAt time.Time `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
}

// TableName 会将 Order 的表名重写为 `order`
func (Order) TableName() string {
	return "order"
}
