package convertors

import (
	apimodels "github.com/mksmstpck/restoracio/internal/APImodels"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/internal/models"
)

// MenuDBToDTO converts a Menu database model to a Menu DTO.
func MenuDBToDTO(menuDB *models.Menu) dto.Menu {
	if menuDB == nil {
		return dto.Menu{}
	}
	dishesDTO := make([]*dto.Dish, len(menuDB.Dishes))
	for i, dishDB := range menuDB.Dishes {
		dish := DishDBToDTO(dishDB)
		dishesDTO[i] = &dish
	}
	return dto.Menu{
		ID:           menuDB.ID,
		Name:         menuDB.Name,
		Description:  menuDB.Description,
		Dishes:       dishesDTO,
		QRCodeID:     menuDB.QRCodeID,
		RestaurantID: menuDB.RestaurantID,
	}
}

// MenuDTOToDB converts a Menu DTO to a Menu database model.
func MenuDTOToDB(menuDTO *dto.Menu) models.Menu {
	if menuDTO == nil {
		return models.Menu{}
	}
	dishesDB := make([]*models.Dish, len(menuDTO.Dishes))
	for i, dishDTO := range menuDTO.Dishes {
		dish := DishDTOToDB(dishDTO)
		dishesDB[i] = &dish
	}
	return models.Menu{
		ID:           menuDTO.ID,
		Name:         menuDTO.Name,
		Description:  menuDTO.Description,
		Dishes:       dishesDB,
		QRCodeID:     menuDTO.QRCodeID,
		RestaurantID: menuDTO.RestaurantID,
	}
}

// MenuDTOToResponse converts a Menu DTO to a Menu response.
func MenuDTOToResponse(menuDTO *dto.Menu) *apimodels.MenuResponse {
	if menuDTO == nil {
		return nil
	}
	dishes := make([]*apimodels.DishResponse, len(menuDTO.Dishes))
	for i, dishDTO := range menuDTO.Dishes {
		dish := DishDTOToResponse(dishDTO)
		dishes[i] = dish
	}
	return &apimodels.MenuResponse{
		ID:           menuDTO.ID,
		Name:         menuDTO.Name,
		Description:  menuDTO.Description,
		Dishes:       dishes,
		QRCodeBytes:  menuDTO.QRCode,
		RestaurantID: menuDTO.RestaurantID,
	}
}

// MenuRequestToDTO converts a MenuRequest to a Menu DTO.
func MenuRequestToDTO(req *apimodels.MenuRequest) dto.Menu {
	if req == nil {
		return dto.Menu{}
	}
	return dto.Menu{
		Name:         req.Name,
		Description:  req.Description,
		RestaurantID: "",
		Dishes:       []*dto.Dish{},
		QRCodeID:     "",
		QRCode:       []byte{},
	}
}
