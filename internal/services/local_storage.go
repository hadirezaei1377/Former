package services

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	BasePath string
}

func NewLocalStorage(basePath string) *LocalStorage {
	return &LocalStorage{
		BasePath: basePath,
	}
}

// upload file in local
func (l *LocalStorage) Upload(file multipart.File, fileName string) error {
	filePath := filepath.Join(l.BasePath, fileName)
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	return err
}

// download file from local
func (l *LocalStorage) Download(fileName string) ([]byte, error) {
	filePath := filepath.Join(l.BasePath, fileName)
	return os.ReadFile(filePath)
}
