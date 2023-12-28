package models

type Dish struct {
	ID          string   `bun:",pk" json:"id"`
	Name        string   `bun:"name"`
	Type        string   `bun:"type"`
	Category    string   `bun:"category"`
	Price       int      `bun:"price"`
	Curency     string   `bun:"currency"`
	MassGrams   int      `bun:"mass_grams"`
	Ingredients []string `bun:",array"`
	Description string   `bun:"description"`
	MenuID      string   `bun:"menu_id"`
}