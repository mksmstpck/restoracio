package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s Services) AdminCreateService(admin dto.Admin) error {
	adminExists, err := s.db.Admin.GetByEmail(s.ctx, admin.Email)
	if err != nil {
		if err.Error() != models.ErrAdminNotFound {
			log.Error(err)
			return err
		}
	}
	if adminExists.ID != "" {
		log.Error(models.ErrAdminAlreadyExists)
		return errors.New(models.ErrAdminAlreadyExists)
	}

	admin.ID = uuid.NewUUID().String()
	admin.PasswordHash, admin.Salt = utils.PasswordHash(admin.Password)

	err = s.db.Admin.CreateOne(s.ctx, admin)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("admin created")
	return nil
}

func (s Services) AdminGetByIDService(id uuid.UUID) (dto.Admin, error) {
	adminAny, err := s.cache.Get(id)
	if adminAny != nil {
		log.Info("admin found")
		return adminAny.(dto.Admin), nil
	}
	if err != nil {
		log.Error(err)
		return dto.Admin{}, err
	}
	admin, err := s.db.Admin.GetByID(s.ctx, id)
	if err != nil {
		if err.Error() == models.ErrAdminNotFound {
			log.Error(models.ErrAdminNotFound)
			return dto.Admin{}, errors.New(models.ErrAdminNotFound)
		}
		log.Error(err)
		return dto.Admin{}, err
	}

	s.cache.Set(uuid.Parse(admin.ID), admin)
	log.Info("admin found")
	return admin, nil
}

func (s Services) AdminGetByEmailService(email string) (dto.Admin, error) {
	admin, err := s.db.Admin.GetByEmail(s.ctx, email)
	if err != nil {
		if err.Error() == models.ErrAdminNotFound {
			log.Error(models.ErrAdminNotFound)
			return dto.Admin{}, errors.New(models.ErrAdminNotFound)
		}
		log.Error(err)
		return dto.Admin{}, err
	}

	s.cache.Set(uuid.Parse(admin.ID), admin)
	log.Info("admin found")
	return admin, nil
}

func (s Services) AdminGetWithPasswordByIdService(id uuid.UUID) (dto.Admin, error) {
	admin, err := s.db.Admin.GetWithPasswordByID(s.ctx, id)
	if err != nil {
		if err.Error() == models.ErrAdminNotFound {
			log.Error(models.ErrAdminNotFound)
			return dto.Admin{}, errors.New(models.ErrAdminNotFound)
		}
		log.Error(err)
		return dto.Admin{}, err
	}

	log.Info("password found")
	return admin, nil
}

func (s Services) AdminUpdateService(admin dto.Admin, adminID uuid.UUID) error {
	admin.ID = adminID.String()
	err := s.db.Admin.UpdateOne(s.ctx, admin)
	if err != nil {
		if err.Error() == models.ErrAdminNotFound {
			log.Error(models.ErrAdminNotFound)
			return errors.New(models.ErrAdminNotFound)
		}
		log.Error(err)
		return err
	}
	s.cache.Set(uuid.Parse(admin.ID), admin)
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
