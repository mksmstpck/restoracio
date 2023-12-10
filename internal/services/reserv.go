package services

import (
	"errors"
	"time"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Services) ReservCreateService(reserv dto.ReservAPI, admin dto.Admin) (dto.ReservDB, error) {
	if admin.Restaurant == nil {
		log.Info(dto.ErrRestaurantNotFound)
		return dto.ReservDB{}, errors.New(dto.ErrRestaurantNotFound)
	}
	if !TableExists(admin.Restaurant.Tables, reserv.TableID) {
		log.Info(dto.ErrTableNotFound)
		return dto.ReservDB{}, errors.New(dto.ErrTableNotFound)
	}
	var reservDB dto.ReservDB
	reservDB.ReservationTime = time.Date(
		reserv.Year,
		time.Month(reserv.Month),
		reserv.Day, reserv.Hour,
		reserv.Minute,
		reserv.Second,
		0,
		time.UTC,
	)
	reservDB.ReserverName = reserv.ReserverName
	reservDB.ReserverPhone = reserv.ReserverPhone
	reservDB.TableID = reserv.TableID
	reservDB.RestaurantID = admin.Restaurant.ID
	reserv.ID = uuid.NewUUID().String()

	reservDB, err := s.db.Reserv.CreateOne(s.ctx, reservDB)
	if err != nil {
		log.Error(err)
		return dto.ReservDB{}, err
	}

	s.cache.Set(uuid.Parse(reservDB.ID), reservDB)

	log.Info("reservation created")
	return reservDB, nil
}

func (s *Services) ReservGetByIDService(id uuid.UUID, admin dto.Admin) (dto.ReservDB, error) {
	if admin.Restaurant == nil {
		log.Info(dto.ErrRestaurantNotFound)
		return dto.ReservDB{}, errors.New(dto.ErrRestaurantNotFound)
	}

	reservAny, err := s.cache.Get(id)
	if reservAny != nil {
		log.Info("reservation found")
		return reservAny.(dto.ReservDB), nil
	}
	if err != nil {
		log.Error(err)
		return dto.ReservDB{}, err
	}

	reserv, err := s.db.Reserv.GetByID(s.ctx, id, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Error(err)
		return dto.ReservDB{}, err
	}
	log.Info("reservation found")
	return reserv, nil
}

func (s *Services) ReservGetAllInRestaurantService(admin dto.Admin) ([]dto.ReservDB, error) {
	if admin.Restaurant == nil {
		log.Info(dto.ErrRestaurantNotFound)
		return nil, errors.New(dto.ErrRestaurantNotFound)
	}
	reservs, err := s.db.Reserv.GetAllInRestaurant(s.ctx, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("reservations found")
	return reservs, nil
}

func (s *Services) ReservUpdateService(reserv dto.ReservAPI, admin dto.Admin) error {
	if admin.Restaurant == nil {
		log.Info(dto.ErrRestaurantNotFound)
		return errors.New(dto.ErrRestaurantNotFound)
	}
	if !TableExists(admin.Restaurant.Tables, reserv.TableID) {
		log.Info(dto.ErrTableNotFound)
		return errors.New(dto.ErrTableNotFound)
	}
	var reservDB dto.ReservDB
	reservDB.RestaurantID = admin.Restaurant.ID
	reservDB.ReservationTime = time.Date(
		reserv.Year,
		time.Month(reserv.Month),
		reserv.Day, reserv.Hour,
		reserv.Minute,
		reserv.Second,
		0,
		time.UTC,
	)
	reservDB.ReserverName = reserv.ReserverName
	reservDB.ReserverPhone = reserv.ReserverPhone
	reservDB.TableID = reserv.TableID
	reservDB.ID = reserv.ID
	err := s.db.Reserv.UpdateOne(s.ctx, reservDB)
	if err != nil {
		log.Error(err)
		return err
	}

	s.cache.Set(uuid.Parse(reservDB.ID), reservDB)

	log.Info("reservation updated")
	return nil
}

func (s *Services) ReservDeleteService(id uuid.UUID, admin dto.Admin) error {
	if admin.Restaurant == nil {
		log.Info(dto.ErrRestaurantNotFound)
		return errors.New(dto.ErrRestaurantNotFound)
	}
	log.Print(id)
	log.Print(uuid.Parse(admin.Restaurant.ID))
	err := s.db.Reserv.DeleteOne(s.ctx, id, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Error(err)
		return err
	}

	s.cache.Delete(id)

	log.Info("reservation deleted")
	return nil
}