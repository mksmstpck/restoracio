package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mksmstpck/restoracio/backend/internal/models"
	"github.com/mksmstpck/restoracio/backend/utils"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (d *ReservDB) CreateOne(ctx context.Context, reserv models.ReservDB) (models.ReservDB, error) {
	_, err := d.db.
		NewInsert().
		Model(&reserv).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return models.ReservDB{}, err
	}
	log.Print(reserv)
	log.Info("reservation created")
	return reserv, nil
}

func (d ReservDB) GetByID(ctx context.Context, id uuid.UUID, restaurantID uuid.UUID) (models.ReservDB, error) {
	var reserv models.ReservDB
	err := d.db.
		NewSelect().
		Model(&reserv).
		Where("id = ? AND restaurant_id = ?", id, restaurantID).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(utils.ErrReservationNotFound)
			return models.ReservDB{}, errors.New(utils.ErrReservationNotFound)
		}
		log.Error(err)
		return models.ReservDB{}, err
	}
	log.Info("reservation found")
	return reserv, nil
}

func (d ReservDB) GetAllInRestaurant(ctx context.Context, restauranID uuid.UUID) ([]models.ReservDB, error) {
	var reservs []models.ReservDB
	err := d.db.
		NewSelect().
		Model(&reservs).
		Where("restaurant_id = ?", restauranID).
		Scan(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("reservations found")
	return reservs, nil
}

func (d ReservDB) UpdateOne(ctx context.Context, reserv models.ReservDB) error {
	log.Print(reserv.ID)
	log.Print(reserv.RestaurantID)
	res, err := d.db.
		NewUpdate().
		Model(&reserv).
		Where("id = ? AND restaurant_id = ?", reserv.ID, reserv.RestaurantID).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return err
	}
	if count == 0 {
		log.Error(utils.ErrReservationNotFound)
		return errors.New(utils.ErrReservationNotFound)
	}
	log.Info("reservation updated")
	return nil
}

func (d ReservDB) DeleteOne(ctx context.Context, id uuid.UUID, restaurantID uuid.UUID) error {
	res, err := d.db.
		NewDelete().
		Model(&models.ReservDB{}).
		Where("id = ? AND restaurant_id = ?", id, restaurantID).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return err
	}
	if count == 0 {
		log.Error(utils.ErrReservationNotFound)
		return errors.New(utils.ErrReservationNotFound)
	}
	log.Info("reservation deleted")
	return nil
}

func (d ReservDB) DeleteAllByRestaurant(ctx context.Context, restaurantID uuid.UUID) (error) {
	res, err := d.db.
		NewDelete().
		Model(&models.ReservDB{}).
		Where("restaurant_id = ?", restaurantID).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return err
	}
	if count == 0 {
		log.Error(utils.ErrReservationNotFound)
		return errors.New(utils.ErrReservationNotFound)
	}
	log.Info("reservation deleted")
	return nil
}
