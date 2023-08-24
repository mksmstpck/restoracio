package services

import (
	"errors"
	"time"

	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/patrickmn/go-cache"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Services) ReservCreateService(reserv models.ReservAPI, admin models.Admin) (models.ReservDB, error) {
	if admin.Restaurant == nil {
		log.Info(utils.ErrRestaurantNotFound)
		return models.ReservDB{}, errors.New(utils.ErrRestaurantNotFound)
	}
	if !TableExists(admin.Restaurant.Tables, reserv.TableID) {
		log.Info(utils.ErrTableNotFound)
		return models.ReservDB{}, errors.New(utils.ErrTableNotFound)
	}
	var reservDB models.ReservDB
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

	reservDB, err := s.db.Reserv.CreateOne(s.ctx, reservDB)
	if err != nil {
		log.Error(err)
		return models.ReservDB{}, err
	}

	s.cache.Set(reservDB.ID, reservDB, cache.DefaultExpiration)

	log.Info("reservation created")
	return reservDB, nil
}

func (s *Services) ReservGetByIDService(id uuid.UUID, admin models.Admin) (models.ReservDB, error) {
	if admin.Restaurant == nil {
		log.Info(utils.ErrRestaurantNotFound)
		return models.ReservDB{}, errors.New(utils.ErrRestaurantNotFound)
	}

	reserv, exist := s.cache.Get(id.String())
	if exist {
		log.Info("reservation found")
		return reserv.(models.ReservDB), nil
	}

	reserv, err := s.db.Reserv.GetByID(s.ctx, id, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Error(err)
		return models.ReservDB{}, err
	}
	log.Info("reservation found")
	return reserv.(models.ReservDB), nil
}

func (s *Services) ReservGetAllInRestaurantService(admin models.Admin) ([]models.ReservDB, error) {
	if admin.Restaurant == nil {
		log.Info(utils.ErrRestaurantNotFound)
		return nil, errors.New(utils.ErrRestaurantNotFound)
	}
	reservs, err := s.db.Reserv.GetAllInRestaurant(s.ctx, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("reservations found")
	return reservs, nil
}

func (s *Services) ReservUpdateService(reserv models.ReservAPI, admin models.Admin) error {
	if admin.Restaurant == nil {
		log.Info(utils.ErrRestaurantNotFound)
		return errors.New(utils.ErrRestaurantNotFound)
	}
	if !TableExists(admin.Restaurant.Tables, reserv.TableID) {
		log.Info(utils.ErrTableNotFound)
		return errors.New(utils.ErrTableNotFound)
	}
	var reservDB models.ReservDB
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

	s.cache.Set(reservDB.ID, reservDB, cache.DefaultExpiration)

	log.Info("reservation updated")
	return nil
}

func (s *Services) ReservDeleteService(id uuid.UUID, admin models.Admin) error {
	if admin.Restaurant == nil {
		log.Info(utils.ErrRestaurantNotFound)
		return errors.New(utils.ErrRestaurantNotFound)
	}
	log.Print(id)
	log.Print(uuid.Parse(admin.Restaurant.ID))
	err := s.db.Reserv.DeleteOne(s.ctx, id, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Error(err)
		return err
	}

	s.cache.Delete(id.String())

	log.Info("reservation deleted")
	return nil
}