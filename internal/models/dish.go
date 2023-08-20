package models

type Dish struct {
	ID          string   `bun:",pk" json:"id"`
	Name        string   `json:"name" binding:"required"`
	Type        string   `json:"type" binding:"required"`
	Category    string   `json:"category" binding:"required"`
	Price       int      `json:"price" binding:"required"`
	Curency     string   `json:"currency" binding:"required"`
	MassGrams   int      `json:"mass_grams" binding:"required" bun:"mass_grams"`
	Ingredients []string `json:"ingredients" binding:"required" bun:",array"`
	Description string   `json:"description" binding:"required"`
	MenuID      string   `json:"menu_id" bun:"menu_id"`
}
