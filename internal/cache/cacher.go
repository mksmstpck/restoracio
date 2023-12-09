package cache

import (
	"context"
	"encoding/json"

	"github.com/pborman/uuid"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func (r *Cache) Get(id uuid.UUID) (any, error) {
	var res any
	val, err := r.client.Get(context.TODO(), id.String()).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return res, nil
}

func (r *Cache) Set(key uuid.UUID, value interface{}) error {
	val, err := json.Marshal(value)
	if err != nil {
		log.Error(err)
		return err
	}
	return r.client.Set(context.TODO(), key.String(), val, r.exp).Err()
}

func (r *Cache) Delete(key uuid.UUID) error {
	return r.client.Del(context.TODO(), key.String()).Err()
}