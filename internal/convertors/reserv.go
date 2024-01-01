package convertors

import (
	"time"

	apimodels "github.com/mksmstpck/restoracio/internal/APImodels"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/internal/models"
)

func ReservDBToDTO(reserv *models.Reserv) dto.Reserv {
	if reserv == nil {
		return dto.Reserv{}
	}
	return dto.Reserv{
		ID:              reserv.ID,
		ReservationTime: reserv.ReservationTime,
		ReserverName:    reserv.ReserverName,
		ReserverPhone:   reserv.ReserverPhone,
		TableID:         reserv.TableID,
		RestaurantID:    reserv.RestaurantID,
	}
}

func ReservDTOToDB(reserv *dto.Reserv) models.Reserv {
	if reserv == nil {
		return models.Reserv{}
	}
	return models.Reserv{
		ID:              reserv.ID,
		ReservationTime: reserv.ReservationTime,
		ReserverName:    reserv.ReserverName,
		ReserverPhone:   reserv.ReserverPhone,
		TableID:         reserv.TableID,
		RestaurantID:    reserv.RestaurantID,
	}
}

// ReservDTOToResponse converts a Reservation DTO to a Reservation API response.
func ReservDTOToResponse(reservDTO *dto.Reserv) *apimodels.ReservResponse {
	if reservDTO == nil {
		return nil
	}
	return &apimodels.ReservResponse{
		ID:              reservDTO.ID,
		ReservationTime: reservDTO.ReservationTime,
		ReserverName:    reservDTO.ReserverName,
		ReserverPhone:   reservDTO.ReserverPhone,
		TableID:         reservDTO.TableID,
		RestaurantID:    reservDTO.RestaurantID,
	}
}

// ReservRequestToDTO converts a Reservation API request to a Reservation DTO.
func ReservRequestToDTO(req *apimodels.ReservRequest) dto.Reserv {
	if req == nil {
		return dto.Reserv{}
	}
	reservationTime := time.Date(
		req.Year, time.Month(req.Month), req.Day,
		req.Hour, req.Minute, req.Second, 0, time.UTC,
	)
	return dto.Reserv{
		ID:              req.ID,
		ReservationTime: reservationTime,
		ReserverName:    req.ReserverName,
		ReserverPhone:   req.ReserverPhone,
		TableID:         req.TableID,
		RestaurantID:    "",
	}
}
