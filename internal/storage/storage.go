package storage

import "github.com/vioblex/storage/internal/file"

type Storage struct {}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(fileName string, blob []byte) (*file.File, error)  {
	return file.NewFile(fileName, blob)
}