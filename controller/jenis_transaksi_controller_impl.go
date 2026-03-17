package controller

import (
	"net/http"
	"sims-ppob/helper"
	"sims-ppob/model/web"
	"sims-ppob/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type JenisTransaksiControllerImpl struct {
	JenisTransaksiService service.JenisTransaksiService
}

func NewJenisTransaksiController(jenisTransaksi service.JenisTransaksiService) JenisTransaksiController {
	return &JenisTransaksiControllerImpl{
		JenisTransaksiService: jenisTransaksi,
	}
}

// Create implements [JenisTransaksiController].
func (controller *JenisTransaksiControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	jenisTransaksiCreateRequest := web.JenisTransaksiCreateRequest{}
	helper.ReadFromRequestBody(request, &jenisTransaksiCreateRequest)

	jenisTransaksiResponse := controller.JenisTransaksiService.Create(request.Context(), jenisTransaksiCreateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success create jenis transaksi",
		Data:    jenisTransaksiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Update implements [JenisTransaksiController].
func (controller *JenisTransaksiControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	jenisTransaksiUpdateRequest := web.JenisTransaksiUpdateRequest{}
	helper.ReadFromRequestBody(request, &jenisTransaksiUpdateRequest)

	jenisTransaksiId := params.ByName("jenisTransaksiId")
	id, err := strconv.Atoi(jenisTransaksiId)
	helper.PanicIfError(err)

	jenisTransaksiUpdateRequest.Jenistransaksi_id = id

	jenisTransaksiResponse := controller.JenisTransaksiService.Update(request.Context(), jenisTransaksiUpdateRequest)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success update jenis transaksi",
		Data:    jenisTransaksiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// FindById implements [JenisTransaksiController].
func (controller *JenisTransaksiControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	jenisTransaksiId := params.ByName("jenisTransaksiId")
	id, err := strconv.Atoi(jenisTransaksiId)
	helper.PanicIfError(err)

	jenisTransaksiResponse := controller.JenisTransaksiService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success get jenis transaksi",
		Data:    jenisTransaksiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// FindAll implements [JenisTransaksiController].
func (controller *JenisTransaksiControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	page, _ := strconv.Atoi(request.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(request.URL.Query().Get("limit"))
	helper.DefaultPageLimit(&page, &limit)

	jenisTransaksiResponse, paging := controller.JenisTransaksiService.FindAll(request.Context(), limit, page)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success get all jenis transaksi",
		Data:    jenisTransaksiResponse,
		Paging:  paging,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Delete implements [JenisTransaksiController].
func (controller *JenisTransaksiControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	jenisTransaksiId := params.ByName("jenisTransaksiId")
	id, err := strconv.Atoi(jenisTransaksiId)
	helper.PanicIfError(err)

	controller.JenisTransaksiService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "Success delete data jenis transaksi",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
