package config

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/puny-activity/authentication/pkg/environment"
)

type Config struct {
	App    *AppConfig    `group:"App args" namespace:"app" env-namespace:"APP"`
	Logger *LoggerConfig `group:"Logger args" namespace:"logger" env-namespace:"LOGGER"`
	DB     *DBConfig     `group:"Db args" namespace:"database" env-namespace:"DATABASE"`
	HTTP   *HTTPConfig   `group:"Http args" namespace:"http" env-namespace:"HTTP"`
}

type AppConfig struct {
	Name                string `long:"name" env:"NAME" default:"authentication"`
	Environment         string `long:"environment" env:"ENVIRONMENT"`
	environmentInternal environment.Environment
}

type LoggerConfig struct {
	Level string `long:"level" env:"LEVEL" default:"info"`
}

type DBConfig struct {
	Host          string `long:"host" env:"HOST"`
	Port          int    `long:"port" env:"PORT"`
	Name          string `long:"name" env:"NAME"`
	User          string `long:"user" env:"USER"`
	Password      string `long:"password" env:"PASSWORD"`
	MigrationPath string `long:"migration-path" env:"MIGRATION_PATH" default:"internal/infrastructure/database/postgres/migration"`
}

type HTTPConfig struct {
	Host string `long:"host" env:"HOST"`
	Port string `long:"port" env:"PORT"`
}

func Parse() (*Config, error) {
	var config Config

	p := flags.NewParser(&config, flags.HelpFlag|flags.PassDoubleDash)
	_, err := p.ParseArgs([]string{})
	if err != nil {
		return nil, err
	}

	config.App.environmentInternal = environment.New(config.App.Environment)

	return &config, nil
}

func (c *Config) GetConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable default_query_exec_mode=cache_describe",
		c.DB.Host, c.DB.Port, c.DB.User, c.DB.Name, c.DB.Password)
}

func (c *Config) GetMigrationsPath() string {
	return c.DB.MigrationPath
}

func (c *Config) IsProduction() bool {
	return c.App.environmentInternal == environment.Production
}

func (c *Config) IsTest() bool {
	return c.App.environmentInternal == environment.Test
}

func (c *Config) IsDevelopment() bool {
	return c.App.environmentInternal == environment.Development
}

func (c *Config) IsLocal() bool {
	return c.App.environmentInternal == environment.Local
}
