package handlers

import (
	"log"
	"math/rand"
	"strconv"
	"strings"

	"github.com/druidamix/go_server/controllers"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	hUser := c.Request.Header.Get("user")
	hPass := c.Request.Header.Get("pass")

	user, err := controllers.GetUserFromDbByPass(hUser, hPass)

	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
	}

	token, err := controllers.GenerateRedundantToken(user.User)

	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}
	// if first time login correctly, we return 206 (change password)
	if user.First_login == 0 {
		c.JSON(206, gin.H{"redundant_token": token})
		return
	}

	c.JSON(200, gin.H{"redundant_token": token})
}

func RefreshToken(c *gin.Context) {

	hRedundant := c.Request.Header.Get("redundant")
	hUser := c.Request.Header.Get("user")

	if len(hRedundant) < 1 || len(hUser) < 1 {
		c.Status(400)
		return
	}

	token, err := controllers.GenerateJwtToken(hUser, hRedundant)

	if err != nil {
		log.Println("-- error generating token")
		c.Status(400)
		return
	}

	c.String(200, token)
}

// returns demo data
func GetStations(c *gin.Context) {

	data := (`[
        {"stationName":"Station 1","stationAddr":"Plaça esglesia nº8","stationCode":"stcode1"},
        {"stationName":"Station 2","stationAddr":"Prudenci Murillo nº2","stationCode":"stcode2"}
    ]`)

	c.String(200, data)

}

// return demo data
func KpiRealtime(c *gin.Context) {
	q := []string{"{",
		`"day_power":`, strconv.Itoa(rand.Intn(10)),
		",",
		`"total_power":`, strconv.Itoa(rand.Intn(100)),
		"}"}

	c.String(200, strings.Join(q, ""))
}

func UpdateUserPassword(c *gin.Context) {
	hUser := c.Request.Header.Get("user")
	hPass := c.Request.Header.Get("pass")

	err := controllers.UpdateUserPassword(hUser, hPass)

	if err != nil {
		c.Status(404)
		return
	}

	c.String(200, "Register Updated")
}
