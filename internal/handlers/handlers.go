package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mksmstpck/restoracio/internal/services"
)

type Handlers struct {
	gin           *gin.Engine
	service       services.Servicer
	accessSecret  []byte
	refreshSecret []byte
	accessExp     time.Duration
	refreshExp    time.Duration
}

func NewHandlers(gin *gin.Engine,
	service services.Servicer,
	accessSecret []byte,
	refreshSecret []byte,
	accessExp time.Duration,
	refreshExp time.Duration) *Handlers {
	return &Handlers{
		gin:           gin,
		service:       service,
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
		accessExp:     accessExp,
		refreshExp:    refreshExp,
	}
}

func (h *Handlers) HandleAll() {
	// groups
	admin := h.gin.Group("/admin")
	auth := h.gin.Group("/auth")
	rest := h.gin.Group("/restaurant")
	table := h.gin.Group("/table")

	// middleware
	rest.Use(h.DeserializeUser())

	// admin
	admin.POST("/create", h.adminCreate)
	admin.GET("/get-by-id/:id", h.adminGetByID)
	admin.GET("/get-by-email/:email", h.DeserializeUser(), h.adminGetByEmail)
	admin.GET("/get-me", h.DeserializeUser(), h.adminGetMe)
	admin.PUT("/update", h.DeserializeUser(), h.adminUpdate)
	admin.DELETE("/delete", h.DeserializeUser(), h.adminDelete)

	// auth
	auth.POST("/login", h.login)
	auth.POST("/refresh", h.refresh)

	// restorant
	rest.POST("/create", h.restaurantCreate)
	rest.GET("/get-by-id/:id", h.restaurantGetByID)
	rest.GET("/get-mine", h.restaurantGetMine)
	rest.PUT("/update", h.restaurantUpdate)
	rest.DELETE("/delete", h.restaurantDelete)

	// table
	table.POST("/", h.tableCreate)
	table.GET("/:id", h.tableGetByID)
	table.PUT("/", h.tableUpdate)
	table.DELETE("/", h.tableDelete)
}
