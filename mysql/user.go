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
func InsertUser(user *table.User) (int64, error) {
	result := db.DB.Create(user)
	return result.RowsAffected, result.Error
}

// GetUserByID get user by id
func GetUserByID(id int64) (*table.User, error) {
	var user table.User
	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserLikeEmail(email string) ([]*table.User, error) {
	var users []*table.User
	if err := db.DB.Where("email =", email).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
