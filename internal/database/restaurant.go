package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mksmstpck/restoracio/internal/convertors"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (d *RestDatabase) CreateOne(ctx context.Context, restaurant dto.Restaurant) error {
	rest := convertors.RestaurantDTOToDB(&restaurant)
	_, err := d.db.
		NewInsert().
		Model(&rest).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("restaurant created")
	return nil
}

func (d *RestDatabase) GetByID(ctx context.Context, id uuid.UUID) (dto.Restaurant, error) {
	var restaurant models.Restaurant
	err := d.db.
		NewSelect().
		Model(&restaurant).
		Relation("Tables").
		Where("restaurant.id = ?", id.String()).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(models.ErrRestaurantNotFound)
			return dto.Restaurant{}, errors.New(models.ErrRestaurantNotFound)
		}
		log.Error(err)
		return dto.Restaurant{}, err
	}
	log.Info("restaurant found")
	return convertors.RestaurantDBToDTO(&restaurant), nil
}

func (d *RestDatabase) UpdateOne(ctx context.Context, restaurant dto.Restaurant) error {
	rest := convertors.RestaurantDTOToDB(&restaurant)
	_, err := d.db.
		NewUpdate().
		Model(&rest).
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
		Model(&models.Restaurant{}).
		Where("id = ?", id.String()).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("restaurant deleted")
	return nil
}
