package accountrepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/account"
	"github.com/puny-activity/authentication/internal/entity/account/credential/password"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

type createRepo struct {
	ID           uuid.UUID `db:"id"`
	Email        string    `db:"email"`
	Nickname     string    `db:"nickname"`
	PasswordHash string    `db:"password_hash"`
	RoleCode     string    `db:"role_code"`
	CreatedAt    string    `db:"created_at"`
	LastActiveAt *string   `db:"last_active_at"`
}

func (r Repository) Create(ctx context.Context, accountToCreate account.Account, hashedPasswordToCreate password.Hashed) error {
	return r.create(ctx, r.db, accountToCreate, hashedPasswordToCreate)
}

func (r Repository) CreateTx(ctx context.Context, tx *sqlx.Tx, accountToCreate account.Account, hashedPasswordToCreate password.Hashed) error {
	return r.create(ctx, tx, accountToCreate, hashedPasswordToCreate)
}

func (r Repository) create(ctx context.Context, queryer queryer.Queryer, accountToCreate account.Account, hashedPasswordToCreate password.Hashed) error {
	if accountToCreate.ID == nil {
		return errs.DatabaseUndefinedID
	}

	query := `
INSERT INTO accounts(id, email, nickname, password_hash, role_code, created_at, last_active_at)
VALUES (:id, :email, :nickname, :password_hash, :role_code, :created_at, :last_active_at)
`

	parameter := createRepo{
		ID:           uuid.UUID(*accountToCreate.ID),
		Email:        accountToCreate.Email.String(),
		Nickname:     accountToCreate.Nickname,
		PasswordHash: hashedPasswordToCreate.String(),
		RoleCode:     accountToCreate.Role.Name(),
		CreatedAt:    accountToCreate.CreatedAt.ToDateTimeString(),
		LastActiveAt: nil,
	}

	_, err := queryer.NamedExecContext(ctx, query, parameter)
	if err != nil {
		return werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	return nil
}
