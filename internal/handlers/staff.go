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

// @Summary		StaffCreate
// @Security		JWTAuth
// @Tags			Staff
// @Description	creates a new employee
// @ID				staff-create
// @Accept			json
// @Produce		json
// @Param			input	body		dto.Staff	true	"Staff"
// @Success		201		{object}	dto.Staff
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/staff [post]
func (h *Handlers) staffCreate(c *gin.Context) {
	var staff apimodels.StaffRequest
	if err := c.BindJSON(&staff); err != nil {
		log.Info(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, apimodels.Message{Message: err.Error()})
	}
	err := h.service.StaffCreateService(
		convertors.StaffRequestToDTO(&staff),
		c.MustGet("Admin").(dto.Admin),
	)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
	}
	c.JSON(http.StatusCreated, nil)
}

// @Summary		StaffGetByID
// @Security		JWTAuth
// @Tags			Staff
// @Description	returns an employee by id
// @ID				staff-get-by-id
// @Accept			json
// @Produce		json
// @Success		200		{object}	dto.Staff
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/staff/{id} [get]
// @Param			id	path	string	true	"Staff ID"
func (h *Handlers) staffGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	admin := c.MustGet("Admin").(dto.Admin)
	staff, err := h.service.StaffGetByIDService(id, admin)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
	}
	c.JSON(http.StatusOK, convertors.StaffDTOToResponse(&staff))
}

// @Summary		StaffGetAllInRestaurant
// @Security		JWTAuth
// @Tags			Staff
// @Description	returns all staff in a restaurant
// @ID				staff-get-all
// @Accept			json
// @Produce		json
// @Success		200		{object}	dto.Staff
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/staff [get]
func (h *Handlers) staffGetInRestaurant(c *gin.Context) {
	admin := c.MustGet("Admin").(dto.Admin)
	staff, err := h.service.StaffGetAllInRestaurantService(admin)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
	}
	staffs := make([]*apimodels.StaffResponse, len(staff))
	for i, s := range staff {
		staffs[i] = convertors.StaffDTOToResponse(&s)
	}
	c.JSON(http.StatusOK, staffs)
}

// @Summary		StaffUpdate
// @Security		JWTAuth
// @Tags			Staff
// @Description	updates an employee
// @ID				staff-update
// @Accept			json
// @Produce		json
// @Param			input	body		dto.Staff	true	"Staff"
// @Success		204		{object}	dto.Staff
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/staff [put]
func (h *Handlers) staffUpdate(c *gin.Context) {
	admin := c.MustGet("Admin").(dto.Admin)
	var staff apimodels.StaffRequest
	if err := c.BindJSON(&staff); err != nil {
		log.Info(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, apimodels.Message{Message: err.Error()})
	}
	err := h.service.StaffUpdateService(convertors.StaffRequestToDTO(&staff), admin)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
	}
	c.JSON(http.StatusNoContent, nil)
}

// @Summary		StaffDelete
// @Security		JWTAuth
// @Tags			Staff
// @Description	deletes an employee
// @ID				staff-delete
// @Accept			json
// @Produce		json
// @Success		204		{object}	dto.Staff
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/staff/{id} [delete]
// @Param			id	path	string	true	"Staff ID"
func (h *Handlers) staffDelete(c *gin.Context) {
	id := c.Param("id")
	admin := c.MustGet("Admin").(dto.Admin)
	err := h.service.StaffDeleteService(uuid.Parse(id), admin)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
	}
	c.JSON(http.StatusNoContent, nil)
}
