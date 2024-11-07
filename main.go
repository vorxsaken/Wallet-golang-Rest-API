package main

import (
	"fahmi-wallet/routes"

	"fahmi-wallet/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":5000") // Listen and serve on 0.0.0.0:5000
}
