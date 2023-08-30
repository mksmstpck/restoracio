package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

//	@Summary		AdminCreate
//	@Tags			Admin
//	@Description	creates a new admin
//	@ID				admin-create
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.Admin	true	"Admin"
//	@Success		201		{object}	models.Admin
//	@Failure		400		{object}	models.Message
//	@Failure		500		{object}	models.Message
//	@Failure		default	{object}	models.Message
//	@Router			/admin [post]
func (h *Handlers) adminCreate(c *gin.Context) {
	id := c.Param("id")
	admin, err := h.service.AdminCreateService(uuid.Parse(id))
	if err != nil {
		log.Info("AdminCreate: ", err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	admin.Password = ""
	c.JSON(http.StatusCreated, admin)
}

func (h *Handlers) adminValidate(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		log.Info(err)
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		return
	}
	err := h.service.AdminValidateService(admin)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Message{Message: "Email sent"})
}

//	@Summary		AdminGetMe
//	@Security		JWTAuth
//	@Tags			Admin
//	@Description	returns the logged in admin
//	@ID				admin-get-me
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.Admin
//	@Failure		400		{object}	models.Message
//	@Failure		404		{object}	models.Message
//	@Failure		501		{object}	models.Message
//	@Failure		500		{object}	models.Message
//	@Failure		default	{object}	models.Message
//	@Router			/admin/me [get]
func (h *Handlers) adminGetMe(c *gin.Context) {
	admin := c.MustGet("Admin")
	c.JSON(http.StatusOK, admin.(models.Admin))
}

func (h *Handlers) adminGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	admin, err := h.service.AdminGetByIDService(id)
	if err != nil {
		log.Info("AdminGetByID: ", err)
		c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}

func (h *Handlers) adminGetByEmail(c *gin.Context) {
	email := c.Param("email")
	admin, err := h.service.AdminGetByEmailService(email)
	if err != nil {
		log.Info("AdminGetByEmail: ", err)
		c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
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
//	@Param			input	body		models.Admin	true	"Admin"
//	@Success		204		{object}	models.Admin
//	@Failure		400		{object}	models.Message
//	@Failure		404		{object}	models.Message
//	@Failure		500		{object}	models.Message
//	@Failure		default	{object}	models.Message
//	@Router			/admin [put]
func (h *Handlers) adminUpdate(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		log.Info("AdminUpdate: ", err)
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		return
	}
	authAdmin := c.MustGet("Admin").(models.Admin)
	if err := h.service.AdminUpdateService(admin, uuid.Parse(authAdmin.ID)); err != nil {
		log.Info("AdminUpdate: ", err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, admin)
}

//	@Summary		AdminDelete
//	@Security		JWTAuth
//	@Tags			Admin
//	@Description	creates a new admin
//	@ID				admin-delete
//	@Accept			json
//	@Produce		json
//	@Success		204		{object}	nil
//	@Failure		400		{object}	models.Message
//	@Failure		500		{object}	models.Message
//	@Failure		default	{object}	models.Message
//	@Router			/admin [delete]
func (h *Handlers) adminDelete(c *gin.Context) {
	id := uuid.Parse(c.MustGet("Admin").(models.Admin).Restaurant.ID)
	if err := h.service.AdminDeleteService(id); err != nil {
		log.Info("AdminDelete: ", err)
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.Message{Message: "Admin deleted"})
}
