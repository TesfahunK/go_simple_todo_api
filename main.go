package main

import (
	"simple_api/src/app/routes"
	"simple_api/src/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitialSetups()

	router := gin.Default()

	routes.RegisterAllRoutes(router)

	err := router.Run("localhost:8080")

	if err != nil {

		panic("[Erro] failed to start server" + err.Error())

	}

}
