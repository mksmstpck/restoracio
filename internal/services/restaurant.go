package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Services) RestaurantCreateService(rest dto.RestaurantRequest, admin dto.AdminResponse)  error {
	rest.AdminID = admin.ID
	rest.ID = uuid.NewUUID().String()
	res, err := s.db.Rest.CreateOne(s.ctx, rest)
	if err != nil {
		log.Info(err)
		return dto.Restaurant{}, err
	}

	s.cache.Set(uuid.Parse(res.ID), res)
	s.cache.Set(uuid.Parse(admin.ID), admin)

	log.Info("restaurant created")
	return res, nil
}

func (s *Services) RestaurantGetByIDService(id uuid.UUID) (dto.Restaurant, error) {
	resAny, err := s.cache.Get(id)
	if resAny != nil {
		log.Info("restaurant found")
		return resAny.(dto.Restaurant), nil
	}
	if err != nil {
		log.Info(err)
		return dto.Restaurant{}, err
	}

	res, err := s.db.Rest.GetByID(s.ctx, id)
	if err != nil {
		log.Info(err)
		return dto.Restaurant{}, err
	}

	s.cache.Set(uuid.Parse(res.ID), res)

	log.Info("restaurant found")
	return res, nil
}

func (s *Services) RestaurantUpdateService(rest dto.Restaurant, restID uuid.UUID) error {
	rest.ID = restID.String()
	err := s.db.Rest.UpdateOne(s.ctx, rest)
	if err != nil {
		log.Info(err)
		return err
	}

	s.cache.Set(uuid.Parse(rest.ID), rest)

	log.Info("restaurant updated")
	return nil
}

func (s *Services) RestaurantDeleteService(rest *dto.Restaurant) error {
	if rest == nil {
		log.Info(models.ErrRestaurantNotFound)
		return errors.New(models.ErrRestaurantNotFound)
	}
	err := s.db.Rest.DeleteOne(s.ctx, uuid.Parse(rest.ID))
	if err != nil {
		log.Info(err)
		return err
	}
	
	admin := dto.Admin{ID: rest.AdminID,Restaurant: rest}
	err = s.MenuDeleteService(admin)
	if err != nil {
		log.Info(err)
		return err
	}
	err = s.TableDeleteAllService(admin)
	if err != nil {
		log.Info(err)
		return err
	}
	err = s.StaffDeleteAllService(admin)
	if err != nil {
		log.Info(err)
		return err
	}

	s.cache.Delete(uuid.Parse(rest.ID))
	s.cache.Delete(uuid.Parse(rest.AdminID))

	log.Info("restaurant deleted")
	return nil
}
