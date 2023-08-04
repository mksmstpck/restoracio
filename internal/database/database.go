package database

import (
	"github.com/mksmstpck/restoracio/pkg/models"
	"github.com/pborman/uuid"
	"github.com/uptrace/bun"
)

type AdminDatabase struct {
	db *bun.DB
}

func NewAdminDatabase(db *bun.DB) *AdminDatabase {
	return &AdminDatabase{db: db}
}

type RestDatabase struct {
	db *bun.DB
}

func NewDRatabase(db *bun.DB) *RestDatabase {
	return &RestDatabase{db: db}
}

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
