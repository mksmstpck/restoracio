package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Services) StaffCreateService(staff dto.Staff, admin dto.Admin) (dto.Staff, error) {
	if admin.Restaurant == nil {
		log.Info(dto.ErrRestaurantNotFound)
		return dto.Staff{}, errors.New(dto.ErrRestaurantNotFound)
	}
	staff.RestaurantID = admin.Restaurant.ID
	staff.ID = uuid.NewUUID().String()
	res, err := s.db.Staff.CreateOne(s.ctx, staff)
	if err != nil {
		log.Info(err)
		return dto.Staff{}, err
	}

	err = s.cache.Set(uuid.Parse(res.ID), res)
	if err != nil {
		log.Info(err)
		return dto.Staff{}, err
	}

	log.Info("staff created")
	return res, nil
}

func (s *Services) StaffGetByIDService(id uuid.UUID, admin dto.Admin) (dto.Staff, error) {
	if admin.Restaurant == nil {
		log.Info(dto.ErrRestaurantNotFound)
		return dto.Staff{}, errors.New(dto.ErrRestaurantNotFound)
	}

	staffAny, err := s.cache.Get(id)
	if staffAny != nil {
		log.Info("staff found")
		return staffAny.(dto.Staff), nil
	}
	if err != nil {
		log.Info(err)
		return dto.Staff{}, err
	}

	staff, err := s.db.Staff.GetByID(s.ctx, id, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Info(err)
		return dto.Staff{}, err
	}
	return staff, nil
}

func (s *Services) StaffGetAllInRestaurantService(admin dto.Admin) ([]dto.Staff, error) {
	if admin.Restaurant == nil {
		log.Info(dto.ErrRestaurantNotFound)
		return nil, errors.New(dto.ErrRestaurantNotFound)
	}
	staff, err := s.db.Staff.GetAllInRestaurant(s.ctx, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Info(err)
		return nil, err
	}
	log.Info("staffs found")
	return staff, nil
}

func (s *Services) StaffUpdateService(staff dto.Staff, admin dto.Admin) error {
	if admin.Restaurant == nil {
		log.Info(dto.ErrRestaurantNotFound)
		return errors.New(dto.ErrRestaurantNotFound)
	}
	staff.RestaurantID = admin.Restaurant.ID
	err := s.db.Staff.UpdateOne(s.ctx, staff)
	if err != nil {
		log.Info(err)
		return err
	}

	s.cache.Set(uuid.Parse(staff.ID), staff)

	log.Info("staff updated")
	return nil
}

func (s *Services) StaffDeleteService(id uuid.UUID, admin dto.Admin) error {
	if admin.Restaurant == nil {
		log.Info(dto.ErrRestaurantNotFound)
		return errors.New(dto.ErrRestaurantNotFound)
	}
	err := s.db.Staff.DeleteOne(s.ctx, id, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Info(err)
		return err
	}

	s.cache.Delete(id)

	log.Info("staff deleted")
	return nil
}

func (s *Services) StaffDeleteAllService(admin dto.Admin) error {
	if admin.Restaurant == nil {
		log.Info(dto.ErrRestaurantNotFound)
		return errors.New(dto.ErrRestaurantNotFound)
	}
	err := s.db.Staff.DeleteAll(s.ctx, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Info(err)
		return err
	}
	log.Info("staffs deleted")
	return nil
}