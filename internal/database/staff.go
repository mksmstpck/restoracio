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

func (d *StaffDatabase) CreateOne(ctx context.Context, staff dto.StaffDB) (dto.StaffDB, error) {
	_, err := d.db.
		NewInsert().
		Model(&staff).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return dto.StaffDB{}, err
	}
	log.Print(staff)
	log.Info("staff created")
	return staff, nil
}

func (d *StaffDatabase) GetByID(ctx context.Context, id uuid.UUID, restaurantID uuid.UUID) (dto.StaffDB, error) {
	var staff dto.StaffDB
	err := d.db.
		NewSelect().
		Model(&staff).
		Where("id = ?", id).
		Where("restaurant_id = ?", restaurantID).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info(models.ErrStaffNotFound)
			return dto.StaffDB{}, errors.New(models.ErrStaffNotFound)
		}
		log.Error(err)
		return dto.StaffDB{}, err
	}
	log.Info("staff found")
	return staff, nil
}

func (d *StaffDatabase) GetAllInRestaurant(ctx context.Context, id uuid.UUID) ([]dto.StaffDB, error) {
	var staff []dto.StaffDB
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
		log.Info(models.ErrStaffNotFound)
		return nil, errors.New(models.ErrStaffNotFound)
	}
	log.Info("staffs found")
	return staff, nil
}

func (d *StaffDatabase) UpdateOne(ctx context.Context, staff dto.StaffDB) error {
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
		log.Error(models.ErrStaffNotFound)
		return errors.New(models.ErrStaffNotFound)
	}
	log.Info("staff updated")
	return nil
}

func (d *StaffDatabase) DeleteOne(ctx context.Context, id uuid.UUID, restaurantID uuid.UUID) error {
	res, err := d.db.
		NewDelete().
		Model(&dto.StaffDB{ID: id.String()}).
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
		log.Error(models.ErrStaffNotFound)
		return errors.New(models.ErrStaffNotFound)
	}
	log.Info("staff deleted")
	return nil
}

func (d *StaffDatabase) DeleteAll(ctx context.Context, restaurantID uuid.UUID) error {
	res, err := d.db.
		NewDelete().
		Model(&dto.StaffDB{}).
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
		log.Error(models.ErrStaffNotFound)
		return errors.New(models.ErrStaffNotFound)
	}
	log.Info("staff deleted")
	return nil
}