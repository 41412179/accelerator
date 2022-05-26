package mysql

import (
	"accelerator/entity/db"
	"accelerator/entity/table"
)

func GetShare() (*table.Share, error) {
	var share table.Share
	if err := db.DB.First(&share).Error; err != nil {
		return nil, err
	}
	return &share, nil
}
