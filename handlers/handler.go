package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"

	"github.com/druidamix/go_server/controllers"
	"github.com/druidamix/go_server/database"
	model "github.com/druidamix/go_server/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthRequired(c *gin.Context) error {
	hmacSampleSecret := `Secret`
	tokenString := c.Query("token")
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
	return nil
}

// GetSingleUser from db
func GetSingleUser(c *gin.Context) {
	db := database.DB.Db

	// get id params
	id := c.Query("user")
	var user model.User
	// find single user in the database by id
	db.Find(&user, "user = ?", id)
	if user.ID == 0 {

		c.JSON(200, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, gin.H{"data": user})
}

func Login(c *gin.Context) {

	user := c.Request.Header.Get("user")
	pass := c.Request.Header.Get("pass")
	log.Println("---Entered login")

	dbuser, err := controllers.GetUserFromDbByPass(user, pass)

	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
	}

	token, err := controllers.GenerateRedundantToken(dbuser.User)

	if err != nil {
		c.JSON(404, gin.H{"data": "user not found"})
		return
	}

	if dbuser.First_login == 0 {
		c.JSON(206, gin.H{"redundant_token": token})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func RefreshToken(c *gin.Context) {

	redundant := c.Request.Header.Get("redundant")
	user := c.Request.Header.Get("user")

	log.Println("--Refreshtoken")
	if len(redundant) < 1 || len(user) < 1 {
		c.Status(400)
		return
	}

	token, err := controllers.GenerateJwtToken(user, redundant)

	if err != nil {
		log.Println("-- error generating token")
		c.Status(400)
		return

	}
	c.String(200, token)
}

func GetStations(c *gin.Context) {

	data := (`[
        {"stationName":"Station 1","stationAddr":"Plaça esglesia nº8","stationCode":"stcode1"},
        {"stationName":"Station 2","stationAddr":"Prudenci Murillo nº2","stationCode":"stcode2"}
    ]`)

	c.String(200, data)

}

func KpiRealtime(c *gin.Context) {
	q := []string{"{",
		`"day_power":`, strconv.Itoa(rand.Intn(10)),
		",",
		`"total_power":`, strconv.Itoa(rand.Intn(100)),
		"}"}

	log.Println(strings.Join(q, ""))
	c.String(200, strings.Join(q, ""))
}
