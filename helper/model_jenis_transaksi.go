package helper

import (
	"sims-ppob/model/domain"
	"sims-ppob/model/web"
)

func ToJenisTransaksiResponse(jenisTransaksi domain.JenisTransaksi) web.JenisTransaksiResponse {
	return web.JenisTransaksiResponse{
		Jenistransaksi_id: jenisTransaksi.JenisTransaksi_id,
		Jenis_transakasi:  jenisTransaksi.Jenis_transaksi,
	}
}

func ToJenisTransaksiResponses(jenisTransaksis []domain.JenisTransaksi) []web.JenisTransaksiResponse {
	jenisTransaksiResponses := []web.JenisTransaksiResponse{}
	for _, jenisTransaksi := range jenisTransaksis {
		jenisTransaksiResponses = append(jenisTransaksiResponses, ToJenisTransaksiResponse(jenisTransaksi))
	}

	return jenisTransaksiResponses
}
