package dto

type Table struct {
	ID           string
	Number       int
	Placement    string
	MaxPeople    int
	IsReserved   bool
	IsOccupied   bool
	RestaurantID string
	Reservation  *Reserv
}
