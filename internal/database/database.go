package database

import (
	"github.com/mksmstpck/restoracio/pkg/models"
	"github.com/pborman/uuid"
	"github.com/uptrace/bun"
)

type Database struct {
	db *bun.DB
}

func NewDatabase(db *bun.DB) *Database {
	return &Database{db: db}
}

type Databases interface {
	// admin
	AdminCreate(user models.Admin) (models.Admin, error)
	AdminGetByID(id uuid.UUID) (models.Admin, error)
	AdminGetByEmail(email string) (models.Admin, error)
	AdminGetPasswordById(id uuid.UUID) (string, error)
	AdminUpdate(user models.Admin) error
	AdminDelete(id uuid.UUID) error
	// restaurant
	RestaurantCreate(restaurant models.Restaurant) (models.Restaurant, error)
	RestaurantGetByID(id uuid.UUID) (models.Restaurant, error)
	RestaurantGetByAdminsID(id uuid.UUID) (models.Restaurant, error)
	RestaurantUpdate(restaurant models.Restaurant) error
	RestaurantDelete(id uuid.UUID) error
}
