package domain

type User struct {
	User_id       int
	Email         string
	Nama_depan    string
	Nama_belakang string
	Photo         *string
	Password      string
	Token         *string
}
