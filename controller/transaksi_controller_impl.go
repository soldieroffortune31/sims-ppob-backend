package controller

import (
	"net/http"
	"sims-ppob/helper"
	"sims-ppob/model/web"
	"sims-ppob/service"

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
