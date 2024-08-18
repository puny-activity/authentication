package database

import (
	"errors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"github.com/puny-activity/authentication/pkg/werr"
)

type Config interface {
	ConnectionString() string
	MigrationsPath() string
}

type Database struct {
	cfg Config
	*sqlx.DB
}

func New(cfg Config) (*Database, error) {
	driverName, err := driverByDatabase("postgres")
	if err != nil {
		return nil, werr.WrapSE("failed to get driver by database", err)
	}

	db, err := sqlx.Open(driverName, cfg.ConnectionString())
	if err != nil {
		return nil, werr.WrapSE("failed to open database", err)
	}

	pg := Database{
		cfg: cfg,
		DB:  db,
	}

	return &pg, nil
}

func (d *Database) RunMigrations() error {
	goose.SetTableName("migrations")

	err := goose.SetDialect(d.DriverName())
	if err != nil {
		return werr.WrapSE("failed to set dialect", err)
	}

	err = goose.Up(d.DB.DB, d.cfg.MigrationsPath())
	if err != nil {
		return werr.WrapSE("failed to run migrations", err)
	}

	return nil
}

func driverByDatabase(database string) (string, error) {
	switch database {
	case "postgres":
		return "pgx", nil
	default:
		return "", errors.New("unknown database type")
	}
}
