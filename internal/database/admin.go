package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mksmstpck/restoracio/internal/convertors"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/internal/models"
	log "github.com/sirupsen/logrus"

	"github.com/pborman/uuid"
)

func (d *AdminDatabase) CreateOne(ctx context.Context, admin dto.Admin) error {
	adminDB := convertors.AdminDTOToDB(&admin)
	_, err := d.db.
		NewInsert().
		Model(&adminDB).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("admin created")
	return nil
}

func (d *AdminDatabase) GetByID(ctx context.Context, id uuid.UUID) (dto.Admin, error) {
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
			log.Error(models.ErrAdminNotFound)
			return dto.Admin{}, errors.New(models.ErrAdminNotFound)
		}
		log.Error(err)
		return dto.Admin{}, err
	}
	log.Info("admin found")
	return convertors.AdminDBToDTO(&admin), nil
}

func (d *AdminDatabase) GetByEmail(ctx context.Context, email string) (dto.Admin, error) {
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
			log.Error(models.ErrAdminNotFound)
			return dto.Admin{}, errors.New(models.ErrAdminNotFound)
		}
		log.Error(err)
		return dto.Admin{}, err
	}
	log.Info("admin found")
	return convertors.AdminDBToDTO(&admin), nil
}

func (d *AdminDatabase) GetWithPasswordByID(ctx context.Context, id uuid.UUID) (dto.Admin, error) {
	var admin models.Admin
	err := d.db.NewSelect().
		Model(&admin).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(models.ErrAdminNotFound)
			return dto.Admin{}, errors.New(models.ErrAdminNotFound)
		}
		log.Error(err)
		return dto.Admin{}, err
	}
	log.Info("admin found")
	return convertors.AdminDBToDTO(&admin), nil
}

func (d *AdminDatabase) UpdateOne(ctx context.Context, admin dto.Admin) error {
	adminDB := convertors.AdminDTOToDB(&admin)
	res, err := d.db.
		NewUpdate().
		Model(&adminDB).
		ExcludeColumn("id", "password", "restaurant").
		Where("id = ?", admin.ID).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return err
	}
	if count == 0 {
		log.Error(models.ErrAdminNotFound)
		return errors.New(models.ErrAdminNotFound)
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
		log.Error(err)
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return err
	}
	if count == 0 {
		log.Error(models.ErrAdminNotFound)
		return errors.New(models.ErrAdminNotFound)
	}
	log.Info("admin deleted")
	return nil
}
