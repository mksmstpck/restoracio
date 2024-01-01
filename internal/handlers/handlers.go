package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mksmstpck/restoracio/docs"
	apimodels "github.com/mksmstpck/restoracio/internal/APImodels"
	"github.com/mksmstpck/restoracio/internal/config"
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/internal/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handlers struct {
	service       services.Servicer
	accessSecret  []byte
	refreshSecret []byte
	accessExp     time.Duration
	refreshExp    time.Duration
}

func NewHandlers(
	service services.Servicer,
	accessSecret []byte,
	refreshSecret []byte,
	accessExp time.Duration,
	refreshExp time.Duration) *Handlers {
	return &Handlers{
		service:       service,
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
		accessExp:     accessExp,
		refreshExp:    refreshExp,
	}
}

func (h *Handlers) HandleAll() {
	r := gin.Default()
	// groups
	admin := r.Group("/admin")
	auth := r.Group("/auth")
	rest := r.Group("/restaurant")
	table := r.Group("/table")
	menu := r.Group("/menu")
	dish := r.Group("/dish")
	staff := r.Group("/staff")
	reserv := r.Group("/reserv")

	// middleware
	r.Use(h.corsMiddleware)
	rest.Use(h.authMiddleware)
	staff.Use(h.authMiddleware)
	reserv.Use(h.authMiddleware)

	// admin
	admin.POST("/", h.adminCreate)
	admin.GET("/id/:id", h.adminGetByID)
	admin.GET("/email/:email", h.authMiddleware, h.adminGetByEmail)
	admin.GET("/me", h.authMiddleware, h.adminGetMe)
	admin.PUT("/", h.authMiddleware, h.adminUpdate)
	admin.DELETE("/", h.authMiddleware, h.adminDelete)

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
	table.POST("/", h.authMiddleware, h.tableCreate)
	table.GET("/:id", h.tableGetByID)
	table.GET("/all/:id", h.tableGetAllInRestaurant)
	table.PUT("/", h.authMiddleware, h.tableUpdate)
	table.DELETE("/:id", h.authMiddleware, h.tableDelete)

	// menu
	menu.POST("/", h.authMiddleware, h.menuCreate)
	menu.GET("/:id", h.menuGetByID)
	menu.GET("/qr/:id", h.menuGetWithQrcode)
	menu.PUT("/", h.authMiddleware, h.menuUpdate)
	menu.DELETE("/", h.authMiddleware, h.menuDelete)

	// dish
	dish.POST("/", h.authMiddleware, h.dishCreate)
	dish.GET("/:id", h.dishGetByID)
	dish.GET("/all/:id", h.dishGetAllInMenu)
	dish.PUT("/", h.authMiddleware, h.dishUpdate)
	dish.DELETE("/:id", h.authMiddleware, h.dishDelete)

	// staff
	staff.POST("/", h.staffCreate)
	staff.GET("/:id", h.staffGetByID)
	staff.GET("/", h.staffGetInRestaurant)
	staff.PUT("/", h.staffUpdate)
	staff.DELETE("/:id", h.staffDelete)

	// reservations
	reserv.POST("/", h.reservCreate)
	reserv.GET("/:id", h.reservGetByID)
	reserv.GET("/", h.reservGetAllInRestaurant)
	reserv.PUT("/", h.reservUpdate)
	reserv.DELETE("/:id", h.reservDelete)

	// swagger
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, apimodels.Message{Message: "method not allowed"})
	})
	r.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, apimodels.Message{Message: "route not found"})
	})

	if err := r.Run(config.NewConfig().GinUrl); err != nil {
		log.Fatal(err)
	}
}

func checkStatus(err string) int {
	switch err {
	case models.ErrAdminNotFound:
		return http.StatusNotFound
	case models.ErrAdminAlreadyExists:
		return http.StatusConflict
	case models.ErrRestaurantNotFound:
		return http.StatusNotFound
	case models.ErrMenuAlreadyExists:
		return http.StatusConflict
	case models.ErrDishNotFound:
		return http.StatusNotFound
	case models.ErrDishAlreadyExists:
		return http.StatusConflict
	case models.ErrStaffNotFound:
		return http.StatusNotFound
	case models.ErrReservationNotFound:
		return http.StatusNotFound
	case models.ErrReservationAlreadyExists:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
