package database

import (
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/pborman/uuid"
	"github.com/uptrace/bun"
)

type AdminDatabases interface {
	AdminCreate(user models.Admin) (models.Admin, error)
	AdminGetByID(id uuid.UUID) (models.Admin, error)
	AdminGetByEmail(email string) (models.Admin, error)
	AdminGetPasswordByID(id uuid.UUID) (string, error)
	AdminUpdate(user models.Admin) error
	AdminDelete(id uuid.UUID) error
}

type RestaurantDatabases interface {
	RestaurantCreate(restaurant models.Restaurant) (models.Restaurant, error)
	RestaurantGetByID(id uuid.UUID) (models.Restaurant, error)
	RestaurantGetByAdminsID(id uuid.UUID) (models.Restaurant, error)
	RestaurantUpdate(restaurant models.Restaurant) error
	RestaurantDelete(id uuid.UUID) error
}

type AdminDatabase struct {
	db *bun.DB
}

func NewAdminDatabase(db *bun.DB) AdminDatabases {
	return &AdminDatabase{db: db}
}

type RestDatabase struct {
	db *bun.DB
}

func NewRestaurantDatabase(db *bun.DB) RestaurantDatabases {
	return &RestDatabase{db: db}
}

type Database struct {
	Admin AdminDatabases
	Rest  RestaurantDatabases
}

func NewDatabase(db *bun.DB) *Database {
	return &Database{
		Admin: NewAdminDatabase(db),
		Rest:  NewRestaurantDatabase(db),
	}
}
