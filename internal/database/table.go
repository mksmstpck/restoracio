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

func (d *TableDatabase) CreateOne(ctx context.Context, table dto.Table) error {
	tableDB := convertors.TableDTOToDB(&table)
	_, err := d.db.
		NewInsert().
		Model(&tableDB).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("table created")
	return nil
}

func (d *TableDatabase) GetByID(ctx context.Context, id uuid.UUID) (dto.Table, error) {
	var table models.Table
	err := d.db.
		NewSelect().
		Model(&table).
		Where("id = ?", id.String()).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(models.ErrTableNotFound)
			return dto.Table{}, errors.New(models.ErrTableNotFound)
		}
		log.Error(err)
		return dto.Table{}, err
	}
	log.Info("table found")
	return convertors.TableDBToDTO(&table), nil
}

func (d *TableDatabase) GetAllInRestaurant(ctx context.Context, id uuid.UUID) ([]dto.Table, error) {
	var tables []models.Table
	err := d.db.
		NewSelect().
		Model(&tables).
		Where("restaurant_id = ?", id.String()).
		Scan(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	tablesDTO := make([]dto.Table, len(tables))
	for i, tableDB := range tables {
		tablesDTO[i] = convertors.TableDBToDTO(&tableDB)
	}
	log.Info("tables found")
	return tablesDTO, nil
}

func (d *TableDatabase) UpdateOne(ctx context.Context, table dto.Table) error {
	tableDB := convertors.TableDTOToDB(&table)
	_, err := d.db.
		NewUpdate().
		Model(&tableDB).
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
		Model(&models.Table{}).
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
		Model(&models.Table{}).
		Where("restaurant_id = ?", id).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("tables deleted")
	return nil
}
