package main

import (
	"exchangeapp/config"
	"exchangeapp/router"
)

func main() {
	config.InitConfig()
	router := router.SetupRouter()
	port := config.AppConfig.App.Port
	if port == "" {
		port = ":8080"
	}
	router.Run(port)
}
