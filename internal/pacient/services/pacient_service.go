package pacient

import (
	commondtos "github.com/nelsonmarro/kyber-med/common/commondtos"
	"github.com/nelsonmarro/kyber-med/internal/pacient/dtos"
)

type PacientService interface {
	GetPacientsByCursor(cursor string, limit int, sortOrder string) ([]dtos.PacientDto, commondtos.PaginationInfo, error)
}
