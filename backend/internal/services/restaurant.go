package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/backend/internal/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Services) RestaurantCreateService(rest models.Restaurant, admin models.Admin) (models.Restaurant, error) {
	rest.AdminID = admin.ID
	rest.ID = uuid.NewUUID().String()
	res, err := s.db.Rest.CreateOne(s.ctx, rest)
	if err != nil {
		log.Info(err)
		return models.Restaurant{}, err
	}

	s.cache.Set(uuid.Parse(res.ID), res)
	s.cache.Set(uuid.Parse(admin.ID), admin)

	log.Info("restaurant created")
	return res, nil
}

func (s *Services) RestaurantGetByIDService(id uuid.UUID) (models.Restaurant, error) {
	res, err := s.cache.RestaurantGet(id)
	if res.ID != "" {
		log.Info("restaurant found")
		return res, nil
	}
	if err != nil {
		log.Info(err)
		return models.Restaurant{}, err
	}

	res, err = s.db.Rest.GetByID(s.ctx, id)
	if err != nil {
		log.Info(err)
		return models.Restaurant{}, err
	}

	s.cache.Set(uuid.Parse(res.ID), res)

	log.Info("restaurant found")
	return res, nil
}

func (s *Services) RestaurantUpdateService(rest models.Restaurant, restID uuid.UUID) error {
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

func (s *Services) RestaurantDeleteService(rest *models.Restaurant) error {
	if rest == nil {
		log.Info("restaurant not found")
		return errors.New("restaurant not found")
	}
	err := s.db.Rest.DeleteOne(s.ctx, uuid.Parse(rest.ID))
	if err != nil {
		log.Info(err)
		return err
	}
	
	admin := models.Admin{ID: rest.AdminID,Restaurant: rest}
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
