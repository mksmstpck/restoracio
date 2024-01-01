package dto

type Menu struct {
	ID           string
	Name         string
	Description  string
	Dishes       []*Dish
	QRCodeID     string
	QRCode       []byte
	RestaurantID string
}
