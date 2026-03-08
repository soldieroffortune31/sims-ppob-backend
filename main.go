package main

import (
	"fmt"
	"net/http"
	"sims-ppob/app"
	"sims-ppob/controller"
	"sims-ppob/exception"
	"sims-ppob/helper"
	"sims-ppob/middleware"
	"sims-ppob/repository"
	"sims-ppob/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {

	db := app.NewDb()
	validate := validator.New()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	authMiddleware := middleware.AuthMiddleware(userRepository, db)

	router := httprouter.New()

	//public
	router.POST("/api/user", userController.Create)
	router.POST("/api/login", userController.Login)

	//private
	router.PUT("/api/user/:userId", authMiddleware(userController.Update))
	router.GET("/api/user/:userId", authMiddleware(userController.FindById))
	router.GET("/api/user", authMiddleware(userController.FindAll))
	router.DELETE("/api/user/:userId", authMiddleware(userController.Delete))

	router.PanicHandler = exception.ErrorHandler

	// CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// server
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: corsHandler.Handler(router),
	}

	fmt.Println("Server running on http://localhost:8000")

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
