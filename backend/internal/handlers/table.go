package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/backend/internal/models"
	"github.com/mksmstpck/restoracio/backend/utils"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

//	@Summary		TableCreate
//	@Security		JWTAuth
//	@Tags			Table
//	@Description	creates a new table
//	@ID				table-create
//	@Accept			json
//	@Produce		json
//	@Param			input	body		models.Table	true	"Table"
//	@Success		201		{object}	models.Table
//	@Failure		default	{object}	models.Message
//	@Router			/table [post]
func (h *Handlers) tableCreate(c *gin.Context){
	admin := c.MustGet("Admin")
	var t models.Table
	if err := c.ShouldBindJSON(&t); err != nil{
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Info("handlers.tableCreate: ", err)
		return
	}
	t, err := h.service.TableCreateService(t, admin.(models.Admin))
	if err != nil{
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
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
//	@Success		200		{object}	models.Table
//	@Failure		default	{object}	models.Message
//	@Router			/table/{id} [get]
//	@Param			id	path	string	true	"Table ID"
func (h *Handlers) tableGetByID(c *gin.Context){
	id := uuid.Parse(c.Param("id"))
	t, err := h.service.TableGetByIDService(id)
	if err != nil{
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		log.Info("handlers.tableGetByID: ", err)
		return
	}
	log.Info("tableGetByID: ", t)
	c.JSON(http.StatusOK, t)
}

//	@Summary		TableGetAllInRestaurant
//	@Security		JWTAuth
//	@Tags			Table
//	@Description	returns all tables in a restaurant
//	@ID				table-get-all-in-restaurant
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.Table
//	@Failure		default	{object}	models.Message
//	@Router			/table/all/{id} [get]
//	@Param			id	path	string	true	"Table ID"
func (h *Handlers) tableGetAllInRestaurant(c *gin.Context){
	id := uuid.Parse(c.Param("id"))
	t, err := h.service.TableGetAllInRestaurantService(id)
	if err != nil{
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
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
//	@Param			input	body		models.Table	true	"Table"
//	@Success		204		{object}	nil
//	@Failure		default	{object}	models.Message
//	@Router			/table [put]
func (h *Handlers) tableUpdate(c *gin.Context){
	admin := c.MustGet("Admin").(models.Admin)
	var t models.Table
	if err := c.ShouldBindJSON(&t); err != nil{
		c.JSON(http.StatusBadRequest, models.Message{Message: err.Error()})
		log.Info(err)
		return
	}
	err := h.service.TableUpdateService(t, admin)
	if err != nil{
		if err.Error() == utils.ErrTableNotFound{
			log.Info("table not found")
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		if err.Error() == utils.ErrRestaurantNotFound{
			log.Info("restaurant not found")
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
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
//	@Failure		default	{object}	models.Message
//	@Router			/table/{id} [delete]
//	@Param			id	path	string	true	"Table ID"
func (h *Handlers) tableDelete(c *gin.Context){
	id := uuid.Parse(c.Param("id"))
	admin := c.MustGet("Admin").(models.Admin)
	err := h.service.TableDeleteService(id, admin)
	if err != nil{
		if err.Error() == utils.ErrTableNotFound{
			log.Info("table not found")
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		if err.Error() == utils.ErrRestaurantNotFound{
			log.Info("restaurant not found")
			c.JSON(http.StatusNotFound, models.Message{Message: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, models.Message{Message: err.Error()})
		log.Info("handlers.tableDelete: ", err)
		return
	}
	log.Info("tableDelete: ", id)
	c.JSON(http.StatusNoContent, nil)
}