package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (d *RestDatabase) CreateOne(
	ctx context.Context,
	restaurant dto.Restaurant,
	) (dto.Restaurant, error) {
	_, err := d.db.
		NewInsert().
		Model(&restaurant).
		Exec(ctx)
	if err != nil {
		log.Error("database.RestaurantCreate: ", err)
		return dto.Restaurant{}, err
	}
	log.Info("restaurant created")
	return restaurant, nil
}

func (d *RestDatabase) GetByID(ctx context.Context, id uuid.UUID) (dto.Restaurant, error) {
	var restaurant dto.Restaurant
	err := d.db.
		NewSelect().
		Model(&restaurant).
		Relation("Tables").
		Where("restaurant.id = ?", id.String()).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("restaurant not found")
			return dto.Restaurant{}, errors.New("restaurant not found")
		}
		log.Error("database.RestaurantGetByID: ", err)
		return dto.Restaurant{}, err
	}
	log.Info("restaurant found")
	return restaurant, nil
}

func (d *RestDatabase) UpdateOne(ctx context.Context, restaurant dto.Restaurant) error {
	_, err := d.db.
		NewUpdate().
		Model(&restaurant).
		ExcludeColumn("admin_id", "id").
		Where("id = ?", restaurant.ID).
		Exec(ctx)
	if err != nil {
		log.Error("database.RestaurantUpdate: ", err)
		return err
	}
	log.Info("restaurant updated")
	return nil
}

func (d *RestDatabase) DeleteOne(ctx context.Context, id uuid.UUID) error {
	_, err := d.db.
		NewDelete().
		Model(&dto.Restaurant{}).
		Where("id = ?", id.String()).
		Exec(ctx)
	if err != nil {
		log.Error("database.RestaurantDelete: ", err)
		return err
	}
	log.Info("restaurant deleted")
	return nil
}
