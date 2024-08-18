package config

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/puny-activity/authentication/pkg/base/environment"
	"github.com/puny-activity/authentication/pkg/werr"
)

type Config struct {
	Logger *LoggerConfig `group:"Logger args" namespace:"logger" env-namespace:"LOGGER" required:"true"`
	HTTP   *HTTPConfig   `group:"Http args" namespace:"http" env-namespace:"HTTP" required:"true"`
	App    *AppConfig    `group:"App args" namespace:"app" env-namespace:"APP" required:"true"`
	DB     *DBConfig     `group:"Db args" namespace:"database" env-namespace:"DATABASE" required:"true"`
}

type LoggerConfig struct {
	Level string `long:"level" env:"LEVEL" default:"info"`
}

type HTTPConfig struct {
	Host string `long:"host" env:"HOST" required:"true"`
	Port string `long:"port" env:"PORT" required:"true"`
}

type AppConfig struct {
	Name                string `long:"name" env:"NAME" default:"authorization"`
	Environment         string `long:"environment" env:"ENVIRONMENT" default:"local"`
	RefreshToken        *Token `group:"Refresh token args" namespace:"refresh-token" env-namespace:"REFRESH_TOKEN" required:"true"`
	AccessToken         *Token `group:"Access token args" namespace:"access-token" env-namespace:"ACCESS_TOKEN" required:"true"`
	environmentInternal environment.Environment
}

type Token struct {
	SecretKeyValue string `long:"secret-key" env:"SECRET_KEY" required:"true"`
	TTLSecondValue int    `long:"ttl-second" env:"TTL_SECOND" required:"true"`
}

func (c *Token) SecretKey() string {
	return c.SecretKeyValue
}

func (c *Token) TTLSecond() int {
	return c.TTLSecondValue
}

type DBConfig struct {
	Host          string `long:"host" env:"HOST" required:"true"`
	Port          int    `long:"port" env:"PORT" required:"true"`
	Name          string `long:"name" env:"NAME" required:"true"`
	User          string `long:"user" env:"USER" required:"true"`
	Password      string `long:"password" env:"PASSWORD" required:"true"`
	MigrationPath string `long:"migration-path" env:"MIGRATION_PATH" default:"migration"`
}

func (c *DBConfig) ConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable default_query_exec_mode=cache_describe",
		c.Host, c.Port, c.User, c.Name, c.Password)
}

func (c *DBConfig) MigrationsPath() string {
	return c.MigrationPath
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
