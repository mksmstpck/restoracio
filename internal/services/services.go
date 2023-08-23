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
	// table
	TableCreateService(table models.Table, admin models.Admin) (models.Table, error)
	TableGetByIDService(id uuid.UUID) (models.Table, error)
	TableGetAllInRestaurantService(id uuid.UUID) ([]models.Table, error)
	TableUpdateService(table models.Table, admin models.Admin) error
	TableDeleteService(id uuid.UUID, admin models.Admin) error
	// menu
	MenuCreateService(menu models.Menu, admin models.Admin) (models.Menu, error)
	MenuGetByIDService(id uuid.UUID) (models.Menu, error)
	MenuUpdateService(menu models.Menu, admin models.Admin) error
	MenuDeleteService(admin models.Admin) error
	// dish
	DishCreateService(dish models.Dish, admin models.Admin) (models.Dish, error)
	DishGetByIDService(id uuid.UUID) (models.Dish, error)
	DishGetAllInMenuService(id uuid.UUID) ([]models.Dish, error)
	DishUpdateService(dish models.Dish, admin models.Admin) error
	DishDeleteService(id uuid.UUID, admin models.Admin) error
	// staff
	StaffCreateService(staff models.Staff, admin models.Admin) (models.Staff, error)
	StaffGetByIDService(id uuid.UUID, admin models.Admin) (models.Staff, error)
	StaffGetAllInRestaurantService(admin models.Admin) ([]models.Staff, error)
	StaffUpdateService(staff models.Staff, admin models.Admin) error
	StaffDeleteService(id uuid.UUID, admin models.Admin) error
}
