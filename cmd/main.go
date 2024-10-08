package main

import (
	"log"
	"net/http"

	"former/internal/handlers"
	"formere/internal/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/upload", handlers.UploadFileHandler).Methods("POST")

	r.HandleFunc("/download", handlers.DownloadFileHandler).Methods("GET")

	r.PathPrefix("/private/").Handler(middlewares.AuthMiddleware(http.StripPrefix("/private/", http.FileServer(http.Dir("./private/uploads/")))))

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
