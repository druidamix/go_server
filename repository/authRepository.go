// Package repository acces database data
package repository

import (
	"fmt"

	model "github.com/druidamix/go_server/model"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// UpdateJwtSecret stores the new secret key on db, then is returned
func (a *AuthRepository) UpdateJwtSecret(user string, redundant string, token string) (string, error) {

	userDb := model.User{User: user}
	dbRes := a.db.Model(&userDb).Where("user=? and redundant_token=?", user, redundant).Updates(model.User{Token: token})

	if dbRes.RowsAffected < 1 {
		return "", fmt.Errorf("Not found")
	}
	return token, nil
}

// Stores the reduntat token on db, then returns it
func (a *AuthRepository) SaveRedundantToken(user string, redundant string) error {

	dbRes := a.db.Model(&model.User{}).Where("user=?", user).Update("redundant_token", redundant)

	if dbRes.RowsAffected < 1 {
		return fmt.Errorf("Not found")
	}
	return nil
}

// GetJwtSecret obtains token by user
func (a *AuthRepository) GetJwtSecret(user string) (string, error) {

	var dbUser model.User

	res := a.db.Where("user = ?", user).Find(&dbUser)
	if res.RowsAffected < 1 {
		return "", fmt.Errorf("not found")
	}

	return dbUser.Token, nil
}
