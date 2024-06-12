package config

type Config interface {
	GetConnectionString() string
	GetMigrationsPath() string
}
