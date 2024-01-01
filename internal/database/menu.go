package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mksmstpck/restoracio/internal/convertors"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (d *MenuDatabase) CreateOne(ctx context.Context, menu dto.Menu) error {
	menuDB := convertors.MenuDTOToDB(&menu)
	_, err := d.db.
		NewInsert().
		Model(menuDB).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("menu created")
	return nil
}

func (d *MenuDatabase) GetByID(ctx context.Context, id uuid.UUID) (dto.Menu, error) {
	var menu models.Menu
	err := d.db.
		NewSelect().
		Model(&menu).
		Relation("Dishes").
		Where("id = ?", id.String()).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(models.ErrMenuNotFound)
			return dto.Menu{}, errors.New(models.ErrMenuNotFound)
		}
		log.Error(err)
		return dto.Menu{}, err
	}
	log.Info("menu found")
	return convertors.MenuDBToDTO(&menu), nil
}

func (d *MenuDatabase) UpdateOne(ctx context.Context, menu dto.Menu) error {
	menuDB := convertors.MenuDTOToDB(&menu)
	res, err := d.db.
		NewUpdate().
		Model(&menuDB).
		ExcludeColumn("restaurant_id", "id", "qrcode").
		Where("id = ?", menu.ID).
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
		log.Error(models.ErrMenuNotFound)
		return errors.New(models.ErrMenuNotFound)
	}
	log.Info("menu updated")
	return nil
}

func (d *MenuDatabase) DeleteOne(ctx context.Context, menu dto.Menu) error {
	menuDB := convertors.MenuDTOToDB(&menu)
	res, err := d.db.
		NewDelete().
		Model(&menuDB).
		Where("id = ?", menu.ID).
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
		log.Error(models.ErrMenuNotFound)
		return errors.New(models.ErrMenuNotFound)
	}
	log.Info("menu deleted")
	return nil
}
