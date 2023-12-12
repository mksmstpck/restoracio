package dto

import (
	"time"

	"github.com/uptrace/bun"
)

type ReservDB struct {
	bun.BaseModel 			  `bun:"table:reserv"`
	ID              string    `bun:",pk"`
	ReservationTime time.Time `bun:"reservation_time"`
	ReserverName    string    `bun:"reserver_name"`
	ReserverPhone   string    `bun:"reserver_phone"`
	TableID         string    `bun:"table_id"`
	RestaurantID    string    `bun:"restaurant_id"`
}

type ReservRequest struct {
	ID        string 	 `json:"id"`
	Year      int    	 `json:"year" binding:"required"`
	Month     int     	 `json:"month" binding:"required"`
	Day       int     	 `json:"day" binding:"required"`
	Hour      int    	 `json:"hour" binding:"required"`
	Minute    int    	 `json:"minute" binding:"required"`
	Second    int    	 `json:"second"`
	TableID   string 	 `json:"table_id" binding:"required"`
	ReserverName string  `json:"reserver_name" binding:"required"`
	ReserverPhone string `json:"reserver_phone" binding:"required"`
}

type ReservResponse struct {
	ID        string 	 `json:"id"`
	Year      int    	 `json:"year" binding:"required"`
	Month     int     	 `json:"month" binding:"required"`
	Day       int     	 `json:"day" binding:"required"`
	Hour      int    	 `json:"hour" binding:"required"`
	Minute    int    	 `json:"minute" binding:"required"`
	Second    int    	 `json:"second"`
	TableID   string 	 `json:"table_id" binding:"required"`
	ReserverName string  `json:"reserver_name" binding:"required"`
	ReserverPhone string `json:"reserver_phone" binding:"required"`

}