package config

import "fmt"

var Env Config

type Config struct {
	App struct {
		Name      string `env:"name"`
		Env       string `env:"env"`
		Version   string `env:"version"`
		Port      int32  `env:"port"`
		APIPrefix string `env:"api_prefix"`
	} `env:"app"`

	Database struct {
		Host     string `env:"host"`
		Port     int32  `env:"port"`
		User     string `env:"user"`
		Password string `env:"password"`
		Name     string `env:"name"`
	} `env:"database"`

	Context struct {
		Timeout string `env:"timeout"`
	} `env:"context"`

	Auth struct {
		AccessTokenExpiration string `env:"access_token_expiration"`
	} `env:"auth"`

	JWTKey       string `env:"jwt_key"`
	MigrationURL string `env:"migration_url"`
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
	)
}

func (c *Config) DatabaseUrl() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
	)
}
