package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck обработчик для проверки статуса сервиса
// @Summary Проверка здоровья сервиса
// @Description Возвращает статус работы сервиса
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "available",
		"message": "Service is running",
	})
}