package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Services) DishCreateService(dish models.Dish, admin models.Admin) (models.Dish, error) {
	if admin.Restaurant == nil {
		log.Info(models.ErrRestaurantNotFound)
		return models.Dish{}, errors.New(models.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		log.Info(models.ErrMenuNotFound)
		return models.Dish{}, errors.New(models.ErrMenuNotFound)
	}

	dish.ID = uuid.NewUUID().String()
	dish.MenuID = admin.Restaurant.Menu.ID

	dish, err := s.db.Dish.CreateOne(s.ctx, dish)
	if err != nil {
		log.Error(err)
		return models.Dish{}, err
	}
	s.cache.Set(uuid.Parse(dish.ID), dish)

	log.Info("dish created")
	return dish, nil
}

func (s *Services) DishGetByIDService(id uuid.UUID) (models.Dish, error) {
	dish, err := s.cache.DishGet(id)
	if dish.ID != "" {
		log.Info("dish found")
		return dish, nil
	}
	if err != nil {
		log.Error(err)
		return models.Dish{}, err
	}
	dish, err = s.db.Dish.GetByID(s.ctx, id)
	if err != nil {
		log.Error(err)
		return models.Dish{}, err
	}

	s.cache.Set(uuid.Parse(dish.ID), dish)

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
		log.Info(models.ErrRestaurantNotFound)
		return errors.New(models.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		log.Info(models.ErrMenuNotFound)
		return errors.New(models.ErrMenuNotFound)
	}

	dish.MenuID = admin.Restaurant.Menu.ID

	err := s.db.Dish.UpdateOne(s.ctx, dish)
	if err != nil {
		log.Error(err)
		return err
	}

	s.cache.Set(uuid.Parse(dish.ID), dish)

	log.Info("dish updated")
	return nil
}

func (s *Services) DishDeleteService(id uuid.UUID, admin models.Admin) error {
	if admin.Restaurant == nil {
		log.Info(models.ErrRestaurantNotFound)
		return errors.New(models.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		log.Info(models.ErrMenuNotFound)
		return errors.New(models.ErrMenuNotFound)
	}

	err := s.db.Dish.DeleteOne(s.ctx, id, uuid.Parse(admin.Restaurant.Menu.ID))
	if err != nil {
		log.Error(err)
		return err
	}

	s.cache.Delete(id)

	log.Info("dish deleted")
	return nil
}

func (s *Services) DishDeleteAllService(admin models.Admin) error {
	if admin.Restaurant == nil {
		log.Info(models.ErrRestaurantNotFound)
		return errors.New(models.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		log.Info(models.ErrMenuNotFound)
		return errors.New(models.ErrMenuNotFound)
	}
	err := s.db.Dish.DeleteAll(s.ctx, uuid.Parse(admin.Restaurant.Menu.ID))
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("dishes deleted")
	return nil
}