package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (h *Handlers) staffCreate(c *gin.Context) {
	staff := models.Staff{}
	if err := c.BindJSON(&staff); err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	staff, err := h.service.StaffCreateService(staff, c.MustGet("Admin").(models.Admin))
	if err != nil {
		if err.Error() == utils.ErrRestaurantNotFound {
			log.Info(utils.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		log.Info(err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, staff)
}

func (h *Handlers) staffGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	admin := c.MustGet("Admin").(models.Admin)
	staff, err := h.service.StaffGetByIDService(id, admin)
	if err != nil {
		if err.Error() == utils.ErrStaffNotFound {
			log.Info(utils.ErrStaffNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		if err.Error() == utils.ErrRestaurantNotFound {
			log.Info(utils.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		log.Info(err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, staff)
}

func (h *Handlers) staffGetInRestaurant(c *gin.Context) {
	admin := c.MustGet("Admin").(models.Admin)
	staff, err := h.service.StaffGetAllInRestaurantService(admin)
	if err != nil {
		if err.Error() == utils.ErrRestaurantNotFound {
			log.Info(utils.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		if err.Error() == utils.ErrStaffNotFound {
			log.Info(utils.ErrStaffNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		log.Info(err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, staff)
}

func (h *Handlers) staffUpdate(c *gin.Context) {
	admin := c.MustGet("Admin").(models.Admin)
	var staff models.Staff 
	if err := c.BindJSON(&staff); err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	err := h.service.StaffUpdateService(staff, admin)
	if err != nil {
		if err.Error() == utils.ErrRestaurantNotFound {
			log.Info(utils.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		if err.Error() == utils.ErrStaffNotFound {
			log.Info(utils.ErrStaffNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		log.Error(err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (h *Handlers) staffDelete(c *gin.Context) {
	id := c.Param("id")
	admin := c.MustGet("Admin").(models.Admin)
	err := h.service.StaffDeleteService(uuid.Parse(id), admin)
	if err != nil {
		if err.Error() == utils.ErrRestaurantNotFound {
			log.Info(utils.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		log.Info(err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}