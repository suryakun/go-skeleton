package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/suryakun/skeleton-go/models"
	http_service "github.com/suryakun/skeleton-go/user/delivery/http"
	"github.com/suryakun/skeleton-go/user/mocks"
	"golang.org/x/net/context"
)

func TestFetch(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	mockService := new(mocks.Service)
	mockListUser := make([]models.User, 0)
	mockListUser = append(mockListUser, mockUser)
	limit := 1
	offset := 2

	ctx := context.Background()
	e := echo.New()
	mockService.On("All", ctx, limit, offset).Return(mockListUser, nil)
	req, err := http.NewRequest(echo.GET, "/users?limit=1&offset=2", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := http_service.UserHandler{UserService: mockService}
	err = handler.FetchUser(c)
	require.NoError(t, err)
	mockService.AssertExpectations(t)
}

func TestStore(t *testing.T) {
	mockUser := models.User{
		Name:  "test",
		Email: "test@t.cc",
		Phone: "2345234",
	}

	tempMock := mockUser
	tempMock.ID = 0

	j, err := json.Marshal(tempMock)
	assert.NoError(t, err)
	mockService := new(mocks.Service)
	mockService.On("Store", mock.Anything, mockUser).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/users", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	handler := http_service.UserHandler{UserService: mockService}
	err = handler.Store(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockService := new(mocks.Service)
	num := int(mockUser.ID)

	mockService.On("Get", mock.Anything, int64(num)).Return(&mockUser, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/users/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := http_service.UserHandler{
		UserService: mockService,
	}
	err = handler.GetByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}
