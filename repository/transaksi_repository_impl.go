package repository

import (
	"context"
	"database/sql"
	"sims-ppob/helper"
	"sims-ppob/model/domain"
	"sims-ppob/model/web"
)

type TransaksiRepositoryImpl struct {
	BaseRepository
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
func (t *TransaksiRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, transaksi web.TransaksiQueryRequest) ([]domain.Transaksi, int) {
	baseQuery := "SELECT a.transaksi_id, a.userbalance_id, a.user_id, a.saldo_terakhir, a.saldo_masuk, a.saldo_keluar, a.saldo_sekarang, a.jenistransaksi_id, a.tgl_transaksi, b.jenis_transaksi FROM transaksi_t a JOIN jenistransaksi_m b ON a.jenistransaksi_id = b.jenistransaksi_id"
	countQuery := "SELECT COUNT(*) FROM transaksi_t a JOIN jenistransaksi_m b ON a.jenistransaksi_id = b.jenistransaksi_id"

	builder := helper.NewFilterBuilder()

	if transaksi.JenisTransaksi != nil {
		builder.Add("a.jenistransaksi_id = ?", transaksi.JenisTransaksi)
	}

	if transaksi.TglTransaksi != nil {
		start, end, err := helper.ConvertDateToUTCRange(*transaksi.TglTransaksi)
		if err != nil {
			helper.PanicIfError(err)
		}

		builder.AddRaw("a.tgl_transaksi BETWEEN ? AND ?", start, end)
	}

	where, args := builder.BuildWhere()

	offset := helper.GetOffset(transaksi.Page, transaksi.Limit)

	rows, total := t.QueryWithPagination(ctx, tx, baseQuery, countQuery, where, args, transaksi.Limit, offset)
	defer rows.Close()

	var result []domain.Transaksi
	for rows.Next() {
		t := domain.Transaksi{}
		err := rows.Scan(&t.Transaksi_id, &t.Userbalance_id, &t.User_id, &t.Saldo_terakhir, &t.Saldo_masuk, &t.Saldo_keluar, &t.Saldo_sekarang, &t.Jenistransaksi_id, &t.Tgl_transaksi, &t.JenisTransaksi.Jenis_transaksi)
		helper.PanicIfError(err)
		result = append(result, t)
	}

	return result, total
}
