package handlers

import (
	"Former/internal/config"
	"fmt"
	"net/http"

	"github.com/benmanns/goworker"
)

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

	compress := r.FormValue("compress")
	if compress == "true" {
		err = goworker.Enqueue(&goworker.Job{
			Queue: "compress_queue",
			Payload: goworker.Payload{
				Class: "Compress",
				Args:  []interface{}{handler.Filename, handler.Filename},
			},
		})
		if err != nil {
			http.Error(w, "Failed to enqueue compression job", http.StatusInternalServerError)
			return
		}
	}

	fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
}
