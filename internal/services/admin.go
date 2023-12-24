package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
	"github.com/mksmstpck/restoracio/utils/convertors"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)


func (s Services) AdminCreateService(admin dto.AdminRequest) error {
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

	adminDB := convertors.AdminRequestToDB(admin)

	_, err = s.db.Admin.CreateOne(s.ctx, adminDB)
	if err != nil {
		log.Error(err)
		return err
	}

	s.cache.Set(uuid.Parse(adminDB.ID), convertors.AdminDBToResponse(adminDB))
	log.Info("admin created")
	return nil
}

func (s Services) AdminGetByIDService(id uuid.UUID) (*dto.AdminResponse, error) {
	adminAny, err := s.cache.Get(id)
	if adminAny != nil {
		log.Info("admin found")
		return adminAny.(*dto.AdminResponse), nil
	}
	if err != nil {
		log.Error(err)
		return nil, err
	}
	adminDB, err := s.db.Admin.GetByID(s.ctx, id)
	if err != nil {
		if err.Error() == models.ErrAdminNotFound {
			log.Error(models.ErrAdminNotFound)
			return nil, errors.New(models.ErrAdminNotFound)
		}
		log.Error(err)
		return nil, err
	}

	adminResponse := convertors.AdminDBToResponse(adminDB)

	s.cache.Set(uuid.Parse(adminResponse.ID), adminResponse)
	log.Info("admin found")
	return &adminResponse, nil
}

func (s Services) AdminGetByEmailService(email string) (*dto.AdminResponse, error) {
	adminDB, err := s.db.Admin.GetByEmail(s.ctx, email)
	if err != nil {
		if err.Error() == models.ErrAdminNotFound {
			log.Error(models.ErrAdminNotFound)
			return nil, errors.New(models.ErrAdminNotFound)
		}
		log.Error(err)
		return nil, err
	}

	adminResponse := convertors.AdminDBToResponse(adminDB)

	s.cache.Set(uuid.Parse(adminResponse.ID), adminResponse)
	log.Info("admin found")
	return &adminResponse, nil
}

func (s Services) AdminGetWithPasswordByIdService(id uuid.UUID) (*dto.AdminResponseWishPass, error) {
	adminDB, err := s.db.Admin.GetWithPasswordByID(s.ctx, id)
	if err != nil {
		if err.Error() == models.ErrAdminNotFound {
			log.Error(models.ErrAdminNotFound)
			return nil, errors.New(models.ErrAdminNotFound)
		}
		log.Error(err)
		return nil, err
	}

	adminResponse := convertors.AdminDBToResponseWishPass(adminDB)

	log.Info("password found")
	return &adminResponse,nil
}

func (s Services) AdminUpdateService(AdminRequest dto.AdminRequest, adminID uuid.UUID) error {
	adminDB := convertors.AdminRequestToDB(AdminRequest)

	adminDB.ID = adminID.String()
	s.cache.Set(uuid.Parse(adminDB.ID), convertors.AdminDBToResponse(adminDB))
	err := s.db.Admin.UpdateOne(s.ctx, adminDB)
	if err != nil {
		if err.Error() == models.ErrAdminNotFound {
			log.Error(models.ErrAdminNotFound)
			return errors.New(models.ErrAdminNotFound)
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
