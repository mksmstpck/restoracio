package convertors

import (
	apimodels "github.com/mksmstpck/restoracio/internal/APImodels"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/internal/models"
)

func TableDBToDTO(table *models.Table) dto.Table {
	if table == nil {
		return dto.Table{}
	}
	reserv := ReservDBToDTO(table.Reservation)
	return dto.Table{
		ID:           table.ID,
		Number:       table.Number,
		Placement:    table.Placement,
		MaxPeople:    table.MaxPeople,
		IsReserved:   table.IsReserved,
		IsOccupied:   table.IsOccupied,
		RestaurantID: table.RestaurantID,
		Reservation:  &reserv,
	}
}

func TableDTOToDB(table *dto.Table) models.Table {
	if table == nil {
		return models.Table{}
	}
	reserv := ReservDTOToDB(table.Reservation)
	return models.Table{
		ID:           table.ID,
		Number:       table.Number,
		Placement:    table.Placement,
		MaxPeople:    table.MaxPeople,
		IsReserved:   table.IsReserved,
		IsOccupied:   table.IsOccupied,
		RestaurantID: table.RestaurantID,
		Reservation:  &reserv,
	}
}

// TableDTOToResponse converts a Table DTO to a Table API response.
func TableDTOToResponse(tableDTO *dto.Table) *apimodels.TableResponse {
	if tableDTO == nil {
		return nil
	}
	reserv := ReservDTOToResponse(tableDTO.Reservation)
	return &apimodels.TableResponse{
		ID:           tableDTO.ID,
		Number:       tableDTO.Number,
		Placement:    tableDTO.Placement,
		MaxPeople:    tableDTO.MaxPeople,
		IsReserved:   tableDTO.IsReserved,
		IsOccupied:   tableDTO.IsOccupied,
		RestaurantID: tableDTO.RestaurantID,
		Reservation:  reserv,
	}
}

// TableRequestToDTO converts a Table API request to a Table DTO.
func TableRequestToDTO(req *apimodels.TableRequest) dto.Table {
	return dto.Table{
		Number:       req.Number,
		Placement:    req.Placement,
		MaxPeople:    req.MaxPeople,
		IsReserved:   req.IsReserved,
		IsOccupied:   req.IsOccupied,
		RestaurantID: "",
	}
}
