package domain

type TransaksiQuery struct {
	JenisTransaksi *int
	TglTransaksi   *string
	Page           int
	Limit          int
}
