package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/pkg/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (h *Handlers) adminCreate(c *gin.Context) {
	var a models.Admin
	if err := c.ShouldBindJSON(&a); err != nil {
		log.Info("AdminCreate: ", err)
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		return
	}
	admin, err := h.service.AdminCreateService(a)
	if err != nil {
		log.Info("AdminCreate: ", err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	admin.Password = ""
	c.JSON(http.StatusOK, admin)
}

func (h *Handlers) adminGetMe(c *gin.Context) {
	admin := c.MustGet("Admin")
	c.JSON(http.StatusOK, admin.(models.Admin))
}

func (h *Handlers) adminGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	admin, err := h.service.AdminGetByIDService(id)
	if err != nil {
		log.Info("AdminGetByID: ", err)
		c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}

func (h *Handlers) adminGetByEmail(c *gin.Context) {
	email := c.Param("email")
	admin, err := h.service.AdminGetByEmailService(email)
	if err != nil {
		log.Info("AdminGetByEmail: ", err)
		c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}

func (h *Handlers) adminUpdate(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		log.Info("AdminUpdate: ", err)
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		return
	}
	authAdmin := c.MustGet("Admin").(models.Admin)
	if err := h.service.AdminUpdateService(admin, uuid.Parse(authAdmin.ID)); err != nil {
		log.Info("AdminUpdate: ", err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}

func (h *Handlers) adminDelete(c *gin.Context) {
	id := uuid.Parse(c.MustGet("Admin").(models.Admin).Restaurant.ID)
	if err := h.service.AdminDeleteService(id); err != nil {
		log.Info("AdminDelete: ", err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Message{Message: "Admin deleted"})
}
