package routes

import (
	"github.com/druidamix/go_server/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	//grouping
	//api := app.Group("/api")

	//routes
	app.POST("/login", handlers.Login)
	app.POST("/login/update_register", handlers.UpdateUserPassword)

	app.POST("/refresh_token", handlers.RefreshToken)
	app.Use(handlers.AuthMiddelware())

	app.POST("/stations", handlers.GetStations)
	app.POST("/kpi_realtime", handlers.KpiRealtime)

}
