package web

type UserResponse struct {
	User_id       int     `json:"user_id"`
	Email         string  `json:"email"`
	Nama_depan    string  `json:"nama_depan"`
	Nama_belakang string  `json:"nama_belakang"`
	Photo         *string `json:"photo"`
}
