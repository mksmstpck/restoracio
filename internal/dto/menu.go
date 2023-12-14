package dto

import "github.com/uptrace/bun"

type MenuDB struct {
	bun.BaseModel 		   `bun:"table:menu"`
	ID           string    `bun:"id,pk"`
	Name         string    `bun:"name"`
	Description  string    `bun:"description"`
	Dishes       []*DishDB `bun:"rel:has-many,join:id=menu_id"`
	QRCodeID     string    `bun:"qrcode"`
	RestaurantID string    `bun:"restaurant_id"`
}

type MenuRequest struct {
	Name         string  `json:"name" binding:"required"`
	Description  string  `json:"description" binding:"required"`
}

type MenuResponse struct {
	ID           string  	`json:"id"`
	Name         string  	`json:"name"`
	Description  string  	`json:"description"`
	Dishes       []*DishDB  `json:"dish"`
	QRCodeBytes  []byte 	`json:"qrcode"`
	RestaurantID string  	`json:"restaurant_id"`
}