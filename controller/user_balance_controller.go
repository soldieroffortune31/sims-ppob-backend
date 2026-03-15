package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserBalanceController interface {
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByUserId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
