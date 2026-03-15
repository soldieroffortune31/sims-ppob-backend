package web

type LoginResponse struct {
	User_id int    `json:"user_id"`
	Token   string `json:"token"`
}
