package web

import "time"

type TransaksiRequest struct {
	User_id           int       `validate:"required" json:"user_id"`
	Saldo_masuk       int64     `validate:"required: min:0" json:"saldo_masuk"`
	Saldo_keluar      int64     `validate:"required min:0" json:"saldo_keluar"`
	JenisTransaksi_id int       `validate:"required" json:"jenistransaksi_id"`
	Tgl_transaksi     time.Time `validate:"required" json:"tgl_transaksi"`
}

type TransaksiQueryRequest struct {
	JenisTransaksi *int    `json:"jenistransaksi"`
	TglTransaksi   *string `json:"tgltransaksi"`
	Page           int     `json:"page"`
	Limit          int     `json:"limit"`
}
