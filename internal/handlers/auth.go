package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/pborman/uuid"
)

func (h *Handlers) login(c *gin.Context) {
	var creds *models.Login

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Info("handlers.LogInByEmail: ", err)
		return
	}

	admin, err := h.service.AdminGetByEmailService(creds.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Message{Message: err.Error()})
		log.Error("handlers.LogInByEmail: ", err)
		return
	}

	password, err := h.service.AdminGetPasswordByIdService(uuid.Parse(admin.ID))
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Message{Message: err.Error()})
		log.Error("handlers.LogInByUsername: ", err)
		return
	}

	if ok := utils.CheckPasswordHash(creds.Password, password); ok != true {
		c.JSON(http.StatusBadRequest, models.Message{Message: "invalid password"})
		log.Info("handlers.LogInByEmail: invalid password")
		return
	}

	jwt, err := utils.CreateJWTs(
		h.refreshExp,
		h.accessExp,
		h.refreshSecret,
		h.accessSecret,
		uuid.UUID(admin.ID),
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Message{Message: err.Error()})
	}
	c.Header("access", jwt.Access)
	c.Header("refresh", jwt.Refresh)
	c.Set("Admin", admin)

	c.JSON(http.StatusNoContent, nil)
	log.Info("handlers.LogInByEmail: user logged in")
}

func (h *Handlers) refresh(c *gin.Context) {
	admin_id, err := utils.ValidateJWT(c.Request.Header.Get("refresh"), h.refreshSecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Error("handlers.Refresh: ", err)
		return
	}
	admin, err := h.service.AdminGetByIDService(admin_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Error("handlers.Refresh: ", err)
		return
	}

	jwt, err := utils.CreateJWTs(
		h.refreshExp,
		h.accessExp,
		h.refreshSecret,
		h.accessSecret,
		uuid.UUID(admin.ID),
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Message{Message: err.Error()})
	}

	c.Header("access", jwt.Access)
	c.Header("refresh", jwt.Refresh)

	c.JSON(http.StatusNoContent, nil)
	log.Info("handlers.Refresh: user logged in")
}
