package services

import (
	"github.com/mksmstpck/restoracio/pkg/models"
	"github.com/pborman/uuid"
)

func (s *Services) RestaurantCreateService(rest models.Restaurant, adminID uuid.UUID) (models.Restaurant, error) {
	rest.AdminID = adminID.String()
	res, err := s.db.RestaurantCreate(rest)
	if err != nil {
		return models.Restaurant{}, err
	}
	return res, nil
}

func (s *Services) RestaurantGetByIDService(id uuid.UUID) (models.Restaurant, error) {
	res, err := s.db.RestaurantGetByID(id)
	if err != nil {
		return models.Restaurant{}, err
	}
	return res, nil
}

func (s *Services) RestaurantGetByAdminsIDService(id uuid.UUID) (models.Restaurant, error) {
	res, err := s.db.RestaurantGetByAdminsID(id)
	if err != nil {
		return models.Restaurant{}, err
	}
	return res, nil
}

func (s *Services) RestaurantUpdateService(rest models.Restaurant) error {
	err := s.db.RestaurantUpdate(rest)
	if err != nil {
		return err
	}
	return nil
}

func (s *Services) RestaurantDeleteService(id uuid.UUID) error {
	err := s.db.RestaurantDelete(id)
	if err != nil {
		return err
	}
	return nil
}
