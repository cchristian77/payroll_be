package attendance

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cchristian77/payroll_be/domain/enums"
	"github.com/cchristian77/payroll_be/repository"
	"github.com/cchristian77/payroll_be/util/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
	"testing"
)

func TestNewService_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := repository.NewMockRepository(ctrl)
	writeDB, _, err := mock.NewMockDB()
	if err != nil {
		t.Fatal(err)
	}

	attendanceService, err := NewService(repoMock, writeDB)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, attendanceService)
	assert.Implements(t, (*Service)(nil), attendanceService)
}

type AttendanceServiceTestSuite struct {
	suite.Suite
	repo    *repository.MockRepository
	writeDB *gorm.DB
	sqlMock sqlmock.Sqlmock
	ctx     context.Context

	attendanceService Service
}

func (suite *AttendanceServiceTestSuite) Before(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var err error

	suite.ctx = context.Background()
	suite.ctx = context.WithValue(suite.ctx, enums.AuthUserCtxKey, mock.InitUserDomain())
	suite.repo = repository.NewMockRepository(ctrl)
	suite.writeDB, suite.sqlMock, err = mock.NewMockDB()
	if err != nil {
		t.Fatal(err)
	}

	suite.attendanceService, err = NewService(suite.repo, suite.writeDB)
	if err != nil {
		t.Fatal(err)
	}
}

func (suite *AttendanceServiceTestSuite) After(t *testing.T) {}

func TestSuiteRunAttendanceService(t *testing.T) {
	suite.Run(t, new(AttendanceServiceTestSuite))
}
