package pacient

import (
	commondtos "github.com/nelsonmarro/kyber-med/common/commondtos"
	pDtos "github.com/nelsonmarro/kyber-med/internal/pacient/dtos"
)

type PacientService interface {
	GetPacientsByCursor(cursor string, limit int, sortOrder string) ([]pDtos.PacientDto, commondtos.PaginationInfo, error)
	CreatePacient(pacientDto pDtos.UpsertPacientDto, userID string) (*pDtos.PacientDto, error)
	GetPacientByID(id string) (*pDtos.PacientDto, error)
}
