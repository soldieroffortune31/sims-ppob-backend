package repository

import (
	"context"
	"database/sql"
	"sims-ppob/model/domain"
	"sims-ppob/model/web"
)

type TransaksiRepository interface {
	Save(ctx context.Context, tx *sql.Tx, transaksi domain.Transaksi) domain.Transaksi
	FindAll(ctx context.Context, tx *sql.Tx, transaksi web.TransaksiQueryRequest) ([]domain.Transaksi, int)
}
