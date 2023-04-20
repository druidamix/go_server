package main

import (
	"github.com/druidamix/go_server/database"
	"github.com/druidamix/go_server/repository"
	"github.com/druidamix/go_server/route"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	app := gin.Default()

	aRepo := repository.NewAuthRepository(database.DB.Db)
	uRepo := repository.NewUserRespository(database.DB.Db)
	route.SetupRoutes(app, aRepo, uRepo)
	app.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	app.Run(":3333")
}
