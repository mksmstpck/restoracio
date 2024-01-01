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

// @Summary		TableCreate
// @Security		JWTAuth
// @Tags			Table
// @Description	creates a new table
// @ID				table-create
// @Accept			json
// @Produce		json
// @Param			input	body		dto.Table	true	"Table"
// @Success		201		{object}	dto.Table
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/table [post]
func (h *Handlers) tableCreate(c *gin.Context) {
	admin := c.MustGet("Admin")
	var t apimodels.TableRequest
	if err := c.ShouldBindJSON(&t); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, apimodels.Message{Message: err.Error()})
		log.Info(err)
	}
	err := h.service.TableCreateService(convertors.TableRequestToDTO(&t), admin.(dto.Admin))
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
	}
	log.Info("tablec created")
	c.JSON(http.StatusCreated, nil)
}

// @Summary		TableGetByID
// @Security		JWTAuth
// @Tags			Table
// @Description	returns a table by id
// @ID				table-get-by-id
// @Accept			json
// @Produce		json
// @Success		200		{object}	dto.Table
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/table/{id} [get]
// @Param			id	path	string	true	"Table ID"
func (h *Handlers) tableGetByID(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	t, err := h.service.TableGetByIDService(id)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
	}
	log.Info("Got table")
	c.JSON(http.StatusOK, convertors.TableDTOToResponse(&t))
}

// @Summary		TableGetAllInRestaurant
// @Security		JWTAuth
// @Tags			Table
// @Description	returns all tables in a restaurant
// @ID				table-get-all-in-restaurant
// @Accept			json
// @Produce		json
// @Success		200		{object}	dto.Table
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/table/all/{id} [get]
// @Param			id	path	string	true	"Table ID"
func (h *Handlers) tableGetAllInRestaurant(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	t, err := h.service.TableGetAllInRestaurantService(id)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
	}
	log.Info("tables were gotten")
	tables := make([]*apimodels.TableResponse, len(t))
	for i, table := range t {
		tables[i] = convertors.TableDTOToResponse(&table)
	}
	c.JSON(http.StatusOK, tables)
}

// @Summary		TableUpdate
// @Security		JWTAuth
// @Tags			Table
// @Description	updates a table
// @ID				table update
// @Accept			json
// @Produce		json
// @Param			input	body		dto.Table	true	"Table"
// @Success		204		{object}	nil
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/table [put]
func (h *Handlers) tableUpdate(c *gin.Context) {
	admin := c.MustGet("Admin").(dto.Admin)
	var t apimodels.TableRequest
	if err := c.ShouldBindJSON(&t); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, apimodels.Message{Message: err.Error()})
		log.Info(err)
	}
	err := h.service.TableUpdateService(convertors.TableRequestToDTO(&t), admin)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
	}
	log.Info("tableUpdate: ", t)
	c.JSON(http.StatusNoContent, nil)
}

// @Summary		TableDelete
// @Security		JWTAuth
// @Tags			Table
// @Description	deletes a table
// @ID				table-delete
// @Accept			json
// @Produce		json
// @Success		204		{object}	nil
// @Failure		default	{object}apimodels.apimodels.Message
// @Router			/table/{id} [delete]
// @Param			id	path	string	true	"Table ID"
func (h *Handlers) tableDelete(c *gin.Context) {
	id := uuid.Parse(c.Param("id"))
	admin := c.MustGet("Admin").(dto.Admin)
	err := h.service.TableDeleteService(id, admin)
	if err != nil {
		log.Info(err)
		status := checkStatus(err.Error())
		c.AbortWithStatusJSON(status, apimodels.Message{Message: err.Error()})
	}
	log.Info("tableDelete: ", id)
	c.JSON(http.StatusNoContent, nil)
}
