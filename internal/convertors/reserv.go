package convertors

import (
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
)

func ReservDBToDTO(reserv *models.Reserv) *dto.Reserv {
	return &dto.Reserv{
		ID:              reserv.ID,
		ReservationTime: reserv.ReservationTime,
		ReserverName:    reserv.ReserverName,
		ReserverPhone:   reserv.ReserverPhone,
		TableID:         reserv.TableID,
		RestaurantID:    reserv.RestaurantID,
	}
}
