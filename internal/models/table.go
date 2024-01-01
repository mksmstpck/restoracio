package models

type Table struct {
	ID           string       `bun:",pk"`
	Number       int          `bun:"number"`
	Placement    string       `bun:"placement"`
	MaxPeople    int          `bun:"max_people"`
	IsReserved   bool         `bun:"is_reserved"`
	IsOccupied   bool         `bun:"is_occupied"`
	RestaurantID string       `bun:"restaurant_id"`
	Reservation  *Reserv	  `bun:"rel:has-one"`
}
