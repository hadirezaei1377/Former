package services

import (
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/disintegration/imaging"
)

func CompressImage(file multipart.File, fileName string, quality int) (string, error) {
	img, err := imaging.Decode(file)
	if err != nil {
		return "", err
	}

	compressedFileName := fmt.Sprintf("compressed_%s", fileName)
	compressedFilePath := filepath.Join("/tmp", compressedFileName)
	err = imaging.Save(imaging.Resize(img, 800, 600, imaging.Lanczos), compressedFilePath)
	if err != nil {
		return "", err
	}

	return compressedFilePath, nil
}
