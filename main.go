package main

import (
	"jwt-authentication-golang/controllers"
	"jwt-authentication-golang/database"
	"jwt-authentication-golang/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connet("data source=.;initial catalog=GolangJWTDB;integrated security=True;MultipleActiveResultSets=True;")
	database.Migrate()

	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	// router.SetTrustedProxies([]string{"127.0.0.1"})

	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
