package convertors

import (
	apimodels "github.com/mksmstpck/restoracio/internal/APImodels"
	"github.com/mksmstpck/restoracio/internal/dto"
	"github.com/mksmstpck/restoracio/internal/models"
)

func StaffDBToDTO(staff *models.Staff) dto.Staff {
	if staff == nil {
		return dto.Staff{}
	}
	return dto.Staff{
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

func StaffDTOToDB(staff *dto.Staff) models.Staff {
	if staff == nil {
		return models.Staff{}
	}
	return models.Staff{
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

// StaffDTOToResponse converts a Staff DTO to a StaffResponse.
func StaffDTOToResponse(staffDTO *dto.Staff) *apimodels.StaffResponse {
	if staffDTO == nil {
		return nil
	}
	return &apimodels.StaffResponse{
		ID:           staffDTO.ID,
		Name:         staffDTO.Name,
		Age:          staffDTO.Age,
		Gender:       staffDTO.Gender,
		Email:        staffDTO.Email,
		Phone:        staffDTO.Phone,
		Position:     staffDTO.Position,
		RestaurantID: staffDTO.RestaurantID,
	}
}

// StaffRequestToDTO converts a StaffRequest to a Staff DTO.
func StaffRequestToDTO(staffReq *apimodels.StaffRequest) dto.Staff {
	if staffReq == nil {
		return dto.Staff{}
	}
	return dto.Staff{
		Name:         staffReq.Name,
		Age:          staffReq.Age,
		Gender:       staffReq.Gender,
		Email:        staffReq.Email,
		Phone:        staffReq.Phone,
		Position:     staffReq.Position,
		RestaurantID: "",
	}
}
