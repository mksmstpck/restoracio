package services

import (
	"github.com/mksmstpck/restoracio/internal/database"
	"github.com/mksmstpck/restoracio/pkg/models"
	"github.com/pborman/uuid"
)

type Services struct {
	db database.Databases
}

func NewServices(db *database.Database) *Services {
	return &Services{
		db: db,
	}
}

type Servicer interface {
	// admin
	AdminCreateService(admin models.Admin) (models.Admin, error)
	// restaurant
	RestaurantCreateService(rest models.Restaurant, adminID uuid.UUID) (models.Restaurant, error)
	RestaurantGetByIDService(id uuid.UUID) (models.Restaurant, error)
	RestaurantGetByAdminsIDService(id uuid.UUID) (models.Restaurant, error)
	RestaurantUpdateService(rest models.Restaurant) error
	RestaurantDeleteService(id uuid.UUID) error
}
