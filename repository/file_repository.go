package repository

import "io"

type FileRepository interface {
	Save(path string, file io.Reader) error
}
