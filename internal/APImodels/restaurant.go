package apimodels

type RestaurantRequest struct {
	Name     string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
}

type RestaurantResponse struct {
	ID       string           `json:"id"`
	Name     string           `json:"name"`
	Location string           `json:"location"`
	AdminID  string           `json:"admin_id"`
	Staff    []*StaffResponse `json:"staff"`
	Menu     *MenuResponse    `json:"menu"`
	Tables   []*TableResponse `json:"table"`
}
