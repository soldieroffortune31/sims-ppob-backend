package service

import (
	"context"
	"sims-ppob/model/web"
)

type TransaksiService interface {
	Save(ctx context.Context, request web.TransaksiRequest) web.TransaksiResponse
	Find(ctx context.Context, request web.TransaksiQueryRequest) web.PagedResponse[web.TransaksiResponse]
}
