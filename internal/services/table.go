package services

import (
	"github.com/mksmstpck/restoracio/internal/models"
	"github.com/patrickmn/go-cache"
	"github.com/pborman/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *Services) TableCreateService(table models.Table, admin models.Admin) (models.Table, error) {
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
	table, err := s.db.Table.GetByID(s.ctx, id)
	if err != nil {
		log.Info("TableGetByID: ", err)
		return models.Table{}, err
	}
	return table, nil
}

func (s *Services) TableUpdateService(table models.Table) error {
	err := s.db.Table.UpdateOne(s.ctx, table)
	if err != nil {
		log.Info("TableUpdate: ", err)
		return err
	}
	log.Info("table updated")
	return nil
}

func (s *Services) TableDeleteService(id uuid.UUID) error {
	err := s.db.Table.DeleteOne(s.ctx, id)
	if err != nil {
		log.Info("TableDelete: ", err)
		return err
	}
	log.Info("table deleted")
	return nil
}