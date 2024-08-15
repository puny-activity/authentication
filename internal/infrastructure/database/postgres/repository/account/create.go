package account

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/puny-activity/authentication/internal/entity/account"
	"github.com/puny-activity/authentication/internal/errs"
	"github.com/puny-activity/authentication/pkg/queryer"
	"github.com/puny-activity/authentication/pkg/werr"
)

type createDBParameter struct {
	ID           string  `db:"id"`
	Email        string  `db:"email"`
	Nickname     string  `db:"nickname"`
	PasswordHash string  `db:"password_hash"`
	RoleCode     string  `db:"role_code"`
	CreatedAt    string  `db:"created_at"`
	LastOnline   *string `db:"last_online"`
}

func (r Repository) Create(ctx context.Context, accountToCreate account.ToCreateWithHashedPassword) error {
	return r.create(ctx, r.db, accountToCreate)
}

func (r Repository) CreateTx(ctx context.Context, tx *sqlx.Tx, accountToCreate account.ToCreateWithHashedPassword) error {
	return r.create(ctx, tx, accountToCreate)
}

func (r Repository) create(ctx context.Context, queryer queryer.Queryer, accountToCreate account.ToCreateWithHashedPassword) error {
	query := `
INSERT INTO accounts(id, email, nickname, password_hash, role_code, created_at, last_online)
VALUES (:id, :email, :nickname, :password_hash, :role_code, :created_at, :last_online)
`

	parameter := createDBParameter{
		ID:           accountToCreate.ID.String(),
		Email:        accountToCreate.Email,
		Nickname:     accountToCreate.Nickname,
		PasswordHash: accountToCreate.PasswordHash,
		RoleCode:     accountToCreate.Role.Name(),
		CreatedAt:    accountToCreate.CreatedAt.ToDateTimeString(),
		LastOnline:   nil,
	}

	_, err := queryer.NamedExecContext(ctx, query, parameter)
	if err != nil {
		return werr.WrapES(errs.DatabaseFailedToExecuteQuery, err.Error())
	}

	return nil
}
