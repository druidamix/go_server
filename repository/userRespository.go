// Package repository is in charege of managing database data
package repository

import (
	"fmt"

	model "github.com/druidamix/go_server/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRespository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetUser return an model.User
func (u *UserRepository) GetUser(user string, password string) (model.User, error) {

	dbUser := model.User{}

	res := u.db.Where("user = ? AND password  = ?", user, password).First(&dbUser)

	if res.Error != nil {
		return model.User{}, fmt.Errorf("Error")
	}

	return dbUser, nil
}

// UpdateUserPass updates the user password
func (u *UserRepository) UpdateUserPass(user string, newPassword string) error {

	rowsAffected := u.db.Where("user=?", user).
		Updates(model.User{Password: newPassword, First_login: 1}).RowsAffected

	if rowsAffected < 1 {
		return fmt.Errorf("Not found")
	}

	return nil
}
