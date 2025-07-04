package api

import (
	"fmt"
	"github.com/cchristian77/payroll_be/entrypoint/admin"
	attendanceEntrypoint "github.com/cchristian77/payroll_be/entrypoint/attendance"
	authEntrypoint "github.com/cchristian77/payroll_be/entrypoint/auth"
	overtimeEntrypoint "github.com/cchristian77/payroll_be/entrypoint/overtime"
	reimbursementEntrypoint "github.com/cchristian77/payroll_be/entrypoint/reimbursement"
	"github.com/cchristian77/payroll_be/entrypoint/user"
	"github.com/cchristian77/payroll_be/repository"
	"github.com/cchristian77/payroll_be/service/attendance"
	"github.com/cchristian77/payroll_be/service/auth"
	"github.com/cchristian77/payroll_be/service/overtime"
	payrollPeriod "github.com/cchristian77/payroll_be/service/payroll_period"
	"github.com/cchristian77/payroll_be/service/payslip"
	"github.com/cchristian77/payroll_be/service/reimbursement"
	"github.com/cchristian77/payroll_be/shared/external/database"
	"github.com/cchristian77/payroll_be/util"
	"github.com/cchristian77/payroll_be/util/logger"
	utilMiddleware "github.com/cchristian77/payroll_be/util/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func InitRouter() *echo.Echo {
	router := echo.New()

	// Config CORS
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:          middleware.DefaultSkipper,
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderXCSRFToken},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	// Setup panic recovery for the router
	router.Use(middleware.Recover())

	router.Use(utilMiddleware.RequestID())

	// Config Rate Limiter allows 500 requests/sec
	router.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(500)))

	// Config Validator to Router
	router.Validator = util.InitValidator()

	// Register RequestLog to Router Middleware
	router.Use(logger.RequestLog)

	// Register HTTP Error Handler function
	router.HTTPErrorHandler = util.ErrorHandler

	// Register API Routes
	registerRoutes(router)

	return router
}

func registerRoutes(router *echo.Echo) {
	//timeout, _ := time.ParseDuration(config.Env.Context.Timeout)
	router.GET("/healthcheck", func(ec echo.Context) error {
		return ec.JSON(http.StatusOK, map[string]string{
			"message": "Server is running",
		})
	})

	// initialize DB
	db := database.ConnectToDB()
	if db == nil {
		logger.L().Fatal("Can't connect to Postgres!")
	}

	gormDB, err := database.OpenGormDB(db)
	if err != nil {
		logger.L().Fatal(fmt.Sprintf("gorm driver errror: %v", err))
	}

	repository := repository.NewRepository(gormDB)

	// Initialize all service layers
	authService, err := auth.NewService(repository, gormDB)
	if err != nil {
		logger.L().Fatal(fmt.Sprintf("auth service initialization error: %v", err))
	}

	attendanceService, err := attendance.NewService(repository, gormDB)
	if err != nil {
		logger.L().Fatal(fmt.Sprintf("attendance service initialization error: %v", err))
	}

	overtimeService, err := overtime.NewService(repository, gormDB)
	if err != nil {
		logger.L().Fatal(fmt.Sprintf("overtime service initialization error: %v", err))
	}

	reimbursementService, err := reimbursement.NewService(repository, gormDB)
	if err != nil {
		logger.L().Fatal(fmt.Sprintf("reimbursement service initialization error: %v", err))
	}

	payrollPeriodService, err := payrollPeriod.NewService(repository, gormDB)
	if err != nil {
		logger.L().Fatal(fmt.Sprintf("payroll period service initialization error: %v", err))
	}

	payslipService, err := payslip.NewService(repository, gormDB)
	if err != nil {
		logger.L().Fatal(fmt.Sprintf("payslip.go service initialization error: %v", err))
	}

	utilMiddleware.InitAuthorization(authService)

	// initialize all controller layers
	authController := authEntrypoint.NewController(authService)
	attendanceController := attendanceEntrypoint.NewController(attendanceService)
	overtimeController := overtimeEntrypoint.NewController(overtimeService)
	reimbursementController := reimbursementEntrypoint.NewController(reimbursementService)
	adminController := admin.NewController(payrollPeriodService, payslipService)
	userController := user.NewController(payslipService)

	// register all routes
	authController.RegisterRoutes(router)
	attendanceController.RegisterRoutes(router)
	overtimeController.RegisterRoutes(router)
	reimbursementController.RegisterRoutes(router)
	adminController.RegisterRoutes(router)
	userController.RegisterRoutes(router)
}
