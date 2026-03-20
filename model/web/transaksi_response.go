package web

import "time"

type TransaksiResponse struct {
	Transaksi_id      int                    `json:"transaksi_id"`
	Userbalance_id    int                    `json:"userbalance_id"`
	User_id           int                    `json:"user_id"`
	Saldo_terakhir    int64                  `json:"saldo_terakhir"`
	Saldo_masuk       int64                  `json:"saldo_masuk"`
	Saldo_keluar      int64                  `json:"saldo_keluar"`
	Saldo_sekarang    int64                  `json:"saldo_sekarang"`
	Jenistransaksi_id int                    `json:"jenistransaksi_id"`
	Tgl_transaksi     time.Time              `json:"tgl_transaksi"`
	JenisTransaksi    JenisTransaksiResponse `json:"jenis_transaksi"`
}
