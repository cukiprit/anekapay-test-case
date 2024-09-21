package animal

import (
	"net/http"
	"strconv"

	"github.com/cukiprit/anekapay-test-case/internal/db"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateAnimal(ctx *gin.Context) {
	var newAnimal Animal
	if err := ctx.ShouldBindJSON(&newAnimal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.service.Create(&newAnimal); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, newAnimal)
}

func (h *Handler) UpdateAnimal(ctx *gin.Context) {
	var updatedAnimal Animal
	if err := ctx.ShouldBindJSON(&updatedAnimal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.service.Update(&updatedAnimal); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	ctx.JSON(http.StatusOK, updatedAnimal)
}

func (h *Handler) DeleteAnimal(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.service.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (h *Handler) GetAllAnimals(ctx *gin.Context) {
	animals, err := h.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	ctx.JSON(http.StatusOK, animals)
}

func (h *Handler) GetAnimalByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	animal, err := h.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, animal)
}

func SetupRoutes(r *gin.Engine) {
	dbRepo := NewDBRepository(db.Database)
	service := NewService(dbRepo)
	handler := NewHandler(service)

	r.POST("/animals", handler.CreateAnimal)
	r.PUT("/animals", handler.UpdateAnimal)
	r.DELETE("/animals/:id", handler.DeleteAnimal)
	r.GET("/animals", handler.GetAllAnimals)
	r.GET("/animals/:id", handler.GetAnimalByID)
}
