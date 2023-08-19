package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (h *Handlers) menuCreate(c *gin.Context) {
	admin := c.MustGet("Admin")

	var m models.Menu
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Info(err)
		return
	}
	m, err := h.service.MenuCreateService(m, admin.(models.Admin))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		log.Info(err)
		return
	}
	log.Info("menu created")
	c.JSON(http.StatusOK, m)
}

func (h *Handlers) menuGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	menu, err := h.service.MenuGetByIDService(id)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, menu)
	log.Info("menu found")
}

func (h *Handlers) menuUpdate(c *gin.Context) {
	admin := c.MustGet("Admin").(models.Admin)
	var m models.Menu
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Info(err)
		return
	}
	err := h.service.MenuUpdateService(m, admin)
	if err != nil {
		if err.Error() == utils.ErrMenuNotFound {
			log.Info(utils.ErrMenuNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		if err.Error() == utils.ErrRestaurantNotFound {
			log.Info(utils.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		log.Info(err)
		return
	}
	log.Info("menu updated")
}

func (h *Handlers) menuDelete(c *gin.Context) {
	admin := c.MustGet("Admin").(models.Admin)
	err := h.service.MenuDeleteService(admin)
	if err != nil {
		if err.Error() == utils.ErrMenuNotFound {
			log.Info(utils.ErrMenuNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		if err.Error() == utils.ErrRestaurantNotFound {
			log.Info(utils.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		log.Info(err)
		return
	}
	log.Info("menu deleted")
	c.JSON(http.StatusOK, models.Message{Message: "menu deleted"})
}