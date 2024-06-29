package config

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/puny-activity/authentication/pkg/base/environment"
	"github.com/puny-activity/authentication/pkg/werr"
)

type Config struct {
	App    *AppConfig    `group:"App args" namespace:"app" env-namespace:"APP" required:"true"`
	Logger *LoggerConfig `group:"Logger args" namespace:"logger" env-namespace:"LOGGER" required:"true"`
	DB     *DBConfig     `group:"Db args" namespace:"database" env-namespace:"DATABASE" required:"true"`
	HTTP   *HTTPConfig   `group:"Http args" namespace:"http" env-namespace:"HTTP" required:"true"`
}

type AppConfig struct {
	Name                string `long:"name" env:"NAME" default:"authentication"`
	Environment         string `long:"environment" env:"ENVIRONMENT" default:"local"`
	RefreshToken        *Token `group:"Refresh token args" namespace:"refresh-token" env-namespace:"REFRESH_TOKEN" required:"true"`
	AccessToken         *Token `group:"Access token args" namespace:"access-token" env-namespace:"ACCESS_TOKEN" required:"true"`
	environmentInternal environment.Environment
}

type Token struct {
	SecretKey      string `long:"secret-key" env:"SECRET_KEY" required:"true"`
	DurationSecond int    `long:"duration" env:"DURATION" required:"true"`
}

type LoggerConfig struct {
	Level string `long:"level" env:"LEVEL" default:"info"`
}

type DBConfig struct {
	DatabaseName  string `long:"database" env:"DATABASE" required:"true"`
	Host          string `long:"host" env:"HOST" required:"true"`
	Port          int    `long:"port" env:"PORT" required:"true"`
	Name          string `long:"name" env:"NAME" required:"true"`
	User          string `long:"user" env:"USER" required:"true"`
	Password      string `long:"password" env:"PASSWORD" required:"true"`
	MigrationPath string `long:"migration-path" env:"MIGRATION_PATH" default:"internal/infrastructure/database/postgres/migration"`
}

type HTTPConfig struct {
	Host string `long:"host" env:"HOST" required:"true"`
	Port string `long:"port" env:"PORT" required:"true"`
}

func Parse() (*Config, error) {
	var config Config

	p := flags.NewParser(&config, flags.HelpFlag|flags.PassDoubleDash)
	_, err := p.ParseArgs([]string{})
	if err != nil {
		return nil, err
	}

	config.App.environmentInternal, err = environment.New(config.App.Environment)
	if err != nil {
		return nil, werr.WrapSE("failed to parse environment", err)
	}

	return &config, nil
}

func (c *Config) Database() string {
	return c.DB.DatabaseName
}

func (c *Config) ConnectionString() string {
	switch c.DB.DatabaseName {
	case "postgres":
		return fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable default_query_exec_mode=cache_describe",
			c.DB.Host, c.DB.Port, c.DB.User, c.DB.Name, c.DB.Password)
	default:
		return ""
	}
}

func (c *Config) MigrationsPath() string {
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
