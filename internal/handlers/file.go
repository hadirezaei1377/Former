package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"former/internal/config"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // limitation for file volume

	// we recieve the file from request
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error Retrieving the File", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileType := r.FormValue("type")
	var uploadDir string
	cfg := config.LoadConfig()

	if fileType == "public" {
		uploadDir = cfg.PublicDir
	} else {
		uploadDir = cfg.PrivateDir
	}

	filePath := filepath.Join(uploadDir, handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to save the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
}

func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("filename")
	if fileName == "" {
		http.Error(w, "Filename is missing", http.StatusBadRequest)
		return
	}

	cfg := config.LoadConfig()
	filePath := filepath.Join(cfg.PublicDir, fileName)

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/octet-stream")
	io.Copy(w, file)
}
