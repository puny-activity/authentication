package account

import (
	"github.com/puny-activity/authentication/pkg/txmanager"
	"github.com/rs/zerolog"
)

type UseCase struct {
	accountRepo accountRepository
	txManager   txmanager.Transactor
	log         *zerolog.Logger
}

func New(accountRepo accountRepository, txManager txmanager.Transactor, log *zerolog.Logger) *UseCase {
	return &UseCase{
		accountRepo: accountRepo,
		txManager:   txManager,
		log:         log,
	}
}
