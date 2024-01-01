package services

import (
	"bytes"
	"errors"
	"io"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Services) MenuCreateService(menu dto.Menu, admin dto.Admin) error {
	if admin.Restaurant.Menu != nil {
		return errors.New(models.ErrMenuAlreadyExists)
	}
	if admin.Restaurant == nil {
		return errors.New(models.ErrRestaurantNotFound)
	}

	menu.ID = uuid.NewUUID().String()
	menu.RestaurantID = admin.Restaurant.ID

	menu.ID = uuid.NewUUID().String()
	qrcode, err := utils.QrGenerate("/menu/" + menu.ID)
	if err != nil {
		log.Error(err)
		return err
	}

	qrcodeID, err := s.bucket.UploadFromStream(menu.ID, io.Reader(bytes.NewReader(qrcode)))
	if err != nil {
		log.Error(err)
		return err
	}

	menu.QRCodeID = qrcodeID.Hex()

	err = s.db.Menu.CreateOne(s.ctx, menu)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("menu created")
	return nil
}

func (s *Services) MenuGetWithQrcodeService(id uuid.UUID) (dto.Menu, error) {
	menu, err := s.db.Menu.GetByID(s.ctx, id)
	if err != nil {
		log.Error(err)
		return dto.Menu{}, err
	}
	fileBuffer := bytes.NewBuffer(nil)
	qrID, err := primitive.ObjectIDFromHex(menu.QRCodeID)
	if err != nil {
		log.Error(err)
		return dto.Menu{}, err
	}

	if _, err := s.bucket.DownloadToStream(qrID, fileBuffer); err != nil {
		log.Error(err)
		return dto.Menu{}, err
	}

	menu.QRCode = fileBuffer.Bytes()

	s.cache.Set(uuid.Parse(menu.ID), menu)

	return menu, nil
}

func (s *Services) MenuGetByIDService(id uuid.UUID) (dto.Menu, error) {
	menuAny, err := s.cache.Get(id)
	if menuAny != nil {
		log.Info("menu found")
		return menuAny.(dto.Menu), nil
	}
	if err != nil {
		log.Error(err)
		return dto.Menu{}, err
	}
	menu, err := s.db.Menu.GetByID(s.ctx, id)
	if err != nil {
		log.Error(err)
		return dto.Menu{}, err
	}

	s.cache.Set(uuid.Parse(menu.ID), menu)

	log.Info("menu found")
	return menu, nil
}

func (s *Services) MenuUpdateService(menu dto.Menu, admin dto.Admin) error {
	if admin.Restaurant == nil {
		return errors.New(models.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		return errors.New(models.ErrMenuNotFound)
	}

	menu.RestaurantID = admin.Restaurant.ID
	menu.ID = admin.Restaurant.Menu.ID

	err := s.db.Menu.UpdateOne(s.ctx, menu)
	if err != nil {
		log.Error(err)
		return err
	}

	s.cache.Set(uuid.Parse(menu.ID), menu)

	log.Info("menu updated")
	return nil
}

func (s *Services) MenuDeleteService(admin dto.Admin) error {
	if admin.Restaurant == nil {
		return errors.New(models.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		return errors.New(models.ErrMenuNotFound)
	}

	err := s.db.Menu.DeleteOne(s.ctx, *admin.Restaurant.Menu)
	if err != nil {
		log.Error(err)
		return err
	}
	err = s.DishDeleteAllService(admin)
	if err != nil {
		log.Error(err)
		return err
	}

	s.cache.Delete(uuid.Parse(admin.Restaurant.Menu.ID))

	if err := s.DishDeleteAllService(admin); err != nil {
		log.Error(err)
		return err
	}

	if err := s.bucket.Delete(admin.Restaurant.Menu.QRCodeID); err != nil {
		log.Error(err)
		return err
	}

	log.Info("menu deleted")
	return nil
}
