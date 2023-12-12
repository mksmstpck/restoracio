package dto

type MenuDB struct {
	ID           string    `json:"id" bun:"id,pk"`
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	Dishes       []*DishDB `json:"dish" bun:"rel:has-many,join:id=menu_id"`
	QRCodeID     string    `json:"qrcodeID" bun:"qrcode"`
	QRCodeBytes  []byte    `json:"qrcode" bun:"-"`
	RestaurantID string    `json:"restaurant_id" bun:"restaurant_id"`
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