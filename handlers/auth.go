package handlers

import (
	"fmt"
	"net/http"

	"github.com/druidamix/go_demo_2/controllers"
	"github.com/druidamix/go_demo_2/database"
	model "github.com/druidamix/go_demo_2/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func getJwtSecretKey(user string) (string, error) {
	db := database.DB.Db

	var _user model.User

	res := db.Where("user = ?", user).Find(&_user)
	if res.RowsAffected < 1 {
		return "", fmt.Errorf("not found")
	}
	return _user.Token, nil
}

func AuthMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Request.Header.Get("user")
		jwtToken := c.Request.Header.Get("authorization")

		// get the secrety key from the user
		secretKey, err := getJwtSecretKey(user)

		if err != nil {
			c.Status(400)
			c.Abort()
			return
		}

		// Initialize a new instance of `Claims`
		claims := &controllers.Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": err})

				return
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": err})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": err})
			return
		}

		c.Next()

	}
}
