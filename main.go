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
	userBalanceRepository := repository.NewUserBalanceRepository()
	userService := service.NewUserService(userRepository, userBalanceRepository, db, validate)
	userController := controller.NewUserController(userService)

	fileRepository := repository.NewFileRepository()
	fileService := service.NewFileService(fileRepository)
	fileController := controller.NewFileController(fileService)

	authMiddleware := middleware.AuthMiddleware(userRepository, db)

	jenisTransaksiRepository := repository.NewJenisTransaksi()
	jenisTransaksiService := service.NewJenisTransaksiService(jenisTransaksiRepository, db, validate)
	jenisTransaksiController := controller.NewJenisTransaksiController(jenisTransaksiService)

	transaksiRepository := repository.NewTransaksiRepository()
	transaksiService := service.NewTransaksiService(transaksiRepository, userRepository, userBalanceRepository, jenisTransaksiRepository, db, validate)
	transaksiController := controller.NewTransaksiController(transaksiService)

	router := httprouter.New()

	//public
	router.POST("/api/user", userController.Create)
	router.POST("/api/login", userController.Login)
	router.POST("/api/upload", fileController.Upload)
	router.ServeFiles("/uploads/*filepath", http.Dir("./uploads"))

	//private
	router.PUT("/api/user/:userId", authMiddleware(userController.Update))
	router.GET("/api/user/:userId", authMiddleware(userController.FindById))
	router.GET("/api/user", authMiddleware(userController.FindAll))
	router.DELETE("/api/user/:userId", authMiddleware(userController.Delete))

	router.POST("/api/jenis-transaksi", authMiddleware(jenisTransaksiController.Create))
	router.PUT("/api/jenis-transaksi/:jenisTransaksiId", authMiddleware(jenisTransaksiController.Update))
	router.GET("/api/jenis-transaksi/:jenisTransaksiId", authMiddleware(jenisTransaksiController.FindById))
	router.GET("/api/jenis-transaksi", authMiddleware(jenisTransaksiController.FindAll))
	router.DELETE("/api/jenis-transaksi/:jenisTransaksiId", authMiddleware(jenisTransaksiController.Delete))

	router.POST("/api/transaksi", authMiddleware(transaksiController.Create))

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
