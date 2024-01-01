package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apimodels "github.com/mksmstpck/restoracio/internal/APImodels"
	"github.com/mksmstpck/restoracio/internal/convertors"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

// @Summary		RestaurantCreate
// @Security		JWTAuth
// @Tags			Restaurant
// @Description	creates a new restaurant
// @ID				rest-create
// @Accept			json
// @Produce		json
// @Param			input	body		dto.Restaurant	true	"Restaurant"
// @Success		201		{object}	dto.Restaurant
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/restaurant [post]
func (h *Handlers) restaurantCreate(c *gin.Context) {
	admin := c.MustGet("Admin")
	var r apimodels.RestaurantRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.Info(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apimodels.Message{Message: err.Error()})
		return
	}
	err := h.service.RestaurantCreateService(convertors.RestaurantRequestToDTO(&r), admin.(dto.Admin))
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	log.Info("restaurant created")
	c.JSON(http.StatusCreated, nil)
}

// @Summary		RestaurantGetMine
// @Security		JWTAuth
// @Tags			Restaurant
// @Description	returns an admin's restaurant
// @ID				rest-get-mine
// @Accept			json
// @Produce		json
// @Success		200		{object}	dto.Restaurant
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/restaurant [get]
func (h *Handlers) restaurantGetMine(c *gin.Context) {
	admin := c.MustGet("Admin").(dto.Admin)
	rest, err := h.service.RestaurantGetByIDService(uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
	}
	c.JSON(http.StatusOK, convertors.RestaurantDTOToResponse(&rest))
}

// @Summary		RestaurantGetByID
// @Security		JWTAuth
// @Tags			Restaurant
// @Description	returns a restaurant by id
// @ID				rest-get-by-id
// @Accept			json
// @Produce		json
// @Success		200		{object}	dto.Restaurant
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/restaurant/{id} [get]
// @Param			id	path	string	true	"Restaurant ID"
func (h *Handlers) restaurantGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	rest, err := h.service.RestaurantGetByIDService(id)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
	}
	c.JSON(http.StatusOK, convertors.RestaurantDTOToResponse(&rest))
	log.Info("restaurant found")
}

// @Summary		RestaurantUpdate
// @Security		JWTAuth
// @Tags			Restaurant
// @Description	updates a restaurant
// @ID				rest-update
// @Accept			json
// @Produce		json
// @Param			input	body		dto.Restaurant	true	"Restaurant"
// @Success		204		{object}	nil
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/restaurant [put]
func (h *Handlers) restaurantUpdate(c *gin.Context) {
	var r apimodels.RestaurantRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.Info(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apimodels.Message{Message: err.Error()})
	}
	admin := c.MustGet("Admin").(dto.Admin)
	err := h.service.RestaurantUpdateService(
		convertors.RestaurantRequestToDTO(&r),
		admin,
	)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
	}
	log.Info("restaurant updated")
	c.JSON(http.StatusNoContent, nil)
}

// @Summary		RestaurantDelete
// @Security		JWTAuth
// @Tags			Restaurant
// @Description	deletes a restaurant
// @ID				rest-delete
// @Accept			json
// @Produce		json
// @Success		204		{object}	nil
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/restaurant [delete]
func (h *Handlers) restaurantDelete(c *gin.Context) {
	restaurant := c.MustGet("Admin").(dto.Admin).Restaurant
	if err := h.service.RestaurantDeleteService(restaurant); err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
	}
	log.Info("restaurant deleted")
	c.JSON(http.StatusNoContent, nil)
}
