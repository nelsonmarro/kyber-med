package pacient

import (
	commondtos "github.com/nelsonmarro/kyber-med/common/commondtos"
)

type PacientService interface {
	GetPacientsByCursor(cursor string, limit int, sortOrder string) ([]pDtos.PacientDto, commondtos.PaginationInfo, error)
}
