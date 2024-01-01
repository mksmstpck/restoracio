package convertors

import (
	apimodels "github.com/mksmstpck/restoracio/internal/APImodels"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/internal/models"
)

func DishDBToDTO(dish *models.Dish) dto.Dish {
	if dish == nil {
		return dto.Dish{}
	}
	return dto.Dish{
		ID:          dish.ID,
		Name:        dish.Name,
		Price:       dish.Price,
		Curency:     dish.Curency,
		MassGrams:   dish.MassGrams,
		Type:        dish.Type,
		Category:    dish.Category,
		Ingredients: dish.Ingredients,
		Description: dish.Description,
		MenuID:      dish.MenuID,
	}
}

func DishDTOToDB(dish *dto.Dish) models.Dish {
	if dish == nil {
		return models.Dish{}
	}
	return models.Dish{
		ID:          dish.ID,
		Name:        dish.Name,
		Price:       dish.Price,
		Curency:     dish.Curency,
		MassGrams:   dish.MassGrams,
		Type:        dish.Type,
		Category:    dish.Category,
		Ingredients: dish.Ingredients,
		Description: dish.Description,
		MenuID:      dish.MenuID,
	}
}

// DishDTOToResponse converts a Dish DTO to a DishResponse.
func DishDTOToResponse(dishDTO *dto.Dish) *apimodels.DishResponse {
	if dishDTO == nil {
		return nil
	}
	return &apimodels.DishResponse{
		ID:          dishDTO.ID,
		Name:        dishDTO.Name,
		Price:       dishDTO.Price,
		Curency:     dishDTO.Curency,
		MassGrams:   dishDTO.MassGrams,
		Type:        dishDTO.Type,
		Category:    dishDTO.Category,
		Ingredients: dishDTO.Ingredients,
		Description: dishDTO.Description,
		MenuID:      dishDTO.MenuID,
	}
}

// DishRequestToDTO converts a DishRequest to a Dish DTO.
func DishRequestToDTO(req *apimodels.DishRequest) dto.Dish {
	if req == nil {
		return dto.Dish{}
	}
	return dto.Dish{
		Name:        req.Name,
		Price:       req.Price,
		Curency:     req.Curency,
		MassGrams:   req.MassGrams,
		Type:        req.Type,
		Category:    req.Category,
		Ingredients: req.Ingredients,
		Description: req.Description,
		MenuID:      "",
	}
}
