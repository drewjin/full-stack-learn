package router

import (
	"exchangeapp/controllers"
	"exchangeapp/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	auth := router.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}
	api := router.Group("/api")
	api.GET("/exchange_rates", controllers.GetExchangeRates)
	api.Use(middlewares.AuthMiddleWare())
	{
		api.POST("/exchange_rates", controllers.CreateExchangeRate)
	}
	return router
}
