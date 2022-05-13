package table

import (
	"time"
)

// Node 节点
type Node struct {
	Id        int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(45);NOT NULL" json:"name"`  // 节点名称
	Link      string    `gorm:"column:link;type:varchar(512);NOT NULL" json:"link"` // 节点链接
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
	Tag       int64     `gorm:"column:tag;type:bigint(20);NOT NULL" json:"tag"` // 标记：是不是推荐
}

// TableName 会将 Node 的表名重写为 `node`
func (m *Node) TableName() string {
	return "node"
}
