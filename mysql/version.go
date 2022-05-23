package mysql

import (
	"accelerator/entity/db"
	"accelerator/entity/table"
)

func GetVersion() (*table.Version, error) {
	var version *table.Version
	if err := db.DB.First(version).Error; err != nil {
		return nil, err
	}
	return version, nil
}
