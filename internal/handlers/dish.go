package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (h *Handlers) dishCreate(c *gin.Context) {
	admin := c.MustGet("Admin")
	var m models.Dish
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Info(err)
		return
	}
	m, err := h.service.DishCreateService(m, admin.(models.Admin))
	if err != nil {
		if err.Error() == utils.ErrRestaurantNotFound {
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			log.Info(err)
			return
		}
		if err.Error() == utils.ErrMenuNotFound {
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			log.Info(err)
			return
		}
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		log.Info(err)
		return
	}
	log.Info("dish created")
	c.JSON(http.StatusOK, m)
}

func (h *Handlers) dishGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	dish, err := h.service.DishGetByIDService(id)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dish)
	log.Info("dish found")
}

func (h *Handlers) dishGetAllInMenu(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	dishes, err := h.service.DishGetAllInMenuService(id)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dishes)
	log.Info("dishes found")
}

func (h *Handlers) dishUpdate(c *gin.Context) {
	admin := c.MustGet("Admin")
	var m models.Dish
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Info(err)
		return
	}

	err := h.service.DishUpdateService(m, admin.(models.Admin))
	if err != nil {
		if err.Error() == utils.ErrDishNotFound {
			log.Info(utils.ErrDishNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		if err.Error() == utils.ErrRestaurantNotFound {
			log.Info(utils.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		if err.Error() == utils.ErrMenuNotFound {
			log.Info(utils.ErrMenuNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		log.Info(err)
		return
	}
	log.Info("dish updated")
	c.JSON(http.StatusNoContent, nil)
}

func (h *Handlers) dishDelete(c *gin.Context) {
	admin := c.MustGet("Admin")
	id := uuid.Parse(c.Param("id"))
	err := h.service.DishDeleteService(id, admin.(models.Admin))
	if err != nil {
		if err.Error() == utils.ErrDishNotFound {
			log.Info(utils.ErrDishNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		if err.Error() == utils.ErrRestaurantNotFound {
			log.Info(utils.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		if err.Error() == utils.ErrMenuNotFound {
			log.Info(utils.ErrMenuNotFound)
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		log.Info(err)
		return
	}
	log.Info("dish deleted")
	c.JSON(http.StatusNoContent, nil)
}