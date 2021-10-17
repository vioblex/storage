package storage

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vioblex/storage/internal/file"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}



func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(fileName string, blob []byte) (*file.File, error)  {
	newFile, err := file.NewFile(fileName, blob)
	if err != nil {
		return nil, err
	}

	s.files[newFile.ID] = newFile
	return newFile, nil
}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	foundFile, ok := s.files[fileID]
	if !ok {
		return nil, fmt.Errorf("file %v not found!", fileID)
	}

	return foundFile, nil
}