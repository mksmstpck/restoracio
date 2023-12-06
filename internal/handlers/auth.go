package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/pborman/uuid"
)

//	@Summary		Login
//	@Tags			Auth
//	@Description	logs in an admin
//	@ID				login
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.Login	true	"Login"
//	@Success		204		{object}	nil
//	@Failure		401		{object}	models.Message
//	@Failure		500		{object}	models.Message
//	@Failure		default	{object}	models.Message
//	@Router			/auth/login [post]
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

	admin, err = h.service.AdminGetWithPasswordByIdService(uuid.Parse(admin.ID))
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Message{Message: err.Error()})
		log.Error("handlers.LogInByUsername: ", err)
		return
	}

	if ok := utils.CheckPasswordHash(creds.Password, admin.Password, admin.Salt); ok != true {
		c.JSON(http.StatusBadRequest, models.Message{Message: "invalid password"})
		log.Info("handlers.LogInByEmail: invalid password")
		return
	}

	access, err := utils.CreateJWT(h.accessExp, h.accessSecret, uuid.Parse(admin.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Error("handlers.LogInByEmail: ", err)
		return
	}

	refresh, err := utils.CreateJWT(h.refreshExp, h.refreshSecret, uuid.Parse(admin.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Error("handlers.LogInByEmail: ", err)
		return
	}
	c.Header("access", access)
	c.Header("refresh", refresh)
	c.Set("Admin", admin)

	c.JSON(http.StatusNoContent, nil)
	log.Info("handlers.LogInByEmail: user logged in")
}

//	@Summary		Refresh
//	@Tags			Auth
//	@Description	gives a new access and refresh token
//	@ID				refresh
//	@Accept			json
//	@Produce		json
//	@Param			refresh	header		string	true	"Refresh token"
//	@Success		204		{object}	nil
//	@Failure		401		{object}	models.Message
//	@Failure		500		{object}	models.Message
//	@Failure		default	{object}	models.Message
//	@Router			/auth/refresh [post]
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

	access, err := utils.CreateJWT(h.accessExp, h.accessSecret, uuid.Parse(admin.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Error("handlers.LogInByUsername: ", err)
		return
	}

	refresh, err := utils.CreateJWT(h.refreshExp, h.refreshSecret, uuid.Parse(admin.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Error("handlers.LogInByUsername: ", err)
		return
	}
	c.Header("access", access)
	c.Header("refresh", refresh)

	c.JSON(http.StatusNoContent, nil)
	log.Info("handlers.Refresh: user logged in")
}