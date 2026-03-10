package service

import (
	"fmt"
	"io"
	"sims-ppob/model/web"
	"sims-ppob/repository"
	"time"
)

type FileServiceImpl struct {
	FileRepository repository.FileRepository
}

func NewFileService(fileRepository repository.FileRepository) FileService {
	return &FileServiceImpl{
		FileRepository: fileRepository,
	}
}

// Upload implements [FileService].
func (service *FileServiceImpl) Upload(file io.Reader, filename string) web.UploadFileResponse {

	newFileName := fmt.Sprintf("%d_%s", time.Now().Unix(), filename)

	path := "./uploads/" + newFileName

	err := service.FileRepository.Save(path, file)
	if err != nil {
		panic(err)
	}

	return web.UploadFileResponse{
		FileName: newFileName,
	}
}
