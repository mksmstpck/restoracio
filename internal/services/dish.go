package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
	"github.com/mksmstpck/restoracio/utils/convertors"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Services) DishCreateService(dishRequest dto.DishRequest, admin dto.AdminResponse) error {
	if admin.Restaurant == nil {
		log.Info(models.ErrRestaurantNotFound)
		return errors.New(models.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		log.Info(models.ErrMenuNotFound)
		return errors.New(models.ErrMenuNotFound)
	}

	dishDB := convertors.DishRequestToDB(dishRequest, admin.Restaurant.Menu.ID)

	dish, err := s.db.Dish.CreateOne(s.ctx, dishDB)
	if err != nil {
		log.Error(err)
		return err
	}
	s.cache.Set(uuid.Parse(dish.ID), convertors.DishDBToResponse(dishDB))

	log.Info("dish created")
	return nil
}

func (s *Services) DishGetByIDService(id uuid.UUID) (*dto.DishResponse, error) {
	dishAny, err := s.cache.Get(id)
	if dishAny != nil {
		log.Info("dish found")
		return dishAny.(*dto.DishResponse), nil
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	dish, err := s.db.Dish.GetByID(s.ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	dishResponse := convertors.DishDBToResponse(dish)

	s.cache.Set(uuid.Parse(dish.ID), dishResponse)

	log.Info("dish found")
	return &dishResponse, nil
}

func (s *Services) DishGetAllInMenuService(id uuid.UUID) ([]dto.DishResponse, error) {
	dishes, err := s.db.Dish.GetAllInMenu(s.ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("dishes found")
	return convertors.DishDBSliceToResponseSlice(dishes), nil
}

func (s *Services) DishUpdateService(dish dto.DishRequest, dishID uuid.UUID, admin dto.AdminResponse) error {
	if admin.Restaurant == nil {
		log.Info(models.ErrRestaurantNotFound)
		return errors.New(models.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		log.Info(models.ErrMenuNotFound)
		return errors.New(models.ErrMenuNotFound)
	}

	dishDB := convertors.DishRequestToDB(dish, admin.Restaurant.Menu.ID)
	dishDB.ID = dishID.String()

	err := s.db.Dish.UpdateOne(s.ctx, dishDB)
	if err != nil {
		log.Error(err)
		return err
	}

	s.cache.Set(uuid.Parse(dishDB.ID), convertors.DishDBToResponse(dishDB))

	log.Info("dish updated")
	return nil
}

func (s *Services) DishDeleteService(id uuid.UUID, admin dto.AdminResponse) error {
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

func (s *Services) DishDeleteAllService(admin dto.AdminResponse) error {
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