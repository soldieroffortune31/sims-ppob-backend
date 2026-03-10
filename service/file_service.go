package service

import (
	"io"
	"sims-ppob/model/web"
)

type FileService interface {
	Upload(file io.Reader, filename string) web.UploadFileResponse
}
