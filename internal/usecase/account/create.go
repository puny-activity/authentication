package account

import (
	"context"
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/account"
	"github.com/puny-activity/authentication/internal/entity/role"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/werr"
)

func (u *UseCase) SignUp(ctx context.Context, account account.ToCreate) error {
	accountID := uuid.New()
	account.ID = &accountID
	account.CreatedAt = carbon.Now()

	err := account.Validate()
	if err != nil {
		return werr.WrapSE("failed to validate account", err)
	}

	err = u.txManager.Transaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		isEmailTaken, err := u.accountRepo.IsEmailTakenTx(ctx, tx, account.Email)
		if err != nil {
			return werr.WrapSE("failed to check if email taken", err)
		}
		if isEmailTaken {
			return errs.EmailAlreadyTaken
		}

		isNicknameTaken, err := u.accountRepo.IsNicknameTakenTx(ctx, tx, account.Nickname)
		if err != nil {
			return werr.WrapSE("failed to check if nickname taken", err)
		}
		if isNicknameTaken {
			return errs.NicknameAlreadyTaken
		}

		accountsCount, err := u.accountRepo.AccountsCountTx(ctx, tx)
		if err != nil {
			return werr.WrapSE("failed to count accounts", err)
		}
		if accountsCount == 0 {
			account.Role = role.God
		} else {
			account.Role = role.User
		}

		accountWithHashedPassword, err := account.HashPassword()
		if err != nil {
			return werr.WrapSE("failed to hash password", err)
		}

		err = u.accountRepo.CreateTx(ctx, tx, accountWithHashedPassword)
		if err != nil {
			return werr.WrapSE("failed to create account", err)
		}

		return nil
	})
	if err != nil {
		return err
	}

	u.log.Debug().Str("nickname", account.Nickname).Msg("account created")
	return nil
}
