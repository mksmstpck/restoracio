package convertors

import (
	apimodels "github.com/mksmstpck/restoracio/internal/APImodels"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/internal/models"
)

// Converts an Admin database model to an Admin DTO.
func AdminDBToDTO(admin *models.Admin) dto.Admin {
	restaurant := RestaurantDBToDTO(admin.Restaurant)
	return dto.Admin{
		ID:           admin.ID,
		Name:         admin.Name,
		Email:        admin.Email,
		PasswordHash: admin.PasswordHash,
		Salt:         admin.Salt,
		Restaurant:   &restaurant,
	}
}

// Converts an Admin DTO to an Admin database model.
func AdminDTOToDB(adminDTO *dto.Admin) models.Admin {
	restaurant := RestaurantDTOToDB(adminDTO.Restaurant)
	return models.Admin{
		ID:           adminDTO.ID,
		Name:         adminDTO.Name,
		Email:        adminDTO.Email,
		PasswordHash: adminDTO.PasswordHash,
		Salt:         adminDTO.Salt,
		Restaurant:   &restaurant,
	}
}

// AdminDTOToResponse converts an Admin DTO to an Admin response.
func AdminDTOToResponse(adminDTO *dto.Admin) *apimodels.AdminResponse {
	restaurant := RestaurantDTOToResponse(adminDTO.Restaurant)
	return &apimodels.AdminResponse{
		ID:         adminDTO.ID,
		Name:       adminDTO.Name,
		Email:      adminDTO.Email,
		Restaurant: restaurant,
	}
}

// AdminRequestToDTO converts an AdminRequest to an Admin DTO.
func AdminRequestToDTO(adminReq *apimodels.AdminRequest) dto.Admin {
	return dto.Admin{
		Name:     adminReq.Name,
		Email:    adminReq.Email,
		Password: adminReq.Password,
	}
}
