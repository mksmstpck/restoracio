package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Services) TableCreateService(table dto.Table, admin dto.Admin) (dto.Table, error) {
	table.ID = uuid.NewUUID().String()
	if admin.Restaurant == nil {
		return dto.Table{}, errors.New("create restaurant first")
	}
	table.RestaurantID = admin.Restaurant.ID
	res, err := s.db.Table.CreateOne(s.ctx, table)
	if err != nil {
		log.Info("TableCreate: ", err)
		return dto.Table{}, err
	}
	s.cache.Set(uuid.Parse(res.ID), res)
	log.Info("table created")
	return res, nil
}

func (s *Services) TableGetByIDService(id uuid.UUID) (dto.Table, error) {
	tableAny, err := s.cache.Get(id)
	if tableAny != nil{
		log.Info("table found")
		return tableAny.(dto.Table), nil
	}
	if err != nil {
		log.Info("TableGetByID: ", err)
		return dto.Table{}, err
	}

	table, err := s.db.Table.GetByID(s.ctx, id)
	if err != nil {
		log.Info("TableGetByID: ", err)
		return dto.Table{}, err
	}
	s.cache.Set(uuid.Parse(table.ID), table)
	return table, nil
}

func (s *Services) TableGetAllInRestaurantService(id uuid.UUID) ([]dto.Table, error) {
	tables, err := s.db.Table.GetAllInRestaurant(s.ctx, id)
	if err != nil {
		log.Info("TableGetAllInRestaurant: ", err)
		return nil, err
	}
	return tables, nil
}

func (s *Services) TableUpdateService(table dto.Table, admin dto.Admin) error {
	if admin.Restaurant == nil {
		log.Info("create restaurant first")
		return errors.New(models.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Tables == nil {
		log.Info("create tables first")
		return errors.New(models.ErrTableNotFound)
	}
	table.RestaurantID = admin.Restaurant.ID

	if !TableExists(admin.Restaurant.Tables, table.ID) {
		log.Info(models.ErrTableNotFound)
		return errors.New(models.ErrTableNotFound)
	}

	err := s.db.Table.UpdateOne(s.ctx, table)
	if err != nil {
		log.Info("TableUpdate: ", err)
		return err
	}

	s.cache.Set(uuid.Parse(table.ID), table)

	log.Info("table updated")
	return nil
}

func (s *Services) TableDeleteService(id uuid.UUID, admin dto.Admin) error {
	if admin.Restaurant == nil {
		log.Info("create restaurant first")
		return errors.New(models.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Tables == nil {
		log.Info("create tables first")
		return errors.New(models.ErrTableNotFound)
	}
	if !TableExists(admin.Restaurant.Tables, id.String()) {
		log.Info(models.ErrTableNotFound)
		return errors.New(models.ErrTableNotFound)
	}
	err := s.db.Table.DeleteOne(s.ctx, id)
	if err != nil {
		log.Info("TableDelete: ", err)
		return err
	}

	s.cache.Delete(id)

	log.Info("table deleted")
	return nil
}

func (s *Services) TableDeleteAllService(admin dto.Admin) error {
	if admin.Restaurant == nil {
		log.Info("create restaurant first")
		return errors.New(models.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Tables == nil {
		log.Info("create tables first")
		return errors.New(models.ErrTableNotFound)
	}
	err := s.db.Table.DeleteAll(s.ctx, uuid.Parse(admin.Restaurant.ID))
	if err != nil {
		log.Info("TableDeleteAll: ", err)
		return err
	}
	return nil
}

func TableExists(tables []*dto.Table, id string) bool {
	for _, table := range tables {
		if table.ID == id {
			return true
		}
	}
	return false
}