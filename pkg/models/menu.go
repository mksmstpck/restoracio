package models

type Menu struct {
	ID           string  `json:"id" bun:"id,pk"`
	Name         string  `json:"name" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	DishIDs      []*Dish `json:"dish" bun:"rel:has-many,join:id=menu_id"`
	RestaurantID string  `json:"restaurant_id" bun:"rel:restaurant_id"`
}
