package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (d *RestDatabase) CreateOne(
	ctx context.Context,
	restaurant dto.RestaurantDB,
	) (dto.RestaurantDB, error) {
	_, err := d.db.
		NewInsert().
		Model(&restaurant).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return dto.RestaurantDB{}, err
	}
	log.Info("restaurant created")
	return restaurant, nil
}

func (d *RestDatabase) GetByID(ctx context.Context, id uuid.UUID) (dto.RestaurantDB, error) {
	var restaurant dto.RestaurantDB
	err := d.db.
		NewSelect().
		Model(&restaurant).
		Relation("Tables").
		Where("restaurant.id = ?", id.String()).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(models.ErrRestaurantNotFound)
			return dto.RestaurantDB{}, errors.New(models.ErrRestaurantNotFound)
		}
		log.Error(err)
		return dto.RestaurantDB{}, err
	}
	log.Info("restaurant found")
	return restaurant, nil
}

func (d *RestDatabase) UpdateOne(ctx context.Context, restaurant dto.RestaurantDB) error {
	_, err := d.db.
		NewUpdate().
		Model(&restaurant).
		ExcludeColumn("admin_id", "id", "staff", "menu", "tables").
		Where("id = ?", restaurant.ID).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("restaurant updated")
	return nil
}

func (d *RestDatabase) DeleteOne(ctx context.Context, id uuid.UUID) error {
	_, err := d.db.
		NewDelete().
		Model(&dto.RestaurantDB{}).
		Where("id = ?", id.String()).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("restaurant deleted")
	return nil
}
