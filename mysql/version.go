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

func EditVersion(version string, url, desc string, codeID, id int64) (*table.Version, error) {
	var versionObj table.Version
	if err := db.DB.Model(&versionObj).Where("id = ?", id).Update("version", version).Update("url", url).Update("desc", desc).Update("code_id", codeID).Error; err != nil {
		return nil, err
	}
	versionObj.Id = int(id)
	return &versionObj, nil
}
