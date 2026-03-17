package web

type JenisTransaksiCreateRequest struct {
	Jenis_transaksi string `validate:"required,max=100" json:"jenis_transaksi"`
}
