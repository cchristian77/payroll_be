package api

import (
	"fmt"
	attendanceEntrypoint "github.com/cchristian77/payroll_be/entrypoint/attendance"
	authEntrypoint "github.com/cchristian77/payroll_be/entrypoint/auth"
	overtimeEntrypoint "github.com/cchristian77/payroll_be/entrypoint/overtime"
	"github.com/cchristian77/payroll_be/repository"
	"github.com/cchristian77/payroll_be/service/attendance"
	"github.com/cchristian77/payroll_be/service/auth"
	"github.com/cchristian77/payroll_be/service/overtime"
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
	router.Use(middleware.Recover())

	// Config Rate Limiter allows 500 requests/sec
	router.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(500)))

	// Config Validator to Router
	router.Validator = util.RegisterValidator()

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

	db := database.ConnectToDB()
	if db == nil {
		logger.Fatal("Can't connect to Postgres!")
	}

	gormDB, err := database.OpenGormDB(db)
	if err != nil {
		logger.Fatal(fmt.Sprintf("gorm driver errror: %v", err))
	}

	repository := repository.NewRepository(gormDB)

	// Initialize all service layers
	authService, err := auth.NewService(repository)
	if err != nil {
		logger.Fatal(fmt.Sprintf("auth service initialization error: %v", err))
	}

	attendanceService, err := attendance.NewService(repository)
	if err != nil {
		logger.Fatal(fmt.Sprintf("attendance service initialization error: %v", err))
	}

	overtimeService, err := overtime.NewService(repository)
	if err != nil {
		logger.Fatal(fmt.Sprintf("overtime service initialization error: %v", err))
	}

	utilMiddleware.InitAuthorization(authService)

	authController := authEntrypoint.NewController(authService)
	attendanceController := attendanceEntrypoint.NewController(attendanceService)
	overtimeController := overtimeEntrypoint.NewController(overtimeService)

	// register all routes
	authController.RegisterRoutes(router)
	attendanceController.RegisterRoutes(router)
	overtimeController.RegisterRoutes(router)
}
