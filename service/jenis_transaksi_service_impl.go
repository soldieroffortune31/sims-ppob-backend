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

type JenisTransaksiServiceImpl struct {
	JenisTransaksiRepository repository.JenisTransaksiRepository
	DB                       *sql.DB
	Validate                 *validator.Validate
}

func NewJenisTransaksiService(jenisTransaksiRepository repository.JenisTransaksiRepository, DB *sql.DB, validate *validator.Validate) JenisTransaksiService {
	return &JenisTransaksiServiceImpl{
		JenisTransaksiRepository: jenisTransaksiRepository,
		DB:                       DB,
		Validate:                 validate,
	}
}

// Create implements [JenisTransaksiService].
func (service *JenisTransaksiServiceImpl) Create(ctx context.Context, request web.JenisTransaksiCreateRequest) web.JenisTransaksiResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	jenisTransaksi := domain.JenisTransaksi{
		Jenis_transaksi: request.Jenis_transakasi,
	}

	jenisTransaksi = service.JenisTransaksiRepository.Save(ctx, tx, jenisTransaksi)

	return helper.ToJenisTransaksiResponse(jenisTransaksi)
}

// Update implements [JenisTransaksiService].
func (service *JenisTransaksiServiceImpl) Update(ctx context.Context, request web.JenisTransaksiUpdateRequest) web.JenisTransaksiResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	jenisTransaksi, err := service.JenisTransaksiRepository.FindById(ctx, tx, request.Jenistransaksi_id)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	jenisTransaksi.Jenis_transaksi = request.Jenis_transakasi

	jenisTransaksi = service.JenisTransaksiRepository.Update(ctx, tx, jenisTransaksi)

	return helper.ToJenisTransaksiResponse(jenisTransaksi)
}

// FindById implements [JenisTransaksiService].
func (service *JenisTransaksiServiceImpl) FindById(ctx context.Context, jenisTransaksiId int) web.JenisTransaksiResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	jenisTransaksi, err := service.JenisTransaksiRepository.FindById(ctx, tx, jenisTransaksiId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToJenisTransaksiResponse(jenisTransaksi)
}

// FindAll implements [JenisTransaksiService].
func (service *JenisTransaksiServiceImpl) FindAll(ctx context.Context, limit int, page int) ([]web.JenisTransaksiResponse, web.Paging) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	offset := (page - 1) * limit

	totalData := service.JenisTransaksiRepository.Count(ctx, tx)

	jenisTransaksi := service.JenisTransaksiRepository.FindAll(ctx, tx, limit, offset)

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

	return helper.ToJenisTransaksiResponses(jenisTransaksi), paging

}

// Delete implements [JenisTransaksiService].
func (service *JenisTransaksiServiceImpl) Delete(ctx context.Context, jenisTransaksiId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	service.JenisTransaksiRepository.Delete(ctx, tx, jenisTransaksiId)
}
