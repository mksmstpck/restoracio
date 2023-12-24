package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
	"github.com/mksmstpck/restoracio/utils/convertors"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Services) ReservCreateService(reserv dto.ReservRequest, admin dto.AdminResponse) error {
	if admin.Restaurant == nil {
		log.Info(models.ErrRestaurantNotFound)
		return errors.New(models.ErrRestaurantNotFound)
	}
	if !TableExists(admin.Restaurant.Tables, reserv.TableID) {
		log.Info(models.ErrTableNotFound)
		return errors.New(models.ErrTableNotFound)
	}

	reservDB := convertors.ReservRequestToDB(reserv, admin.Restaurant.ID)

	reservDB, err := s.db.Reserv.CreateOne(s.ctx, reservDB)
	if err != nil {
		log.Error(err)
		return err
	}

	s.cache.Set(uuid.Parse(reservDB.ID), reservDB)

	log.Info("reservation created")
	return nil
}

func (s *Services) ReservGetByIDService(id uuid.UUID, admin dto.AdminResponse) (*dto.ReservDB, error) {
	if admin.Restaurant == nil {
		log.Info(models.ErrRestaurantNotFound)
		return nil, errors.New(models.ErrRestaurantNotFound)
	}

	reservAny, err := s.cache.Get(id)
	if reservAny != nil {
		log.Info("reservation found")
		return reservAny.(*dto.ReservDB), nil
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}

	reserv, err := s.db.Reserv.GetByID(s.ctx, id, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("reservation found")
	return &reserv, nil
}

func (s *Services) ReservGetAllInRestaurantService(admin dto.AdminResponse) ([]dto.ReservDB, error) {
	if admin.Restaurant == nil {
		log.Info(models.ErrRestaurantNotFound)
		return nil, errors.New(models.ErrRestaurantNotFound)
	}
	reservs, err := s.db.Reserv.GetAllInRestaurant(s.ctx, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("reservations found")
	return reservs, nil
}

func (s *Services) ReservUpdateService(reserv dto.ReservRequest, admin dto.AdminResponse) error {
	if admin.Restaurant == nil {
		log.Info(models.ErrRestaurantNotFound)
		return errors.New(models.ErrRestaurantNotFound)
	}
	if !TableExists(admin.Restaurant.Tables, reserv.TableID) {
		log.Info(models.ErrTableNotFound)
		return errors.New(models.ErrTableNotFound)
	}
	
	reservDB := convertors.ReservRequestToDB(reserv, admin.Restaurant.ID)

	err := s.db.Reserv.UpdateOne(s.ctx, reservDB)
	if err != nil {
		log.Error(err)
		return err
	}

	s.cache.Set(uuid.Parse(reservDB.ID), reservDB)

	log.Info("reservation updated")
	return nil
}

func (s *Services) ReservDeleteService(id uuid.UUID, admin dto.AdminResponse) error {
	if admin.Restaurant == nil {
		log.Info(models.ErrRestaurantNotFound)
		return errors.New(models.ErrRestaurantNotFound)
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