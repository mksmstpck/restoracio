package models

type Dish struct {
	ID          string   `bun:",pk" json:"id"`
	Name        string   `json:"name" binding:"required"`
	Type        string   `json:"type" binding:"required"`
	Category    string   `json:"category" binding:"required"`
	Price       int      `json:"price" binding:"required"`
	MassGrams   int      `json:"mass_grams" binding:"required"`
	Ingredients []string `json:"ingredients" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Photo       string   `json:"photo"`
}
