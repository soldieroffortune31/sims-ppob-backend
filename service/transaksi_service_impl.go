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

type TransaksiServiceImpl struct {
	TransaksiRepository      repository.TransaksiRepository
	UserRepository           repository.UserRepository
	UserBalanceRepository    repository.UserBalanceRepository
	JenisTransaksiRepository repository.JenisTransaksiRepository
	DB                       *sql.DB
	Validate                 *validator.Validate
}

func NewTransaksiBalance(transaksiRepository repository.TransaksiRepository, userRepository repository.UserRepository, userBalanceRepository repository.UserBalanceRepository, jenisTransaksiRepository repository.JenisTransaksiRepository, DB *sql.DB, validate *validator.Validate) TransaksiService {
	return &TransaksiServiceImpl{
		TransaksiRepository:      transaksiRepository,
		UserRepository:           userRepository,
		UserBalanceRepository:    userBalanceRepository,
		JenisTransaksiRepository: jenisTransaksiRepository,
		DB:                       DB,
		Validate:                 validate,
	}
}

// Save implements [TransaksiService].
func (service *TransaksiServiceImpl) Save(ctx context.Context, request web.TransaksiRequest) web.TransaksiResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.User_id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	jenisTransaksi, err := service.JenisTransaksiRepository.FindById(ctx, tx, request.JenisTransaksi_id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	userBalance, err := service.UserBalanceRepository.FindByUserId(ctx, tx, request.User_id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if request.Saldo_masuk == 0 && request.Saldo_keluar == 0 {
		panic(exception.NewBadRequestError("transaction must not be 0"))
	}

	saldoSekarang := userBalance.Balance
	saldoMasuk := request.Saldo_masuk
	saldoKeluar := request.Saldo_keluar

	if saldoMasuk > 0 {
		saldoSekarang = saldoSekarang + saldoMasuk
	}

	if saldoKeluar > 0 {
		saldoSekarang = saldoSekarang - saldoKeluar
	}

	// jika saldo minus maka gak boleh
	if saldoSekarang < 0 {
		panic(exception.NewBadRequestError("Insufficient Balance"))
	}

	transaksi := domain.Transaksi{
		Userbalance_id:    userBalance.Userbalance_id,
		User_id:           user.User_id,
		Saldo_terakhir:    userBalance.Balance,
		Saldo_masuk:       saldoMasuk,
		Saldo_keluar:      saldoKeluar,
		Saldo_sekarang:    saldoSekarang,
		Jenistransaksi_id: jenisTransaksi.JenisTransaksi_id,
		Tgl_transaksi:     request.Tgl_transaksi,
	}

	transaksi = service.TransaksiRepository.Save(ctx, tx, transaksi)

	userBalance.Balance = saldoSekarang
	userBalance = service.UserBalanceRepository.Update(ctx, tx, userBalance)

	return web.TransaksiResponse{
		Transaksi_id:      transaksi.Transaksi_id,
		Userbalance_id:    transaksi.Userbalance_id,
		User_id:           transaksi.User_id,
		Saldo_terakhir:    transaksi.Saldo_terakhir,
		Saldo_masuk:       transaksi.Saldo_masuk,
		Saldo_keluar:      transaksi.Saldo_keluar,
		Saldo_sekarang:    transaksi.Saldo_keluar,
		Jenistransaksi_id: transaksi.Transaksi_id,
		Tgl_transaksi:     transaksi.Tgl_transaksi,
	}

}

// Find implements [TransaksiService].
func (service *TransaksiServiceImpl) Find(ctx context.Context, request domain.TransaksiQuery) web.PagedResponse[web.TransaksiResponse] {
	panic("unimplemented")
}
