package models

type Restaurant struct {
	ID       string   	`bun:",pk"`
	Name     string   	`bun:"name"`
	Location string  	`bun:"location"`
	AdminID  string  	`bun:"admin_id"`
	Staff    []*Staff	`bun:"staff_ids,rel:has-many,join:id=restaurant_id"`
	Menu     *Menu  	`bun:"menu_id,rel:has-one,join:id=restaurant_id"`
	Tables   []*Table	`bun:"table_ids,rel:has-many,join:id=restaurant_id"`
}