package dto

import "github.com/uptrace/bun"

type RestaurantDB struct {
	bun.BaseModel 		`bun:"table:restaurant"`
	ID       string   	`json:"id" bun:",pk"`
	Name     string   	`json:"name" binding:"required" bun:"name"`
	Location string   	`json:"location" binding:"required" bun:"location"`
	AdminID  string  	`json:"admin_id" bun:"admin_id"`
	Staff    []*StaffDB `json:"staff" bun:"staff_ids,rel:has-many,join:id=restaurant_id"`
	Menu     *MenuDB  	`json:"menu" bun:"menu_id,rel:has-one,join:id=restaurant_id"`
	Tables   []*TableDB 	`json:"table" bun:"table_ids,rel:has-many,join:id=restaurant_id"`
}

type RestaurantRequest struct {
	Name     string   `json:"name" binding:"required"`
	Location string   `json:"location" binding:"required"`
}


type RestaurantResponse struct {
	ID       string   	`json:"id"`
	Name     string   	`json:"name"`
	Location string   	`json:"location"`
	AdminID  string   	`json:"admin_id"`
	Staff    []*StaffDB `json:"staff"`
	Menu     *MenuDB  	`json:"menu"`
	Tables   []*TableDB 	`json:"table"`
}