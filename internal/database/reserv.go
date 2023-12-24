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

func (d *ReservDB) CreateOne(ctx context.Context, reserv dto.ReservDB) (dto.ReservDB, error) {
	_, err := d.db.
		NewInsert().
		Model(&reserv).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return dto.ReservDB{}, err
	}
	log.Print(reserv)
	log.Info("reservation created")
	return reserv, nil
}

func (d ReservDB) GetByID(ctx context.Context, id uuid.UUID, restaurantID uuid.UUID) (dto.ReservDB, error) {
	var reserv dto.ReservDB
	err := d.db.
		NewSelect().
		Model(&reserv).
		Where("id = ? AND restaurant_id = ?", id, restaurantID).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(models.ErrReservationNotFound)
			return dto.ReservDB{}, errors.New(models.ErrReservationNotFound)
		}
		log.Error(err)
		return dto.ReservDB{}, err
	}
	log.Info("reservation found")
	return reserv, nil
}

func (d ReservDB) GetAllInRestaurant(ctx context.Context, restauranID uuid.UUID) ([]dto.ReservDB, error) {
	var reservs []dto.ReservDB
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

func (d ReservDB) UpdateOne(ctx context.Context, reserv dto.ReservDB) error {
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
		log.Error(models.ErrReservationNotFound)
		return errors.New(models.ErrReservationNotFound)
	}
	log.Info("reservation updated")
	return nil
}

func (d ReservDB) DeleteOne(ctx context.Context, id uuid.UUID, restaurantID uuid.UUID) error {
	res, err := d.db.
		NewDelete().
		Model(&dto.ReservDB{}).
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
		log.Error(models.ErrReservationNotFound)
		return errors.New(models.ErrReservationNotFound)
	}
	log.Info("reservation deleted")
	return nil
}

func (d ReservDB) DeleteAllByRestaurant(ctx context.Context, restaurantID uuid.UUID) (error) {
	res, err := d.db.
		NewDelete().
		Model(&dto.ReservDB{}).
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
		log.Error(models.ErrReservationNotFound)
		return errors.New(models.ErrReservationNotFound)
	}
	log.Info("reservation deleted")
	return nil
}
