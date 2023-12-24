package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

type ReservRequest struct {
	ID        string 	 `json:"id"`
	Year      int    	 `json:"year" binding:"required"`
	Month     int     	 `json:"month" binding:"required"`
	Day       int     	 `json:"day" binding:"required"`
	Hour      int    	 `json:"hour" binding:"required"`
	Minute    int    	 `json:"minute" binding:"required"`
	Second    int    	 `json:"second"`
	TableID   string 	 `json:"table_id" binding:"required"`
	ReserverName string  `json:"reserver_name" binding:"required"`
	ReserverPhone string `json:"reserver_phone" binding:"required"`
}

type ReservResponse struct {
	ID              string    `json:"id"`
	ReservationTime time.Time `json:"reservation_time"`
	ReserverName    string    `json:"reserver_name"`
	ReserverPhone   string    `json:"reserver_phone"`
	TableID         string    `json:"table_id"`
	RestaurantID    string    `bun:"restaurant_id"`
}

//	@Summary		ReservationCreate
//	@Security		JWTAuth
//	@Tags			Reservation
//	@Description	creates a new reservation
//	@ID				reserv-create
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dto.ReservAPI	true	"Reservation"
//	@Success		201		{object}	dto.ReservDB
//	@Failure		default	{object}	dto.Message
//	@Router			/reserv [post]
func (h *Handlers) reservCreate(c *gin.Context) {
	var reserv dto.ReservAPI
	if err := c.BindJSON(&reserv); err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	reservDB, err := h.service.ReservCreateService(reserv, c.MustGet("Admin").(dto.Admin))
	if err != nil {
		if err.Error() == models.ErrRestaurantNotFound {
			log.Info(models.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		log.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, reservDB)
}

//	@Summary		ReservationGetByID
//	@Security		JWTAuth
//	@Tags			Reservation
//	@Description	returns a reservation by id
//	@ID				reserv-get-by-id
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	dto.ReservDB
//	@Failure		default	{object}	dto.Message
//	@Router			/reserv/{id} [get]
//	@Param			id	path	string	true	"Reservation ID"
func (h *Handlers) reservGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	reservDB, err := h.service.ReservGetByIDService(id, c.MustGet("Admin").(dto.Admin))
	if err != nil {
		if err.Error() == models.ErrReservationNotFound {
			log.Info(models.ErrReservationNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		if err.Error() == models.ErrRestaurantNotFound {
			log.Info(models.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		log.Error(err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	log.Info("reservation found")
	c.JSON(http.StatusOK, reservDB)
}

//	@Summary		ReservationGetAllInRestaurant
//	@Security		JWTAuth
//	@Tags			Reservation
//	@Description	returns all reservations in a restaurant
//	@ID				reserv-arst
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	[]dto.ReservDB
//	@Failure		default	{object}	dto.Message
//	@Router			/reserv [get]
func (h *Handlers) reservGetAllInRestaurant(c *gin.Context) {
	reservDB, err := h.service.ReservGetAllInRestaurantService(c.MustGet("Admin").(dto.Admin))
	if err != nil {
		if err.Error() == models.ErrRestaurantNotFound {
			log.Info(models.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		log.Info(err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	log.Info("reservations were gotten")
	c.JSON(http.StatusOK, reservDB)
}

//	@Summary		ReservationUpdate
//	@Security		JWTAuth
//	@Tags			Reservation
//	@Description	updates a reservation
//	@ID				reserv-update
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dto.ReservAPI	true	"Reservation"
//	@Success		204		{object}	nil
//	@Failure		default	{object}	dto.Message
//	@Router			/reserv [put]
func (h *Handlers) reservUpdate(c *gin.Context) {
	var reserv dto.ReservAPI
	if err := c.BindJSON(&reserv); err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	err := h.service.ReservUpdateService(reserv, c.MustGet("Admin").(dto.Admin))
	if err != nil {
		if err.Error() == models.ErrReservationNotFound {
			log.Info(models.ErrReservationNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		if err.Error() == models.ErrRestaurantNotFound {
			log.Info(models.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		if err.Error() == models.ErrTableNotFound {
			log.Info(models.ErrTableNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		log.Info(err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	log.Info("reservation updated")
	c.JSON(http.StatusNoContent, nil)
}


//	@Summary		ReservationDelete
//	@Security		JWTAuth
//	@Tags			Reservation
//	@Description	deletes a  reservation
//	@ID				reserv-delete
//	@Accept			json
//	@Produce		json
//	@Success		204		{object}	nil
//	@Failure		default	{object}	dto.Message
//	@Router			/reserv/{id} [delete]
//	@Param			id	path	string	true	"Reservation ID"
func (h *Handlers) reservDelete(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	err := h.service.ReservDeleteService(id, c.MustGet("Admin").(dto.Admin))
	if err != nil {
		if err.Error() == models.ErrReservationNotFound {
			log.Info(models.ErrReservationNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		if err.Error() == models.ErrRestaurantNotFound {
			log.Info(models.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		log.Info(err)
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		return
	}
	log.Info("reservation deleted")
	c.JSON(http.StatusNoContent, nil)
}