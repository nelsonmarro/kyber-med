package handlers

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	commondtos "github.com/nelsonmarro/kyber-med/common/commondtos"
	"github.com/nelsonmarro/kyber-med/internal/pacient/dtos"

	"github.com/nelsonmarro/kyber-med/internal/pacient/services"
)

type PacientHttpHandlerTestSuite struct {
	suite.Suite
	pacients   []dtos.PacientDto
	pagination commondtos.PaginationInfo
	fiberApp   fiber.App
	router     fiber.Router
}

func (suite *PacientHttpHandlerTestSuite) setupFiberServer() {
	suite.fiberApp = *fiber.New()
	api := suite.fiberApp.Group("/api")
	suite.router = api.Group("/v1")
}

func (suite *PacientHttpHandlerTestSuite) SetupTest() {
	suite.setupFiberServer()

	suite.pacients = []dtos.PacientDto{
		{
			BaseDto: commondtos.BaseDto{
				ID:        "p1",
				CreatedAt: 1000,
			},
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@doe.com",
		},
		{
			BaseDto: commondtos.BaseDto{
				ID:        "p2",
				CreatedAt: 2000,
			},
			FirstName: "Jane",
			LastName:  "Smith",
			Email:     "jane@smith.com",
		},
	}

	suite.pagination = commondtos.PaginationInfo{
		NextCursor: "abc",
		PrevCursor: "",
	}
}

func (suite *PacientHttpHandlerTestSuite) TestGetPacientsByCursor_Success() {
	mockService := services.NewMockPacientService(suite.T())

	mockService.EXPECT().GetPacientsByCursor("someCursor", 5, "asc").Return(suite.pacients, suite.pagination, nil)

	pacientHandler := NewPacientHttpHandler(mockService)

	suite.router.Get("/pacients", pacientHandler.GetPacientsByCursor)

	reqUrl := "/api/v1/pacients?per_page=5&sort_order=asc&cursor=someCursor"

	req := httptest.NewRequest("GET", reqUrl, nil)

	// Perform the test
	resp, err := suite.fiberApp.Test(req)
	require.NoError(suite.T(), err)

	require.Equal(suite.T(), fiber.StatusOK, resp.StatusCode)

	// parse the resp
	var response struct {
		Success    bool                      `json:"success"`
		Data       []dtos.PacientDto         `json:"data"`
		Pagination commondtos.PaginationInfo `json:"pagination"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	require.NoError(suite.T(), err)

	// Verificar contenido
	require.True(suite.T(), response.Success)
	require.Len(suite.T(), response.Data, 2)
	require.Equal(suite.T(), "p1", response.Data[0].ID)
	require.Equal(suite.T(), "abc", response.Pagination.NextCursor)
	require.Empty(suite.T(), response.Pagination.PrevCursor)

	// Asegurar que mock se llamó como se esperaba
	mockService.AssertCalled(suite.T(), "GetPacientsByCursor", "someCursor", 5, "asc")
}

func (suite *PacientHttpHandlerTestSuite) TestGetPacientsByCursor_ServiceError() {
	mockService := services.NewMockPacientService(suite.T())

	mockService.EXPECT().GetPacientsByCursor("", 10, "asc").Return(nil, commondtos.PaginationInfo{}, fmt.Errorf("some error from service"))

	pacientHandler := NewPacientHttpHandler(mockService)

	suite.router.Get("/pacients", pacientHandler.GetPacientsByCursor)

	reqUrl := "/api/v1/pacients"
	req := httptest.NewRequest("GET", reqUrl, nil)

	resp, err := suite.fiberApp.Test(req)
	require.NoError(suite.T(), err)

	require.Equal(suite.T(), fiber.StatusInternalServerError, resp.StatusCode)

	// Decodificar respuesta de error
	var response map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&response)

	require.False(suite.T(), response["success"].(bool))
	require.Equal(suite.T(), "some error from service", response["error"])

	// Verificar mock
	mockService.AssertCalled(suite.T(), "GetPacientsByCursor", "", 10, "asc")
}

func TestPacientServiceTestSuite(t *testing.T) {
	suite.Run(t, new(PacientHttpHandlerTestSuite))
}
