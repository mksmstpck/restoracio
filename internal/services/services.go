package services

import (
	"context"

	"github.com/mksmstpck/restoracio/internal/cache"
	"github.com/mksmstpck/restoracio/internal/database"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/pborman/uuid"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type Services struct {
	db     *database.Database
	cache  cache.Cacher
	bucket *gridfs.Bucket
	ctx    context.Context
}

func NewServices(
	ctx context.Context,
	db *database.Database,
	cache cache.Cacher,
	bucket *gridfs.Bucket,
) Servicer {
	return &Services{
		db:     db,
		cache:  cache,
		bucket: bucket,
		ctx:    ctx,
	}
}

type Servicer interface {
	// admin
	AdminCreateService(dto.Admin) error
	AdminGetByIDService(id uuid.UUID) (dto.Admin, error)
	AdminGetByEmailService(email string) (dto.Admin, error)
	AdminGetWithPasswordByIdService(id uuid.UUID) (dto.Admin, error)
	AdminUpdateService(admin dto.Admin, adminID uuid.UUID) error
	AdminDeleteService(id uuid.UUID) error
	// restaurant
	RestaurantCreateService(rest dto.Restaurant, admin dto.Admin) error
	RestaurantGetByIDService(id uuid.UUID) (dto.Restaurant, error)
	RestaurantUpdateService(rest dto.Restaurant, admin dto.Admin) error
	RestaurantDeleteService(rest *dto.Restaurant) error
	// table
	TableCreateService(table dto.Table, admin dto.Admin) error
	TableGetByIDService(id uuid.UUID) (dto.Table, error)
	TableGetAllInRestaurantService(id uuid.UUID) ([]dto.Table, error)
	TableUpdateService(table dto.Table, admin dto.Admin) error
	TableDeleteService(id uuid.UUID, admin dto.Admin) error
	// menu
	MenuCreateService(menu dto.Menu, admin dto.Admin) error
	MenuGetWithQrcodeService(id uuid.UUID) (dto.Menu, error)
	MenuGetByIDService(id uuid.UUID) (dto.Menu, error)
	MenuUpdateService(menu dto.Menu, admin dto.Admin) error
	MenuDeleteService(admin dto.Admin) error
	// dish
	DishCreateService(dish dto.Dish, admin dto.Admin) error
	DishGetByIDService(id uuid.UUID) (dto.Dish, error)
	DishGetAllInMenuService(id uuid.UUID) ([]dto.Dish, error)
	DishUpdateService(dish dto.Dish, admin dto.Admin) error
	DishDeleteService(id uuid.UUID, admin dto.Admin) error
	DishDeleteAllService(admin dto.Admin) error
	// staff
	StaffCreateService(staff dto.Staff, admin dto.Admin) error
	StaffGetByIDService(id uuid.UUID, admin dto.Admin) (dto.Staff, error)
	StaffGetAllInRestaurantService(admin dto.Admin) ([]dto.Staff, error)
	StaffUpdateService(staff dto.Staff, admin dto.Admin) error
	StaffDeleteService(id uuid.UUID, admin dto.Admin) error
	// reservation
	ReservCreateService(reserv dto.Reserv, admin dto.Admin) error
	ReservGetByIDService(id uuid.UUID, admin dto.Admin) (dto.Reserv, error)
	ReservGetAllInRestaurantService(admin dto.Admin) ([]dto.Reserv, error)
	ReservUpdateService(reserv dto.Reserv, admin dto.Admin) error
	ReservDeleteService(id uuid.UUID, admin dto.Admin) error
}
