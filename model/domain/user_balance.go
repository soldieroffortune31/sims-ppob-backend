package domain

import "time"

type UserBalance struct {
	Userbalance_id int
	User_id        int
	Balance        int64
	Created_at     time.Time
	Update_at      time.Time
	Deleted_at     time.Time
}
