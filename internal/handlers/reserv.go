package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (h *Handlers) reservCreate(c *gin.Context) {
	var reserv models.ReservAPI
	if err := c.BindJSON(&reserv); err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	reservDB, err := h.service.ReservCreateService(reserv, c.MustGet("Admin").(models.Admin))
	if err != nil {
		if err.Error() == utils.ErrRestaurantNotFound {
			log.Info(utils.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		log.Error(err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, reservDB)
}

func (h *Handlers) reservGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	reservDB, err := h.service.ReservGetByIDService(id, c.MustGet("Admin").(models.Admin))
	if err != nil {
		if err.Error() == utils.ErrReservationNotFound {
			log.Info(utils.ErrReservationNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		if err.Error() == utils.ErrRestaurantNotFound {
			log.Info(utils.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		log.Error(err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	log.Info("reservation found")
	c.JSON(http.StatusOK, reservDB)
}

func (h *Handlers) reservGetAllInRestaurant(c *gin.Context) {
	reservDB, err := h.service.ReservGetAllInRestaurantService(c.MustGet("Admin").(models.Admin))
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
	log.Info("reservations were gotten")
	c.JSON(http.StatusOK, reservDB)
}

func (h *Handlers) reservUpdate(c *gin.Context) {
	var reserv models.ReservAPI
	if err := c.BindJSON(&reserv); err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	err := h.service.ReservUpdateService(reserv, c.MustGet("Admin").(models.Admin))
	if err != nil {
		if err.Error() == utils.ErrReservationNotFound {
			log.Info(utils.ErrReservationNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		if err.Error() == utils.ErrRestaurantNotFound {
			log.Info(utils.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		if err.Error() == utils.ErrTableNotFound {
			log.Info(utils.ErrTableNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		log.Info(err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	log.Info("reservation updated")
	c.JSON(http.StatusNoContent, nil)
}

func (h *Handlers) reservDelete(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	err := h.service.ReservDeleteService(id, c.MustGet("Admin").(models.Admin))
	if err != nil {
		if err.Error() == utils.ErrReservationNotFound {
			log.Info(utils.ErrReservationNotFound)
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
	log.Info("reservation deleted")
	c.JSON(http.StatusNoContent, nil)
}