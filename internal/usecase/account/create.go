package account

import (
	"context"
	"github.com/golang-module/carbon"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity"
	"github.com/puny-activity/authentication/internal/interr"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (u *UseCase) SignUp(ctx context.Context, account entity.AccountCreateRequest) error {
	err := account.Validate()
	if err != nil {
		return err
	}

	err = u.txManager.Transaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		isUsernameTaken, err := u.accountRepo.IsUsernameTakenTx(ctx, tx, account.Username)
		if err != nil {
			return werr.WrapSE("failed to check if username taken", err)
		}
		if isUsernameTaken {
			return interr.UsernameAlreadyTaken
		}

		accountWithHashedPassword, err := account.HashPassword()
		if err != nil {
			return werr.WrapSE("failed to hash password", err)
		}

		accountWithHashedPassword.GenerateID()
		accountWithHashedPassword.CreatedAt = carbon.Now()
		err = u.accountRepo.CreateTx(ctx, tx, accountWithHashedPassword)
		if err != nil {
			return werr.WrapSE("failed to create account", err)
		}

		accountsCount, err := u.accountRepo.AccountsCountTx(ctx, tx)
		if err != nil {
			return werr.WrapSE("failed to count accounts", err)
		}
		if accountsCount == 1 {
			err = u.roleRepo.AssignTx(ctx, tx, *accountWithHashedPassword.ID, entity.RoleAdmin)
			if err != nil {
				return werr.WrapSE("failed to assign admin role", err)
			}
		}

		err = u.roleRepo.AssignTx(ctx, tx, *accountWithHashedPassword.ID, entity.RoleUser)
		if err != nil {
			return werr.WrapSE("failed to assign user role", err)
		}

		return nil
	})
	if err != nil {
		return err
	}

	u.log.Debug().Str("username", account.Username).Msg("account created")
	return nil
}
