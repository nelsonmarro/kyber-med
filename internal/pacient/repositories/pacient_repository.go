package pacient

import (
	commondtos "github.com/nelsonmarro/kyber-med/common/commondtos"
	pEntities "github.com/nelsonmarro/kyber-med/internal/pacient/entities"
)

type PacientRepository interface {
	FindByCursor(cursor string, limit int, sortOrder string) (data []pEntities.Pacient, pagination commondtos.PaginationInfo, err error)
	CreatePacient(pacient *pEntities.Pacient, userID string) (*pEntities.Pacient, error)
}
