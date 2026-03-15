package helper

import (
	"sims-ppob/model/domain"
	"sims-ppob/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		User_id:       user.User_id,
		Email:         user.Email,
		Nama_depan:    user.Nama_depan,
		Nama_belakang: user.Nama_belakang,
		Photo:         user.Photo,
		UserBalanceResponse: web.UserBalanceInfo{
			Balance: user.UserBalance.Balance,
		},
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	userResponses := []web.UserResponse{}
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}
