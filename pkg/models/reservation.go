package models

import "time"

type Reservation struct {
	ID              string    `json:"id" bun:",pk"`
	ReservationTime time.Time `json:"reservation_time"`
	TableID         string    `json:"table_id"`
}
