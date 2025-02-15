package pacient

import (
	commondtos "github.com/nelsonmarro/kyber-med/common/commondtos"
	pDtos "github.com/nelsonmarro/kyber-med/internal/pacient/dtos"
	pRepo "github.com/nelsonmarro/kyber-med/internal/pacient/repositories"
)

type pacientServiceImpl struct {
	pacientRepository pRepo.PacientRepository
}

func NewPacientService(pacientRepository pRepo.PacientRepository) PacientService {
	return &pacientServiceImpl{
		pacientRepository: pacientRepository,
	}
}

func (s *pacientServiceImpl) GetPacientsByCursor(cursor string, limit int, sortOrder string) ([]pDtos.PacientDto, commondtos.PaginationInfo, error) {
	if limit < 1 || limit > 100 {
		limit = 10
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "asc"
	}

	pacientSliceDd, pagination, err := s.pacientRepository.FindByCursor(cursor, limit, sortOrder)

	var pacientSliceDto []pDtos.PacientDto
	for _, pacient := range pacientSliceDd {
		pacientSliceDto = append(pacientSliceDto, pDtos.PacientDto{
			BaseDto:       commondtos.BaseDto{ID: pacient.ID, CreatedAt: pacient.CreatedAt},
			FirstName:     pacient.FirstName,
			LastName:      pacient.LastName,
			Email:         pacient.Email,
			IDCard:        pacient.IDCard,
			PhoneNumber:   pacient.PhoneNumber,
			DateOfBirth:   pacient.DateOfBirth,
			Gender:        pacient.Gender,
			Address:       pacient.Address,
			Age:           pacient.Age,
			Height:        pacient.Height,
			Weight:        pacient.Weight,
			TargetWeight:  pacient.TargetWeight,
			ActivityLevel: pacient.ActivityLevel,
			DietaryGoal:   pacient.DietaryGoal,
			TargetDate:    pacient.TargetDate,
		})
	}

	return pacientSliceDto, pagination, err
}

func (s *pacientServiceImpl) CreatePacient(pacientDto pDtos.PacientDto) (pDtos.PacientDto, error) {
	panic("unimplemented")
}

func (s *pacientServiceImpl) GetPacientByID(id string) (pDtos.PacientDto, error) {
	panic("unimplemented")
}
