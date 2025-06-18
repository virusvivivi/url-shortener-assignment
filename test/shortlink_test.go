// test: unit test cho shortlink
package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"urlshortener/handler"
	"urlshortener/model"
	"urlshortener/storage"
)

func TestCreateAndGetShortlink(t *testing.T) {
	// Tạo database tạm cho test
	dbPath := "test.db"
	store, err := storage.NewSQLiteStore(dbPath)
	if err != nil {
		t.Fatalf("lỗi tạo test database: %v", err)
	}
	defer func() {
		store.Close()
		os.Remove(dbPath)
	}()

	h := handler.NewShortlinkHandler(store)

	// Tạo shortlink
	body := []byte(`{"original_url": "https://example.com"}`)
	req := httptest.NewRequest("POST", "/api/shortlinks", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	h.CreateShortlink(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var resp map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid response: %v", err)
	}
	id := resp["id"]
	if id == "" {
		t.Fatal("missing id in response")
	}

	// Lấy chi tiết shortlink
	req2 := httptest.NewRequest("GET", "/api/shortlinks/"+id, nil)
	rec2 := httptest.NewRecorder()
	h.GetShortlink(rec2, req2)
	if rec2.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec2.Code)
	}
	var detail map[string]interface{}
	if err := json.Unmarshal(rec2.Body.Bytes(), &detail); err != nil {
		t.Fatalf("invalid detail response: %v", err)
	}
	if detail["original_url"] != "https://example.com" {
		t.Errorf("expected original_url to match")
	}
	if _, err := time.Parse(time.RFC3339, detail["created_at"].(string)); err != nil {
		t.Errorf("invalid created_at format")
	}
}

func TestRedirectShortlink(t *testing.T) {
	// Tạo database tạm cho test
	dbPath := "test_redirect.db"
	store, err := storage.NewSQLiteStore(dbPath)
	if err != nil {
		t.Fatalf("lỗi tạo test database: %v", err)
	}
	defer func() {
		store.Close()
		os.Remove(dbPath)
	}()

	h := handler.NewShortlinkHandler(store)
	link := &model.Shortlink{ID: "abc123", OriginalURL: "https://golang.org", CreatedAt: time.Now()}
	store.Save(link)

	// Test redirect
	req := httptest.NewRequest("GET", "/shortlinks/abc123", nil)
	rec := httptest.NewRecorder()
	h.RedirectShortlink(rec, req)
	if rec.Code != http.StatusFound {
		t.Fatalf("expected 302, got %d", rec.Code)
	}
	loc := rec.Header().Get("Location")
	if loc != "https://golang.org" {
		t.Errorf("expected redirect to golang.org, got %s", loc)
	}
} 