package service

import (
	"context"
	"database/sql"
	"sims-ppob/exception"
	"sims-ppob/helper"
	"sims-ppob/model/domain"
	"sims-ppob/model/web"
	"sims-ppob/repository"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	exist, err := service.UserRepository.IsEmailExist(ctx, tx, request.Email)

	if err != nil {
		panic(err)
	}

	if exist {
		panic(exception.NewConflictError("Email sudah terdaftar"))
	}

	if request.Password != request.Password_repeat {
		panic(exception.NewBadRequestError("Password tidak sama"))
	}

	hashedPassword := helper.HashPassord(request.Password)

	user := domain.User{
		Email:         request.Email,
		Nama_depan:    request.Nama_depan,
		Nama_belakang: request.Nama_belakang,
		Photo:         request.Photo,
		Password:      hashedPassword,
	}

	user = service.UserRepository.Save(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// cek data ada atau tidak by user_id
	user, err := service.UserRepository.FindById(ctx, tx, request.User_id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// cek email, apakah sudah digunakan atau belum
	exist, err := service.UserRepository.IsEmailExistByIdAndEmail(ctx, tx, request.User_id, request.Email)

	if err != nil {
		panic(err)
	}

	if exist {
		panic(exception.NewConflictError("Email sudah terdaftar"))
	}

	user.Email = request.Email
	user.Nama_depan = request.Nama_depan
	user.Nama_belakang = request.Nama_belakang
	user.Photo = request.Photo

	user = service.UserRepository.Update(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := service.UserRepository.FindAll(ctx, tx)

	return helper.ToUserResponses(user)
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepository.Delete(ctx, tx, user)
}
