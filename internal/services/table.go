package services

import (
	"errors"

	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/mksmstpck/restoracio/utils"
	"github.com/patrickmn/go-cache"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Services) TableCreateService(table models.Table, admin models.Admin) (models.Table, error) {
	if admin.Restaurant == nil {
		return models.Table{}, errors.New("create restaurant first")
	}
	table.RestaurantID = admin.Restaurant.ID
	res, err := s.db.Table.CreateOne(s.ctx, table)
	if err != nil {
		log.Info("TableCreate: ", err)
		return models.Table{}, err
	}
	s.cache.Set(res.ID, res, cache.DefaultExpiration)
	log.Info("table created")
	return res, nil
}

func (s *Services) TableGetByIDService(id uuid.UUID) (models.Table, error) {
	table, exist := s.cache.Get(id.String())
	if exist {
		log.Info("table found")
		return table.(models.Table), nil
	}

	table, err := s.db.Table.GetByID(s.ctx, id)
	if err != nil {
		log.Info("TableGetByID: ", err)
		return models.Table{}, err
	}
	return table.(models.Table), nil
}

func (s *Services) TableGetAllInRestaurantService(id uuid.UUID) ([]models.Table, error) {
	tables, err := s.db.Table.GetAllInRestaurant(s.ctx, id)
	if err != nil {
		log.Info("TableGetAllInRestaurant: ", err)
		return nil, err
	}
	return tables, nil
}

func (s *Services) TableUpdateService(table models.Table, admin models.Admin) error {
	if admin.Restaurant == nil {
		log.Info("create restaurant first")
		return errors.New(utils.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Tables == nil {
		log.Info("create tables first")
		return errors.New(utils.ErrTableNotFound)
	}
	table.RestaurantID = admin.Restaurant.ID

	if !TableExists(admin.Restaurant.Tables, table.ID) {
		log.Info("table not found")
		return errors.New("table not found")
	}

	err := s.db.Table.UpdateOne(s.ctx, table)
	if err != nil {
		log.Info("TableUpdate: ", err)
		return err
	}

	s.cache.Set(table.ID, table, cache.DefaultExpiration)

	log.Info("table updated")
	return nil
}

func (s *Services) TableDeleteService(id uuid.UUID, admin models.Admin) error {
	if admin.Restaurant == nil {
		log.Info("create restaurant first")
		return errors.New(utils.ErrRestaurantNotFound)
	}
	if admin.Restaurant.Tables == nil {
		log.Info("create tables first")
		return errors.New(utils.ErrTableNotFound)
	}
	if !TableExists(admin.Restaurant.Tables, id.String()) {
		log.Info("table not found")
		return errors.New("table not found")
	}
	err := s.db.Table.DeleteOne(s.ctx, id)
	if err != nil {
		log.Info("TableDelete: ", err)
		return err
	}

	s.cache.Delete(id.String())

	log.Info("table deleted")
	return nil
}

func TableExists(tables []*models.Table, id string) bool {
	for _, table := range tables {
		if table.ID == id {
			return true
		}
	}
	return false
}