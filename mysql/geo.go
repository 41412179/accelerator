package mysql

import (
	"accelerator/entity/db"
	"accelerator/entity/table"
)

// GetGeos 获取地理位置列表
func GetGeos() ([]*table.Geo, error) {
	var geos []*table.Geo
	err := db.DB.Find(&geos).Error
	return geos, err
}
