package web

type TransaksiQueryRequest struct {
	JenisTransaksi *int    `json:"jenistransaksi"`
	TglTransaksi   *string `json:"tgltransaksi"`
	Page           int     `json:"page"`
	Limit          int     `json:"limit"`
}
