package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Menu struct {
	ID           string  `json:"id" bun:"id,pk"`
	Name         string  `json:"name" binding:"required"`
	Description  string  `json:"description" binding:"required"`
	Dishes      []*Dish `json:"dish" bun:"rel:has-many,join:id=menu_id"`
	QRCodeID       primitive.ObjectID  `json:"qrcodeID" bun:"qrcode"`
	QRCodeBytes []byte `json:"qrcode"`
	RestaurantID string  `json:"restaurant_id" bun:"restaurant_id"`
}
