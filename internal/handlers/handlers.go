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
	menu := h.gin.Group("/menu")

	// middleware
	rest.Use(h.DeserializeUser())

	// admin
	admin.POST("/", h.adminCreate)
	admin.GET("/id/:id", h.adminGetByID)
	admin.GET("/email/:email", h.DeserializeUser(), h.adminGetByEmail)
	admin.GET("/me", h.DeserializeUser(), h.adminGetMe)
	admin.PUT("/", h.DeserializeUser(), h.adminUpdate)
	admin.DELETE("/", h.DeserializeUser(), h.adminDelete)

	// auth
	auth.POST("/login", h.login)
	auth.POST("/refresh", h.refresh)

	// restorant
	rest.POST("/", h.restaurantCreate)
	rest.GET("/:id", h.restaurantGetByID)
	rest.GET("/mine", h.restaurantGetMine)
	rest.PUT("/", h.restaurantUpdate)
	rest.DELETE("/", h.restaurantDelete)

	// table
	table.POST("/", h.DeserializeUser(),  h.tableCreate)
	table.GET("/:id", h.tableGetByID)
	table.GET("/all/:id", h.tableGetAllInRestaurant)
	table.PUT("/", h.DeserializeUser(), h.tableUpdate)
	table.DELETE("/:id", h.DeserializeUser(), h.tableDelete)

	// menu
	menu.POST("/", h.DeserializeUser(), h.menuCreate)
	menu.GET("/:id", h.menuGetByID)
	menu.PUT("/", h.DeserializeUser(), h.menuUpdate)
	menu.DELETE("/", h.DeserializeUser(), h.menuDelete)
}
