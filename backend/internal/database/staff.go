package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mksmstpck/restoracio/backend/internal/models"
	"github.com/mksmstpck/restoracio/backend/utils"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (d *StaffDatabase) CreateOne(ctx context.Context, staff models.Staff) (models.Staff, error) {
	_, err := d.db.
		NewInsert().
		Model(&staff).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return models.Staff{}, err
	}
	log.Print(staff)
	log.Info("staff created")
	return staff, nil
}

func (d *StaffDatabase) GetByID(ctx context.Context, id uuid.UUID, restaurantID uuid.UUID) (models.Staff, error) {
	var staff models.Staff
	err := d.db.
		NewSelect().
		Model(&staff).
		Where("id = ?", id).
		Where("restaurant_id = ?", restaurantID).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info(utils.ErrStaffNotFound)
			return models.Staff{}, errors.New(utils.ErrStaffNotFound)
		}
		log.Error(err)
		return models.Staff{}, err
	}
	log.Info("staff found")
	return staff, nil
}

func (d *StaffDatabase) GetAllInRestaurant(ctx context.Context, id uuid.UUID) ([]models.Staff, error) {
	var staff []models.Staff
	err := d.db.
		NewSelect().
		Model(&staff).
		Where("restaurant_id = ?", id).
		Scan(ctx)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if len(staff) == 0 {
		log.Info(utils.ErrStaffNotFound)
		return nil, errors.New(utils.ErrStaffNotFound)
	}
	log.Info("staffs found")
	return staff, nil
}

func (d *StaffDatabase) UpdateOne(ctx context.Context, staff models.Staff) error {
	log.Print(staff.ID)
	log.Print(staff.RestaurantID)
	res, err := d.db.NewUpdate().
	Model(&staff).
	Where("id = ? AND restaurant_id = ?", staff.ID, staff.RestaurantID).
	ExcludeColumn("id", "restaurant_id").
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
		log.Error("staff not found")
		return errors.New("staff not found")
	}
	log.Info("staff updated")
	return nil
}

func (d *StaffDatabase) DeleteOne(ctx context.Context, id uuid.UUID, restaurantID uuid.UUID) error {
	res, err := d.db.
		NewDelete().
		Model(&models.Staff{ID: id.String()}).
		Where("id = ? AND restaurant_id = ?", id, restaurantID).
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
		log.Error("staff not found")
		return errors.New("staff not found")
	}
	log.Info("staff deleted")
	return nil
}

func (d *StaffDatabase) DeleteAll(ctx context.Context, restaurantID uuid.UUID) error {
	res, err := d.db.
		NewDelete().
		Model(&models.Staff{}).
		Where("restaurant_id = ?", restaurantID).
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
		log.Error("staff not found")
		return errors.New("staff not found")
	}
	log.Info("staff deleted")
	return nil
}