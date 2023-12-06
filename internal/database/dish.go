package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (d *DishDatabase) CreateOne(ctx context.Context, dish models.Dish) (models.Dish, error) {
	_, err := d.db.
		NewInsert().
		Model(&dish).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return models.Dish{}, err
	}
	log.Info("dish created")
	return dish, nil
}

func (d *DishDatabase) GetByID(ctx context.Context, id uuid.UUID) (models.Dish, error) {
	var dish models.Dish
	err := d.db.
		NewSelect().
		Model(&dish).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(models.ErrDishNotFound)
			return models.Dish{}, errors.New(models.ErrDishNotFound)
		}
		log.Error(err)
		return models.Dish{}, err
	}
	log.Info("dish found")
	return dish, nil
}

func (d *DishDatabase) GetAllInMenu(ctx context.Context, id uuid.UUID) ([]models.Dish, error) {
	var dishes []models.Dish
	err := d.db.
		NewSelect().
		Model(&dishes).
		Where("menu_id = ?", id.String()).
		Scan(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("dishes found")
	return dishes, nil
}

func (d *DishDatabase) UpdateOne(ctx context.Context, dish models.Dish) error {
	res, err := d.db.
		NewUpdate().
		Model(&dish).
		Where("id = ?", dish.ID).
		Where("menu_id = ?", dish.MenuID).
		ExcludeColumn("id").
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
		log.Error(models.ErrDishNotFound)
		return errors.New(models.ErrDishNotFound)
	}
	log.Info("dish updated")
	return nil
}

func (d *DishDatabase) DeleteOne(ctx context.Context, id uuid.UUID, menuID uuid.UUID) error {
	res, err := d.db.
		NewDelete().
		Model(&models.Dish{ID: id.String()}).
		Where("id = ?", id).
		Where("menu_id = ?", menuID).
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
		log.Error(models.ErrDishNotFound)
		return errors.New(models.ErrDishNotFound)
	}
	log.Info("dish deleted")
	return nil
}

func (d *DishDatabase) DeleteAll(ctx context.Context, menuID uuid.UUID) (error) {
	res, err := d.db.
	NewDelete().
	Model(&models.Dish{MenuID: menuID.String()}).
	Where("menu_id = ?", menuID).
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
		log.Error(models.ErrDishNotFound)
		return errors.New(models.ErrDishNotFound)
	}
	log.Info("dish deleted")
	return nil
}