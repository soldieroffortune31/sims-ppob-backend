package web

type UserBalanceRequest struct {
	User_id int   `validate:"required" json:"user_id"`
	Balance int64 `validate:"required" json:"balance"`
}
