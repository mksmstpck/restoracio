package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (d *StaffDatabase) CreateOne(ctx context.Context, staff dto.Staff) (dto.Staff, error) {
	_, err := d.db.
		NewInsert().
		Model(&staff).
		Exec(ctx)
	if err != nil {
		log.Error(err)
		return dto.Staff{}, err
	}
	log.Print(staff)
	log.Info("staff created")
	return staff, nil
}

func (d *StaffDatabase) GetByID(ctx context.Context, id uuid.UUID, restaurantID uuid.UUID) (dto.Staff, error) {
	var staff dto.Staff
	err := d.db.
		NewSelect().
		Model(&staff).
		Where("id = ?", id).
		Where("restaurant_id = ?", restaurantID).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info(dto.ErrStaffNotFound)
			return dto.Staff{}, errors.New(dto.ErrStaffNotFound)
		}
		log.Error(err)
		return dto.Staff{}, err
	}
	log.Info("staff found")
	return staff, nil
}

func (d *StaffDatabase) GetAllInRestaurant(ctx context.Context, id uuid.UUID) ([]dto.Staff, error) {
	var staff []dto.Staff
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
		log.Info(dto.ErrStaffNotFound)
		return nil, errors.New(dto.ErrStaffNotFound)
	}
	log.Info("staffs found")
	return staff, nil
}

func (d *StaffDatabase) UpdateOne(ctx context.Context, staff dto.Staff) error {
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
		Model(&dto.Staff{ID: id.String()}).
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
		Model(&dto.Staff{}).
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