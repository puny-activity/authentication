package accountrepo

import (
	"context"
	"github.com/golang-module/carbon"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/account"
	"github.com/puny-activity/authentication/internal/entity/account/credential/email"
	"github.com/puny-activity/authentication/internal/entity/role"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/util"
	"github.com/puny-activity/authentication/pkg/werr"
)

type getRepo struct {
	ID           uuid.UUID `db:"id"`
	Email        string    `db:"email"`
	Nickname     string    `db:"nickname"`
	RoleCode     string    `db:"role_code"`
	CreatedAt    string    `db:"created_at"`
	LastActiveAt *string   `db:"last_active_at"`
}

func (r Repository) Get(ctx context.Context, accountID account.ID) (account.Account, error) {
	return r.get(ctx, r.db, accountID)
}

func (r Repository) GetTx(ctx context.Context, tx *sqlx.Tx, accountID account.ID) (account.Account, error) {
	return r.get(ctx, tx, accountID)
}

func (r Repository) get(ctx context.Context, queryer queryer.Queryer, accountID account.ID) (account.Account, error) {
	query := `
SELECT id,
       email,
       nickname,
       role_code,
       created_at,
       last_active_at
FROM accounts
WHERE id = $1
`

	var accountRepo getRepo
	err := queryer.GetContext(ctx, &accountRepo, query, uuid.UUID(accountID))
	if err != nil {
		return account.Account{}, werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	accountEmail, err := email.New(accountRepo.Email)
	if err != nil {
		return account.Account{}, werr.WrapSE("failed to construct email", err)
	}

	accountRole, err := role.New(accountRepo.RoleCode)
	if err != nil {
		return account.Account{}, werr.WrapSE("failed to construct role", err)
	}

	createdAt := carbon.Parse(accountRepo.CreatedAt)
	if createdAt.Error != nil {
		return account.Account{}, werr.WrapSE("failed to construct created at", err)
	}

	var lastOnline *carbon.Carbon = nil
	if accountRepo.LastActiveAt != nil {
		lastOnline = util.ToPointer(carbon.Parse(*accountRepo.LastActiveAt))
		if lastOnline.Error != nil {
			return account.Account{}, werr.WrapSE("failed to construct last online", err)
		}
	}

	accountResponse := account.Account{
		ID:         util.ToPointer(account.ID(accountRepo.ID)),
		Email:      accountEmail,
		Nickname:   accountRepo.Nickname,
		Role:       accountRole,
		CreatedAt:  createdAt,
		LastActive: lastOnline,
	}

	return accountResponse, nil
}
