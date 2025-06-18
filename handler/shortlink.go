// handler: Xử lý HTTP cho shortlink
package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"urlshortener/model"
	"urlshortener/storage"
	"urlshortener/util"
)

type ShortlinkHandler struct {
	store storage.Store
}

func NewShortlinkHandler(store storage.Store) *ShortlinkHandler {
	return &ShortlinkHandler{store: store}
}

// Tạo shortlink (POST)
func (h *ShortlinkHandler) CreateShortlink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		OriginalURL string `json:"original_url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if !util.IsValidURL(req.OriginalURL) {
		http.Error(w, "Invalid original_url", http.StatusBadRequest)
		return
	}
	// Check trùng URL
	if existing := h.store.FindByOriginalURL(req.OriginalURL); existing != nil {
		resp := map[string]string{
			"id":       existing.ID,
			"short_url": "http://localhost:8081/shortlinks/" + existing.ID,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}
	id := util.GenerateShortCode(6, 8)
	shortlink := &model.Shortlink{
		ID:          id,
		OriginalURL: req.OriginalURL,
		CreatedAt:   time.Now().UTC(),
	}
	
	// Lưu vào database
	if err := h.store.Save(shortlink); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	
	resp := map[string]string{
		"id":       id,
		"short_url": "http://localhost:8081/shortlinks/" + id,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Lấy chi tiết shortlink (GET)
func (h *ShortlinkHandler) GetShortlink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/api/shortlinks/")
	if id == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}
	shortlink := h.store.FindByID(id)
	if shortlink == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	resp := map[string]interface{}{
		"id":           shortlink.ID,
		"original_url": shortlink.OriginalURL,
		"created_at":   shortlink.CreatedAt.Format(time.RFC3339),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Redirect shortlink (GET)
func (h *ShortlinkHandler) RedirectShortlink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/shortlinks/")
	if id == "" {
		http.Error(w, "Missing id", http.StatusBadRequest)
		return
	}
	shortlink := h.store.FindByID(id)
	if shortlink == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, shortlink.OriginalURL, http.StatusFound)
} 