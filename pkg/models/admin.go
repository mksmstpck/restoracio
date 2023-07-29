package models

type Admin struct {
	ID           string `json:"_id" bun:"id,pk"`
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	RestaurantID string `json:"restaurant_id"`
}
