package domain

import "time"

type Transaksi struct {
	Transaksi_id      int
	Userbalance_id    int
	User_id           int
	Saldo_terakhir    int64
	Saldo_masuk       int64
	Saldo_keluar      int64
	Saldo_sekarang    int64
	Jenistransaksi_id int
	Tgl_transaksi     time.Time
	Created_at        time.Time
	Update_at         time.Time
	Deleted_at        time.Time
	JenisTransaksi    JenisTransaksi
}
