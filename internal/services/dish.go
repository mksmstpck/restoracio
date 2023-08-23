package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Services) DishCreateService(dish models.Dish, admin models.Admin) (models.Dish, error) {
	if admin.Restaurant == nil {
		log.Info(utils.ErrRestaurantNotFound)
		return models.Dish{}, errors.New(utils.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		log.Info(utils.ErrMenuNotFound)
		return models.Dish{}, errors.New(utils.ErrMenuNotFound)
	}

	dish.MenuID = admin.Restaurant.Menu.ID

	dish, err := s.db.Dish.CreateOne(s.ctx, dish)
	if err != nil {
		log.Error(err)
		return models.Dish{}, err
	}
	log.Info("dish created")
	return dish, nil
}

func (s *Services) DishGetByIDService(id uuid.UUID) (models.Dish, error) {
	dish, err := s.db.Dish.GetByID(s.ctx, id)
	if err != nil {
		log.Error(err)
		return models.Dish{}, err
	}
	log.Info("dish found")
	return dish, nil
}

func (s *Services) DishGetAllInMenuService(id uuid.UUID) ([]models.Dish, error) {
	dishes, err := s.db.Dish.GetAllInMenu(s.ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("dishes found")
	return dishes, nil
}

func (s *Services) DishUpdateService(dish models.Dish, admin models.Admin) error {
	if admin.Restaurant == nil {
		log.Info(utils.ErrRestaurantNotFound)
		return errors.New(utils.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		log.Info(utils.ErrMenuNotFound)
		return errors.New(utils.ErrMenuNotFound)
	}

	dish.MenuID = admin.Restaurant.Menu.ID

	err := s.db.Dish.UpdateOne(s.ctx, dish)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("dish updated")
	return nil
}

func (s *Services) DishDeleteService(id uuid.UUID, admin models.Admin) error {
	if admin.Restaurant == nil {
		log.Info(utils.ErrRestaurantNotFound)
		return errors.New(utils.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		log.Info(utils.ErrMenuNotFound)
		return errors.New(utils.ErrMenuNotFound)
	}

	err := s.db.Dish.DeleteOne(s.ctx, uuid.UUID(id), uuid.UUID(admin.Restaurant.Menu.ID))
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("dish deleted")
	return nil
}