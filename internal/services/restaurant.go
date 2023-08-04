package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/pkg/models"
	"github.com/patrickmn/go-cache"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Services) RestaurantCreateService(rest models.Restaurant, admin models.Admin) (models.Restaurant, error) {
	rest.AdminID = admin.ID
	res, err := s.restdb.RestaurantCreate(rest)
	if err != nil {
		log.Info("RestaurantCreate: ", err)
		return models.Restaurant{}, err
	}

	s.cache.Set(res.ID, res, cache.DefaultExpiration)
	s.cache.Set(admin.ID, admin, cache.DefaultExpiration)

	log.Info("restaurant created")
	return res, nil
}

func (s *Services) RestaurantGetByIDService(id uuid.UUID) (models.Restaurant, error) {
	res, exist := s.cache.Get(id.String())
	if exist {
		log.Info("restaurant found")
		return res.(models.Restaurant), nil
	}
	res, err := s.restdb.RestaurantGetByID(id)
	if err != nil {
		log.Info("RestaurantGetByID: ", err)
		return models.Restaurant{}, err
	}

	s.cache.Set(res.(models.Restaurant).ID, res, cache.DefaultExpiration)

	log.Info("restaurant found")
	return res.(models.Restaurant), nil
}

func (s *Services) RestaurantUpdateService(rest models.Restaurant, restID uuid.UUID) error {
	rest.ID = restID.String()
	err := s.restdb.RestaurantUpdate(rest)
	if err != nil {
		log.Info("RestaurantUpdate: ", err)
		return err
	}

	s.cache.Set(rest.ID, rest, cache.DefaultExpiration)

	log.Info("restaurant updated")
	return nil
}

func (s *Services) RestaurantDeleteService(rest *models.Restaurant) error {
	if rest == nil {
		log.Info("restaurant not found")
		return errors.New("restaurant not found")
	}
	err := s.restdb.RestaurantDelete(uuid.Parse(rest.ID))
	if err != nil {
		log.Info("RestaurantDelete: ", err)
		return err
	}

	s.cache.Delete(rest.ID)
	s.cache.Delete(rest.AdminID)

	log.Info("restaurant deleted")
	return nil
}
