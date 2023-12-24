package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

type TableRequest struct {
	Number       int          `json:"number" binding:"required"`
	Placement    string       `json:"placement" binding:"required"`
	MaxPeople    int          `json:"max_people" binding:"required"`
	IsReserved   bool         `json:"is_reserved" binding:"required"`
	IsOccupied   bool         `json:"is_occupied" binding:"required"`
}

type TableResponse struct {
	ID           string       	 `json:"id"`
	Number       int          	 `json:"number"`
	Placement    string       	 `json:"placement"`
	MaxPeople    int             `json:"max_people"`
	IsReserved   bool            `json:"is_reserved"`
	IsOccupied   bool            `json:"is_occupied"`
	RestaurantID string          `json:"restaurant_id"`
	Reservation  *ReservResponse `json:"reservation"`
}

//	@Summary		TableCreate
//	@Security		JWTAuth
//	@Tags			Table
//	@Description	creates a new table
//	@ID				table-create
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dto.Table	true	"Table"
//	@Success		201		{object}	dto.Table
//	@Failure		default	{object}	dto.Message
//	@Router			/table [post]
func (h *Handlers) tableCreate(c *gin.Context){
	admin := c.MustGet("Admin")
	var t dto.Table
	if err := c.ShouldBindJSON(&t); err != nil{
		c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
		log.Info("handlers.tableCreate: ", err)
		return
	}
	t, err := h.service.TableCreateService(t, admin.(dto.Admin))
	if err != nil{
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		log.Info("handlers.tableCreate: ", err)
		return
	}
	log.Info("tablec created")
	c.JSON(http.StatusCreated, t)
}

//	@Summary		TableGetByID
//	@Security		JWTAuth
//	@Tags			Table
//	@Description	returns a table by id
//	@ID				table-get-by-id
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	dto.Table
//	@Failure		default	{object}	dto.Message
//	@Router			/table/{id} [get]
//	@Param			id	path	string	true	"Table ID"
func (h *Handlers) tableGetByID(c *gin.Context){
	id := uuid.Parse(c.Param("id"))
	t, err := h.service.TableGetByIDService(id)
	if err != nil{
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		log.Info("handlers.tableGetByID: ", err)
		return
	}
	log.Info("Got table")
	c.JSON(http.StatusOK, t)
}

//	@Summary		TableGetAllInRestaurant
//	@Security		JWTAuth
//	@Tags			Table
//	@Description	returns all tables in a restaurant
//	@ID				table-get-all-in-restaurant
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	dto.Table
//	@Failure		default	{object}	dto.Message
//	@Router			/table/all/{id} [get]
//	@Param			id	path	string	true	"Table ID"
func (h *Handlers) tableGetAllInRestaurant(c *gin.Context){
	id := uuid.Parse(c.Param("id"))
	t, err := h.service.TableGetAllInRestaurantService(id)
	if err != nil{
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		log.Info(err)
		return
	}
	log.Info("tables were gotten")
	c.JSON(http.StatusOK, t)
}

//	@Summary		TableUpdate
//	@Security		JWTAuth
//	@Tags			Table
//	@Description	updates a table
//	@ID				table update
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dto.Table	true	"Table"
//	@Success		204		{object}	nil
//	@Failure		default	{object}	dto.Message
//	@Router			/table [put]
func (h *Handlers) tableUpdate(c *gin.Context){
	admin := c.MustGet("Admin").(dto.Admin)
	var t dto.Table
	if err := c.ShouldBindJSON(&t); err != nil{
		c.JSON(http.StatusBadRequest, dto.Message{Message: err.Error()})
		log.Info(err)
		return
	}
	err := h.service.TableUpdateService(t, admin)
	if err != nil{
		if err.Error() == models.ErrTableNotFound{
			log.Info(models.ErrTableNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		if err.Error() == models.ErrRestaurantNotFound{
			log.Info(models.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		log.Info(err)
		return
	}
	log.Info("tableUpdate: ", t)
	c.JSON(http.StatusNoContent, nil)
}

//	@Summary		TableDelete
//	@Security		JWTAuth
//	@Tags			Table
//	@Description	deletes a table
//	@ID				table-delete
//	@Accept			json
//	@Produce		json
//	@Success		204		{object}	nil
//	@Failure		default	{object}	dto.Message
//	@Router			/table/{id} [delete]
//	@Param			id	path	string	true	"Table ID"
func (h *Handlers) tableDelete(c *gin.Context){
	id := uuid.Parse(c.Param("id"))
	admin := c.MustGet("Admin").(dto.Admin)
	err := h.service.TableDeleteService(id, admin)
	if err != nil{
		if err.Error() == models.ErrTableNotFound{
			log.Info(models.ErrTableNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		if err.Error() == models.ErrRestaurantNotFound{
			log.Info(models.ErrRestaurantNotFound)
			c.JSON(http.StatusNotFound, dto.Message{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
		log.Info("handlers.tableDelete: ", err)
		return
	}
	log.Info("tableDelete: ", id)
	c.JSON(http.StatusNoContent, nil)
}