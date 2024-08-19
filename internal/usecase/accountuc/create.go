package accountuc

import (
	"context"
	"github.com/golang-module/carbon"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/account"
	"github.com/puny-activity/authentication/internal/entity/account/credential/password"
	"github.com/puny-activity/authentication/internal/entity/role"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (u *UseCase) Create(ctx context.Context, accountToCreate account.Account, passwordToCreate password.Password) error {
	accountToCreate = accountToCreate.GenerateID()
	accountToCreate.CreatedAt = carbon.Now()
	accountToCreate.LastActive = nil

	err := u.txManager.Transaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		isEmailTaken, err := u.accountRepo.IsEmailTakenTx(ctx, tx, accountToCreate.Email)
		if err != nil {
			return werr.WrapSE("failed to check if email taken", err)
		}
		if isEmailTaken {
			return errs.EmailAlreadyTaken
		}

		isNicknameTaken, err := u.accountRepo.IsNicknameTakenTx(ctx, tx, accountToCreate.Nickname)
		if err != nil {
			return werr.WrapSE("failed to check if nickname taken", err)
		}
		if isNicknameTaken {
			return errs.NicknameAlreadyTaken
		}

		accountsCount, err := u.accountRepo.CountTx(ctx, tx)
		if err != nil {
			return werr.WrapSE("failed to count accounts", err)
		}
		if accountsCount == 0 {
			accountToCreate.Role = role.God
		} else {
			accountToCreate.Role = role.User
		}

		hashedPassword, err := passwordToCreate.Hash()
		if err != nil {
			return werr.WrapSE("failed to hash password", err)
		}

		err = u.accountRepo.CreateTx(ctx, tx, accountToCreate, hashedPassword)
		if err != nil {
			return werr.WrapSE("failed to create account", err)
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
