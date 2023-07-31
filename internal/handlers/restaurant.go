package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/pkg/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (h *Handlers) restaurantCreate(c *gin.Context) {
	admin, exists := c.Get("Admin")
	if !exists {
		c.JSON(http.StatusBadRequest, models.Message{Message: "Admin not found"})
		log.Info("RestaurantCreate: ", exists)
		return
	}
	var r models.Restaurant
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Info("RestaurantCreate: ", err)
		return
	}
	r, err := h.service.RestaurantCreateService(r, uuid.Parse(admin.(models.Admin).ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		log.Info("RestaurantCreate: ", err)
		return
	}
	log.Info("RestaurantCreate: ", r)
	c.JSON(http.StatusOK, r)
}

func (h *Handlers) restaurantGetMine(c *gin.Context) {
	admin := c.MustGet("Admin").(models.Admin)
	if admin.ID == "" {
		c.JSON(http.StatusBadRequest, models.Message{Message: "Admin not found"})
		return
	}
	rest, err := h.service.RestaurantGetByIDService(uuid.Parse(admin.RestaurantID))
	if err != nil {
		log.Info("RestaurantGetByID: ", err)
		c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, rest)
}

func (h *Handlers) restaurantGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	rest, err := h.service.RestaurantGetByIDService(id)
	if err != nil {
		log.Info("RestaurantGetByID: ", err)
		c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, rest)
	log.Info("RestaurantGetByID: restaurant found")
}

func (h *Handlers) restaurantGetByAdminsID(c *gin.Context) {
	id := c.Param("id")
	rest, err := h.service.RestaurantGetByAdminsIDService(uuid.UUID(id))
	if err != nil {
		log.Info("RestaurantGetByAdminsID: ", err)
		c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, rest)
	log.Info("RestaurantGetByAdminsID: restaurant found")
}

func (h *Handlers) restaurantUpdate(c *gin.Context) {
	var r models.Restaurant
	if err := c.ShouldBindJSON(&r); err != nil {
		log.Info("RestaurantUpdate: ", err)
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		return
	}
	if err := h.service.RestaurantUpdateService(r); err != nil {
		log.Info("RestaurantUpdate: ", err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	log.Info("RestaurantUpdate: restaurant updated")
	c.JSON(http.StatusOK, models.Message{Message: "Restaurant updated"})
}

func (h *Handlers) restaurantDelete(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	if err := h.service.RestaurantDeleteService(id); err != nil {
		log.Info("RestaurantDelete: ", err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	log.Info("RestaurantDelete: restaurant deleted")
	c.JSON(http.StatusOK, models.Message{Message: "Restaurant deleted"})
}
