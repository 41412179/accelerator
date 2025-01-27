package table

import "time"

type Good struct {
	Id          int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	Duration    int64     `gorm:"column:duration;type:bigint(20);NOT NULL" json:"duration"` // 用户时长
	Price       int64     `gorm:"column:price;type:bigint(20);NOT NULL" json:"price"`       // 实际单价，单位：分
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
	OriginPrice int64     `gorm:"column:origin_price;type:bigint(20);NOT NULL" json:"origin_price"`
	Radio       string    `gorm:"column:radio;type:varchar(45);NOT NULL" json:"radio"` // 折扣
}

func (m *Good) TableName() string {
	return "good"
}
