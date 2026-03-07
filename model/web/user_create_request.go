package web

type UserCreateRequest struct {
	Email           string  `validate:"required,max=100" json:"email"`
	Nama_depan      string  `validate:"required,min=1,max=100" json:"nama_depan"`
	Nama_belakang   string  `validate:"required,max=100" json:"nama_belakang"`
	Photo           *string `json:"photo"`
	Password        string  `validate:"required,min=1,max=100" json:"password"`
	Password_repeat string  `validate:"required,min=1,max=100" json:"password_repeat"`
}
