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

// @Summary		DishCreate
// @Security		JWTAuth
// @Tags			Dish
// @Description	creates a new dish
// @ID				dish-create
// @Accept			json
// @Produce		json
// @Param			input	body		dto.Dish	true	"Dish"
// @Success		201		{object}	dto.Dish
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/dish [post]
func (h *Handlers) dishCreate(c *gin.Context) {
	admin := c.MustGet("Admin")
	var m apimodels.DishRequest
	if err := c.ShouldBindJSON(&m); err != nil {
		log.Info(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apimodels.Message{Message: err.Error()})
		return
	}
	err := h.service.DishCreateService(convertors.DishRequestToDTO(&m), admin.(dto.Admin))
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	log.Info("dish created")
	c.JSON(http.StatusCreated, nil)
}

// @Summary		DishGetByID
// @Tags			Dish
// @Description	returns a dish by id
// @ID				dish-get-by-id
// @Accept			json
// @Produce		json
// @Success		200		{object}	dto.Dish
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/dish$id [get]
func (h *Handlers) dishGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	dish, err := h.service.DishGetByIDService(id)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, convertors.DishDTOToResponse(&dish))
	log.Info("dish found")
}

// @Summary		DishGetAllInMenu
// @Tags			Dish
// @Description	returns all dishes in a menu
// @ID				dish-get-all-in-menu
// @Accept			json
// @Produce		json
// @Success		200		{object}	[]dto.Dish
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/dish/all/{id} [get]
// @Param			id	path	string	true	"Menu ID"
func (h *Handlers) dishGetAllInMenu(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	dishes, err := h.service.DishGetAllInMenuService(id)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	dishesRes := make([]*apimodels.DishResponse, len(dishes))
	for i, d := range dishes {
		dishesRes[i] = convertors.DishDTOToResponse(&d)
	}
	c.JSON(http.StatusOK, dishesRes)
	log.Info("dishes found")
}

// @Summary		DishUpdate
// @Security		JWTAuth
// @Tags			Dish
// @Description	updates a dish
// @ID				dish-update
// @Accept			json
// @Produce		json
// @Param			input	body		dto.Dish	true	"Dish"
// @Success		204		{object}	nil
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/dish [put]
func (h *Handlers) dishUpdate(c *gin.Context) {
	admin := c.MustGet("Admin")
	var m apimodels.DishRequest
	if err := c.ShouldBindJSON(&m); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, apimodels.Message{Message: err.Error()})
		log.Info(err)
		return
	}

	err := h.service.DishUpdateService(convertors.DishRequestToDTO(&m), admin.(dto.Admin))
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	log.Info("dish updated")
	c.JSON(http.StatusNoContent, nil)
}

// @Summary		DishDelete
// @Security		JWTAuth
// @Tags			Dish
// @Description	deletes a dish
// @ID				dish-delete
// @Accept			json
// @Produce		json
// @Success		204		{object}	nil
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/dish/{id} [delete]
// @Param			id	path	string	true	"Dish ID"
func (h *Handlers) dishDelete(c *gin.Context) {
	admin := c.MustGet("Admin")
	id := uuid.Parse(c.Param("id"))
	err := h.service.DishDeleteService(id, admin.(dto.Admin))
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	log.Info("dish deleted")
	c.JSON(http.StatusNoContent, nil)
}
