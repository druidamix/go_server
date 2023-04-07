package controllers

import (
	"fmt"
	"log"
	"strconv"

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

func UpdateUserPassword(user string, newPassword string) error {
	db := database.DB.Db
	log.Println("user: " + user)
	//_user := model.User{User: user}

	rowsAffected := db.Where("user=?", user).Updates(model.User{Password: newPassword, First_login: 1}).RowsAffected
	//res := db.Where("user = ?", user).Select("password", "first_login").Updates(model.User{Password: newPassword, First_login: 1})

	log.Println("--rows affected: " + strconv.Itoa(int(rowsAffected)))
	if rowsAffected < 1 {
		return fmt.Errorf("user not found")
	}
	return nil
}
