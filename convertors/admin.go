package convertors

import (
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/models"
)

func AdminDBToDTO(admin *models.Admin) *dto.Admin {
	return &dto.Admin{
		ID:           admin.ID,
		Name:         admin.Name,
		Email:        admin.Email,
		PasswordHash: admin.PasswordHash,
		Salt:         admin.Salt,
		Restaurant:   RestaurantDBToDTO(admin.Restaurant),
	}
}
