package controller

import (
	"net/http"
	"sims-ppob/helper"
	"sims-ppob/model/web"
	"sims-ppob/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

// Login implements [UserController].
func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web.LoginRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	loginReponse := controller.UserService.Login(request.Context(), loginRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Login success",
		Data:    loginReponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.UserService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success create user",
		Data:    categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Update implements [UserController].
func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.User_id = id

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success update user",
		Data:    userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// FindById implements [UserController].
func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success get user data",
		Data:    userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// FindAll implements [UserController].
func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	page, _ := strconv.Atoi(request.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(request.URL.Query().Get("limit"))

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	userResponses, paging := controller.UserService.FindAll(request.Context(), page, limit)

	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success get all user data",
		Data:    userResponses,
		Paging:  paging,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Delete implements [UserController].
func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success remove data user",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
