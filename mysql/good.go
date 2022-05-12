package mysql

import (
	"accelerator/entity/db"
	"accelerator/entity/table"
)

// GetGoods 获取商品列表
func GetGoods() ([]*table.Good, error) {
	var goods []*table.Good
	err := db.DB.Find(&goods).Error
	return goods, err
}

func GetGoodByID(id int64) (*table.Good, error) {
	var good table.Good
	err := db.DB.Where("id = ?", id).First(&good).Error
	return &good, err
}
