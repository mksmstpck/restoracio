package convertors

import (
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
)

func DishDBToDTO(dish *models.Dish) *dto.Dish {
	return &dto.Dish{
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
