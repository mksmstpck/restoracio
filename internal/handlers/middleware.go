package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
)

func (h *Handlers) authMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.Message{Message: "Token not found"})
		return
	}
	admin_id, err := utils.ValidateJWT(token, h.accessSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.Message{Message: err.Error()})
		return
	}

	admin, err := h.service.AdminGetByIDService(admin_id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.Message{Message: err.Error()})
		return
	}
	c.Set("Admin", admin)
	c.Next()
}

func (h *Handlers) corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "restoracio.fly.dev")
	c.Header("Access-Control-Allow-Methods", "restoracio.fly.dev")
	c.Header("Access-Control-Allow-Headers", "restoracio.fly.dev")
  
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}