package repository

import (
	"io"
	"os"
	"sims-ppob/helper"
)

type FileRepositoryImpl struct{}

func NewFileRepository() FileRepository {
	return &FileRepositoryImpl{}
}

func (r *FileRepositoryImpl) Save(path string, file io.Reader) error {

	dst, err := os.Create(path)
	if err != nil {
		helper.PanicIfError(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	return err
}
