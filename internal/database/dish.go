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

func (d *DishDatabase) CreateOne(ctx context.Context, dish dto.Dish) error {
	dishDB := convertors.DishDTOToDB(&dish)
	_, err := d.db.
		NewInsert().
		Model(&dishDB).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("dish created")
	return nil
}

func (d *DishDatabase) GetByID(ctx context.Context, id uuid.UUID) (dto.Dish, error) {
	var dish models.Dish
	err := d.db.
		NewSelect().
		Model(&dish).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error(models.ErrDishNotFound)
			return dto.Dish{}, errors.New(models.ErrDishNotFound)
		}
		log.Error(err)
		return dto.Dish{}, err
	}
	log.Info("dish found")
	return convertors.DishDBToDTO(&dish), nil
}

func (d *DishDatabase) GetAllInMenu(ctx context.Context, id uuid.UUID) ([]dto.Dish, error) {
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
	dishesDTO := make([]dto.Dish, len(dishes))
	for i, dishDB := range dishes {
		dishesDTO[i] = convertors.DishDBToDTO(&dishDB)
	}
	log.Info("dishes found")
	return dishesDTO, nil
}

func (d *DishDatabase) UpdateOne(ctx context.Context, dish dto.Dish) error {
	dishDB := convertors.DishDTOToDB(&dish)
	res, err := d.db.
		NewUpdate().
		Model(&dishDB).
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

func (d *DishDatabase) DeleteAll(ctx context.Context, menuID uuid.UUID) error {
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
