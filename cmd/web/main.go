package main

import (
	"database/sql"
	"fmt"
	"github.com/cchristian77/payroll_be/util/config"
	"github.com/cchristian77/payroll_be/util/database"
	"github.com/cchristian77/payroll_be/util/logger"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
)

var k = koanf.New(".")

type Server struct {
	DB     *sql.DB
	GormDB *gorm.DB
	Router *echo.Echo
}

func init() {
	// Load Config JSON
	if err := k.Load(file.Provider("./env.json"), json.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	if err := k.UnmarshalWithConf("env", &config.Env, koanf.UnmarshalConf{Tag: "env"}); err != nil {
		log.Fatalf("failed to read env.json file: %v", err)
	}

	log.Println("Starting service on port", config.Env.App.Port)
}

func main() {
	// Initialized Logger
	log := logger.Init()
	defer log.Sync()

	db := database.ConnectToDB()
	if db == nil {
		logger.Fatal("Can't connect to Postgres!")
	}

	gormDB, err := database.OpenGormDB(db)
	if err != nil {
		logger.Fatal(fmt.Sprintf("gorm driver errror: %v", err))
	}

	app := Server{
		DB:     db,
		GormDB: gormDB,
		Router: echo.New(),
	}
	defer app.DB.Close()

	// Run application
	app.Router.Logger.Fatal(
		app.Router.Start(fmt.Sprintf(":%d", config.Env.App.Port)),
	)
}
