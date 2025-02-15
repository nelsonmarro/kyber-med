package pacient

import (
	"fmt"

	commondtos "github.com/nelsonmarro/kyber-med/common/commondtos"
	"github.com/nelsonmarro/kyber-med/common/commonhelpers"
	"github.com/nelsonmarro/kyber-med/internal/database"
	pEntities "github.com/nelsonmarro/kyber-med/internal/pacient/entities"
	uRepo "github.com/nelsonmarro/kyber-med/internal/user/repositories"
)

type pacientRepository struct {
	db             database.Database
	userRepository uRepo.UserRepository
}

func NewPacientRepository(db database.Database) PacientRepository {
	return &pacientRepository{db: db}
}

func (r *pacientRepository) FindByCursor(cursor string, limit int, sortOrder string) (data []pEntities.Pacient, pagination commondtos.PaginationInfo, err error) {
	// build base query
	db := r.db.GetDb()
	query := db.Model(&pEntities.Pacient{})

	isFirstPage := cursor == ""
	pointsNext := false

	if cursor != "" {
		decoded, errCursor := commonhelpers.DecodeCursor(cursor)
		if errCursor != nil && cursor != "" {
			err = fmt.Errorf("invalid cursor: %w", errCursor)
			return
		}

		pointsNext = decoded.PointsNext

		operator, order := commonhelpers.GetPaginationOperator(pointsNext, sortOrder)
		whereStr := fmt.Sprintf("(created_at %s ? OR (created_at = ? AND id %s ?))", operator, operator)
		query = query.Where(whereStr, decoded.CreatedAt, decoded.CreatedAt, decoded.ID)
		if order != "" {
			sortOrder = order
		}
	}

	query.Order("created_at " + sortOrder + ", id " + sortOrder).Limit(limit + 1).Find(&data)
	hasPagination := len(data) > limit

	if hasPagination {
		data = data[:limit]
	}

	if !isFirstPage && !pointsNext {
		data = commonhelpers.Reverse(data)
	}

	pagination = commonhelpers.CalculatePagination(isFirstPage, hasPagination, limit, data, pointsNext)

	return data, pagination, nil
}

func (r *pacientRepository) CreatePacient(pacient *pEntities.Pacient, userID string) (*pEntities.Pacient, error) {
	db := r.db.GetDb()

	userDb, err := r.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	if userDb == nil {
		return nil, fmt.Errorf("user not found")
	}

	pacient.UserID = userDb.ID
	result := db.Create(&pacient)
	if result != nil && result.Error != nil {
		return nil, result.Error
	}

	return pacient, nil
}

func (r *pacientRepository) GetPacientByID(pacientID string) (*pEntities.Pacient, error) {
	db := r.db.GetDb()

	var pacient pEntities.Pacient
	result := db.First(&pacient, pacientID)
	if result.Error != nil {
		return nil, fmt.Errorf("Paciente no encontrado", result.Error)
	}

	return &pacient, nil
}
