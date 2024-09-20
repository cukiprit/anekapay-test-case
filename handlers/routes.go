package handlers

import "github.com/gin-gonic/gin"

func SetupRoutes(routes *gin.Engine) {
	routes.POST("/animals", CreateAnimal)
	routes.PUT("/animals", UpdateAnimal)
	routes.DELETE("/animals/:id", DeleteAnimal)
	routes.GET("/animals", GetAllAnimals)
	routes.GET("/animals/:id", GetAnimalByID)
}
