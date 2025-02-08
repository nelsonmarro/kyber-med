package pacient

import (
	commondtos "github.com/nelsonmarro/kyber-med/common/commondtos"
)

type pacientServiceImpl struct {
	pacientRepository PacientRepository
}

func NewPacientService(pacientRepository PacientRepository) PacientService {
	return &pacientServiceImpl{
		pacientRepository: pacientRepository,
	}
}

func (s *pacientServiceImpl) GetPacientsByCursor(cursor string, limit int, sortOrder string) ([]PacientDto, commondtos.PaginationInfo, error) {
	if limit < 1 || limit > 100 {
		limit = 10
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "asc"
	}

	pacientSliceDd, pagination, err := s.pacientRepository.FindByCursor(cursor, limit, sortOrder)

	var pacientSliceDto []PacientDto
	for _, pacient := range pacientSliceDd {
		pacientSliceDto = append(pacientSliceDto, PacientDto{
			BaseDto:               commondtos.BaseDto{ID: pacient.ID, CreatedAt: pacient.CreatedAt},
			FirstName:             pacient.FirstName,
			LastName:              pacient.LastName,
			Email:                 pacient.Email,
			IDCard:                pacient.IDCard,
			PhoneNumber:           pacient.PhoneNumber,
			DateOfBirth:           pacient.DateOfBirth,
			Gender:                pacient.Gender,
			Address:               pacient.Address,
			EmergencyContactName:  pacient.EmergencyContactName,
			EmergencyContactPhone: pacient.EmergencyContactPhone,
		})
	}

	return pacientSliceDto, pagination, err
}
