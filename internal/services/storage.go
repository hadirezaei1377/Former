package services

// TODO: implement codes related to foreign storages

import "mime/multipart"

type Storage interface {
	Upload(file multipart.File, fileName string) error
	Download(fileName string) ([]byte, error)
}
