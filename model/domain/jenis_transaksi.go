package domain

import "time"

type JenisTransaksi struct {
	JenisTransaksi_id int
	Jenis_transaksi   string
	Created_at        time.Time
	Updated_at        time.Time
	Deleted_at        time.Time
}
