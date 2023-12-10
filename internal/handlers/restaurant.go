package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

//	@Summary		RestaurantCreate
//	@Security		JWTAuth
//	@Tags			Restaurant
//	@Description	creates a new restaurant
//	@ID				rest-create
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dto.Restaurant	true	"Restaurant"
//	@Success		201		{object}	dto.Restaurant
//	@Failure		default	{object}	dto.Message
//	@Router			/restaurant [post]
func (h *Handlers) restaurantCreate(c *gin.Context) {
	admin, exists := c.Get("Admin")
	if !exists {
		c.JSON(http.StatusBadRequest, dto.Message{Message: "Admin not found"})
		log.Info("RestaurantCreate: ", exists)
		return
	}
	var r dto.Restaurant
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
		log.Info("RestaurantCreate: ", err)
		return
	}
	r, err := h.service.RestaurantCreateService(r, admin.(dto.Admin))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		log.Info("RestaurantCreate: ", err)
		return
	}
	log.Info("RestaurantCreate: ", r)
	c.JSON(http.StatusCreated, r)
}

//	@Summary		RestaurantGetMine
//	@Security		JWTAuth
//	@Tags			Restaurant
//	@Description	returns an admin's restaurant
//	@ID				rest-get-mine
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	dto.Restaurant
//	@Failure		default	{object}	dto.Message
//	@Router			/restaurant [get]
func (h *Handlers) restaurantGetMine(c *gin.Context) {
	admin := c.MustGet("Admin").(dto.Admin)
	rest, err := h.service.RestaurantGetByIDService(uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Info("RestaurantGetByID: ", err)
		c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, rest)
}

//	@Summary		RestaurantGetByID
//	@Security		JWTAuth
//	@Tags			Restaurant
//	@Description	returns a restaurant by id
//	@ID				rest-get-by-id
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	dto.Restaurant
//	@Failure		default	{object}	dto.Message
//	@Router			/restaurant/{id} [get]
//	@Param			id	path	string	true	"Restaurant ID"
func (h *Handlers) restaurantGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	rest, err := h.service.RestaurantGetByIDService(id)
	if err != nil {
		log.Info("RestaurantGetByID: ", err)
		c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, rest)
	log.Info("RestaurantGetByID: restaurant found")
}

//	@Summary		RestaurantUpdate
//	@Security		JWTAuth
//	@Tags			Restaurant
//	@Description	updates a restaurant
//	@ID				rest-update
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dto.Restaurant	true	"Restaurant"
//	@Success		204		{object}	nil
//	@Failure		default	{object}	dto.Message
//	@Router			/restaurant [put]
func (h *Handlers) restaurantUpdate(c *gin.Context) {
	var r dto.Restaurant
	if err := c.ShouldBindJSON(&r); err != nil {
		log.Info("RestaurantUpdate: ", err)
		c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
		return
	}
	admin := c.MustGet("Admin").(dto.Admin)
	if err := h.service.RestaurantUpdateService(r, uuid.UUID(admin.Restaurant.ID)); err != nil {
		log.Info("RestaurantUpdate: ", err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	log.Info("RestaurantUpdate: restaurant updated")
	c.JSON(http.StatusNoContent, nil)
}

//	@Summary		RestaurantDelete
//	@Security		JWTAuth
//	@Tags			Restaurant
//	@Description	deletes a restaurant
//	@ID				rest-delete
//	@Accept			json
//	@Produce		json
//	@Success		204		{object}	nil
//	@Failure		default	{object}	dto.Message
//	@Router			/restaurant [delete]
func (h *Handlers) restaurantDelete(c *gin.Context) {
	restaurant := c.MustGet("Admin").(dto.Admin).Restaurant
	if err := h.service.RestaurantDeleteService(restaurant); err != nil {
		log.Info("RestaurantDelete: ", err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	log.Info("RestaurantDelete: restaurant deleted")
	c.JSON(http.StatusNoContent, nil)
}
