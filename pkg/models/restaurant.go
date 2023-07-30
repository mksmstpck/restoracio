package models

type Restaurant struct {
	ID       string   `json:"id" bun:",pk"`
	Name     string   `json:"name" binding:"required"`
	Location string   `json:"location" binding:"required"`
	StaffIDs []string `json:"staff" binding:"required"`
	DishIDs  []string `json:"dish" binding:"required"`
	TableIDs []string `json:"table" binding:"required"`
	AdminID  string   `json:"admin_id"`
}
