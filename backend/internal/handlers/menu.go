package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/backend/internal/models"
	"github.com/mksmstpck/restoracio/backend/utils"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

//	@Summary		MenuCreate
//	@Security		JWTAuth
//	@Tags			Menu
//	@Description	creates a new menu
//	@ID				menu-create
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.Menu	true	"Menu"
//	@Success		201		{object}	models.Menu
//	@Failure		default	{object}	models.Message
//	@Router			/menu [post]
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
	c.JSON(http.StatusCreated, m)
}

//	@Summary		MenuGetWithQrcode
//	@Tags			Menu
//	@Description	returns a menu by id with qrcode
//	@ID				menu-get-by-id-with-qrcode
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.Menu
//	@Failure		default	{object}	models.Message
//	@Router			/menu/qr/{id} [get]
//	@Param			id	path	string	true	"Menu ID"
func (h *Handlers) menuGetWithQrcode(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	menu, err := h.service.MenuGetWithQrcodeService(id)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, menu)
	log.Info("menu found")
}

//	@Summary		MenuGetByID
//	@Tags			Menu
//	@Description	returns a menu by id
//	@ID				menu-get-by-id
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.Menu
//	@Failure		default	{object}	models.Message
//	@Router			/menu/{id} [get]
//	@Param			id	path	string	true	"Menu ID"
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

//	@Summary		MenuUpdate
//	@Security		JWTAuth
//	@Tags			Menu
//	@Description	updates a menu
//	@ID				menu-update
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.Menu	true	"Menu"
//	@Success		204		{object}	nil
//	@Failure		default	{object}	models.Message
//	@Router			/menu [put]
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
	c.JSON(http.StatusNoContent, nil)
	log.Info("menu updated")
}

//	@Summary		MenuDelete
//	@Security		JWTAuth
//	@Tags			Menu
//	@Description	deletes a menu
//	@ID				menu delete
//	@Accept			json
//	@Produce		json
//	@Success		201		{object}	models.Menu
//	@Failure		default	{object}	models.Message
//	@Router			/menu{id} [delete]
//	@Param			id	path	string	true	"Menu ID"
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
	c.JSON(http.StatusNoContent, nil)
}