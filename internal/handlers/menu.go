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

// @Summary		MenuCreate
// @Security		JWTAuth
// @Tags			Menu
// @Description	creates a new menu
// @ID				menu-create
// @Accept			json
// @Produce		json
// @Param			input	body		dto.Menu	true	"Menu"
// @Success		201		{object}	dto.Menu
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/menu [post]
func (h *Handlers) menuCreate(c *gin.Context) {
	admin := c.MustGet("Admin")

	var m apimodels.MenuRequest
	if err := c.ShouldBindJSON(&m); err != nil {
		log.Info(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apimodels.Message{Message: err.Error()})
		return
	}
	err := h.service.MenuCreateService(convertors.MenuRequestToDTO(&m), admin.(dto.Admin))
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	log.Info("menu created")
	c.JSON(http.StatusCreated, m)
}

// @Summary		MenuGetWithQrcode
// @Tags			Menu
// @Description	returns a menu by id with qrcode
// @ID				menu-get-by-id-with-qrcode
// @Accept			json
// @Produce		json
// @Success		200		{object}	dto.Menu
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/menu/qr/{id} [get]
// @Param			id	path	string	true	"Menu ID"
func (h *Handlers) menuGetWithQrcode(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	menu, err := h.service.MenuGetWithQrcodeService(id)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, convertors.MenuDTOToResponse(&menu))
	log.Info("menu found")
}

// @Summary		MenuGetByID
// @Tags			Menu
// @Description	returns a menu by id
// @ID				menu-get-by-id
// @Accept			json
// @Produce		json
// @Success		200		{object}	dto.Menu
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/menu/{id} [get]
// @Param			id	path	string	true	"Menu ID"
func (h *Handlers) menuGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	menu, err := h.service.MenuGetByIDService(id)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, convertors.MenuDTOToResponse(&menu))
	log.Info("menu found")
}

// @Summary		MenuUpdate
// @Security		JWTAuth
// @Tags			Menu
// @Description	updates a menu
// @ID				menu-update
// @Accept			json
// @Produce		json
// @Param			input	body		dto.Menu	true	"Menu"
// @Success		204		{object}	nil
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/menu [put]
func (h *Handlers) menuUpdate(c *gin.Context) {
	admin := c.MustGet("Admin").(dto.Admin)
	var m apimodels.MenuRequest
	if err := c.ShouldBindJSON(&m); err != nil {
		log.Info(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apimodels.Message{Message: err.Error()})
		return
	}
	err := h.service.MenuUpdateService(convertors.MenuRequestToDTO(&m), admin)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
	log.Info("menu updated")
}

// @Summary		MenuDelete
// @Security		JWTAuth
// @Tags			Menu
// @Description	deletes a menu
// @ID				menu delete
// @Accept			json
// @Produce		json
// @Success		201		{object}	dto.Menu
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/menu{id} [delete]
// @Param			id	path	string	true	"Menu ID"
func (h *Handlers) menuDelete(c *gin.Context) {
	admin := c.MustGet("Admin").(dto.Admin)
	err := h.service.MenuDeleteService(admin)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	log.Info("menu deleted")
	c.JSON(http.StatusNoContent, nil)
}
