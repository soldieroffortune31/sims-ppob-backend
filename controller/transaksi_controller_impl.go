package controller

import (
	"net/http"
	"sims-ppob/helper"
	"sims-ppob/model/web"
	"sims-ppob/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TransaksiControllerImpl struct {
	TransaksiService service.TransaksiService
}

func NewTransaksiController(transaksiService service.TransaksiService) TransaksiController {
	return &TransaksiControllerImpl{
		TransaksiService: transaksiService,
	}
}

// Create implements [TransaksiController].
func (controller *TransaksiControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transaksiRequest := web.TransaksiRequest{}
	helper.ReadFromRequestBody(request, &transaksiRequest)

	transaksiResponse := controller.TransaksiService.Save(request.Context(), transaksiRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success create transaction",
		Data:    transaksiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// FindAll implements [TransaksiController].
func (controller *TransaksiControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	query := request.URL.Query()

	page := 1
	limit := 10

	if (query.Get("page")) != "" {
		val, _ := strconv.Atoi(query.Get("page"))
		page = val
	}

	if (query.Get("limit")) != "" {
		val, _ := strconv.Atoi(query.Get("limit"))
		limit = val
	}

	var jenis *int
	if query.Get("jenistransaksi") != "" {
		val, _ := strconv.Atoi(query.Get("jenistransaksi"))
		jenis = &val
	}

	var tgl *string
	if query.Get("tgltransaksi") != "" {
		val := query.Get("tgltransaksi")
		tgl = &val
	}

	req := web.TransaksiQueryRequest{
		JenisTransaksi: jenis,
		TglTransaksi:   tgl,
		Page:           page,
		Limit:          limit,
	}

	transaksiResponse := controller.TransaksiService.Find(request.Context(), req)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success get transactions data",
		Data:    transaksiResponse.Data,
		Paging:  transaksiResponse.Meta,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
