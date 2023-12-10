package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mksmstpck/restoracio/internal/dto"
	log "github.com/sirupsen/logrus"

	"github.com/pborman/uuid"
)

func (d *AdminDatabase) CreateOne(ctx context.Context, admin dto.Admin) (dto.Admin, error) {
	_, err := d.db.
		NewInsert().
		Model(&admin).
		Exec(ctx)
	if err != nil {
		log.Error("database.AdminCreate: ", err)
		return dto.Admin{}, err
	}
	log.Info("admin created")
	return admin, nil
}

func (d *AdminDatabase) GetByID(ctx context.Context, id uuid.UUID) (dto.Admin, error) {
	var admin dto.Admin
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
			return dto.Admin{}, errors.New("admin not found")
		}
		log.Error("database.AdminGetByID: ", err)
		return dto.Admin{}, err
	}
	log.Info("admin found")
	return admin, nil
}

func (d *AdminDatabase) GetByEmail(ctx context.Context, email string) (dto.Admin, error) {
	var admin dto.Admin
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
			return dto.Admin{}, errors.New("admin not found")
		}
		log.Error("database.AdminGetByEmail: ", err)
		return dto.Admin{}, err
	}
	log.Info("admin found")
	return admin, nil
}

func (d *AdminDatabase) GetWithPasswordByID(ctx context.Context, id uuid.UUID) (dto.Admin, error) {
	var admin dto.Admin
	err := d.db.NewSelect().
		Model(&admin).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("admin not found")
			return dto.Admin{}, errors.New("admin not found")
		}
		log.Error("database.AdminGetPasswordById: ", err)
		return dto.Admin{},err
	}
	log.Info("admin found")
	return admin, nil
}

func (d *AdminDatabase) UpdateOne(ctx context.Context, admin dto.Admin) error {
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
		Model(&dto.Admin{}).
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
