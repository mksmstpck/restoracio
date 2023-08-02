package services

import (
	"github.com/mksmstpck/restoracio/internal/database"
	"github.com/mksmstpck/restoracio/pkg/models"
	"github.com/patrickmn/go-cache"
	"github.com/pborman/uuid"
)

type Services struct {
	admindb *database.AdminDatabase
	restdb  *database.RestDatabase
	cache   *cache.Cache
}

func NewServices(
	admindb *database.AdminDatabase,
	restdb *database.RestDatabase,
	cache *cache.Cache,
) *Services {
	return &Services{
		admindb: admindb,
		restdb:  restdb,
		cache:   cache,
	}
}

type Servicer interface {
	// admin
	AdminCreateService(admin models.Admin) (models.Admin, error)
	AdminGetByIDService(id uuid.UUID) (models.Admin, error)
	AdminGetByEmailService(email string) (models.Admin, error)
	AdminGetPasswordByIdService(id uuid.UUID) (string, error)
	AdminUpdateService(admin models.Admin) error
	AdminDeleteService(id uuid.UUID) error
	// restaurant
	RestaurantCreateService(rest models.Restaurant, admin models.Admin) (models.Restaurant, error)
	RestaurantGetByIDService(id uuid.UUID) (models.Restaurant, error)
	RestaurantGetByAdminsIDService(id uuid.UUID) (models.Restaurant, error)
	RestaurantUpdateService(rest models.Restaurant) error
	RestaurantDeleteService(id uuid.UUID) error
}
