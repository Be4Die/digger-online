package handlers

import (
	"net/http"
	"sync"

	"github.com/Be4Die/digger-online/server-orchestrator-service/internal/domain"
	"github.com/gin-gonic/gin"
)

type ServerHandler struct {
	servers map[string]domain.Server
	mu      sync.RWMutex
}

func NewServerHandler() *ServerHandler {
	return &ServerHandler{
		servers: make(map[string]domain.Server),
	}
}

func (h *ServerHandler) CreateServer(c *gin.Context) {
	var server domain.Server

	if err := c.ShouldBindJSON(&server); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data: " + err.Error()})
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	if _, exists := h.servers[server.ID]; exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Server with this ID already exists"})
		return
	}

	h.servers[server.ID] = server
	c.JSON(http.StatusCreated, server)
}

func (h *ServerHandler) GetAvailableServer(c *gin.Context) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.servers) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "no available servers"})
		return
	}

	for _, server := range h.servers {
		c.JSON(http.StatusOK, server)
		return
	}
}