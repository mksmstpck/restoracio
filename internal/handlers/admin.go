package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

type adminRequest struct {
	Name       string      `json:"name" binding:"required"`
	Email      string      `json:"email" binding:"required"`
	Password   string      `json:"password" binding:"required"`
}

type adminResponse struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	Email      	   string        `json:"email"`
	Restaurant     *RestaurantDB `json:"restaurant"`
}

//	@Summary		AdminCreate
//	@Tags			Admin
//	@Description	creates a new admin
//	@ID				admin-create
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dto.Admin	true	"Admin"
//	@Success		201		{object}	dto.Admin
//	@Failure		400		{object}	dto.Message
//	@Failure		500		{object}	dto.Message
//	@Failure		default	{object}	dto.Message
//	@Router			/admin [post]
func (h *Handlers) adminCreate(c *gin.Context) {
	var a dto.Admin
	if err := c.ShouldBindJSON(&a); err != nil {
		log.Info(err)
		c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
		return
	}
	
	admin, err := h.service.AdminCreateService(a)
	if err != nil {
		if err == errors.New(models.ErrAdminAlreadyExists) {
			c.JSON(http.StatusConflict, dto.Message{Message: err.Error()})
			return
		}
		log.Info(err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, admin)
}

//	@Summary		AdminGetMe
//	@Security		JWTAuth
//	@Tags			Admin
//	@Description	returns the logged in admin
//	@ID				admin-get-me
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	dto.Admin
//	@Failure		400		{object}	dto.Message
//	@Failure		404		{object}	dto.Message
//	@Failure		501		{object}	dto.Message
//	@Failure		500		{object}	dto.Message
//	@Failure		default	{object}	dto.Message
//	@Router			/admin/me [get]
func (h *Handlers) adminGetMe(c *gin.Context) {
	admin := c.MustGet("Admin")
	c.JSON(http.StatusOK, admin.(dto.Admin))
}

func (h *Handlers) adminGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	admin, err := h.service.AdminGetByIDService(id)
	if err != nil {
		if err == errors.New(models.ErrAdminNotFound) {
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		log.Info(err)
		c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}

func (h *Handlers) adminGetByEmail(c *gin.Context) {
	email := c.Param("email")
	admin, err := h.service.AdminGetByEmailService(email)
	if err != nil {
		if err == errors.New(models.ErrAdminNotFound) {
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		log.Info(err)
		c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}

//	@Summary		AdminUpdate
//	@Security		JWTAuth
//	@Tags			Admin
//	@Description	creates a new admin
//	@ID				admin-update
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dto.Admin	true	"Admin"
//	@Success		204		{object}	dto.Admin
//	@Failure		400		{object}	dto.Message
//	@Failure		404		{object}	dto.Message
//	@Failure		500		{object}	dto.Message
//	@Failure		default	{object}	dto.Message
//	@Router			/admin [put]
func (h *Handlers) adminUpdate(c *gin.Context) {
	var admin dto.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		log.Info("AdminUpdate: ", err)
		c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
		return
	}
	authAdmin := c.MustGet("Admin").(dto.Admin)
	if err := h.service.AdminUpdateService(admin, uuid.Parse(authAdmin.ID)); err != nil {
		if err == errors.New(models.ErrAdminNotFound) {
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		log.Info("AdminUpdate: ", err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

//	@Summary		AdminDelete
//	@Security		JWTAuth
//	@Tags			Admin
//	@Description	creates a new admin
//	@ID				admin-delete
//	@Accept			json
//	@Produce		json
//	@Success		204		{object}	nil
//	@Failure		400		{object}	dto.Message
//	@Failure		500		{object}	dto.Message
//	@Failure		default	{object}	dto.Message
//	@Router			/admin [delete]
func (h *Handlers) adminDelete(c *gin.Context) {
	id := uuid.Parse(c.MustGet("Admin").(dto.Admin).Restaurant.ID)
	if err := h.service.AdminDeleteService(id); err != nil {
		if err == errors.New(models.ErrAdminNotFound) {
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		log.Info("AdminDelete: ", err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
