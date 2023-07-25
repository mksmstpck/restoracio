package database

import (
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/mksmstpck/restoracio/pkg/models"
	"github.com/mksmstpck/restoracio/pkg/utils"
	"github.com/pborman/uuid"
)

func (d *Database) AdminCreate(admin *models.Admin) (*models.Admin, error) {
	admin.Password = utils.PasswordHash(admin.Password)
	res, err := d.admin_coll.InsertOne(context.Background(), admin)
	if err != nil {
		log.Error("database.AdminCreate: ", err)
		return nil, err
	}
	admin.ID = res.InsertedID.(primitive.ObjectID)
	log.Info("admin created")
	return admin, nil
}

func (d *Database) AdminGetByID(id uuid.UUID) (*models.Admin, error) {
	admin := &models.Admin{}
	err := d.admin_coll.FindOne(context.Background(), bson.M{"_id": id}).Decode(admin)
	if err != nil {
		log.Error("database.AdminGetByID: ", err)
		return nil, err
	}
	log.Info("admin found")
	return admin, nil
}

func (d *Database) AdminGetByEmail(email string) (*models.Admin, error) {
	admin := &models.Admin{}
	err := d.admin_coll.
		FindOne(context.Background(), bson.M{"email": email}).Decode(admin)
	if err != nil {
		log.Error("database.AdminGetByEmail: ", err)
		return nil, err
	}
	log.Info("admin found")
	return admin, nil
}

func (d *Database) AdminUpdate(admin *models.Admin) error {
	admin.Password = utils.PasswordHash(admin.Password)
	res, err := d.admin_coll.UpdateOne(context.Background(), bson.M{"_id": admin.ID}, bson.M{"$set": admin})
	if err != nil {
		log.Error("database.AdminUpdate: ", err)
		return err
	}
	if res.ModifiedCount == 0 {
		log.Error("admin not found")
		return errors.New("admin not found")
	}
	log.Info("admin updated")
	return nil
}

func (d *Database) AdminDelete(id uuid.UUID) error {
	res, err := d.admin_coll.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		log.Error("database.AdminDelete: ", err)
		return err
	}
	if res.DeletedCount == 0 {
		log.Error("admin not found")
		return errors.New("admin not found")
	}
	log.Info("admin deleted")
	return nil
}
