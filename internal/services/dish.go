package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Services) DishCreateService(dish dto.Dish, admin dto.Admin) (dto.Dish, error) {
	if admin.Restaurant == nil {
		log.Info(dto.ErrRestaurantNotFound)
		return dto.Dish{}, errors.New(dto.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		log.Info(dto.ErrMenuNotFound)
		return dto.Dish{}, errors.New(dto.ErrMenuNotFound)
	}

	dish.ID = uuid.NewUUID().String()
	dish.MenuID = admin.Restaurant.Menu.ID

	dish, err := s.db.Dish.CreateOne(s.ctx, dish)
	if err != nil {
		log.Error(err)
		return dto.Dish{}, err
	}
	s.cache.Set(uuid.Parse(dish.ID), dish)

	log.Info("dish created")
	return dish, nil
}

func (s *Services) DishGetByIDService(id uuid.UUID) (dto.Dish, error) {
	dishAny, err := s.cache.Get(id)
	if dishAny != nil {
		log.Info("dish found")
		return dishAny.(dto.Dish), nil
	}
	if err != nil {
		log.Error(err)
		return dto.Dish{}, err
	}
	dish, err := s.db.Dish.GetByID(s.ctx, id)
	if err != nil {
		log.Error(err)
		return dto.Dish{}, err
	}

	s.cache.Set(uuid.Parse(dish.ID), dish)

	log.Info("dish found")
	return dish, nil
}

func (s *Services) DishGetAllInMenuService(id uuid.UUID) ([]dto.Dish, error) {
	dishes, err := s.db.Dish.GetAllInMenu(s.ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("dishes found")
	return dishes, nil
}

func (s *Services) DishUpdateService(dish dto.Dish, admin dto.Admin) error {
	if admin.Restaurant == nil {
		log.Info(dto.ErrRestaurantNotFound)
		return errors.New(dto.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		log.Info(dto.ErrMenuNotFound)
		return errors.New(dto.ErrMenuNotFound)
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

func (s *Services) DishDeleteService(id uuid.UUID, admin dto.Admin) error {
	if admin.Restaurant == nil {
		log.Info(dto.ErrRestaurantNotFound)
		return errors.New(dto.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		log.Info(dto.ErrMenuNotFound)
		return errors.New(dto.ErrMenuNotFound)
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

func (s *Services) DishDeleteAllService(admin dto.Admin) error {
	if admin.Restaurant == nil {
		log.Info(dto.ErrRestaurantNotFound)
		return errors.New(dto.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		log.Info(dto.ErrMenuNotFound)
		return errors.New(dto.ErrMenuNotFound)
	}
	err := s.db.Dish.DeleteAll(s.ctx, uuid.Parse(admin.Restaurant.Menu.ID))
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("dishes deleted")
	return nil
}