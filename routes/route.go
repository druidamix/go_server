package routes

import (
	"log"

	"github.com/druidamix/go_server/handlers"
	"github.com/druidamix/go_server/repositories"
	"github.com/druidamix/go_server/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine, aRepo *repositories.AuthRepository, uRepo *repositories.UserRepository) {

	//routes
	app.POST("/login", func(c *gin.Context) {

		hUser := c.Request.Header.Get("user")
		hPass := c.Request.Header.Get("pass")

		user, err := services.GetUser(hUser, hPass, uRepo)

		if err != nil {
			c.JSON(404, gin.H{"error": "Not found"})
		}

		token, err := services.SaveRedundantToken(hUser, aRepo)

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
	})

	app.POST("/login/update_register", func(c *gin.Context) {

		hUser := c.Request.Header.Get("user")
		hPass := c.Request.Header.Get("pass")

		err := services.UpdateUserPass(hUser, hPass, uRepo)

		if err != nil {
			c.Status(404)
			return
		}

		c.String(200, "Register Updated")
	})
	app.POST("/refresh_token", func(c *gin.Context) {

		hRedundant := c.Request.Header.Get("redundant")
		hUser := c.Request.Header.Get("user")

		if len(hRedundant) < 1 || len(hUser) < 1 {
			c.Status(400)
			return
		}

		token, err := services.GenerateJwtToken(hUser, hRedundant, aRepo)

		if err != nil {
			log.Println("-- error generating token")
			c.Status(400)
			return
		}

		c.String(200, token)
	})

	app.Use(handlers.AuthMiddelware())

	app.POST("/stations", handlers.GetStations)
	app.POST("/kpi_realtime", handlers.KpiRealtime)

}
