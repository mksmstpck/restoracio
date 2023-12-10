package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (d *MenuDatabase) CreateOne(ctx context.Context, menu dto.Menu) (dto.Menu, error) {
	_, err := d.db.
		NewInsert().
		Model(&menu).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return dto.Menu{}, err
	}
	log.Info("menu created")
	return menu, nil
}

func (d *MenuDatabase) GetByID(ctx context.Context, id uuid.UUID) (dto.Menu, error) {
	var menu dto.Menu
	err := d.db.
		NewSelect().
		Model(&menu).
		Relation("Dishes").
		Where("id = ?", id.String()).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("menu not found")
			return dto.Menu{}, errors.New(dto.ErrMenuNotFound)
		}
		log.Error(err)
		return dto.Menu{}, err
	}
	log.Info("menu found")
	return menu, nil
}

func (d *MenuDatabase) UpdateOne(ctx context.Context, menu dto.Menu) error {
	res, err := d.db.
		NewUpdate().
		Model(&menu).
		ExcludeColumn("restaurant_id", "id").
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
		log.Error(dto.ErrMenuNotFound)
		return errors.New(dto.ErrMenuNotFound)
	}
	log.Info("menu updated")
	return nil
}

func (d *MenuDatabase) DeleteOne(ctx context.Context, menu dto.Menu) error {
	res, err := d.db.
		NewDelete().
		Model(&menu).
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
		log.Error(dto.ErrMenuNotFound)
		return errors.New(dto.ErrMenuNotFound)
	}
	log.Info("menu deleted")
	return nil
}