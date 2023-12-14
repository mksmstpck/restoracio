package dto

import "github.com/uptrace/bun"

type TableDB struct {
	bun.BaseModel 			  `bun:"table:table"`
	ID           string       `bun:",pk"`
	Number       int          `bun:"number"`
	Placement    string       `bun:"placement"`
	MaxPeople    int          `bun:"max_people"`
	IsReserved   bool         `bun:"is_reserved"`
	IsOccupied   bool         `bun:"is_occupied"`
	RestaurantID string       `bun:"restaurant_id"`
	Reservation  *ReservDB	  `bun:"rel:has-one"`
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