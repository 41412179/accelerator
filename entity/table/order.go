package table

import "time"

// Order 订单
type Order struct {
	Id        int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	UserId    int64     `gorm:"column:user_id;type:bigint(20);NOT NULL" json:"user_id"` // 用户id
	GoodId    int       `gorm:"column:good_id;type:int(11);NOT NULL" json:"good_id"`    // 订单类型：0.5美刀/1天，,2美刀/7天，6美刀/30天，12美刀/90天，20美刀/180天，30美刀/360天\r
	Status    int       `gorm:"column:status;type:int(11);NOT NULL" json:"status"`      // 订单状态：0：未支付，1:已支付
	StartTime time.Time `gorm:"column:start_time;type:datetime" json:"start_time"`      // 订单开始时间
	EndTime   time.Time `gorm:"column:end_time;type:datetime" json:"end_time"`          // 订单结束时间
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
	PaymentId int64     `gorm:"column:payment_id;type:bigint(20);NOT NULL" json:"payment_id"`
	PayType   string    `gorm:"column:pay_type;type:varchar(20);NOT NULL" json:"pay_type"` // 支付方式
}

// TableName 会将 Order 的表名重写为 `order`
func (Order) TableName() string {
	return "order"
}
