package services

import (
	"context"

	"github.com/mksmstpck/restoracio/internal/database"
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/patrickmn/go-cache"
	"github.com/pborman/uuid"
)

type Services struct {
	db      *database.Database
	cache   *cache.Cache
	ctx context.Context
}

func NewServices(
	ctx context.Context,
	db *database.Database,
	cache *cache.Cache,
) Servicer {
	return &Services{
		db:      db,
		cache:   cache,
		ctx:     ctx,
	}
}

type Servicer interface {
	// admin
	AdminCreateService(admin models.Admin) (models.Admin, error)
	AdminGetByIDService(id uuid.UUID) (models.Admin, error)
	AdminGetByEmailService(email string) (models.Admin, error)
	AdminGetPasswordByIdService(id uuid.UUID) (string, error)
	AdminUpdateService(admin models.Admin, adminID uuid.UUID) error
	AdminDeleteService(id uuid.UUID) error
	// restaurant
	RestaurantCreateService(rest models.Restaurant, admin models.Admin) (models.Restaurant, error)
	RestaurantGetByIDService(id uuid.UUID) (models.Restaurant, error)
	RestaurantUpdateService(rest models.Restaurant, restID uuid.UUID) error
	RestaurantDeleteService(rest *models.Restaurant) error
}
