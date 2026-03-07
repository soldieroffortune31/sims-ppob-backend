package main

import (
	"fmt"
	"net/http"
	"sims-ppob/app"
	"sims-ppob/controller"
	"sims-ppob/exception"
	"sims-ppob/helper"
	"sims-ppob/repository"
	"sims-ppob/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDb()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	router := httprouter.New()

	router.POST("/api/user", userController.Create)
	router.PUT("/api/user/:userId", userController.Update)
	router.GET("/api/user/:userId", userController.FindById)
	router.GET("/api/user", userController.FindAll)
	router.DELETE("/api/user/:userId", userController.Delete)

	router.POST("/api/login", userController.Login)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("Server running on http://localhost:3000")

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
