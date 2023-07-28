package services

import (
	"github.com/mksmstpck/restoracio/internal/database"
	"github.com/mksmstpck/restoracio/pkg/models"
)

type Services struct {
	db database.Databases
}

func NewServices(db *database.Database) *Services {
	return &Services{
		db: db,
	}
}

type Servicer interface {
	AdminCreateService(admin models.Admin) (models.Admin, error)
}
