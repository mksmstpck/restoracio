package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/pkg/models"
	"github.com/mksmstpck/restoracio/pkg/utils"
)

func (h *Handlers) DeserializeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		admin_id, err := utils.ValidateJWT(token, h.access_secret)

		admin, err := h.db.AdminGetByID(admin_id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, models.Message{Message: err.Error()})
			return
		}
		c.Set("Admin", admin)
		c.Next()
	}
}