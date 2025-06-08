package mock

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cchristian77/payroll_be/domain"
	"github.com/cchristian77/payroll_be/domain/enums"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

/*
This file provides functionality to create instances of the specified required structs for unit testing purposes.
This ensures that tests have consistent and predictable data without the need for creating these objects manually in each test case.
*/

/*
 * ============================= MOCKING =============================
 */

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, fmt.Errorf("Error occurs when opening a stub database connection : %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	}), &gorm.Config{})

	if err != nil {
		return nil, nil, fmt.Errorf("Error occurs when opening gorm database : %v", err)
	}

	return gormDB, mock, err
}

func NewEchoContext() echo.Context {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	return e.NewContext(req, rec)
}

/*
 * ============================= DOMAIN =============================
 */

func InitUserDomain() *domain.User {
	now := time.Now()

	return &domain.User{
		ID:         1,
		CreatedAt:  now,
		UpdatedAt:  now,
		Username:   "test_username",
		FullName:   "test_full_name",
		Role:       enums.USERRole,
		BaseSalary: 5000000,
	}
}

func InitAttendanceDomain() *domain.Attendance {
	now := time.Now()

	return &domain.Attendance{
		BaseModel: domain.BaseModel{
			ID:        1,
			CreatedAt: now,
			UpdatedAt: now,
		},
		UserID:   1,
		Date:     now,
		CheckIn:  now,
		CheckOut: nil,
	}
}

func InitOvertimeDomain() *domain.Overtime {
	now := time.Now()

	return &domain.Overtime{
		AttendanceID: 1,
		CreatedAt:    now,
		UpdatedAt:    now,
		CreatedBy:    1,
		UpdatedBy:    nil,
		Date:         now,
		Duration:     3,
		UserID:       1,
	}
}

func InitReimbursementDomain() *domain.Reimbursement {
	now := time.Now()

	return &domain.Reimbursement{
		BaseModel: domain.BaseModel{
			ID:        1,
			CreatedAt: now,
			UpdatedAt: now,
			CreatedBy: 1,
		},
		UserID:      1,
		Description: "test_desc",
		Amount:      7500,
		Status:      enums.PENDINGReimbursementStatus,
	}
}
