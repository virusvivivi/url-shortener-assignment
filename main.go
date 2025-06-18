// main: RESTful URL shortener
package main

import (
	"log"
	"net/http"

	"urlshortener/handler"
	"urlshortener/storage"
)

func main() {
	// Khởi tạo SQLite store
	store, err := storage.NewSQLiteStore("shortlinks.db")
	if err != nil {
		log.Fatalf("lỗi khởi tạo database: %v", err)
	}
	defer store.Close()

	h := handler.NewShortlinkHandler(store)

	http.HandleFunc("/api/shortlinks", h.CreateShortlink)
	http.HandleFunc("/api/shortlinks/", h.GetShortlink)
	http.HandleFunc("/shortlinks/", h.RedirectShortlink)

	log.Println("Server chạy ở http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("lỗi chạy server: %v", err)
	}
} 