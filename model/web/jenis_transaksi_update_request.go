package web

type JenisTransaksiUpdateRequest struct {
	Jenistransaksi_id int    `json:"jenistransaksi_id"`
	Jenis_transakasi  string `json:"jenis_transaksi"`
}
