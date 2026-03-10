package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type FileController interface {
	Upload(write http.ResponseWriter, request *http.Request, _ httprouter.Params)
}
