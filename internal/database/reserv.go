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

func (d *ReservDB) CreateOne(ctx context.Context, reserv dto.Reserv) error {
	reservDB := convertors.ReservDTOToDB(&reserv)
	_, err := d.db.
		NewInsert().
		Model(&reservDB).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Print(reserv)
	log.Info("reservation created")
	return nil
}

func (d ReservDB) GetByID(ctx context.Context, id uuid.UUID, restaurantID uuid.UUID) (dto.Reserv, error) {
	var reserv models.Reserv
	err := d.db.
		NewSelect().
		Model(&reserv).
		Where("id = ? AND restaurant_id = ?", id, restaurantID).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(models.ErrReservationNotFound)
			return dto.Reserv{}, errors.New(models.ErrReservationNotFound)
		}
		log.Error(err)
		return dto.Reserv{}, err
	}
	log.Info("reservation found")
	return convertors.ReservDBToDTO(&reserv), nil
}

func (d ReservDB) GetAllInRestaurant(ctx context.Context, restauranID uuid.UUID) ([]dto.Reserv, error) {
	var reservs []models.Reserv
	err := d.db.
		NewSelect().
		Model(&reservs).
		Where("restaurant_id = ?", restauranID).
		Scan(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	reservsDTO := make([]dto.Reserv, len(reservs))
	for i, reservDB := range reservs {
		reservsDTO[i] = convertors.ReservDBToDTO(&reservDB)
	}
	log.Info("reservations found")
	return reservsDTO, nil
}

func (d ReservDB) UpdateOne(ctx context.Context, reserv dto.Reserv) error {
	reservDB := convertors.ReservDTOToDB(&reserv)
	res, err := d.db.
		NewUpdate().
		Model(&reservDB).
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
		Model(&models.Reserv{}).
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

func (d ReservDB) DeleteAllByRestaurant(ctx context.Context, restaurantID uuid.UUID) error {
	res, err := d.db.
		NewDelete().
		Model(&dto.Reserv{}).
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
