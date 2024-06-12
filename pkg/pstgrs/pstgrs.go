package pstgrs

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

type Config interface {
	GetConnectionString() string
	GetMigrationsPath() string
}

type Postgres struct {
	*sqlx.DB
}

func New(cfg Config) (*Postgres, error) {
	db, err := sqlx.Open("pgx", cfg.GetConnectionString())
	if err != nil {
		return nil, err
	}

	pg := Postgres{
		db,
	}

	goose.SetTableName("migrations")

	err = goose.SetDialect("postgres")
	if err != nil {
		return nil, fmt.Errorf("failed to set postgres dialect: %w", err)
	}

	err = goose.Up(db.DB, cfg.GetMigrationsPath())
	if err != nil {
		return nil, fmt.Errorf("failed to perform migrations: %w", err)
	}

	return &pg, nil
}
