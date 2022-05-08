package mysql

// "accelerator/table"
import (
	"accelerator/entity/db"
	"accelerator/entity/table"
)

// GetLoginByEmail get login info by email
func GetLoginByEmail(email string) (*table.User, error) {
	var user table.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
