package service

import (
	"fmt"

	model "github.com/druidamix/go_server/model"
	"github.com/druidamix/go_server/repository"
)

// Getuser returns user model
func GetUser(user string, password string, rep *repository.UserRepository) (model.User, error) {
	dbUser, err := rep.GetUser(user, password)
	if err != nil {
		return model.User{}, fmt.Errorf("error")
	}
	return dbUser, nil
}

// UpdateUserPass updates user pass after first login
func UpdateUserPass(user string, newPass string, rep *repository.UserRepository) error {
	err := rep.UpdateUserPass(user, newPass)
	if err != nil {
		return fmt.Errorf("error updating")
	}
	return nil
}
