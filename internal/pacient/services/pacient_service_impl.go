package services

import (
	commondtos "github.com/nelsonmarro/kyber-med/common/commondtos"
	"github.com/nelsonmarro/kyber-med/internal/pacient/dtos"
	"github.com/nelsonmarro/kyber-med/internal/pacient/repositories"
)

type pacientServiceImpl struct {
	pacientRepository repositories.PacientRepository
}

func NewPacientServiceImpl(pacientRepository repositories.PacientRepository) PacientService {
	return &pacientServiceImpl{
		pacientRepository: pacientRepository,
	}
}

func (s *pacientServiceImpl) GetPacientsByCursor(cursor string, limit int, sortOrder string) ([]dtos.PacientDto, commondtos.PaginationInfo, error) {
	if limit < 1 || limit > 100 {
		limit = 10
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "asc"
	}

	pacientSliceDd, pagination, err := s.pacientRepository.FindByCursor(cursor, limit, sortOrder)

	var pacientSliceDto []dtos.PacientDto
	for _, pacient := range pacientSliceDd {
		pacientSliceDto = append(pacientSliceDto, dtos.PacientDto{
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
