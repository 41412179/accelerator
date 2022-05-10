package mysql

import (
	"accelerator/entity/db"
	"accelerator/entity/table"
)

// GetNodes 获取节点列表
func GetNodes() ([]*table.Node, error) {
	var nodes []*table.Node
	err := db.DB.Find(&nodes).Error
	return nodes, err
}
