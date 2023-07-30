package models

type Restaurant struct {
	ID       string   `json:"id" bun:",pk"`
	Name     string   `json:"name" binding:"required"`
	Location string   `json:"location" binding:"required"`
	Staff    []string `json:"staff" binding:"required"`
	Dish     []string `json:"dish" binding:"required"`
	Table    []string `json:"table" binding:"required"`
	AdminID  []string `json:"admin_id"`
}
