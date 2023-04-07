package controllers

import (
	"fmt"

	"github.com/druidamix/go_server/database"
	model "github.com/druidamix/go_server/models"
)

// Returns user using pass or redundant token
func getUserFromdb(user string, argument string, isPass bool) (model.User, error) {

	db := database.DB.Db
	dbUser := model.User{}

	var whereArgument string
	if isPass {
		whereArgument = "password"

	} else {
		whereArgument = "redundant_token"
	}

	res := db.Where("user = ? AND "+whereArgument+" = ?", user, argument).First(&dbUser)

	if res.Error != nil {
		return dbUser, fmt.Errorf("Error")
	}

	return dbUser, nil

}

// Returns user from db pasing user and pass. Returns user and error
func GetUserFromDbByPass(user string, pass string) (model.User, error) {
	return getUserFromdb(user, pass, true)

}

// Returns user from db passing user an reduntant token. Returns user and error
func GetUserFromDbByRedundant(user string, redundant string) (model.User, error) {
	return getUserFromdb(user, redundant, false)
}

func UpdateUserPassword(user string, newPassword string) (bool, error) {
	db := database.DB.Db

	res := db.Where("user = ?", user).Select("password", "first_login").Updates(model.User{Password: newPassword, First_login: 1})
	if res.RowsAffected < 1 {
		return false, fmt.Errorf("user not found")
	}
	return true, nil
}
