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

// @Summary		AdminCreate
// @Tags			Admin
// @Description	creates a new admin
// @ID				admin-create
// @Accept			json
// @Produce		json
// @Param			input	body		dto.Admin	true	"Admin"
// @Success		201		{object}	dto.Admin
// @Failure		400		{object}apimodels.apimodels.Message
// @Failure		500		{object}apimodels.apimodels.Message
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/admin [post]
func (h *Handlers) adminCreate(c *gin.Context) {
	var a apimodels.AdminRequest
	if err := c.ShouldBindJSON(&a); err != nil {
		log.Info(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apimodels.Message{Message: err.Error()})
		return
	}

	err := h.service.AdminCreateService(convertors.AdminRequestToDTO(&a))
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

// @Summary		AdminGetMe
// @Security		JWTAuth
// @Tags			Admin
// @Description	returns the logged in admin
// @ID				admin-get-me
// @Accept			json
// @Produce		json
// @Success		200		{object}	dto.Admin
// @Failure		400		{object}apimodels.apimodels.Message
// @Failure		404		{object}apimodels.apimodels.Message
// @Failure		501		{object}apimodels.apimodels.Message
// @Failure		500		{object}apimodels.apimodels.Message
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/admin/me [get]
func (h *Handlers) adminGetMe(c *gin.Context) {
	admin := c.MustGet("Admin").(dto.Admin)
	c.AbortWithStatusJSON(http.StatusOK, convertors.AdminDTOToResponse(&admin))
}

func (h *Handlers) adminGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	admin, err := h.service.AdminGetByIDService(id)
	if err != nil {
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, convertors.AdminDTOToResponse(&admin))
}

func (h *Handlers) adminGetByEmail(c *gin.Context) {
	email := c.Param("email")
	admin, err := h.service.AdminGetByEmailService(email)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, convertors.AdminDTOToResponse(&admin))
}

// @Summary		AdminUpdate
// @Security		JWTAuth
// @Tags			Admin
// @Description	creates a new admin
// @ID				admin-update
// @Accept			json
// @Produce		json
// @Param			input	body		dto.Admin	true	"Admin"
// @Success		204		{object}	dto.Admin
// @Failure		400		{object}apimodels.apimodels.Message
// @Failure		404		{object}apimodels.apimodels.Message
// @Failure		500		{object}apimodels.apimodels.Message
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/admin [put]
func (h *Handlers) adminUpdate(c *gin.Context) {
	var admin apimodels.AdminRequest
	if err := c.ShouldBindJSON(&admin); err != nil {
		log.Info(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, apimodels.Message{Message: err.Error()})
		return
	}
	authAdmin := c.MustGet("Admin").(dto.Admin)
	err := h.service.AdminUpdateService(convertors.AdminRequestToDTO(&admin), uuid.Parse(authAdmin.ID))
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// @Summary		AdminDelete
// @Security		JWTAuth
// @Tags			Admin
// @Description	creates a new admin
// @ID				admin-delete
// @Accept			json
// @Produce		json
// @Success		204		{object}	nil
// @Failure		400		{object}apimodels.apimodels.Message
// @Failure		500		{object}apimodels.apimodels.Message
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/admin [delete]
func (h *Handlers) adminDelete(c *gin.Context) {
	id := uuid.Parse(c.MustGet("Admin").(dto.Admin).ID)
	if err := h.service.AdminDeleteService(id); err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
