package database

import (
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/pborman/uuid"
	"github.com/uptrace/bun"
)

type AdminDatabases interface {
	CreateOne(user models.Admin) (models.Admin, error)
	GetByID(id uuid.UUID) (models.Admin, error)
	GetByEmail(email string) (models.Admin, error)
	GetPasswordByID(id uuid.UUID) (string, error)
	UpdateOne(user models.Admin) error
	DeleteOne(id uuid.UUID) error
}

type RestaurantDatabases interface {
	CreateOne(restaurant models.Restaurant) (models.Restaurant, error)
	GetByID(id uuid.UUID) (models.Restaurant, error)
	GetByAdminsID(id uuid.UUID) (models.Restaurant, error)
	UpdateOne(restaurant models.Restaurant) error
	DeleteOne(id uuid.UUID) error
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
