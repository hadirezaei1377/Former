package handlers

import (
	"fmt"
	"net/http"

	"former/internal/config"
	"former/internal/services"
)

func getStorage(storageType string, cfg *config.Config) (services.Storage, error) {
	switch storageType {
	case "s3":
		return services.NewS3Storage("your-bucket-name", "your-region"), nil
	case "local":
		return services.NewLocalStorage(cfg.PublicDir), nil
	}
	return nil, fmt.Errorf("Invalid storage type")
}

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error Retrieving the File", http.StatusBadRequest)
		return
	}
	defer file.Close()

	storageType := r.FormValue("storage_type")
	cfg := config.LoadConfig()

	storage, err := getStorage(storageType, cfg)
	if err != nil {
		http.Error(w, "Invalid storage type", http.StatusBadRequest)
		return
	}

	err = storage.Upload(file, handler.Filename)
	if err != nil {
		http.Error(w, "Unable to upload file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
}
