package convertors

import (
	apimodels "github.com/mksmstpck/restoracio/internal/APImodels"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/internal/models"
)

// RestaurantDBToDTO converts a Restaurant database model to a Restaurant DTO.
func RestaurantDBToDTO(restaurantDB *models.Restaurant) dto.Restaurant {
	if restaurantDB == nil {
		return dto.Restaurant{}
	}
	tablesDTO := make([]*dto.Table, len(restaurantDB.Tables))
	for i, tableDB := range restaurantDB.Tables {
		table := TableDBToDTO(tableDB)
		tablesDTO[i] = &table
	}

	staffDTO := make([]*dto.Staff, len(restaurantDB.Staff))
	for i, staffDB := range restaurantDB.Staff {
		staff := StaffDBToDTO(staffDB)
		staffDTO[i] = &staff
	}
	return dto.Restaurant{
		ID:       restaurantDB.ID,
		Name:     restaurantDB.Name,
		Location: restaurantDB.Location,
		AdminID:  restaurantDB.AdminID,
		Tables:   tablesDTO,
		Staff:    staffDTO,
	}
}

// RestaurantDTOToDB converts a Restaurant DTO to a Restaurant database model.
func RestaurantDTOToDB(restaurantDTO *dto.Restaurant) models.Restaurant {
	if restaurantDTO == nil {
		return models.Restaurant{}
	}
	tablesDB := make([]*models.Table, len(restaurantDTO.Tables))
	for i, tableDTO := range restaurantDTO.Tables {
		table := TableDTOToDB(tableDTO)
		tablesDB[i] = &table
	}

	staffDB := make([]*models.Staff, len(restaurantDTO.Staff))
	for i, staffDTO := range restaurantDTO.Staff {
		staff := StaffDTOToDB(staffDTO)
		staffDB[i] = &staff
	}
	return models.Restaurant{
		ID:       restaurantDTO.ID,
		Name:     restaurantDTO.Name,
		Location: restaurantDTO.Location,
		AdminID:  restaurantDTO.AdminID,
		Tables:   tablesDB,
		Staff:    staffDB,
	}
}

// RestaurantDTOToResponse converts a Restaurant DTO to a Restaurant response.
func RestaurantDTOToResponse(restaurantDTO *dto.Restaurant) *apimodels.RestaurantResponse {
	if restaurantDTO == nil {
		return nil
	}
	tablesRes := make([]*apimodels.TableResponse, len(restaurantDTO.Tables))
	for i, tableDTO := range restaurantDTO.Tables {
		table := TableDTOToResponse(tableDTO)
		tablesRes[i] = table
	}

	staffRes := make([]*apimodels.StaffResponse, len(restaurantDTO.Staff))
	for i, staffDTO := range restaurantDTO.Staff {
		staff := StaffDTOToResponse(staffDTO)
		staffRes[i] = staff
	}

	return &apimodels.RestaurantResponse{
		ID:       restaurantDTO.ID,
		Name:     restaurantDTO.Name,
		Location: restaurantDTO.Location,
		AdminID:  restaurantDTO.AdminID,
		Staff:    staffRes,
		Tables:   tablesRes,
	}
}

// RestaurantRequestToDTO converts a RestaurantRequest to a Restaurant DTO.
func RestaurantRequestToDTO(request *apimodels.RestaurantRequest) dto.Restaurant {
	if request == nil {
		return dto.Restaurant{}
	}
	return dto.Restaurant{
		Name:     request.Name,
		Location: request.Location,
		AdminID:  "",
	}
}
