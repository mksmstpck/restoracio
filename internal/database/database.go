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
	DeleteAll(ctx context.Context, id uuid.UUID) error
}

type MenuDatabases interface {
	CreateOne(ctx context.Context, menu models.Menu) (models.Menu, error)
	GetByID(ctx context.Context, id uuid.UUID) (models.Menu, error)
	UpdateOne(ctx context.Context, menu models.Menu) error
	DeleteOne(ctx context.Context, menu models.Menu) error
}

type DishDatabases interface {
	CreateOne(ctx context.Context, dish models.Dish) (models.Dish, error)
	GetByID(ctx context.Context, id uuid.UUID) (models.Dish, error)
	GetAllInMenu(ctx context.Context, id uuid.UUID) ([]models.Dish, error)
	UpdateOne(ctx context.Context, dish models.Dish) error
	DeleteOne(ctx context.Context, id uuid.UUID, menuID uuid.UUID) error
	DeleteAll(ctx context.Context, menuID uuid.UUID) error
}

type StaffDatabases interface {
	CreateOne(ctx context.Context, staff models.Staff) (models.Staff, error)
	GetByID(ctx context.Context, id uuid.UUID, restaurantID uuid.UUID) (models.Staff, error)
	GetAllInRestaurant(ctx context.Context, id uuid.UUID) ([]models.Staff, error)
	UpdateOne(ctx context.Context, staff models.Staff) error
	DeleteOne(ctx context.Context, id uuid.UUID, restaurantID uuid.UUID) error
	DeleteAll(ctx context.Context, restaurantID uuid.UUID) error
}

type ReservationDatabases interface {
	CreateOne(ctx context.Context, reserv models.ReservDB) (models.ReservDB, error)
	GetByID(ctx context.Context, id uuid.UUID, restaurantID uuid.UUID) (models.ReservDB, error)
	GetAllInRestaurant(ctx context.Context, id uuid.UUID) ([]models.ReservDB, error)
	UpdateOne(ctx context.Context, reserv models.ReservDB) error
	DeleteOne(ctx context.Context, id uuid.UUID, restaurantID uuid.UUID) error
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

type MenuDatabase struct {
	db *bun.DB
}

func NewMenuDatabase(db *bun.DB) MenuDatabases {
	return &MenuDatabase{db: db}
}

type DishDatabase struct {
	db *bun.DB
}

func NewDishDatabase(db *bun.DB) DishDatabases {
	return &DishDatabase{db: db}
}

type StaffDatabase struct {
	db *bun.DB
}

func NewStaffDatabase(db *bun.DB) StaffDatabases {
	return &StaffDatabase{db: db}
}

type ReservDB struct {
	db *bun.DB
}

func NewReservationDatabase(db *bun.DB) ReservationDatabases {
	return &ReservDB{db: db}
}

type Database struct {
	Admin AdminDatabases
	Rest  RestaurantDatabases
	Table TableDatabases
	Menu  MenuDatabases
	Dish  DishDatabases
	Staff StaffDatabases
	Reserv ReservationDatabases
}

func NewDatabase(db *bun.DB) *Database {
	return &Database{
		Admin: NewAdminDatabase(db),
		Rest:  NewRestaurantDatabase(db),
		Table: NewTableDatabase(db),
		Menu:  NewMenuDatabase(db),
		Dish:  NewDishDatabase(db),
		Staff: NewStaffDatabase(db),
		Reserv: NewReservationDatabase(db),
	}
}
