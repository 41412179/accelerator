package mysql

import (
	"accelerator/entity/db"
	"accelerator/entity/table"
)

func GetVersion() (*table.Version, error) {
	var version table.Version
	if err := db.DB.First(&version).Error; err != nil {
		return nil, err
	}
	return &version, nil
}

func EditVersion(version string, url, desc string, id int64) (*table.Version, error) {
	var versionObj table.Version
	if err := db.DB.Model(&versionObj).Where("id = ?", id).Update("version", version).Update("url", url).Update("desc", desc).Error; err != nil {
		return nil, err
	}
	return &versionObj, nil
}
