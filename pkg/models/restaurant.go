package models

type Restaurant struct {
	ID       string   `json:"id" bun:",pk"`
	Name     string   `json:"name" binding:"required" bun:"name"`
	Location string   `json:"location" binding:"required" bun:"location"`
	Staff    []*Staff `json:"staff" bun:"staff_ids,rel:has-many,join:id=restaurant_id"`
	Menu     *Menu    `json:"dish" bun:"dish_ids,rel:has-one,join:id=restaurant_id"`
	Tables   []*Table `json:"table" bun:"table_ids,rel:has-many,join:id=restaurant_id"`
	AdminID  string   `json:"admin_id" bun:"admin_id"`
}
