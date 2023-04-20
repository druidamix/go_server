// Package handler contain Jwt Middelware handler
package handler

import (
	"net/http"

	"github.com/druidamix/go_server/repository"
	"github.com/druidamix/go_server/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// AuthMiddelware is in charge of verifying  jwt signatures
func AuthMiddelware(aRep *repository.AuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		hUser := c.Request.Header.Get("user")
		hJwToken := c.Request.Header.Get("authorization")

		// get the secrety key from the user
		secretKey, err := service.GetJwtSecret(hUser, aRep)

		if err != nil {
			c.Status(400)
			c.Abort()
			return
		}
		// Initialize a new instance of `Claims`

		claims := &service.Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		token, err := jwt.ParseWithClaims(hJwToken, claims, func(token *jwt.Token) (interface{}, error) {
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
