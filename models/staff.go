package models

type Staff struct {
	ID           string `bun:",pk"`
	Name         string `binding:"required"`
	Age          int    `binding:"required"`
	Gender       string `binding:"required"`
	Email        string `binding:"required"`
	Phone        string `binding:"required"`
	Position     string `binding:"required"`
	RestaurantID string `bun:"restaurant_id"`
}