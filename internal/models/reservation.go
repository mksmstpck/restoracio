package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Reserv struct {
	bun.BaseModel 			  `bun:"table:reserv"`
	ID              string    `bun:",pk"`
	ReservationTime time.Time `bun:"reservation_time"`
	ReserverName    string    `bun:"reserver_name"`
	ReserverPhone   string    `bun:"reserver_phone"`
	TableID         string    `bun:"table_id"`
	RestaurantID    string    `bun:"restaurant_id"`
}