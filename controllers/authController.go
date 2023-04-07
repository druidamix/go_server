package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/druidamix/go_server/database"
	model "github.com/druidamix/go_server/models"

	"github.com/golang-jwt/jwt/v4"
)

// randToken generates a random hex value.
func randToken(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// Stores the reduntat token on db, then returns it
func GenerateRedundantToken(user string) (string, error) {
	redundant_token, _ := randToken(2048)
	db := database.DB.Db

	dbRes := db.Model(&model.User{}).Where("user=?", user).Update("redundant_token", redundant_token)

	if dbRes.RowsAffected < 1 {
		return "", fmt.Errorf("user not found")
	}

	return redundant_token, nil
}

// Stores the new secret key on db then returns it
func updateJwtSecretKey(user string, redundant string) (string, error) {
	token, _ := randToken(250)

	db := database.DB.Db
	userDb := model.User{User: user}
	dbRes := db.Model(&userDb).Where("user=? and redundant_token=?", user, redundant).Updates(model.User{Token: token})

	if dbRes.RowsAffected < 1 {
		log.Println("-- 0 rows affected")
		return "", fmt.Errorf("")
	}
	return token, nil
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

// Returns a jwt token string
func GenerateJwtToken(user string, redundant string) (string, error) {

	jwtSecretkey, _ := updateJwtSecretKey(user, redundant)

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: user,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(jwtSecretkey))

	if err != nil {
		log.Println("Signing error", err)
		return "", fmt.Errorf("error signing token")
	}

	return tokenString, err
}
