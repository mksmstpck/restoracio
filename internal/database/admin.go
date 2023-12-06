package database

import (
	"context"
	"database/sql"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/pborman/uuid"
)

func (d *AdminDatabase) CreateOne(ctx context.Context, admin models.Admin) (models.Admin, error) {
	admin.Password, admin.Salt = utils.PasswordHash(admin.Password)
	_, err := d.db.
		NewInsert().
		Model(&admin).
		Exec(ctx)
	if err != nil {
		log.Error("database.AdminCreate: ", err)
		return models.Admin{}, err
	}
	log.Info("admin created")
	return admin, nil
}

func (d *AdminDatabase) GetByID(ctx context.Context, id uuid.UUID) (models.Admin, error) {
	var admin models.Admin
	err := d.db.
		NewSelect().
		Model(&admin).
		ExcludeColumn("password", "salt").
		Relation("Restaurant").
		Relation("Restaurant.Tables").
		Relation("Restaurant.Menu").
		Where("admin.id = ?", id.String()).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("admin not found")
			return models.Admin{}, errors.New("admin not found")
		}
		log.Error("database.AdminGetByID: ", err)
		return models.Admin{}, err
	}
	log.Info("admin found")
	return admin, nil
}

func (d *AdminDatabase) GetByEmail(ctx context.Context, email string) (models.Admin, error) {
	var admin models.Admin
	err := d.db.NewSelect().
		Model(&admin).
		ExcludeColumn("password", "salt").
		Relation("Restaurant").
		Relation("Restaurant.Tables").
		Where("admin.email = ?", email).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("admin not found")
			return models.Admin{}, errors.New("admin not found")
		}
		log.Error("database.AdminGetByEmail: ", err)
		return models.Admin{}, err
	}
	log.Info("admin found")
	return admin, nil
}

func (d *AdminDatabase) GetWithPasswordByID(ctx context.Context, id uuid.UUID) (models.Admin, error) {
	var admin models.Admin
	err := d.db.NewSelect().
		Model(&admin).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("admin not found")
			return models.Admin{}, errors.New("admin not found")
		}
		log.Error("database.AdminGetPasswordById: ", err)
		return models.Admin{}, err
	}
	log.Info("admin found")
	return admin, nil
}

func (d *AdminDatabase) UpdateOne(ctx context.Context, admin models.Admin) error {
	res, err := d.db.
		NewUpdate().
		Model(&admin).
		ExcludeColumn("id", "password").
		Where("id = ?", admin.ID).
		Exec(ctx)
	if err != nil {
		log.Error("database.AdminUpdate: ", err)
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Error("database.AdminUpdate: ", err)
		return err
	}
	if count == 0 {
		log.Error("admin not found")
		return errors.New("admin not found")
	}
	log.Info("admin updated")
	return nil
}

func (d *AdminDatabase) DeleteOne(ctx context.Context, id uuid.UUID) error {
	res, err := d.db.
		NewDelete().
		Model(&models.Admin{}).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		log.Error("database.AdminDelete: ", err)
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Error("database.AdminDelete: ", err)
		return err
	}
	if count == 0 {
		log.Error("admin not found")
		return errors.New("admin not found")
	}
	log.Info("admin deleted")
	return nil
}
