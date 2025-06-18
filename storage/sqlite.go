// storage: SQLite database
package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"urlshortener/model"
)

type sqliteStore struct {
	db *sql.DB
}

func NewSQLiteStore(dbPath string) (Store, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Tạo bảng nếu chưa có
	if err := createTable(db); err != nil {
		return nil, err
	}

	return &sqliteStore{db: db}, nil
}

func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS shortlinks (
		id TEXT PRIMARY KEY,
		original_url TEXT NOT NULL,
		created_at DATETIME NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_original_url ON shortlinks(original_url);
	CREATE INDEX IF NOT EXISTS idx_created_at ON shortlinks(created_at);
	`
	_, err := db.Exec(query)
	return err
}

func (s *sqliteStore) Save(link *model.Shortlink) error {
	query := `INSERT OR REPLACE INTO shortlinks (id, original_url, created_at) VALUES (?, ?, ?)`
	_, err := s.db.Exec(query, link.ID, link.OriginalURL, link.CreatedAt)
	return err
}

func (s *sqliteStore) FindByID(id string) *model.Shortlink {
	query := `SELECT id, original_url, created_at FROM shortlinks WHERE id = ?`
	var link model.Shortlink
	err := s.db.QueryRow(query, id).Scan(&link.ID, &link.OriginalURL, &link.CreatedAt)
	if err != nil {
		return nil
	}
	return &link
}

func (s *sqliteStore) FindByOriginalURL(url string) *model.Shortlink {
	query := `SELECT id, original_url, created_at FROM shortlinks WHERE original_url = ?`
	var link model.Shortlink
	err := s.db.QueryRow(query, url).Scan(&link.ID, &link.OriginalURL, &link.CreatedAt)
	if err != nil {
		return nil
	}
	return &link
}

func (s *sqliteStore) Close() error {
	return s.db.Close()
} 