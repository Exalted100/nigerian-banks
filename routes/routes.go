package routes

import (
	"github.com/gin-gonic/gin"
)

func AddAppRoutes(r *gin.Engine) {

	r.GET("/health", healthCheck)
	r.POST("/get-banks", getBanks)
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    "SUCCESS",
		"message": "The RESTful server is OK!",
	})
}
