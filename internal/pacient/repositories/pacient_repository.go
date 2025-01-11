package repositories

import (
	commondtos "github.com/nelsonmarro/kyber-med/common/commondtos"
	"github.com/nelsonmarro/kyber-med/internal/pacient/entities"
)

type PacientRepository interface {
	FindByCursor(cursor string, limit int, sortOrder string) (data []entities.Pacient, pagination commondtos.PaginationInfo, err error)
}
