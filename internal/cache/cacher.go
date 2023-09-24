package cache

import (
	"context"
	"encoding/json"

	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/pborman/uuid"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func (r *Cache) AdminGet(id uuid.UUID) (models.Admin, error) {
	var res models.Admin
	val, err := r.client.Get(context.TODO(), id.String()).Result()
	if err == redis.Nil {
		return models.Admin{}, nil
	}
	if err != nil {
		log.Error(err)
		return models.Admin{}, err
	}
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		log.Error(err)
		return models.Admin{}, err
	}
	return res, nil
}

func (r *Cache) DishGet(id uuid.UUID) (models.Dish, error) {
	var res models.Dish
	val, err := r.client.Get(context.TODO(), id.String()).Result()
	if err == redis.Nil {
		return models.Dish{}, nil
	}
	if err != nil {
		log.Error(err)
		return models.Dish{}, err
	}
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		log.Error(err)
		return models.Dish{}, err
	}
	return res, nil
}

func (r *Cache) MenuGet(id uuid.UUID) (models.Menu, error) {
	var res models.Menu
	val, err := r.client.Get(context.TODO(), id.String()).Result()
	if err == redis.Nil {
		return models.Menu{}, nil
	}
	if err != nil {
		log.Error(err)
		return models.Menu{}, err
	}
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		log.Error(err)
		return models.Menu{}, err
	}
	return res, nil
}

func (r *Cache) RestaurantGet(id uuid.UUID) (models.Restaurant, error) {
	var res models.Restaurant
	val, err := r.client.Get(context.TODO(), id.String()).Result()
	if err == redis.Nil {
		return models.Restaurant{}, nil
	}
	if err != nil {
		log.Error(err)
		return models.Restaurant{}, err
	}
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		log.Error(err)
		return models.Restaurant{}, err
	}
	return res, nil
}

func (r *Cache) TableGet(id uuid.UUID) (models.Table, error) {
	var res models.Table
	val, err := r.client.Get(context.TODO(), id.String()).Result()
	if err == redis.Nil {
		return models.Table{}, nil
	}
	if err != nil {
		log.Error(err)
		return models.Table{}, err
	}
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		log.Error(err)
		return models.Table{}, err
	}
	return res, nil
}

func (r *Cache) StaffGet(id uuid.UUID) (models.Staff, error) {
	var res models.Staff
	val, err := r.client.Get(context.TODO(), id.String()).Result()
	if err == redis.Nil {
		return models.Staff{}, nil
	}
	if err != nil {
		log.Error(err)
		return models.Staff{}, err
	}
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		log.Error(err)
		return models.Staff{}, err
	}
	return res, nil
}

func (r *Cache) ReservGet(id uuid.UUID) (models.ReservDB, error) {
	var res models.ReservDB
	val, err := r.client.Get(context.TODO(), id.String()).Result()
	if err == redis.Nil {
		return models.ReservDB{}, nil
	}
	if err != nil {
		log.Error(err)
		return models.ReservDB{}, err
	}
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		log.Error(err)
		return models.ReservDB{}, err
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