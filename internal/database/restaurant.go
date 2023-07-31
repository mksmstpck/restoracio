package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mksmstpck/restoracio/pkg/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (d *RestDatabase) RestaurantCreate(restaurant models.Restaurant) (models.Restaurant, error) {
	restaurant.ID = uuid.NewUUID().String()
	_, err := d.db.
		NewInsert().
		Model(&restaurant).
		Exec(context.Background())
	if err != nil {
		log.Error("database.RestaurantCreate: ", err)
		return models.Restaurant{}, err
	}
	log.Info("restaurant created")
	return restaurant, nil
}

func (d *RestDatabase) RestaurantGetByID(id uuid.UUID) (models.Restaurant, error) {
	var restaurant models.Restaurant
	err := d.db.
		NewSelect().
		Model(&restaurant).
		Where("id = ?", id.String()).
		Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("restaurant not found")
			return models.Restaurant{}, errors.New("restaurant not found")
		}
		log.Error("database.RestaurantGetByID: ", err)
		return models.Restaurant{}, err
	}
	log.Info("restaurant found")
	return restaurant, nil
}

func (d *RestDatabase) RestaurantGetByAdminsID(id uuid.UUID) (models.Restaurant, error) {
	var restaurant models.Restaurant
	err := d.db.
		NewSelect().
		Model(&restaurant).
		Where("admin_id = ?", id.String()).
		Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("restaurant not found")
			return models.Restaurant{}, errors.New("restaurant not found")
		}
		log.Error("database.RestaurantGetByID: ", err)
		return models.Restaurant{}, err
	}
	log.Info("restaurant found")
	return restaurant, nil
}

func (d *RestDatabase) RestaurantUpdate(restaurant models.Restaurant) error {
	_, err := d.db.
		NewUpdate().
		Model(&restaurant).
		ExcludeColumn("admin_id", "id").
		Where("id = ?", restaurant.ID).
		Exec(context.Background())
	if err != nil {
		log.Error("database.RestaurantUpdate: ", err)
		return err
	}
	log.Info("restaurant updated")
	return nil
}

func (d *RestDatabase) RestaurantDelete(id uuid.UUID) error {
	_, err := d.db.
		NewDelete().
		Model(&models.Restaurant{}).
		Where("id = ?", id.String()).
		Exec(context.Background())
	if err != nil {
		log.Error("database.RestaurantDelete: ", err)
		return err
	}
	log.Info("restaurant deleted")
	return nil
}
