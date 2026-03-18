package repository

import (
	"context"
	"database/sql"
	"sims-ppob/helper"
	"sims-ppob/model/domain"
)

type TransaksiRepositoryImpl struct {
}

func NewTransaksiRepository() TransaksiRepository {
	return &TransaksiRepositoryImpl{}
}

// Save implements [TransaksiRepository].
func (t *TransaksiRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, transaksi domain.Transaksi) domain.Transaksi {
	SQL := "INSERT INTO transaksi_t (userbalance_id, user_id, saldo_terakhir, saldo_masuk, saldo_keluar, saldo_sekarang, jenistransaksi_id, tgl_transaksi, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, transaksi.Userbalance_id, transaksi.User_id, transaksi.Saldo_terakhir, transaksi.Saldo_masuk, transaksi.Saldo_keluar, transaksi.Saldo_sekarang, transaksi.Jenistransaksi_id, transaksi.Tgl_transaksi, helper.GetTimeUTCNow(), helper.GetTimeUTCNow())
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	transaksi.Transaksi_id = int(id)
	return transaksi
}

// FindAll implements [TransaksiRepository].
func (t *TransaksiRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, transaksi domain.TransaksiQuery) ([]domain.Transaksi, int) {
	panic("unimplemented")
}
