// Route packages manages gin routes
package route

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/druidamix/go_server/handler"
	"github.com/druidamix/go_server/repository"
	"github.com/druidamix/go_server/service"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine, aRepo *repository.AuthRepository, uRepo *repository.UserRepository) {
	app.POST("/login", func(c *gin.Context) {

		user := c.Request.Header.Get("user")
		pass := c.Request.Header.Get("pass")

		dbUser, err := service.GetUser(user, pass, uRepo)

		if err != nil {
			c.JSON(404, gin.H{"error": "Not found"})
		}

		token, err := service.SaveRedundantToken(user, aRepo)

		if err != nil {
			c.JSON(404, gin.H{"error": "Not found"})
			return
		}
		// if first time login correctly, we return 206 (change password)
		if dbUser.First_login == 0 {
			c.JSON(206, gin.H{"redundant_token": token})
			return
		}

		c.JSON(200, gin.H{"redundant_token": token})
	})

	app.POST("/login/update_register", func(c *gin.Context) {

		user := c.Request.Header.Get("user")
		pass := c.Request.Header.Get("pass")

		err := service.UpdateUserPass(user, pass, uRepo)

		if err != nil {
			c.Status(404)
			return
		}

		c.String(200, "Register Updated")
	})

	app.POST("/refresh_token", func(c *gin.Context) {

		redundant := c.Request.Header.Get("redundant")
		user := c.Request.Header.Get("user")

		if len(redundant) < 1 || len(user) < 1 {
			c.Status(400)
			return
		}

		token, err := service.GenerateJwtToken(user, redundant, aRepo)

		if err != nil {
			c.Status(400)
			return
		}

		c.String(200, token)
	})

	app.Use(handler.AuthMiddelware(aRepo))

	app.POST("/stations", func(c *gin.Context) {

		data := (`[
			{"stationName":"Station 1","stationAddr":"Plaça esglesia nº8","stationCode":"stcode1"},
        		{"stationName":"Station 2","stationAddr":"Prudenci Murillo nº2","stationCode":"stcode2"}
    ]`)

		c.String(200, data)
	})

	app.POST("/kpi_realtime", func(c *gin.Context) {

		q := []string{"{",
			`"day_power":`, strconv.Itoa(rand.Intn(10)),
			",",
			`"total_power":`, strconv.Itoa(rand.Intn(100)),
			"}"}

		c.String(200, strings.Join(q, ""))
	})
}
