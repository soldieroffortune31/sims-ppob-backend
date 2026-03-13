package web

import "time"

type UserBalanceResponse struct {
	User_id    int       `json:"user_id"`
	Balance    int64     `json:"balance"`
	Updated_at time.Time `json:"updated_at"`
}
