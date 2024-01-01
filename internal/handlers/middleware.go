package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apimodels "github.com/mksmstpck/restoracio/internal/APImodels"
	"github.com/mksmstpck/restoracio/utils"
)

func (h *Handlers) authMiddleware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, apimodels.Message{Message: "Token not found"})
		return
	}
	adminID, err := utils.ValidateJWT(token, h.accessSecret)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, apimodels.Message{Message: err.Error()})
		return
	}

	admin, err := h.service.AdminGetByIDService(adminID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, apimodels.Message{Message: err.Error()})
		return
	}
	c.Set("Admin", admin)
	c.Next()
}

func (h *Handlers) corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "restoracio.fly.dev")
	c.Header("Access-Control-Allow-Methods", "restoracio.fly.dev")
	c.Header("Access-Control-Allow-Headers", "restoracio.fly.dev")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
