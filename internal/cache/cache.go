package cache

import (
	"time"

	"github.com/pborman/uuid"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
	exp time.Duration
}

type Cacher interface {
	Get(id uuid.UUID) (any, error)
	Set(key uuid.UUID, value interface{}) error
	Delete(key uuid.UUID) error
}

func NewCache(client *redis.Client, exp time.Duration) Cacher {
	return &Cache{
		client: client,
		exp: exp,
	}
}