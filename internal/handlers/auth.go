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
	if c.Request.Method != http.MethodPost {
		err := h.t.Render(c.Writer, "login.html", nil)
		if err != nil {
			log.Error(err)
			return
		}
		c.Next()
	}
	creds := models.Login{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	log.Info(creds)

	admin, err := h.service.AdminGetByEmailService(creds.Email)
	if err != nil {
		h.t.Render(c.Writer, "login.html", models.Message{Success: false, Message: err.Error()})
		log.Error(err)
		return
	}

	password, err := h.service.AdminGetPasswordByIdService(uuid.Parse(admin.ID))
	if err != nil {
		h.t.Render(c.Writer, "login.html", models.Message{Success: false, Message: err.Error()})
		log.Error(err)
		return
	}

	if ok := utils.CheckPasswordHash(creds.Password, password); ok != true {
		h.t.Render(c.Writer, "login.html", models.Message{Success: false, Message: "invalid password"})
		log.Info("invalid password")
		return
	}

	access, err := utils.CreateJWT(h.accessExp, h.accessSecret, uuid.Parse(admin.ID))
	if err != nil {
		h.t.Render(c.Writer, "login.html", models.Message{Success: false, Message: err.Error()})
		log.Error(err)
		return
	}

	c.SetCookie("access", access, int(h.accessExp.Seconds()), "/", "localhost:8080", false, true)

	h.t.Render(c.Writer, "login.html", struct{Status bool}{Status: true})
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