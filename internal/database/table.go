package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (d *TableDatabase) CreateOne(table models.Table) (models.Table, error) {
	table.ID = uuid.NewUUID().String()
	_, err := d.db.
		NewInsert().
		Model(&table).
		Exec(context.Background())
	if err != nil {
		log.Error("database.TableCreate: ", err)
		return models.Table{}, err
	}
	log.Info("table created")
	return table, nil
}

func (d *TableDatabase) GetByID(id uuid.UUID) (models.Table, error) {
	var table models.Table
	err := d.db.
		NewSelect().
		Model(&table).
		Where("id = ?", id.String()).
		Scan(context.Background())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("table not found")
			return models.Table{}, errors.New("table not found")
		}
		log.Error("database.TableGetByID: ", err)
		return models.Table{}, err
	}
	log.Info("table found")
	return table, nil
}

func (d *TableDatabase) UpdateOne(table models.Table) error {
	_, err := d.db.
		NewUpdate().
		Model(&table).
		Where("id = ?", table.ID).
		Exec(context.Background())
	if err != nil {
		log.Error("database.TableUpdate: ", err)
		return err
	}
	log.Info("table updated")
	return nil
}

func (d *TableDatabase) DeleteOne(id uuid.UUID) error {
	_, err := d.db.
		NewDelete().
		Model(models.Table{}).
		Where("id = ?", id).
		Exec(context.Background())
	if err != nil {
		log.Error("database.TableDelete: ", err)
		return err
	}
	log.Info("table deleted")
	return nil
}