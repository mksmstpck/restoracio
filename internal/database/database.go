package database

import (
	"github.com/mksmstpck/restoracio/pkg/models"
	"github.com/pborman/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	admin_coll *mongo.Collection
}

func NewDatabase(admin_coll *mongo.Collection) *Database {
	return &Database{admin_coll: admin_coll}
}

type Databases interface {
	// admin
	AdminCreate(user *models.Admin) (*models.Admin, error)
	AdminGetByID(id uuid.UUID) (*models.Admin, error)
	AdminGetByEmail(email string) (*models.Admin, error)
	AdminUpdate(user *models.Admin) error
	AdminDelete(id uuid.UUID) error
}
