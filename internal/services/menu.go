package services

import (
	"bytes"
	"errors"
	"io"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/mksmstpck/restoracio/utils/convertors"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s Services) MenuCreateService(MenuRequest dto.MenuRequest, admin dto.AdminResponse) error {
	if admin.Restaurant.Menu != nil {
		return errors.New(models.ErrMenuAlreadyExists)
	}

	if admin.Restaurant == nil {
		return errors.New(models.ErrRestaurantNotFound)
	}

	menuDB := convertors.MenuRequestToDB(MenuRequest, admin.Restaurant.ID)
	
	qrcode, err := utils.QrGenerate("/menu/"+menuDB.ID)
	if err != nil {
		log.Error(err)
		return err
	}

	qrcodeID, err := s.bucket.UploadFromStream(menuDB.ID, io.Reader(bytes.NewReader(qrcode)))
	if err != nil {
		log.Error(err)
		return err
	}

	menuDB.QRCodeID = qrcodeID.Hex()

	_, err = s.db.Menu.CreateOne(s.ctx, menuDB)
	if err != nil {
		log.Error(err)
		return err
	}

	s.cache.Set(uuid.Parse(menuDB.ID), convertors.MenuDBToResponse(menuDB))

	log.Info("menu created")
	return nil
}

func (s *Services) MenuGetWithQrcodeService(id uuid.UUID) (*dto.MenuResponse, error) {
	menu, err := s.db.Menu.GetByID(s.ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	fileBuffer := bytes.NewBuffer(nil)
	qrID, err := primitive.ObjectIDFromHex(menu.QRCodeID)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if _, err := s.bucket.DownloadToStream(qrID, fileBuffer); err != nil {
		log.Error(err)
		return nil, err
	}

	menuResponce := convertors.MenuDBToResponse(menu)

	menuResponce.QRCodeBytes = fileBuffer.Bytes()

	s.cache.Set(uuid.Parse(menu.ID), menu)

	return &menuResponce, nil
}

func (s *Services) MenuGetByIDService(id uuid.UUID) (*dto.MenuResponse, error) {
	menuAny, err := s.cache.Get(id)
	if menuAny != nil {
		log.Info("menu found")
		return menuAny.(*dto.MenuResponse), nil
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	menu, err := s.db.Menu.GetByID(s.ctx, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	menuResponce := convertors.MenuDBToResponse(menu)

	s.cache.Set(uuid.Parse(menu.ID), menuResponce)

	log.Info("menu found")
	return &menuResponce, nil
}

func (s *Services) MenuUpdateService(menu dto.MenuRequest, admin dto.AdminResponse) error {
	if admin.Restaurant == nil {
		return errors.New(models.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Menu == nil {
		return errors.New(models.ErrMenuNotFound)
	}

	menuDB := convertors.MenuRequestToDB(menu, admin.Restaurant.Menu.ID)

	menuDB.RestaurantID = admin.Restaurant.ID
	menuDB.ID = admin.Restaurant.Menu.ID

	err := s.db.Menu.UpdateOne(s.ctx, menuDB)
	if err != nil {
		log.Error(err)
		return err
	}

	s.cache.Set(uuid.Parse(menuDB.ID), convertors.MenuDBToResponse(menuDB))

	log.Info("menu updated")
	return nil
}

func (s *Services) MenuDeleteService(admin dto.AdminResponse) error {
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