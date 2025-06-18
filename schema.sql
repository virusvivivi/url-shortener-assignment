-- SQLite schema for URL shortener service
-- Tác giả: xem README.md

CREATE TABLE IF NOT EXISTS shortlinks (
    id TEXT PRIMARY KEY,
    original_url TEXT NOT NULL,
    created_at DATETIME NOT NULL
);

-- Index để tìm kiếm nhanh theo original_url
CREATE INDEX IF NOT EXISTS idx_original_url ON shortlinks(original_url);

-- Index để sắp xếp theo thời gian tạo
CREATE INDEX IF NOT EXISTS idx_created_at ON shortlinks(created_at); 