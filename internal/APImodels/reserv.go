package apimodels

import "time"

type ReservRequest struct {
	ID            string `json:"id"`
	Year          int    `json:"year" binding:"required"`
	Month         int    `json:"month" binding:"required"`
	Day           int    `json:"day" binding:"required"`
	Hour          int    `json:"hour" binding:"required"`
	Minute        int    `json:"minute" binding:"required"`
	Second        int    `json:"second"`
	TableID       string `json:"table_id" binding:"required"`
	ReserverName  string `json:"reserver_name" binding:"required"`
	ReserverPhone string `json:"reserver_phone" binding:"required"`
}

type ReservResponse struct {
	ID              string    `json:"id"`
	ReservationTime time.Time `json:"reservation_time"`
	ReserverName    string    `json:"reserver_name"`
	ReserverPhone   string    `json:"reserver_phone"`
	TableID         string    `json:"table_id"`
	RestaurantID    string    `bun:"restaurant_id"`
}
