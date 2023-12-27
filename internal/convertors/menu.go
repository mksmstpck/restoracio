package convertors

import (
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
)

// MenuDBToDTO converts a Menu database model to a Menu DTO.
func MenuDBToDTO(menuDB *models.Menu) *dto.Menu {
	return &dto.Menu{
		ID:           menuDB.ID,
		Name:         menuDB.Name,
		Description:  menuDB.Description,
		QRCodeID:     menuDB.QRCodeID,
		RestaurantID: menuDB.RestaurantID,
	}
}
