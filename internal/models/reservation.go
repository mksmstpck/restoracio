package models

import "time"

type ReservDB struct {
	ID              string    `json:"id" bun:",pk"`
	ReservationTime time.Time `json:"reservation_time"`
	ReserverName    string    `json:"reserver_name" bun:"reserver_name"`
	ReserverPhone   string    `json:"reserver_phone" bun:"reserver_phone"`
	TableID         string    `json:"table_id" bun:"table_id"`
	RestaurantID    string    `json:"restaurant_id" bun:"restaurant_id"`
}

type ReservAPI struct {
	Year      int    `json:"year"`
	Month     int    `json:"month"`
	Day       int    `json:"day"`
	Hour      int    `json:"hour"`
	Minute    int    `json:"minute"`
	Second    int    `json:"second"`
	ReserverName string `json:"reserver_name"`
}