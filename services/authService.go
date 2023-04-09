package services

import (
	"fmt"
	"log"
	"time"

	"github.com/druidamix/go_server/repositories"
	"github.com/golang-jwt/jwt/v4"
)

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
func GenerateJwtToken(user string, redundant string, rep *repositories.AuthRepository) (string, error) {

	jwtSecretkey, _ := rep.UpdateJwtSecret(user, redundant)

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

func SaveRedundantToken(user string, rep *repositories.AuthRepository) (string, error) {
	token, err := rep.SaveRedundantToken(user)
	if err != nil {
		return "", fmt.Errorf("Error saving")
	}
	return token, nil
}

// GetJwtSecretKey returns jwt secret key
func GetJwtSecret(user string, rep *repositories.AuthRepository) (string, error) {
	secret, err := rep.GetJwtSecretKey(user)

	if err != nil {
		return "", fmt.Errorf("Error obtaining")
	}
	return secret, nil
}
