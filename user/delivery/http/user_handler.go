package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/suryakun/skeleton-go/models"
	"github.com/suryakun/skeleton-go/user"
	validator "gopkg.in/go-playground/validator.v9"
)

// ResponseError ...
type ResponseError struct {
	Message string `json:"message"`
}

// UserHandler ...
type UserHandler struct {
	UserService user.Service
}

// NewUserHandler ...
func NewUserHandler(e *echo.Echo, us user.Service) {
	handler := &UserHandler{
		UserService: us,
	}
	e.GET("/users", handler.FetchUser)
	e.POST("/users", handler.Store)
	e.GET("/users/:id", handler.GetByID)
}

// FetchUser ...
func (u *UserHandler) FetchUser(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	numLimit, _ := strconv.Atoi(limit)
	numOffset, _ := strconv.Atoi(offset)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	listUser, err := u.UserService.All(ctx, numLimit, numOffset)
	if err != nil {
		return c.JSON(403, ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, listUser)
}

// GetByID ...
func (u *UserHandler) GetByID(c echo.Context) error {
	idUser, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}
	id := int64(idUser)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	usr, err := u.UserService.Get(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, usr)
}

// Store ...
func (u *UserHandler) Store(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if ok, err := isRequestValid(&user); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	err = u.UserService.Store(ctx, user)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

func isRequestValid(user *models.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
