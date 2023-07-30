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
	c.JSON(http.StatusOK, admin)
}

func (h *Handlers) adminGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	admin, err := h.db.AdminGetByID(id)
	if err != nil {
		log.Info("AdminGetByID: ", err)
		c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}

func (h *Handlers) adminGetByEmail(c *gin.Context) {
	email := c.Param("email")
	admin, err := h.db.AdminGetByEmail(email)
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
	if err := h.db.AdminUpdate(admin); err != nil {
		log.Info("AdminUpdate: ", err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}

func (h *Handlers) adminDelete(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	if err := h.db.AdminDelete(id); err != nil {
		log.Info("AdminDelete: ", err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Message{Message: "Admin deleted"})
}
