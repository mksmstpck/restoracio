package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (d *DishDatabase) CreateOne(ctx context.Context, dish dto.Dish) (dto.Dish, error) {
	_, err := d.db.
		NewInsert().
		Model(&dish).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return dto.Dish{}, err
	}
	log.Info("dish created")
	return dish, nil
}

func (d *DishDatabase) GetByID(ctx context.Context, id uuid.UUID) (dto.Dish, error) {
	var dish dto.Dish
	err := d.db.
		NewSelect().
		Model(&dish).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(dto.ErrDishNotFound)
			return dto.Dish{}, errors.New(dto.ErrDishNotFound)
		}
		log.Error(err)
		return dto.Dish{}, err
	}
	log.Info("dish found")
	return dish, nil
}

func (d *DishDatabase) GetAllInMenu(ctx context.Context, id uuid.UUID) ([]dto.Dish, error) {
	var dishes []dto.Dish
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

func (d *DishDatabase) UpdateOne(ctx context.Context, dish dto.Dish) error {
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
		log.Error(dto.ErrDishNotFound)
		return errors.New(dto.ErrDishNotFound)
	}
	log.Info("dish updated")
	return nil
}

func (d *DishDatabase) DeleteOne(ctx context.Context, id uuid.UUID, menuID uuid.UUID) error {
	res, err := d.db.
		NewDelete().
		Model(&dto.Dish{ID: id.String()}).
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
		log.Error(dto.ErrDishNotFound)
		return errors.New(dto.ErrDishNotFound)
	}
	log.Info("dish deleted")
	return nil
}

func (d *DishDatabase) DeleteAll(ctx context.Context, menuID uuid.UUID) (error) {
	res, err := d.db.
	NewDelete().
	Model(&dto.Dish{MenuID: menuID.String()}).
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
		log.Error(dto.ErrDishNotFound)
		return errors.New(dto.ErrDishNotFound)
	}
	log.Info("dish deleted")
	return nil
}