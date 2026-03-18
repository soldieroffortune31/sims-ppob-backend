package service

import (
	"context"
	"sims-ppob/model/domain"
	"sims-ppob/model/web"
)

type TransaksiService interface {
	Save(ctx context.Context, request web.TransaksiRequest) web.TransaksiResponse
	Find(ctx context.Context, request domain.TransaksiQuery) web.PagedResponse[web.TransaksiResponse]
}
