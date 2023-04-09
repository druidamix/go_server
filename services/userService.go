package services

import (
	"fmt"

	model "github.com/druidamix/go_server/models"
	"github.com/druidamix/go_server/repositories"
)

// Getuser returns user model
func GetUser(user string, password string, rep *repositories.UserRepository) (model.User, error) {
	mdl, err := rep.GetUser(user, password)
	if err != nil {
		return model.User{}, fmt.Errorf("error")
	}
	return mdl, nil
}

// UpdateUserPass updates user pass after first login
func UpdateUserPass(user string, newPass string, rep *repositories.UserRepository) error {
	err := rep.UpdateUserPass(user, newPass)

	if err != nil {
		return fmt.Errorf("error updating")
	}
	return nil
}
