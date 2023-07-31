package services

import (
	"github.com/mksmstpck/restoracio/internal/database"
	"github.com/mksmstpck/restoracio/pkg/models"
	"github.com/pborman/uuid"
)

type Services struct {
	admindb *database.AdminDatabase
	restdb  *database.RestDatabase
}

func NewServices(admindb *database.AdminDatabase, restdb *database.RestDatabase) *Services {
	return &Services{
		admindb: admindb,
		restdb:  restdb,
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
	RestaurantCreateService(rest models.Restaurant, adminID uuid.UUID) (models.Restaurant, error)
	RestaurantGetByIDService(id uuid.UUID) (models.Restaurant, error)
	RestaurantGetByAdminsIDService(id uuid.UUID) (models.Restaurant, error)
	RestaurantUpdateService(rest models.Restaurant) error
	RestaurantDeleteService(id uuid.UUID) error
}
