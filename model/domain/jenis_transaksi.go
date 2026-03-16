package domain

import "time"

type JenisTransaksi struct {
	JenisTransaksi_id int
	jenis_transaksi   string
	Created_at        time.Time
	Updated_at        time.Time
	Deleted_at        time.Time
}
