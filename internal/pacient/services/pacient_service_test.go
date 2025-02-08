package pacient

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	commondtos "github.com/nelsonmarro/kyber-med/common/commondtos"
	commonentities "github.com/nelsonmarro/kyber-med/common/commonentities"
)

type PacientServiceTestSuite struct {
	suite.Suite
	pacients   []Pacient
	pagination commondtos.PaginationInfo
}

func (suite *PacientServiceTestSuite) SetupTest() {
	suite.pacients = []Pacient{
		{
			BaseEntity: commonentities.BaseEntity{
				ID:        "p1",
				CreatedAt: 1000,
			},
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@doe.com",
		},
		{
			BaseEntity: commonentities.BaseEntity{
				ID:        "p2",
				CreatedAt: 2000,
			},
			FirstName: "Jane",
			LastName:  "Smith",
			Email:     "jane@smith.com",
		},
	}

	suite.pagination = commondtos.PaginationInfo{
		NextCursor: "abc123",
		PrevCursor: "",
	}
}

func (suite *PacientServiceTestSuite) TestGetPacientsByCursor_Success() {
	// crear el mock
	mockRepo := NewMockPacientRepository(suite.T())

	// instanciar el servicio con el mock
	svc := NewPacientService(mockRepo)

	mockRepo.EXPECT().FindByCursor("", 2, "asc").Return(suite.pacients, suite.pagination, nil)

	// llamar al servicio
	dtos, pagination, err := svc.GetPacientsByCursor("", 2, "asc")

	// validar
	require.NoError(suite.T(), err)
	require.Len(suite.T(), dtos, 2)
	assert.Equal(suite.T(), "p1", dtos[0].ID)
	assert.Equal(suite.T(), "John", dtos[0].FirstName)
	assert.Equal(suite.T(), "Doe", dtos[0].LastName)
	assert.Equal(suite.T(), "john@doe.com", dtos[0].Email)

	// chequear la paginacion
	assert.Equal(suite.T(), "abc123", pagination.NextCursor)
	assert.Equal(suite.T(), "", pagination.PrevCursor)

	// verificar que el mock se llamo con los params correctos
	mockRepo.AssertCalled(suite.T(), "FindByCursor", "", 2, "asc")
}

func (suite *PacientServiceTestSuite) TestGetPacientsByCursor_Error() {
	mockRepo := NewMockPacientRepository(suite.T())

	// instanciar el servicio con el mock
	svc := NewPacientService(mockRepo)

	mockRepo.EXPECT().FindByCursor("", 10, "asc").Return([]Pacient{}, commondtos.PaginationInfo{}, errors.New("error to fetch pacients"))

	// llamar al servicio
	dtos, pagination, err := svc.GetPacientsByCursor("", 10, "asc")

	// validar
	require.Error(suite.T(), err)
	assert.Len(suite.T(), dtos, 0)
	assert.Empty(suite.T(), pagination, pagination)
}

func TestPacientServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PacientServiceTestSuite))
}
