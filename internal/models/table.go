package models

type Table struct {
	ID           string       `json:"id" bun:",pk"`
	Number       int          `json:"number" bun:"number"`
	Placement    string       `json:"placement" bun:"placement"`
	MaxPeople    int          `json:"max_people" bun:"max_people"`
	IsReserved   bool         `json:"is_reserved" bun:"is_reserved"`
	IsOccupied   bool         `json:"is_occupied" bun:"is_occupied"`
	RestaurantID string       `json:"restaurant_id" bun:"restaurant_id"`
	Reservation  *Reservation `json:"reservation" bun:"rel:has-one"`
}
