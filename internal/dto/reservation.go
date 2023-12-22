package dto

import (
	"time"
)

type Reserv struct {
	ID              string
	ReservationTime time.Time
	ReserverName    string
	ReserverPhone   string
	TableID         string
	RestaurantID    string
}

