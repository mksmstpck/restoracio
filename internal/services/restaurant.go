package services

import (
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
	admin.RestaurantID = res.ID
	err = s.admindb.AdminUpdate(admin)
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

func (s *Services) RestaurantGetByAdminsIDService(id uuid.UUID) (models.Restaurant, error) {
	res, exist := s.cache.Get(id.String())
	if exist {
		log.Info("restaurant found")
		return res.(models.Restaurant), nil
	}
	res, err := s.restdb.RestaurantGetByAdminsID(id)
	if err != nil {
		log.Info("RestaurantGetByAdminsID: ", err)
		return models.Restaurant{}, err
	}

	s.cache.Set(res.(models.Restaurant).ID, res, cache.DefaultExpiration)

	log.Info("restaurant found")
	return res.(models.Restaurant), nil
}

func (s *Services) RestaurantUpdateService(rest models.Restaurant) error {
	err := s.restdb.RestaurantUpdate(rest)
	if err != nil {
		log.Info("RestaurantUpdate: ", err)
		return err
	}

	s.cache.Set(rest.ID, rest, cache.DefaultExpiration)

	log.Info("restaurant updated")
	return nil
}

func (s *Services) RestaurantDeleteService(id uuid.UUID) error {
	admin, err := s.admindb.AdminGetByID(id)
	if err != nil {
		log.Info("RestaurantDelete: ", err)
		return err
	}
	err = s.restdb.RestaurantDelete(id)
	if err != nil {
		log.Info("RestaurantDelete: ", err)
		return err
	}
	admin.RestaurantID = uuid.NIL.String()
	err = s.admindb.AdminUpdate(admin)

	s.cache.Delete(id.String())

	s.cache.Set(admin.ID, admin, cache.DefaultExpiration)
	s.cache.Set(admin.Email, admin, cache.DefaultExpiration)

	log.Info("restaurant deleted")
	return nil
}
