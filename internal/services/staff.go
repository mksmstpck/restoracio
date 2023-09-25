package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Services) StaffCreateService(staff models.Staff, admin models.Admin) (models.Staff, error) {
	if admin.Restaurant == nil {
		log.Info(utils.ErrRestaurantNotFound)
		return models.Staff{}, errors.New(utils.ErrRestaurantNotFound)
	}
	staff.RestaurantID = admin.Restaurant.ID
	res, err := s.db.Staff.CreateOne(s.ctx, staff)
	if err != nil {
		log.Info(err)
		return models.Staff{}, err
	}

	err = s.cache.Set(uuid.Parse(res.ID), res)
	if err != nil {
		log.Info(err)
		return models.Staff{}, err
	}

	log.Info("staff created")
	return res, nil
}

func (s *Services) StaffGetByIDService(id uuid.UUID, admin models.Admin) (models.Staff, error) {
	if admin.Restaurant == nil {
		log.Info(utils.ErrRestaurantNotFound)
		return models.Staff{}, errors.New(utils.ErrRestaurantNotFound)
	}

	staff, err := s.cache.StaffGet(id)
	if staff.ID != "" {
		log.Info("staff found")
		return staff, nil
	}
	if err != nil {
		log.Info(err)
		return models.Staff{}, err
	}

	staff, err = s.db.Staff.GetByID(s.ctx, id, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Info(err)
		return models.Staff{}, err
	}
	return staff, nil
}

func (s *Services) StaffGetAllInRestaurantService(admin models.Admin) ([]models.Staff, error) {
	if admin.Restaurant == nil {
		log.Info(utils.ErrRestaurantNotFound)
		return nil, errors.New(utils.ErrRestaurantNotFound)
	}
	staff, err := s.db.Staff.GetAllInRestaurant(s.ctx, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Info(err)
		return nil, err
	}
	log.Info("staffs found")
	return staff, nil
}

func (s *Services) StaffUpdateService(staff models.Staff, admin models.Admin) error {
	if admin.Restaurant == nil {
		log.Info(utils.ErrRestaurantNotFound)
		return errors.New(utils.ErrRestaurantNotFound)
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

func (s *Services) StaffDeleteService(id uuid.UUID, admin models.Admin) error {
	if admin.Restaurant == nil {
		log.Info(utils.ErrRestaurantNotFound)
		return errors.New(utils.ErrRestaurantNotFound)
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

func (s *Services) StaffDeleteAllService(admin models.Admin) error {
	if admin.Restaurant == nil {
		log.Info(utils.ErrRestaurantNotFound)
		return errors.New(utils.ErrRestaurantNotFound)
	}
	err := s.db.Staff.DeleteAll(s.ctx, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Info(err)
		return err
	}
	log.Info("staffs deleted")
	return nil
}