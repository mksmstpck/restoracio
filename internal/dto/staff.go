package dto

type Staff struct {
	ID           string `json:"id" bun:",pk"`
	Name         string `json:"name" binding:"required"`
	Age          int    `json:"age" binding:"required"`
	Gender       string `json:"gender" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Position     string `json:"position" binding:"required"`
	RestaurantID string `json:"restaurant_id" bun:"restaurant_id"`
}
