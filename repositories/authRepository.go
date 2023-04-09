package repositories

import (
	"fmt"
	"log"

	"github.com/druidamix/go_server/helpers"
	model "github.com/druidamix/go_server/models"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// updateJwtSecret stores the new secret key on db, then is returned
func (a *AuthRepository) UpdateJwtSecret(user string, redundant string) (string, error) {
	token, _ := helpers.RandToken(250)

	userDb := model.User{User: user}
	dbRes := a.db.Model(&userDb).Where("user=? and redundant_token=?", user, redundant).Updates(model.User{Token: token})

	if dbRes.RowsAffected < 1 {
		log.Println("-- 0 rows affected")
		return "", fmt.Errorf("")
	}
	return token, nil
}

// Stores the reduntat token on db, then returns it
func (a *AuthRepository) SaveRedundantToken(user string) (string, error) {
	redundant_token, _ := helpers.RandToken(2048)

	dbRes := a.db.Model(&model.User{}).Where("user=?", user).Update("redundant_token", redundant_token)

	if dbRes.RowsAffected < 1 {
		return "", fmt.Errorf("Not found")
	}

	return redundant_token, nil
}

// Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GetJwtSecretKey obtains token by user
func (a *AuthRepository) GetJwtSecretKey(user string) (string, error) {

	var _user model.User

	res := a.db.Where("user = ?", user).Find(&_user)
	if res.RowsAffected < 1 {
		return "", fmt.Errorf("not found")
	}

	return _user.Token, nil
}
