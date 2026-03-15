package repository

import (
	"context"
	"database/sql"
	"errors"
	"sims-ppob/helper"
	"sims-ppob/model/domain"
	"time"
)

type UserBalanceRepositoryImpl struct {
}

func NewUserBalanceRepository() UserBalanceRepository {
	return &UserBalanceRepositoryImpl{}
}

// Save implements [UserBalanceRepository].
func (repository *UserBalanceRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, userBalance domain.UserBalance) domain.UserBalance {
	SQL := "INSERT INTO userbalance_m (user_id, balance, created_at, updated_at) VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, userBalance.User_id, userBalance.Balance, time.Now().UTC(), time.Now().UTC())
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	userBalance.Userbalance_id = int(id)
	return userBalance
}

// Update implements [UserBalanceRepository].
func (repository *UserBalanceRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, userBalance domain.UserBalance) domain.UserBalance {
	SQL := "UPDATE userbalance_m set balance = ?, updated_at = ? where user_id = ?"
	_, err := tx.ExecContext(ctx, SQL, userBalance.Balance, time.Now().UTC())
	helper.PanicIfError(err)

	return userBalance
}

// FindByUserId implements [UserBalanceRepository].
func (repository *UserBalanceRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId int) (domain.UserBalance, error) {
	SQL := "select userbalance_id, user_id, balance where user_id = ? AND deleted_at IS NULL"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	userBalance := domain.UserBalance{}
	if rows.Next() {
		err := rows.Scan(&userBalance.Userbalance_id, &userBalance.User_id, &userBalance.Balance)
		helper.PanicIfError(err)
		return userBalance, nil
	} else {
		return userBalance, errors.New("user balance is not found")
	}
}
