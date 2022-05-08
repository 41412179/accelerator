package mysql

// "accelerator/table"
import (
	"accelerator/entity/db"
	"accelerator/entity/table"
)

// GetUserByEmail get login info by email
func GetUserByEmail(email string) (*table.User, error) {
	var user table.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// InsertUser insert user
func InsertUser(user *table.User) error {
	if err := db.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
