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
			log.Error("AdminCreate: ", err)
			return models.Admin{}, err
		}
	}
	if adminExists.ID != "" {
		log.Error("AdminCreate: admin already exists")
		return models.Admin{}, errors.New("admin already exists")
	}
	admin, err = s.db.Admin.CreateOne(s.ctx, admin)
	if err != nil {
		log.Error("AdminCreate: ", err)
		return models.Admin{}, err
	}
	s.cache.Set(uuid.Parse(admin.ID), admin)
	log.Info("admin created")
	return admin, nil
}

func (s *Services) AdminValidateService(admin models.Admin) error {
	adminExists, err := s.db.Admin.GetByEmail(s.ctx, admin.Email)
	if err != nil {
		if err.Error() != "admin not found" {
			log.Error(err)
			return err
		}
	}
	if adminExists.ID != "" {
		log.Error("AdminValidate: admin already exists")
		return errors.New("admin already exists")
	}

	genUUID := uuid.NewUUID()
	err = s.cache.Set(genUUID, admin)
	if err != nil {
		log.Error(err)
		return err
	}

	err = utils.EmailValidator(admin.Email, genUUID)
	if err != nil {
		log.Error(err)
		return err
	}

	log.Info("admin validated")
	return nil
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
		log.Error("AdminGetByEmail: ", err)
		return models.Admin{}, err
	}
	s.cache.Set(uuid.Parse(admin.ID), admin)
	log.Info("admin found")
	return admin, nil
}

func (s Services) AdminGetPasswordByIdService(id uuid.UUID) (string, error) {
	password, err := s.db.Admin.GetPasswordByID(s.ctx, id)
	if err != nil {
		log.Error("AdminGetPasswordById: ", err)
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
		log.Error("AdminUpdate: ", err)
		return err
	}
	log.Info("admin updated")
	return nil
}

func (s Services) AdminDeleteService(id uuid.UUID) error {
	admin, err := s.db.Admin.GetByID(s.ctx, id)
	if err != nil {
		log.Error("AdminDelete: ", err)
		return err
	}
	err = s.db.Admin.DeleteOne(s.ctx, id)
	if err != nil {
		log.Error("AdminDelete: ", err)
		return err
	}
	err = s.db.Rest.DeleteOne(s.ctx, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Error("RestaurantDelete: ", err)
		return err
	}

	s.cache.Delete(uuid.Parse(admin.ID))
	s.cache.Delete(uuid.Parse(admin.Restaurant.ID))

	log.Info("admin deleted")
	return nil
}
