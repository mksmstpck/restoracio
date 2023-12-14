package dto

import "github.com/uptrace/bun"

type StaffDB struct {
	bun.BaseModel  		`bun:"table:staff"`
	ID           string `bun:",pk"`
	Name         string `binding:"required"`
	Age          int    `binding:"required"`
	Gender       string `binding:"required"`
	Email        string `binding:"required"`
	Phone        string `binding:"required"`
	Position     string `binding:"required"`
	RestaurantID string `bun:"restaurant_id"`
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