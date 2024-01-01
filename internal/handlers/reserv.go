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

// @Summary		ReservationCreate
// @Security		JWTAuth
// @Tags			Reservation
// @Description	creates a new reservation
// @ID				reserv-create
// @Accept			json
// @Produce		json
// @Param			input	body		dto.ReservAPI	true	"Reservation"
// @Success		201		{object}	dto.ReservDB
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/reserv [post]
func (h *Handlers) reservCreate(c *gin.Context) {
	var reserv apimodels.ReservRequest
	if err := c.BindJSON(&reserv); err != nil {
		log.Info(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, apimodels.Message{Message: err.Error()})
		return
	}
	err := h.service.ReservCreateService(
		convertors.ReservRequestToDTO(&reserv),
		c.MustGet("Admin").(dto.Admin),
	)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

// @Summary		ReservationGetByID
// @Security		JWTAuth
// @Tags			Reservation
// @Description	returns a reservation by id
// @ID				reserv-get-by-id
// @Accept			json
// @Produce		json
// @Success		200		{object}	dto.ReservDB
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/reserv/{id} [get]
// @Param			id	path	string	true	"Reservation ID"
func (h *Handlers) reservGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	reserv, err := h.service.ReservGetByIDService(id, c.MustGet("Admin").(dto.Admin))
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	log.Info("reservation found")
	c.JSON(http.StatusOK, convertors.ReservDTOToResponse(&reserv))
}

// @Summary		ReservationGetAllInRestaurant
// @Security		JWTAuth
// @Tags			Reservation
// @Description	returns all reservations in a restaurant
// @ID				reserv-arst
// @Accept			json
// @Produce		json
// @Success		200		{object}	[]dto.ReservDB
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/reserv [get]
func (h *Handlers) reservGetAllInRestaurant(c *gin.Context) {
	reserv, err := h.service.ReservGetAllInRestaurantService(c.MustGet("Admin").(dto.Admin))
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	reservRes := make([]*apimodels.ReservResponse, len(reserv))
	for i, r := range reserv {
		reservRes[i] = convertors.ReservDTOToResponse(&r)
	}
	log.Info("reservations were gotten")
	c.JSON(http.StatusOK, reservRes)
}

// @Summary		ReservationUpdate
// @Security		JWTAuth
// @Tags			Reservation
// @Description	updates a reservation
// @ID				reserv-update
// @Accept			json
// @Produce		json
// @Param			input	body		dto.ReservAPI	true	"Reservation"
// @Success		204		{object}	nil
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/reserv [put]
func (h *Handlers) reservUpdate(c *gin.Context) {
	var reserv apimodels.ReservRequest
	if err := c.BindJSON(&reserv); err != nil {
		log.Info(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, apimodels.Message{Message: err.Error()})
		return
	}
	err := h.service.ReservUpdateService(
		convertors.ReservRequestToDTO(&reserv),
		c.MustGet("Admin").(dto.Admin),
	)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	log.Info("reservation updated")
	c.JSON(http.StatusNoContent, nil)
}

// @Summary		ReservationDelete
// @Security		JWTAuth
// @Tags			Reservation
// @Description	deletes a  reservation
// @ID				reserv-delete
// @Accept			json
// @Produce		json
// @Success		204		{object}	nil
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/reserv/{id} [delete]
// @Param			id	path	string	true	"Reservation ID"
func (h *Handlers) reservDelete(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	err := h.service.ReservDeleteService(id, c.MustGet("Admin").(dto.Admin))
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
		return
	}
	log.Info("reservation deleted")
	c.JSON(http.StatusNoContent, nil)
}
