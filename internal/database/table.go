package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (d *TableDatabase) CreateOne(ctx context.Context, table dto.TableDB) (dto.TableDB, error) {
	_, err := d.db.
		NewInsert().
		Model(&table).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return dto.TableDB{}, err
	}
	log.Info("table created")
	return table, nil
}

func (d *TableDatabase) GetByID(ctx context.Context, id uuid.UUID) (dto.TableDB, error) {
	var table dto.TableDB
	err := d.db.
		NewSelect().
		Model(&table).
		Where("id = ?", id.String()).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(models.ErrTableNotFound)
			return dto.TableDB{}, errors.New(models.ErrTableNotFound)
		}
		log.Error(err)
		return dto.TableDB{}, err
	}
	log.Info("table found")
	return table, nil
}

func (d *TableDatabase) GetAllInRestaurant(ctx context.Context, id uuid.UUID) ([]dto.TableDB, error) {
	var tables []dto.TableDB
	err := d.db.
		NewSelect().
		Model(&tables).
		Where("restaurant_id = ?", id.String()).
		Scan(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("tables found")
	return tables, nil
}

func (d *TableDatabase) UpdateOne(ctx context.Context, table dto.TableDB) error {
	_, err := d.db.
		NewUpdate().
		Model(&table).
		ExcludeColumn("restaurant_id", "id").
		Where("id = ?", table.ID).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("table updated")
	return nil
}

func (d *TableDatabase) DeleteOne(ctx context.Context, id uuid.UUID) error {
	_, err := d.db.
		NewDelete().
		Model(&dto.TableDB{}).
		Where("id = ?", id).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("table deleted")
	return nil
}

func (d *TableDatabase) DeleteAll(ctx context.Context, id uuid.UUID) error {
	_, err := d.db.
		NewDelete().
		Model(&dto.TableDB{}).
		Where("restaurant_id = ?", id).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("tables deleted")
	return nil
}