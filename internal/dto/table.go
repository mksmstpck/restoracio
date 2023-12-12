package dto

type TableDB struct {
	ID           string       `json:"id" bun:",pk"`
	Number       int          `json:"number" binding:"required" bun:"number"`
	Placement    string       `json:"placement" binding:"required" bun:"placement"`
	MaxPeople    int          `json:"max_people" binding:"required" bun:"max_people"`
	IsReserved   bool         `json:"is_reserved" bun:"is_reserved"`
	IsOccupied   bool         `json:"is_occupied" binding:"required" bun:"is_occupied"`
	RestaurantID string       `json:"restaurant_id" bun:"restaurant_id"`
	Reservation  *ReservDB	  `json:"reservation" bun:"rel:has-one"`
}

type TableRequest struct {
	Number       int          `json:"number" binding:"required"`
	Placement    string       `json:"placement" binding:"required"`
	MaxPeople    int          `json:"max_people" binding:"required"`
	IsReserved   bool         `json:"is_reserved" binding:"required"`
	IsOccupied   bool         `json:"is_occupied" binding:"required"`
}

type TableResponse struct {
	ID           string       `json:"id"`
	Number       int          `json:"number"`
	Placement    string       `json:"placement"`
	MaxPeople    int          `json:"max_people"`
	IsReserved   bool         `json:"is_reserved"`
	IsOccupied   bool         `json:"is_occupied"`
	RestaurantID string       `json:"restaurant_id"`
	Reservation  *ReservDB	  `json:"reservation"`
}