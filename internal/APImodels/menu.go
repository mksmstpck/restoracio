package apimodels

type MenuRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type MenuResponse struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Dishes       []*DishResponse `json:"dish"`
	QRCodeBytes  []byte          `json:"qrcode"`
	RestaurantID string          `json:"restaurant_id"`
}
