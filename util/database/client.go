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

var counts int64

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

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

func ConnectToDB() *sql.DB {
	// get dsn from env.json
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Env.Database.Host,
		config.Env.Database.Port,
		config.Env.Database.User,
		config.Env.Database.Password,
		config.Env.Database.Name,
	)

	// get dsn from docker
	//dsn := os.Getenv("DSN")
	for {
		connection, err := openDB(dsn)
		if err != nil {
			logger.Info("Postgres not yet ready")
			counts++
		} else {
			logger.Info("Connected to Postgres")
			return connection
		}

		if counts > 10 {
			logger.Error(err.Error())
			return nil
		}

		logger.Info("backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}
