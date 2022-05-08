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

func GetTokenByUserID(userID int64) (*table.Token, error) {
	var token table.Token
	if err := db.DB.Where("user_id = ?", userID).First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}
