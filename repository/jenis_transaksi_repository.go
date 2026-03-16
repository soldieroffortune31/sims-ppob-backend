package repository

import (
	"context"
	"database/sql"
	"sims-ppob/model/domain"
)

type JenisTransaksiRepository interface {
	Save(ctx context.Context, tx *sql.Tx, jenisTransaksi domain.JenisTransaksi) domain.JenisTransaksi
	Update(ctx context.Context, tx *sql.Tx, jenisTransaksi domain.JenisTransaksi) domain.JenisTransaksi
	FindById(ctx context.Context, tx *sql.Tx, jenistransaksiId int) (domain.JenisTransaksi, error)
	Count(ctx context.Context, tx *sql.Tx) int
	FindAll(ctx context.Context, tx *sql.Tx, limit int, offset int) []domain.JenisTransaksi
	Delete(ctx context.Context, tx *sql.Tx, jenistransaksiId int)
}
