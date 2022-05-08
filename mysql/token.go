package mysql

import (
	"accelerator/entity/db"
	"accelerator/entity/table"
)

// InsertToken insert token
func InsertToken(token *table.Token) error {
	if err := db.DB.Create(token).Error; err != nil {
		return err
	}
	return nil
}
