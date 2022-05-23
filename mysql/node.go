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

// DeleteNode 删除节点
func DeleteNode(id int64) error {
	return db.DB.Where("id = ?", id).Delete(&table.Node{}).Error
}

// AddNode 添加节点
func AddNode(node *table.Node) error {
	return db.DB.Create(node).Error
}

// UpdateNode 更新节点
func EditNode(node *table.Node) error {
	return db.DB.Model(&table.Node{}).Where("id = ?", node.Id).Updates(node).Error
}
