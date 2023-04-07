package main

import (
	"github.com/druidamix/go_demo_2/database"
	"github.com/druidamix/go_demo_2/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()
	app := gin.Default()
	//app.Use(logger.New())
	routes.SetupRoutes(app)
	app.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	app.Run(":3333")
}
