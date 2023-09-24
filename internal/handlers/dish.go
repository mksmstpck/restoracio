package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

//	@Summary		DishCreate
//	@Security		JWTAuth
//	@Tags			Dish
//	@Description	creates a new dish
//	@ID				dish-create
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.Dish	true	"Dish"
//	@Success		201		{object}	models.Dish
//	@Failure		default	{object}	models.Message
//	@Router			/dish [post]
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
	c.JSON(http.StatusCreated, m)
}

//	@Summary		DishGetByID
//	@Tags			Dish
//	@Description	returns a dish by id
//	@ID				dish-get-by-id
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.Dish
//	@Failure		default	{object}	models.Message
//	@Router			/dish$id [get]
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

//	@Summary		DishGetAllInMenu
//	@Tags			Dish
//	@Description	returns all dishes in a menu
//	@ID				dish-get-all-in-menu
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	[]models.Dish
//	@Failure		default	{object}	models.Message
//	@Router			/dish/all/{id} [get]
//	@Param			id	path	string	true	"Menu ID"
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

//	@Summary		DishUpdate
//	@Security		JWTAuth
//	@Tags			Dish
//	@Description	updates a dish
//	@ID				dish-update
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.Dish	true	"Dish"
//	@Success		204		{object}	nil
//	@Failure		default	{object}	models.Message
//	@Router			/dish [put]
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

//	@Summary		DishDelete
//	@Security		JWTAuth
//	@Tags			Dish
//	@Description	deletes a dish
//	@ID				dish-delete
//	@Accept			json
//	@Produce		json
//	@Success		204		{object}	nil
//	@Failure		default	{object}	models.Message
//	@Router			/dish/{id} [delete]
//	@Param			id	path	string	true	"Dish ID"
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