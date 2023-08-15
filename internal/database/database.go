package database

import (
	"context"

	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/pborman/uuid"
	"github.com/uptrace/bun"
)

type AdminDatabases interface {
	CreateOne(ctx context.Context, user models.Admin) (models.Admin, error)
	GetByID(ctx context.Context, id uuid.UUID) (models.Admin, error)
	GetByEmail(ctx context.Context, email string) (models.Admin, error)
	GetPasswordByID(ctx context.Context, id uuid.UUID) (string, error)
	UpdateOne(ctx context.Context, user models.Admin) error
	DeleteOne(ctx context.Context, id uuid.UUID) error
}

type RestaurantDatabases interface {
	CreateOne(ctx context.Context, restaurant models.Restaurant) (models.Restaurant, error)
	GetByID(ctx context.Context, id uuid.UUID) (models.Restaurant, error)
	UpdateOne(ctx context.Context, restaurant models.Restaurant) error
	DeleteOne(ctx context.Context, id uuid.UUID) error
}

type TableDatabases interface {
	CreateOne(ctx context.Context, table models.Table) (models.Table, error)
	GetByID(ctx context.Context, id uuid.UUID) (models.Table, error)
	GetAllInRestaurant(ctx context.Context, id uuid.UUID) ([]models.Table, error)
	UpdateOne(ctx context.Context, table models.Table) error
	DeleteOne(ctx context.Context, id uuid.UUID) error
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

type TableDatabase struct {
	db *bun.DB
}

func NewTableDatabase(db *bun.DB) TableDatabases {
	return &TableDatabase{db: db}
}

type Database struct {
	Admin AdminDatabases
	Rest  RestaurantDatabases
	Table TableDatabases
}

func NewDatabase(db *bun.DB) *Database {
	return &Database{
		Admin: NewAdminDatabase(db),
		Rest:  NewRestaurantDatabase(db),
		Table: NewTableDatabase(db),
	}
}
