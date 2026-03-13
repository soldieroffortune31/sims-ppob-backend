package repository

import (
	"context"
	"database/sql"
	"sims-ppob/model/domain"
)

type UserBalanceRepository interface {
	Save(ctx context.Context, tx *sql.Tx, userBalance domain.UserBalance) domain.UserBalance
	Update(ctx context.Context, tx *sql.Tx, userBalance domain.UserBalance) domain.UserBalance
	FindByUserId(ctx context.Context, tx *sql.Tx, userId int) (domain.UserBalance, error)
}
