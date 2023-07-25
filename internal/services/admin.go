package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/pkg/models"
	log "github.com/sirupsen/logrus"
)

func (s Services) AdminCreateService(admin *models.Admin) (*models.Admin, error) {
	admin, err := s.db.AdminGetByEmail(admin.Email)
	if err != nil {
		log.Error("AdminCreate: ", err)
		return nil, err
	}
	if admin != nil {
		log.Error("AdminCreate: admin with this email already exists")
		return nil, errors.New("admin with this email already exists")
	}
	admin, err = s.db.AdminCreate(admin)
	if err != nil {
		log.Error("AdminCreate: ", err)
		return nil, err
	}
	log.Info("admin created")
	return admin, nil
}
