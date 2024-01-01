package models

type Menu struct {
	ID           string    `bun:"id,pk"`
	Name         string    `bun:"name"`
	Description  string    `bun:"description"`
	Dishes       []*Dish   `bun:"rel:has-many,join:id=menu_id"`
	QRCodeID     string    `bun:"qrcode"`
	RestaurantID string    `bun:"restaurant_id"`
}