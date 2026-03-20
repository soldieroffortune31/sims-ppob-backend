package helper

import (
	"sims-ppob/model/domain"
	"sims-ppob/model/web"
)

func ToTransaksiResponse(transaksi domain.Transaksi) web.TransaksiResponse {

	jenisTransaksiResponse := web.JenisTransaksiResponse{
		Jenistransaksi_id: transaksi.Jenistransaksi_id,
		Jenis_transakasi:  transaksi.JenisTransaksi.Jenis_transaksi,
	}

	return web.TransaksiResponse{
		Transaksi_id:      transaksi.Transaksi_id,
		Userbalance_id:    transaksi.Userbalance_id,
		User_id:           transaksi.User_id,
		Saldo_terakhir:    transaksi.Saldo_terakhir,
		Saldo_masuk:       transaksi.Saldo_masuk,
		Saldo_keluar:      transaksi.Saldo_keluar,
		Saldo_sekarang:    transaksi.Saldo_sekarang,
		Jenistransaksi_id: transaksi.Jenistransaksi_id,
		Tgl_transaksi:     transaksi.Tgl_transaksi,
		JenisTransaksi:    jenisTransaksiResponse,
	}
}

func ToTransaksiResponses(transaksis []domain.Transaksi) []web.TransaksiResponse {
	transaksiResponses := []web.TransaksiResponse{}
	for _, transaksi := range transaksis {
		transaksiResponses = append(transaksiResponses, ToTransaksiResponse(transaksi))
	}

	return transaksiResponses
}
