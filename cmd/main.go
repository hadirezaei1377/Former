package main

import (
	"former/internal/handlers"
	"former/internal/workers"
	"net/http"
)

func main() {

	go workers.StartWorkers()

	http.HandleFunc("/upload", handlers.UploadFileHandler)
	http.ListenAndServe(":8080", nil)
}
