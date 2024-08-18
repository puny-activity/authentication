package devicerepo

import (
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/pkg/txmanager"
	"github.com/rs/zerolog"
)

type Repository struct {
	db        *sqlx.DB
	txManager *txmanager.TxManager
	log       *zerolog.Logger
}

func New(db *sqlx.DB, txManager *txmanager.TxManager, log *zerolog.Logger) *Repository {
	return &Repository{
		db:        db,
		txManager: txManager,
		log:       log,
	}
}
