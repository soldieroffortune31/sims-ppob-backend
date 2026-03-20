package exception

import (
	"fmt"
	"net/http"
	"sims-ppob/helper"
	"sims-ppob/model/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	if conflictError(writer, request, err) {
		return
	}

	if badRequestError(writer, request, err) {
		return
	}

	if unauthorizedError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)

}

func unauthorizedError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(UnauthorizedError)

	if ok {
		writer.Header().Set("Content-type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:    http.StatusUnauthorized,
			Message: exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func badRequestError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(BadRequestError)

	if ok {
		writer.Header().Set("Content-type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func conflictError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(ConflictError)

	if ok {
		writer.Header().Set("Content-type", "application/json")
		writer.WriteHeader(http.StatusConflict)

		webResponse := web.WebResponse{
			Code:    http.StatusConflict,
			Message: exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	fmt.Println("ERROR", err)

	webResponse := web.WebResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)

	if ok {
		writer.Header().Set("Content-type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:    http.StatusBadRequest,
			Message: exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}

}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)

	if ok {
		writer.Header().Set("Content-type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webReponse := web.WebResponse{
			Code:    http.StatusNotFound,
			Message: exception.Error,
		}

		helper.WriteToResponseBody(writer, webReponse)

		return true
	} else {
		return false
	}
}
