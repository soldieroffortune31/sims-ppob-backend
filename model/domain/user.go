package domain

import "time"

type User struct {
	User_id       int
	Email         string
	Nama_depan    string
	Nama_belakang string
	Photo         *string
	Password      string
	Token         *string
	UserBalance   UserBalance
	Created_at    time.Time
	Update_at     time.Time
	Deleted_at    *time.Time
}
