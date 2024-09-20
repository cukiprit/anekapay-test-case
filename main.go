package main

import (
	"net/http"

	"github.com/cukiprit/anekapay-test-case/db"
	"github.com/cukiprit/anekapay-test-case/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	routes := gin.Default()

	routes.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Ping",
		})
	})

	handlers.SetupRoutes(routes)

	routes.Run()
}
