package apimodels

type TableRequest struct {
	Number     int    `json:"number" binding:"required"`
	Placement  string `json:"placement" binding:"required"`
	MaxPeople  int    `json:"max_people" binding:"required"`
	IsReserved bool   `json:"is_reserved"`
	IsOccupied bool   `json:"is_occupied"`
}

type TableResponse struct {
	ID           string          `json:"id"`
	Number       int             `json:"number"`
	Placement    string          `json:"placement"`
	MaxPeople    int             `json:"max_people"`
	IsReserved   bool            `json:"is_reserved"`
	IsOccupied   bool            `json:"is_occupied"`
	RestaurantID string          `json:"restaurant_id"`
	Reservation  *ReservResponse `json:"reservation"`
}
