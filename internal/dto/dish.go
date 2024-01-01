package dto

type Dish struct {
	ID          string  
	Name        string  
	Type        string  
	Category    string  
	Price       int     
	Curency     string  
	MassGrams   int     
	Ingredients []string
	Description string
	MenuID      string
}
