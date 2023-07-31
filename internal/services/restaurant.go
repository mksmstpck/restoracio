package services

import (
	"github.com/mksmstpck/restoracio/pkg/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Services) RestaurantCreateService(rest models.Restaurant, adminID uuid.UUID) (models.Restaurant, error) {
	rest.AdminID = adminID.String()
	res, err := s.restdb.RestaurantCreate(rest)
	if err != nil {
		log.Info("RestaurantCreate: ", err)
		return models.Restaurant{}, err
	}
	err = s.admindb.AdminUpdate(models.Admin{RestaurantID: res.ID})
	log.Info("restaurant created")
	return res, nil
}

func (s *Services) RestaurantGetByIDService(id uuid.UUID) (models.Restaurant, error) {
	res, err := s.restdb.RestaurantGetByID(id)
	if err != nil {
		log.Info("RestaurantGetByID: ", err)
		return models.Restaurant{}, err
	}
	log.Info("restaurant found")
	return res, nil
}

func (s *Services) RestaurantGetByAdminsIDService(id uuid.UUID) (models.Restaurant, error) {
	res, err := s.restdb.RestaurantGetByAdminsID(id)
	if err != nil {
		log.Info("RestaurantGetByAdminsID: ", err)
		return models.Restaurant{}, err
	}
	log.Info("restaurant found")
	return res, nil
}

func (s *Services) RestaurantUpdateService(rest models.Restaurant) error {
	err := s.restdb.RestaurantUpdate(rest)
	if err != nil {
		log.Info("RestaurantUpdate: ", err)
		return err
	}
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
	log.Info("restaurant deleted")
	return nil
}
