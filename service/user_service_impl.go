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
	UserRepository        repository.UserRepository
	UserBalanceRepository repository.UserBalanceRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, userBalanceRepository repository.UserBalanceRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository:        userRepository,
		UserBalanceRepository: userBalanceRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

// Login implements [UserService].
func (service *UserServiceImpl) Login(ctx context.Context, request web.LoginRequest) web.LoginResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.Login(ctx, tx, request.Email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	checkPassword := helper.CheckPassword(request.Password, user.Password)
	if !checkPassword {
		panic(exception.NewBadRequestError("Password salah"))
	}

	token, err := helper.GenerateJWT(user.User_id, user.Email)
	if err != nil {
		panic(err)
	}

	service.UserRepository.UpdateToken(ctx, tx, user.User_id, token)

	return web.LoginResponse{
		Token: token,
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
		Photo:         helper.EmptyStringToNil(request.Photo),
		Password:      hashedPassword,
	}

	user = service.UserRepository.Save(ctx, tx, user)

	userBalance := domain.UserBalance{
		User_id: user.User_id,
		Balance: 0,
	}

	service.UserBalanceRepository.Save(ctx, tx, userBalance)

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

func (service *UserServiceImpl) FindAll(ctx context.Context, page int, limit int) ([]web.UserResponse, web.Paging) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	offset := (page - 1) * limit

	totalData := service.UserRepository.Count(ctx, tx)

	user := service.UserRepository.FindAll(ctx, tx, limit, offset)

	totalPage := totalData / limit
	if totalData%limit != 0 {
		totalPage++
	}

	paging := web.Paging{
		Page:      page,
		Limit:     limit,
		TotalPage: totalPage,
		Total:     totalData,
	}

	return helper.ToUserResponses(user), paging
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
