package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/services"
)

type Handlers struct {
	gin            *gin.Engine
	service        services.Servicer
	access_secret  []byte
	refresh_secret []byte
	access_exp     time.Duration
	refresh_exp    time.Duration
}

func NewHandlers(gin *gin.Engine,
	service *services.Services,
	access_secret []byte,
	refresh_secret []byte,
	access_exp time.Duration,
	refresh_exp time.Duration) *Handlers {
	return &Handlers{
		gin:            gin,
		service:        service,
		access_secret:  access_secret,
		refresh_secret: refresh_secret,
		access_exp:     access_exp,
		refresh_exp:    refresh_exp,
	}
}

func (h *Handlers) HandleAll() {
	// groups
	admin := h.gin.Group("/admin")
	auth := h.gin.Group("/auth")
	rest := h.gin.Group("/restaurant")

	// middleware
	rest.Use(h.DeserializeUser())

	//admin
	admin.POST("/create", h.adminCreate)
	admin.GET("/get-by-id/:id", h.DeserializeUser(), h.adminGetByID)
	admin.GET("/get-by-email/:email", h.DeserializeUser(), h.adminGetByEmail)
	admin.GET("/get-me", h.DeserializeUser(), h.adminGetMe)
	admin.POST("/update", h.DeserializeUser(), h.adminUpdate)
	admin.DELETE("/delete", h.DeserializeUser(), h.adminDelete)

	//auth
	auth.POST("/login", h.login)
	auth.POST("/refresh", h.login)

	//restorant
	rest.POST("/create", h.restaurantCreate)
	rest.GET("/get-by-id/:id", h.restaurantCreate)
	rest.GET("/get-by-admins-id/:id", h.restaurantCreate)
	rest.GET("/get-by-email/:email", h.restaurantCreate)
	rest.POST("/update", h.restaurantUpdate)
	rest.DELETE("/delete", h.restaurantDelete)
}
