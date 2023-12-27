package convertors

import (
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
)

// RestaurantDBToDTO converts a Restaurant database model to a Restaurant DTO.
func RestaurantDBToDTO(restaurantDB *models.Restaurant) *dto.Restaurant {
	tablesDTO := make([]*dto.Table, len(restaurantDB.Tables))
	for i, tableDB := range restaurantDB.Tables {
		tablesDTO[i] = TableDBToDTO(tableDB)
	}

	staffDTO := make([]*dto.Staff, len(restaurantDB.Staff))
	for i, staffDB := range restaurantDB.Staff {
		staffDTO[i] = StaffDBToDTO(staffDB)
	}
	return &dto.Restaurant{
		ID:       restaurantDB.ID,
		Name:     restaurantDB.Name,
		Location: restaurantDB.Location,
		AdminID:  restaurantDB.AdminID,
		Tables:   tablesDTO,
		Staff:    staffDTO,
	}
}
