package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/pkg/models"
	log "github.com/sirupsen/logrus"
)

func (s Services) AdminCreateService(admin models.Admin) (models.Admin, error) {
	adminExists, err := s.db.AdminGetByEmail(admin.Email)
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
	admin, err = s.db.AdminCreate(admin)
	if err != nil {
		log.Error("AdminCreate: ", err)
		return models.Admin{}, err
	}
	log.Info("admin created")
	return admin, nil
}
