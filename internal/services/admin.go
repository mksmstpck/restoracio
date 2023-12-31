package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)


func (s Services) AdminCreateService(admin models.Admin) (models.Admin, error) {
	adminExists, err := s.db.Admin.GetByEmail(s.ctx, admin.Email)
	if err != nil {
		if err.Error() != "admin not found" {
			log.Error(err)
			return models.Admin{}, err
		}
	}
	if adminExists.ID != "" {
		log.Error(utils.ErrAdminAlreadyExists)
		return models.Admin{}, errors.New(utils.ErrAdminAlreadyExists)
	}
	admin.ID = uuid.NewUUID().String()
	admin, err = s.db.Admin.CreateOne(s.ctx, admin)
	if err != nil {
		log.Error(err)
		return models.Admin{}, err
	}
	s.cache.Set(uuid.Parse(admin.ID), admin)
	log.Info("admin created")
	return admin, nil
}

func (s Services) AdminGetByIDService(id uuid.UUID) (models.Admin, error) {
	admin, err := s.cache.AdminGet(id)
	if admin.ID != "" {
		log.Info("admin found")
		return admin, nil
	}
	if err != nil {
		log.Error(err)
		return models.Admin{}, err
	}
	admin, err = s.db.Admin.GetByID(s.ctx, id)
	if err != nil {
		if err.Error() == "admin not found" {
			log.Error(utils.ErrAdminNotFound)
			return models.Admin{}, errors.New(utils.ErrAdminNotFound)
		}
		log.Error(err)
		return models.Admin{}, err
	}
	s.cache.Set(uuid.Parse(admin.ID), admin)
	log.Info("admin found")
	return admin, nil
}

func (s Services) AdminGetByEmailService(email string) (models.Admin, error) {
	admin, err := s.db.Admin.GetByEmail(s.ctx, email)
	if err != nil {
		if err.Error() == "admin not found" {
			log.Error(utils.ErrAdminNotFound)
			return models.Admin{}, errors.New(utils.ErrAdminNotFound)
		}
		log.Error(err)
		return models.Admin{}, err
	}
	s.cache.Set(uuid.Parse(admin.ID), admin)
	log.Info("admin found")
	return admin, nil
}

func (s Services) AdminGetPasswordByIdService(id uuid.UUID) (string, error) {
	password, err := s.db.Admin.GetPasswordByID(s.ctx, id)
	if err != nil {
		if err.Error() == "admin not found" {
			log.Error(utils.ErrAdminNotFound)
			return "", errors.New(utils.ErrAdminNotFound)
		}
		log.Error(err)
		return "", err
	}
	log.Info("password found")
	return password, nil
}

func (s Services) AdminUpdateService(admin models.Admin, adminID uuid.UUID) error {
	admin.ID = adminID.String()
	s.cache.Set(uuid.Parse(admin.ID), admin)
	err := s.db.Admin.UpdateOne(s.ctx, admin)
	if err != nil {
		if err.Error() == "admin not found" {
			log.Error(utils.ErrAdminNotFound)
			return errors.New(utils.ErrAdminNotFound)
		}
		log.Error(err)
		return err
	}
	log.Info("admin updated")
	return nil
}

func (s Services) AdminDeleteService(id uuid.UUID) error {
	admin, err := s.db.Admin.GetByID(s.ctx, id)
	if err != nil {
		log.Error(err)
		return err
	}
	err = s.db.Admin.DeleteOne(s.ctx, id)
	if err != nil {
		log.Error(err)
		return err
	}
	err = s.RestaurantDeleteService(admin.Restaurant)
	if err != nil {
		log.Error(err)
		return err
	}

	s.cache.Delete(uuid.Parse(admin.ID))

	log.Info("admin deleted")
	return nil
}
