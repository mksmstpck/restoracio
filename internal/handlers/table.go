package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (h *Handlers) tableCreate(c *gin.Context){
	admin := c.MustGet("Admin")
	var t models.Table
	if err := c.ShouldBindJSON(&t); err != nil{
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Info("handlers.tableCreate: ", err)
		return
	}
	t, err := h.service.TableCreateService(t, admin.(models.Admin))
	if err != nil{
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		log.Info("handlers.tableCreate: ", err)
		return
	}
	log.Info("tablec created")
	c.JSON(http.StatusOK, t)
}

func (h *Handlers) tableGetByID(c *gin.Context){
	id := uuid.Parse(c.Param("id"))
	t, err := h.service.TableGetByIDService(id)
	if err != nil{
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		log.Info("handlers.tableGetByID: ", err)
		return
	}
	log.Info("tableGetByID: ", t)
	c.JSON(http.StatusOK, t)
}

func (h *Handlers) tableGetAllInRestaurant(c *gin.Context){
	id := uuid.Parse(c.Param("id"))
	t, err := h.service.TableGetAllInRestaurantService(id)
	if err != nil{
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		log.Info(err)
		return
	}
	log.Info("tables were gotten")
	c.JSON(http.StatusOK, t)
}

func (h *Handlers) tableUpdate(c *gin.Context){
	admin := c.MustGet("Admin").(models.Admin)
	var t models.Table
	if err := c.ShouldBindJSON(&t); err != nil{
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Info(err)
		return
	}
	err := h.service.TableUpdateService(t, admin)
	if err != nil{
		if err.Error() == utils.ErrTableNotFound{
			log.Info("table not found")
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		if err.Error() == utils.ErrRestaurantNotFound{
			log.Info("restaurant not found")
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		log.Info(err)
		return
	}
	log.Info("tableUpdate: ", t)
	c.JSON(http.StatusNoContent, nil)
}

func (h *Handlers) tableDelete(c *gin.Context){
	id := uuid.Parse(c.Param("id"))
	admin := c.MustGet("Admin").(models.Admin)
	err := h.service.TableDeleteService(id, admin)
	if err != nil{
		if err.Error() == utils.ErrTableNotFound{
			log.Info("table not found")
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		if err.Error() == utils.ErrRestaurantNotFound{
			log.Info("restaurant not found")
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		log.Info("handlers.tableDelete: ", err)
		return
	}
	log.Info("tableDelete: ", id)
	c.JSON(http.StatusNoContent, nil)
}