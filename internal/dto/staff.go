package dto

type StaffDB struct {
	ID           string `json:"id" bun:",pk"`
	Name         string `json:"name" binding:"required"`
	Age          int    `json:"age" binding:"required"`
	Gender       string `json:"gender" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Position     string `json:"position" binding:"required"`
	RestaurantID string `json:"restaurant_id" bun:"restaurant_id"`
}

type StaffRequest struct {
	Name     string	`json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Position string `json:"position" binding:"required"`
}

type StaffResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Gender       string `json:"gender"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Position     string `json:"position"`
	RestaurantID string `json:"restaurant_id"`
}