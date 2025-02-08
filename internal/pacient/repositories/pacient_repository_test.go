package pacient

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	commonentities "github.com/nelsonmarro/kyber-med/common/commonentities"
	"github.com/nelsonmarro/kyber-med/common/commonhelpers"
	entities "github.com/nelsonmarro/kyber-med/internal/pacient/entities"
)

type inMemoryDB struct {
	db *gorm.DB
}

func (i *inMemoryDB) GetDb() *gorm.DB {
	return i.db
}

type PacientRepositoryTestSuite struct {
	suite.Suite
	repo PacientRepository
}

func (suite *PacientRepositoryTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.Nil(suite.T(), err)

	err = db.AutoMigrate(&entities.Pacient{})
	require.Nil(suite.T(), err)

	seedData(db)

	inMem := &inMemoryDB{db: db}
	suite.repo = NewPacientRepository(inMem)
}

func (suite *PacientRepositoryTestSuite) TestFindByCursor_GetFirstPage() {
	data, pagination, err := suite.repo.FindByCursor("", 5, "asc")

	require.Nil(suite.T(), err)
	require.NotEmpty(suite.T(), data)
	require.True(suite.T(), len(data) == 5)

	// check asc sort order
	for i := 0; i < len(data)-1; i++ {
		require.LessOrEqual(suite.T(),
			data[i].CreatedAt,
			data[i+1].CreatedAt,
			"Expected ascending order in CreatedAt",
		)
	}

	require.NotEmpty(suite.T(), pagination.NextCursor)
	require.Empty(suite.T(), pagination.PrevCursor)
}

func (suite *PacientRepositoryTestSuite) TestFindByCursor_GetSecondPage() {
	_, pagination, err := suite.repo.FindByCursor("", 5, "asc")
	require.Nil(suite.T(), err)

	// page 2
	nextCur := pagination.NextCursor
	data2, pagination2, err := suite.repo.FindByCursor(nextCur, 5, "asc")
	require.Nil(suite.T(), err)

	require.NotEmpty(suite.T(), data2)

	// check last item has the ID of nextCur used to fetch second page (desc order)
	decodedCur, _ := commonhelpers.DecodeCursor(nextCur)
	require.Equal(suite.T(), data2[0].CreatedAt, decodedCur.CreatedAt)

	require.Equal(suite.T(), len(data2), 5)

	// check asc sort order
	for i := 0; i < len(data2)-1; i++ {
		require.LessOrEqual(suite.T(),
			data2[i].CreatedAt,
			data2[i+1].CreatedAt,
			"Expected ascending order in CreatedAt",
		)
	}

	require.NotEmpty(suite.T(), pagination2.PrevCursor)
	require.Empty(suite.T(), pagination2.NextCursor)
}

func (suite *PacientRepositoryTestSuite) TestFindByCursor_ReturnToFirstPageFromSecond() {
	// page 1
	_, pagination, err := suite.repo.FindByCursor("", 5, "asc")
	require.Nil(suite.T(), err)

	// page 2
	nextCur := pagination.NextCursor
	_, pagination2, err := suite.repo.FindByCursor(nextCur, 5, "asc")
	require.Nil(suite.T(), err)

	// return to page 1
	prevCursor := pagination2.PrevCursor
	data, pagination, err := suite.repo.FindByCursor(prevCursor, 5, "asc")
	require.Nil(suite.T(), err)

	require.NotEmpty(suite.T(), data)
	require.True(suite.T(), len(data) == 5)

	// check asc sort order
	for i := 0; i < len(data)-1; i++ {
		require.LessOrEqual(suite.T(),
			data[i].CreatedAt,
			data[i+1].CreatedAt,
			"Expected ascending order in CreatedAt",
		)
	}

	decodedPrevCur, _ := commonhelpers.DecodeCursor(prevCursor)
	require.Greater(suite.T(), decodedPrevCur.ID, data[len(data)-1].ID)

	require.NotEmpty(suite.T(), pagination.NextCursor)
	require.Empty(suite.T(), pagination.PrevCursor)
}

func (suite *PacientRepositoryTestSuite) TearDownTest() {
}

func seedData(db *gorm.DB) {
	date, _ := time.Parse("2006-01-02", "1999-01-09")

	pacients := make([]entities.Pacient, 0)

	for i := 0; i < 10; i++ {
		date = date.Add(-time.Duration(21-i) * time.Hour)
		pacients = append(pacients, entities.Pacient{
			BaseEntity:            commonentities.BaseEntity{},
			FirstName:             fmt.Sprintf("Paciente %d", i),
			LastName:              fmt.Sprintf("Last %d", i),
			Email:                 fmt.Sprintf("nelsonmarro%d@gmail.com", i),
			IDCard:                strconv.Itoa(rand.Intn(99999)),
			PhoneNumber:           "0985134196",
			DateOfBirth:           date,
			Gender:                "Masculino",
			Address:               "Quito",
			EmergencyContactName:  "Alieen Torres",
			EmergencyContactPhone: "0999079590",
		})
	}

	db.Migrator().DropTable(&entities.Pacient{})
	db.Migrator().CreateTable(&entities.Pacient{})
	db.CreateInBatches(pacients, 10)
}

func TestPacientRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PacientRepositoryTestSuite))
}
