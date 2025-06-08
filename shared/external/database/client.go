package database

import (
	"database/sql"
	"fmt"
	"github.com/cchristian77/payroll_be/util/config"
	"github.com/cchristian77/payroll_be/util/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// openDB opens a database connection using pgx driver
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// OpenGormDB initializes a gorm.DB instance from sql.DB
func OpenGormDB(sqlDB *sql.DB) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return gormDB, err
}

// ConnectToDB establishes a connection to the PostgreSQL database using the configured DSN and retries on failure.
func ConnectToDB() *sql.DB {
	var counts int

	// get dsn from env.json
	dsn := config.Env.DSN()

	// get dsn from docker
	//dsn := os.Getenv("DSN")
	for {
		connection, err := openDB(dsn)
		if err != nil {
			logger.Warn(fmt.Sprintf("Postgres not yet ready: %v", err))
			counts += 1
		} else {
			logger.Info("Connected to Postgres")
			return connection
		}

		if counts > 10 {
			logger.Error(fmt.Sprintf("Failed to connect to Postgres: %v", err))
			return nil
		}

		logger.Info("backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}
