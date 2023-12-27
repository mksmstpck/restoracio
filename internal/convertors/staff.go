package convertors

import (
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
)

func StaffDBToDTO(staff *models.Staff) *dto.Staff {
	return &dto.Staff{
		ID:           staff.ID,
		RestaurantID: staff.RestaurantID,
		Name:         staff.Name,
		Age:          staff.Age,
		Gender:       staff.Gender,
		Email:        staff.Email,
		Phone:        staff.Phone,
		Position:     staff.Position,
	}
}
