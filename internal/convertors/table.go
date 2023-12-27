package convertors

import (
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
)

func TableDBToDTO(table *models.Table) *dto.Table {
	return &dto.Table{
		ID:           table.ID,
		Number:       table.Number,
		Placement:    table.Placement,
		MaxPeople:    table.MaxPeople,
		IsReserved:   table.IsReserved,
		IsOccupied:   table.IsOccupied,
		RestaurantID: table.RestaurantID,
		Reservation:  ReservDBToDTO(table.Reservation),
	}
}
