package main

import (
	"nigerian-banks/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	routes.AddAppRoutes(router)

	router.Run(":8000")
}
