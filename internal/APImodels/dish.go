package apimodels

type DishRequest struct {
	Name        string   `json:"name" binding:"required"`
	Type        string   `json:"type" binding:"required"`
	Category    string   `json:"category" binding:"required"`
	Price       int      `json:"price" binding:"required"`
	Curency     string   `json:"currency" binding:"required"`
	MassGrams   int      `json:"mass_grams" binding:"required"`
	Ingredients []string `json:"ingredients" binding:"required"`
	Description string   `json:"description" binding:"required"`
}

type DishResponse struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Category    string   `json:"category"`
	Price       int      `json:"price"`
	Curency     string   `json:"currency"`
	MassGrams   int      `json:"mass_grams"`
	Ingredients []string `json:"ingredients"`
	Description string   `json:"description"`
	MenuID      string   `json:"menu_id"`
}
