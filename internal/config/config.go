package config

type Config interface {
	Database() string
	ConnectionString() string
	MigrationsPath() string
}
