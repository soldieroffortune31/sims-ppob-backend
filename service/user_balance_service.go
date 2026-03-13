package service

import (
	"context"
	"sims-ppob/model/web"
)

type UserBalanceService interface {
	Update(ctx context.Context, request web.UserBalanceRequest) web.UserBalanceResponse
	FindByUserId(ctx context.Context, userId int) web.UserBalanceResponse
}
