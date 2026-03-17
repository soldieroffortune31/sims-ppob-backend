package repository

import (
	"context"
	"database/sql"
	"errors"
	"sims-ppob/helper"
	"sims-ppob/model/domain"
)

type JenisTransaksiRepositoryImpl struct {
}

func NewJenisTransaksi() JenisTransaksiRepository {
	return &JenisTransaksiRepositoryImpl{}
}

// Save implements [JenisTransaksiRepository].
func (repository *JenisTransaksiRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, jenisTransaksi domain.JenisTransaksi) domain.JenisTransaksi {
	SQL := "INSERT INTO jenistransaksi_m (jenis_transaksi, created_at, updated_at) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, jenisTransaksi.Jenis_transaksi, helper.GetTimeUTCNow(), helper.GetTimeUTCNow())
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	jenisTransaksi.JenisTransaksi_id = int(id)
	return jenisTransaksi
}

// Update implements [JenisTransaksiRepository].
func (repository *JenisTransaksiRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, jenisTransaksi domain.JenisTransaksi) domain.JenisTransaksi {
	SQL := "UPDATE jenistransaksi_m SET jenis_transaksi = ?, updated_at = ? WHERE jenistransaksi_id = ?"
	_, err := tx.ExecContext(ctx, SQL, jenisTransaksi.Jenis_transaksi, helper.GetTimeUTCNow(), jenisTransaksi.JenisTransaksi_id)
	helper.PanicIfError(err)

	return jenisTransaksi
}

// FindById implements [JenisTransaksiRepository].
func (repository *JenisTransaksiRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, jenistransaksiId int) (domain.JenisTransaksi, error) {
	SQL := "SELECT jenistransaksi_id, jenis_transaksi FROM jenistransaksi_m WHERE jenistransaksi_id = ? AND deleted_at IS NULL"
	rows, err := tx.QueryContext(ctx, SQL, jenistransaksiId)
	helper.PanicIfError(err)
	defer rows.Close()

	jenisTransaksi := domain.JenisTransaksi{}
	if rows.Next() {
		err := rows.Scan(&jenisTransaksi.JenisTransaksi_id, &jenisTransaksi.Jenis_transaksi)
		helper.PanicIfError(err)
		return jenisTransaksi, nil
	} else {
		return jenisTransaksi, errors.New("Jenis transaksi is not found")
	}
}

// Count implements [JenisTransaksiRepository].
func (repository *JenisTransaksiRepositoryImpl) Count(ctx context.Context, tx *sql.Tx) int {
	SQL := "SELECT COUNT(*) FROM jenistransaksi_m WHERE deleted_at IS NULL"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	total := 0
	if rows.Next() {
		err := rows.Scan(&total)
		helper.PanicIfError(err)
	}

	return total
}

// FindAll implements [JenisTransaksiRepository].
func (repository *JenisTransaksiRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, limit int, offset int) []domain.JenisTransaksi {
	SQL := "SELECT jenistransaksi_id, jenis_transaksi FROM jenistransaksi_m WHERE deleted_at IS NULL LIMIT ? OFFSET ?"
	rows, err := tx.QueryContext(ctx, SQL, limit, offset)
	helper.PanicIfError(err)
	defer rows.Close()

	var jenisTransaksis []domain.JenisTransaksi
	for rows.Next() {
		jenisTransaksi := domain.JenisTransaksi{}
		err := rows.Scan(&jenisTransaksi.JenisTransaksi_id, &jenisTransaksi.Jenis_transaksi)
		helper.PanicIfError(err)
		jenisTransaksis = append(jenisTransaksis, jenisTransaksi)
	}

	return jenisTransaksis
}

// Delete implements [JenisTransaksiRepository].
func (repository *JenisTransaksiRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, jenistransaksiId int) {
	SQL := "UPDATE FROM jenistransaksi_m SET deleted_at = ?"
	_, err := tx.ExecContext(ctx, SQL, helper.GetTimeUTCNow())
	helper.PanicIfError(err)
}
