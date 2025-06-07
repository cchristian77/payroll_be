package payroll_period

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cchristian77/payroll_be/repository"
	"github.com/cchristian77/payroll_be/util/mock"
	"github.com/labstack/echo/v4"
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

	payrollPeriodService, err := NewService(repoMock, writeDB)
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, payrollPeriodService)
	assert.Implements(t, (*Service)(nil), payrollPeriodService)
}

type PayrollPeriodTestSuite struct {
	suite.Suite
	ec      echo.Context
	repo    *repository.MockRepository
	writeDB *gorm.DB
	sqlMock sqlmock.Sqlmock

	payrollPeriodService Service
}

func (suite *PayrollPeriodTestSuite) Before(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var err error

	suite.ec = mock.NewEchoContext()
	suite.ec.Set("auth_user", mock.InitUserDomain())
	suite.repo = repository.NewMockRepository(ctrl)
	suite.writeDB, suite.sqlMock, err = mock.NewMockDB()
	if err != nil {
		t.Fatal(err)
	}

	suite.payrollPeriodService, err = NewService(suite.repo, suite.writeDB)
	if err != nil {
		t.Fatal(err)
	}
}

func (suite *PayrollPeriodTestSuite) After(t *testing.T) {}

func TestSuiteRunPayrollPeriodService(t *testing.T) {
	suite.Run(t, new(PayrollPeriodTestSuite))
}
