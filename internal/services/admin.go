package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/pkg/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s Services) AdminCreateService(admin models.Admin) (models.Admin, error) {
	adminExists, err := s.admindb.AdminGetByEmail(admin.Email)
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
	admin, err = s.admindb.AdminCreate(admin)
	if err != nil {
		log.Error("AdminCreate: ", err)
		return models.Admin{}, err
	}
	log.Info("admin created")
	return admin, nil
}

func (s Services) AdminGetByIDService(id uuid.UUID) (models.Admin, error) {
	admin, err := s.admindb.AdminGetByID(id)
	if err != nil {
		log.Error("AdminGetByID: ", err)
		return models.Admin{}, err
	}
	log.Info("admin found")
	return admin, nil
}

func (s Services) AdminGetByEmailService(email string) (models.Admin, error) {
	admin, err := s.admindb.AdminGetByEmail(email)
	if err != nil {
		log.Error("AdminGetByEmail: ", err)
		return models.Admin{}, err
	}
	log.Info("admin found")
	return admin, nil
}

func (s Services) AdminGetPasswordByIdService(id uuid.UUID) (string, error) {
	password, err := s.admindb.AdminGetPasswordByID(id)
	if err != nil {
		log.Error("AdminGetPasswordById: ", err)
		return "", err
	}
	log.Info("password found")
	return password, nil
}

func (s Services) AdminUpdateService(admin models.Admin) error {
	err := s.admindb.AdminUpdate(admin)
	if err != nil {
		log.Error("AdminUpdate: ", err)
		return err
	}
	log.Info("admin updated")
	return nil
}

func (s Services) AdminDeleteService(id uuid.UUID) error {
	admin, err := s.admindb.AdminGetByID(id)
	if err != nil {
		log.Error("AdminDelete: ", err)
		return err
	}
	err = s.admindb.AdminDelete(id)
	if err != nil {
		log.Error("AdminDelete: ", err)
		return err
	}
	err = s.restdb.RestaurantDelete(uuid.Parse(admin.RestaurantID))
	log.Info("admin deleted")
	return nil
}
