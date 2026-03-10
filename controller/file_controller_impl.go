package controller

import (
	"net/http"
	"sims-ppob/exception"
	"sims-ppob/helper"
	"sims-ppob/model/web"
	"sims-ppob/service"

	"github.com/julienschmidt/httprouter"
)

type FileControllerImpl struct {
	FileService service.FileService
}

func NewFileController(fileService service.FileService) FileController {
	return &FileControllerImpl{
		FileService: fileService,
	}
}

// Upload implements [FileController].
func (controller *FileControllerImpl) Upload(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	// limit body request max 5MB
	// r.Body = http.MaxBytesReader(w, r.Body, 5<<20)

	err := request.ParseMultipartForm(10 << 20)
	if err != nil {
		panic(exception.NewBadRequestError("Cannot parse form"))
	}

	file, handler, err := request.FormFile("image")
	if err != nil {
		panic(exception.NewBadRequestError("File error"))
	}
	defer file.Close()

	// / validasi ukuran file
	// if handler.Size > 5<<20 {
	// 	http.Error(w, "Ukuran file maksimal 5MB", http.StatusBadRequest)
	// 	return
	// }

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		panic(exception.NewBadRequestError("Cannot read file"))
	}

	fileType := http.DetectContentType(buffer)

	allowedTypes := map[string]bool{
		"image/jpeg":      true,
		"image/png":       true,
		"application/pdf": true,
	}

	if !allowedTypes[fileType] {
		panic(exception.NewBadRequestError("File type not allowed"))
	}

	// reset pointer file ke awal
	file.Seek(0, 0)

	fileResponse := controller.FileService.Upload(file, handler.Filename)
	webResponse := web.WebResponse{
		Code:    200,
		Message: "success upload file",
		Data:    fileResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
