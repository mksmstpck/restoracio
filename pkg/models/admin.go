package models

type Admin struct {
	ID         string      `json:"id" bun:"id,pk"`
	Name       string      `json:"name" binding:"required"`
	Email      string      `json:"email" binding:"required"`
	Password   string      `json:"password" binding:"required"`
	Restaurant *Restaurant `json:"restaurant" bun:"rel:has-one,join:id=admin_id"`
}
