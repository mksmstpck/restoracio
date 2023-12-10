package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

//	@Summary		StaffCreate
//	@Security		JWTAuth
//	@Tags			Staff
//	@Description	creates a new employee
//	@ID				staff-create
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dto.Staff	true	"Staff"
//	@Success		201		{object}	dto.Staff
//	@Failure		default	{object}	dto.Message
//	@Router			/staff [post]
func (h *Handlers) staffCreate(c *gin.Context) {
	staff := dto.Staff{}
	if err := c.BindJSON(&staff); err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	staff, err := h.service.StaffCreateService(staff, c.MustGet("Admin").(dto.Admin))
	if err != nil {
		if err.Error() == dto.ErrRestaurantNotFound {
			log.Info(dto.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		log.Info(err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, staff)
}

//	@Summary		StaffGetByID
//	@Security		JWTAuth
//	@Tags			Staff
//	@Description	returns an employee by id
//	@ID				staff-get-by-id
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	dto.Staff
//	@Failure		default	{object}	dto.Message
//	@Router			/staff/{id} [get]
//	@Param			id	path	string	true	"Staff ID"
func (h *Handlers) staffGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	admin := c.MustGet("Admin").(dto.Admin)
	staff, err := h.service.StaffGetByIDService(id, admin)
	if err != nil {
		if err.Error() == dto.ErrStaffNotFound {
			log.Info(dto.ErrStaffNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		if err.Error() == dto.ErrRestaurantNotFound {
			log.Info(dto.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		log.Info(err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, staff)
}

//	@Summary		StaffGetAllInRestaurant
//	@Security		JWTAuth
//	@Tags			Staff
//	@Description	returns all staff in a restaurant
//	@ID				staff-get-all
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	dto.Staff
//	@Failure		default	{object}	dto.Message
//	@Router			/staff [get]
func (h *Handlers) staffGetInRestaurant(c *gin.Context) {
	admin := c.MustGet("Admin").(dto.Admin)
	staff, err := h.service.StaffGetAllInRestaurantService(admin)
	if err != nil {
		if err.Error() == dto.ErrRestaurantNotFound {
			log.Info(dto.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		if err.Error() == dto.ErrStaffNotFound {
			log.Info(dto.ErrStaffNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		log.Info(err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, staff)
}

//	@Summary		StaffUpdate
//	@Security		JWTAuth
//	@Tags			Staff
//	@Description	updates an employee
//	@ID				staff-update
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dto.Staff	true	"Staff"
//	@Success		204		{object}	dto.Staff
//	@Failure		default	{object}	dto.Message
//	@Router			/staff [put]
func (h *Handlers) staffUpdate(c *gin.Context) {
	admin := c.MustGet("Admin").(dto.Admin)
	var staff dto.Staff 
	if err := c.BindJSON(&staff); err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	err := h.service.StaffUpdateService(staff, admin)
	if err != nil {
		if err.Error() == dto.ErrRestaurantNotFound {
			log.Info(dto.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		if err.Error() == dto.ErrStaffNotFound {
			log.Info(dto.ErrStaffNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		log.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

//	@Summary		StaffDelete
//	@Security		JWTAuth
//	@Tags			Staff
//	@Description	deletes an employee
//	@ID				staff-delete
//	@Accept			json
//	@Produce		json
//	@Success		204		{object}	dto.Staff
//	@Failure		default	{object}	dto.Message
//	@Router			/staff/{id} [delete]
//	@Param			id	path	string	true	"Staff ID"
func (h *Handlers) staffDelete(c *gin.Context) {
	id := c.Param("id")
	admin := c.MustGet("Admin").(dto.Admin)
	err := h.service.StaffDeleteService(uuid.Parse(id), admin)
	if err != nil {
		if err.Error() == dto.ErrRestaurantNotFound {
			log.Info(dto.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		log.Info(err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}