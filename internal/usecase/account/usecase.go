package account

import (
	"github.com/puny-activity/authentication/pkg/txmanager"
	"github.com/rs/zerolog"
)

type UseCase struct {
	accountRepo accountRepository
	roleRepo    roleRepository
	txManager   txmanager.Transactor
	log         *zerolog.Logger
}

func New(accountRepo accountRepository, roleRepo roleRepository, txManager txmanager.Transactor, log *zerolog.Logger) *UseCase {
	return &UseCase{
		accountRepo: accountRepo,
		roleRepo:    roleRepo,
		txManager:   txManager,
		log:         log,
	}
}
