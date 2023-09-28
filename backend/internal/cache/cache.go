package cache

import (
	"time"

	"github.com/mksmstpck/restoracio/backend/internal/models"
	"github.com/pborman/uuid"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
	exp time.Duration
}

type Cacher interface {
	AdminGet(id uuid.UUID) (models.Admin, error)
	DishGet(id uuid.UUID) (models.Dish, error)
	MenuGet(id uuid.UUID) (models.Menu, error)
	ReservGet(id uuid.UUID) (models.ReservDB, error)
	TableGet(id uuid.UUID) (models.Table, error)
	StaffGet(id uuid.UUID) (models.Staff, error)
	RestaurantGet(id uuid.UUID) (models.Restaurant, error)
	Set(key uuid.UUID, value interface{}) error
	Delete(key uuid.UUID) error
}

func NewCache(client *redis.Client, exp time.Duration) Cacher {
	return &Cache{
		client: client,
		exp: exp,
	}
}