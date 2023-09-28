package models

import (
	"time"

	"github.com/uptrace/bun"
)

type ReservDB struct {
	bun.BaseModel `json:"-" bun:"table:reserv"`
	ID              string    `json:"id" bun:",pk"`
	ReservationTime time.Time `json:"reservation_time"`
	ReserverName    string    `json:"reserver_name" bun:"reserver_name"`
	ReserverPhone   string    `json:"reserver_phone" bun:"reserver_phone"`
	TableID         string    `json:"table_id" bun:"table_id"`
	RestaurantID    string    `json:"restaurant_id" bun:"restaurant_id"`
}

type ReservAPI struct {
	ID        string `json:"id"`
	Year      int    `json:"year" binding:"required"`
	Month     int    `json:"month" binding:"required"`
	Day       int    `json:"day" binding:"required"`
	Hour      int    `json:"hour" binding:"required"`
	Minute    int    `json:"minute" binding:"required"`
	Second    int    `json:"second"`
	TableID   string `json:"table_id" binding:"required"`
	ReserverName string `json:"reserver_name" binding:"required"`
	ReserverPhone string `json:"reserver_phone" binding:"required"`
}