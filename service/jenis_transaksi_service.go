package service

import (
	"context"
	"sims-ppob/model/web"
)

type JenisTransaksiService interface {
	Create(ctx context.Context, request web.JenisTransaksiCreateRequest) web.JenisTransaksiResponse
	Update(ctx context.Context, request web.JenisTransaksiUpdateRequest) web.JenisTransaksiResponse
	FindById(ctx context.Context, jenisTransaksiId int) web.JenisTransaksiResponse
	FindAll(ctx context.Context, limit int, page int) ([]web.JenisTransaksiResponse, web.Paging)
	Delete(ctx context.Context, jenisTransaksiId int)
}
