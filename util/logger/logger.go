package logger

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var logger *zap.Logger

func Get() *zap.Logger {
	if logger == nil {
		initLogger()
	}
	return logger
}

func initLogger() *zap.Logger {
	fileLoggerConfig := zap.NewProductionEncoderConfig()
	fileLoggerConfig.MessageKey = "message"
	fileLoggerConfig.LevelKey = "level"
	fileLoggerConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	fileLoggerConfig.TimeKey = "timestamp"
	fileLoggerConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	fileLoggerConfig.CallerKey = "caller"
	fileLoggerConfig.EncodeCaller = zapcore.ShortCallerEncoder
	fileLoggerConfig.FunctionKey = "func"
	logFile, _ := os.OpenFile("logs/errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	core := zapcore.NewTee(
		// logger to record in warn level (including errors) to errors.log
		zapcore.NewCore(
			zapcore.NewJSONEncoder(fileLoggerConfig),
			zapcore.AddSync(logFile),
			zapcore.WarnLevel,
		),
		// logger to record in debug level in terminal
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			zapcore.AddSync(os.Stdout),
			zapcore.DebugLevel,
		),
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	return logger
}

// RequestLog logs all requests that occurs when service is running
func RequestLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ec echo.Context) error {
		if err := next(ec); err != nil {
			ec.Error(err)
		}

		request := ec.Request()
		response := ec.Response()

		fields := []zapcore.Field{
			zap.Int("status", response.Status),
			zap.String("latency", time.Since(time.Now()).String()),
			zap.String("method", request.Method),
			zap.String("uri", request.RequestURI),
			zap.String("remote_ip", ec.RealIP()),
		}

		statusCode := response.Status
		switch {
		case statusCode >= 500:
			logger.Error("Internal Server Error", fields...)
		case statusCode >= 400:
			logger.Warn("Client-side Error", fields...)
		case statusCode >= 300:
			logger.Info("Redirection", fields...)
		default:
			logger.Debug("Success", fields...)
		}

		return nil
	}
}

func Fatal(message string) {
	logger.Fatal(message)
}

func Error(message string) {
	logger.Error(message)
}

func Warn(message string) {
	logger.Warn(message)
}

func Info(message string) {
	logger.Info(message)
}

func InfoWithCtx(ctx context.Context, message string) {
	logger.Info(message, getZapFieldsFromCtx(ctx)...)
}

func Debug(message string) {
	logger.Debug(message)
}

func getZapFieldsFromCtx(ctx context.Context) []zapcore.Field {
	requestID := requestIDFromContext(ctx)
	authUser := authUserFromContext(ctx)

	var userID uint64
	if authUser != nil {
		userID = authUser.ID
	}

	var fields []zapcore.Field

	if requestID != "" {
		fields = append(fields, zap.String("request_id", requestID))
	}

	if userID != 0 {
		fields = append(fields, zap.Uint64("user_id", userID))
	}

	return fields
}
