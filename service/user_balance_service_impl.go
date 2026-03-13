package service

import (
	"context"
	"database/sql"
	"sims-ppob/exception"
	"sims-ppob/helper"
	"sims-ppob/model/web"
	"sims-ppob/repository"

	"github.com/go-playground/validator/v10"
)

type UserBalanceServiceImpl struct {
	UserBalanceRepository repository.UserBalanceRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewUserBalance(userBalanceRepository repository.UserBalanceRepository, DB *sql.DB, validate *validator.Validate) UserBalanceService {
	return &UserBalanceServiceImpl{
		UserBalanceRepository: userBalanceRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

// FindByUserId implements [UserBalanceService].
func (service *UserBalanceServiceImpl) FindByUserId(ctx context.Context, userId int) web.UserBalanceResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userBalance, err := service.UserBalanceRepository.FindByUserId(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.UserBalanceResponse{
		User_id:    userBalance.User_id,
		Balance:    userBalance.Balance,
		Updated_at: userBalance.Update_at,
	}
}

// Update implements [UserBalanceService].
func (service *UserBalanceServiceImpl) Update(ctx context.Context, request web.UserBalanceRequest) web.UserBalanceResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userBalance, err := service.UserBalanceRepository.FindByUserId(ctx, tx, request.User_id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	userBalance.Balance = request.Balance

	userBalance = service.UserBalanceRepository.Update(ctx, tx, userBalance)

	return web.UserBalanceResponse{
		User_id:    userBalance.User_id,
		Balance:    userBalance.Balance,
		Updated_at: userBalance.Update_at,
	}
}
