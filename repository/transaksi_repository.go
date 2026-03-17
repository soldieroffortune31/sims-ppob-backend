package repository

import (
	"context"
	"database/sql"
	"sims-ppob/model/domain"
)

type TransaksiRepository interface {
	Save(ctx context.Context, tx *sql.Tx, transaksi domain.Transaksi) domain.Transaksi
	FindAll(ctx context.Context, tx *sql.Tx, transaksi domain.TransaksiQuery) ([]domain.Transaksi, int)
}
