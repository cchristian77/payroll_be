package main

import (
	"database/sql"
	"fmt"
	api "github.com/cchristian77/payroll_be/entrypoint"
	"github.com/cchristian77/payroll_be/util/config"
	"github.com/cchristian77/payroll_be/util/logger"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
)

var k = koanf.New(".")

type server struct {
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
	log := logger.Get()
	defer log.Sync()

	router := api.InitRouter()

	router.Logger.Fatal(
		router.Start(fmt.Sprintf(":%d", config.Env.App.Port)),
	)
}
