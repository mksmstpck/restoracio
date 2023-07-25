package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/database"
	"github.com/mksmstpck/restoracio/internal/services"
)

type Handlers struct {
	gin            *gin.Engine
	service        services.Servicer
	db             database.Databases
	access_secret  []byte
	refresh_secret []byte
	access_exp     time.Duration
	refresh_exp    time.Duration
}

func NewHandlers(gin *gin.Engine,
	service *services.Services,
	db database.Databases,
	access_secret []byte,
	refresh_secret []byte,
	access_exp time.Duration,
	refresh_exp time.Duration) *Handlers {
	return &Handlers{
		gin:            gin,
		service:        service,
		db:             db,
		access_secret:  access_secret,
		refresh_secret: refresh_secret,
		access_exp:     access_exp,
		refresh_exp:    refresh_exp,
	}
}

func (h *Handlers) HandleAll() {
	admin := h.gin.Group("/admin")
	admin.POST("/create", h.AdminCreate)
	admin.GET("/get/:id", h.AdminGetByID)
	admin.GET("/getByEmail/:email", h.AdminGetByEmail)
	admin.POST("/update", h.AdminUpdate)
	admin.DELETE("/delete", h.AdminDelete)
}
