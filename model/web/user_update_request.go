package web

type UserUpdateRequest struct {
	User_id       int     `validate:"required"`
	Email         string  `validate:"required"`
	Nama_depan    string  `validate:"required,min=1,max=100" json:"nama_depan"`
	Nama_belakang string  `validate:"required,max=100" json:"nama_belakang"`
	Photo         *string `json:"photo"`
}
