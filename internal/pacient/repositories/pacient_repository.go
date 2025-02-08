package pacient

import (
	commondtos "github.com/nelsonmarro/kyber-med/common/commondtos"
)

type PacientRepository interface {
	FindByCursor(cursor string, limit int, sortOrder string) (data []Pacient, pagination commondtos.PaginationInfo, err error)
}
