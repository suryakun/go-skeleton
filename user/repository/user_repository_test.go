package repository_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/suryakun/skeleton-go/models"
	"github.com/suryakun/skeleton-go/user"
	"github.com/suryakun/skeleton-go/user/repository"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository user.Repository
	user       *models.User
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, mock, err := sqlmock.New()
	s.mock = mock
	require.NoError(s.T(), err)

	sdb, err := gorm.Open("postgres", db)
	s.DB = sdb
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.repository = repository.NewUserRepository(s.DB)
}

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_repository_Get() {
	var (
		id    = int64(1)
		name  = "test-name"
		email = "test@test.com"
		phone = "21251243"
	)

	ctx := context.Background()
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT`)).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone"}).
			AddRow(int64(1), name, email, phone))

	res, err := s.repository.GetByID(ctx, id)
	require.NoError(s.T(), err)
	require.Equal(s.T(), email, res.Email)
}

func (s *Suite) Test_Get_Users() {
	var (
		name  = "test-name"
		email = "test@test.com"
		phone = "21251243"
	)
	ctx := context.Background()
	s.mock.ExpectQuery(regexp.QuoteMeta(`SELECT `)).
		WithArgs().
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "phone"}).AddRow(int64(1), name, email, phone))
	res, err := s.repository.GetUsers(ctx, 1, 3)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), res)
	require.Equal(s.T(), name, res[0].Name)
}

func (s *Suite) Test_Create_User() {
	var (
		name  = "test-name"
		email = "test@test.com"
		phone = "21251243"
	)

	ctx := context.Background()
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`INSERT`)).
		WithArgs().
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	s.mock.ExpectCommit()

	err := s.repository.CreateUser(ctx, models.User{Name: name, Email: email, Phone: phone})
	require.NoError(s.T(), err)

}
