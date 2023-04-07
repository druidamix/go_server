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

	user := c.Request.Header.Get("user")
	pass := c.Request.Header.Get("pass")

	dbuser, err := controllers.GetUserFromDbByPass(user, pass)

	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
	}

	token, err := controllers.GenerateRedundantToken(dbuser.User)

	if err != nil {
		c.JSON(404, gin.H{"data": "Not founds"})
		return
	}

	if dbuser.First_login == 0 {
		c.JSON(206, gin.H{"redundant_token": token})
		return
	}

	c.JSON(200, gin.H{"redundant_token": token})
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

	log.Println(strings.Join(q, ""))
	c.String(200, strings.Join(q, ""))
}

func UpdateUserPassword(c *gin.Context) {
	user := c.Request.Header.Get("user")
	pass := c.Request.Header.Get("pass")

	err := controllers.UpdateUserPassword(user, pass)

	if err != nil {
		c.Status(404)
		return
	}

	c.String(200, "Register Updated")
}
