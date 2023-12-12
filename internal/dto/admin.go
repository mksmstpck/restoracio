package dto

type AdminDB struct {
	ID         string      `bun:"id,pk"`
	Name       string      `bun:"name"`
	Email      string      `bun:"email"`
	Password   string      `bun:"password"`
	Salt       string      `bun:"salt"`
	Restaurant *RestaurantDB `bun:"rel:has-one,join:id=admin_id"`
}

type AdminRequest struct {
	Name       string      `json:"name" binding:"required"`
	Email      string      `json:"email" binding:"required"`
	Password   string      `json:"password" binding:"required"`
}

type AdminResponse struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Email      string      `json:"email"`
	Password   string      `json:"password"`
	Restaurant *RestaurantDB `json:"restaurant"`
}

type AdminLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}