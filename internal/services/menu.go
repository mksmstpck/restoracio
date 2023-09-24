package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s Services) MenuCreateService(menu models.Menu, admin models.Admin) (models.Menu, error) {
	if admin.Restaurant.Menu != nil {
		return models.Menu{}, errors.New(utils.ErrMenuAlreadyExists)
	}

	if admin.Restaurant == nil {
		return models.Menu{}, errors.New(utils.ErrRestaurantNotFound)
	}
	menu.RestaurantID = admin.Restaurant.ID
	
	qrcode, err := utils.QrGenerate("menu/${menu.ID}")
	if err != nil {
		log.Error(err)
		return models.Menu{}, err
	}

	menu.QRCode = qrcode

	menu, err = s.db.Menu.CreateOne(s.ctx, menu)
	if err != nil {
		log.Error(err)
		return models.Menu{}, err
	}

	s.cache.Set(uuid.Parse(menu.ID), menu)

	log.Info("menu created")
	return menu, nil
}

func (s *Services) MenuGetByIDService(id uuid.UUID) (models.Menu, error) {
	menu, err := s.cache.MenuGet(id)
	if menu.ID != "" {
		log.Info("menu found")
		return menu, nil
	}
	if err != nil {
		log.Error(err)
		return models.Menu{}, err
	}
	menu, err = s.db.Menu.GetByID(s.ctx, id)
	if err != nil {
		log.Error(err)
		return models.Menu{}, err
	}

	s.cache.Set(uuid.Parse(menu.ID), menu)

	log.Info("menu found")
	return menu, nil
}

func (s *Services) MenuUpdateService(menu models.Menu, admin models.Admin) error {
	if admin.Restaurant == nil {
		return errors.New(utils.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		return errors.New(utils.ErrMenuNotFound)
	}

	menu.RestaurantID = admin.Restaurant.ID

	err := s.db.Menu.UpdateOne(s.ctx, menu)
	if err != nil {
		log.Error(err)
		return err
	}

	s.cache.Set(uuid.Parse(menu.ID), menu)

	log.Info("menu updated")
	return nil
}

func (s *Services) MenuDeleteService(admin models.Admin) error {
	if admin.Restaurant == nil {
		return errors.New(utils.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		return errors.New(utils.ErrMenuNotFound)
	}

	err := s.db.Menu.DeleteOne(s.ctx, *admin.Restaurant.Menu)
	if err != nil {
		log.Error(err)
		return err
	}

	s.cache.Delete(uuid.Parse(admin.Restaurant.Menu.ID))

	log.Info("menu deleted")
	return nil
}