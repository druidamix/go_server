package handlers

import (
	"net/http"

	"github.com/druidamix/go_server/controllers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// middelware in chargeo of verifying  http requests
func AuthMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		hUser := c.Request.Header.Get("user")
		hJwToken := c.Request.Header.Get("authorization")

		// get the secrety key from the user
		secretKey, err := controllers.GetJwtSecretKey(hUser)

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
