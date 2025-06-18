# Rút gọn URL
RESTful backend để short URLs bằng Go Language

## Tính năng
- Input: long URL, Output: short URL
- Sử dụng SQLite database để lưu trữ dữ liệu

## Hướng dẫn cài đặt

### Yêu cầu
- Go 1.21+
- SQLite3

### Kiểm tra Go đã cài chưa
go version
Nếu chưa cài Go, hãy làm theo hướng dẫn sau:

#### Cài đặt Go trên macOS
brew install go
hoặc tải từ trang chủ: https://golang.org/dl/
# Trên macOS
brew install sqlite3

#### Cài đặt Go trên Windows

Tải file .msi từ https://golang.org/dl/
Chạy file installer
Thêm Go vào PATH

#### Cài đặt Go trên Linux

Ubuntu/Debian
sudo apt update
sudo apt install golang-go

hoặc với (Centos)
sudo yum install golang

### Cài đặt
go mod tidy

### Chạy dịch vụ
go run main.go

Server chạy: http://localhost:8081

Database SQLite sẽ được tạo tự động tại file `shortlinks.db`

## API Endpoints

### 1. Tạo shortlink
- POST /api/shortlinks
- Request JSON:
{ "original_url": "https://example.com" }

- Trả về
{ "id": "abc123", "short_url": "http://localhost:8081/shortlinks/abc123" }

### 2. Lấy chi tiết shortlink
- GET /api/shortlinks/{id}
- Trả về
{ "id": "abc123", "original_url": "https://example.com", "created_at": "2025-06-18" }
  ```

### 3. Chuyển hướng đến URL gốc
- GET /shortlinks/{id}
- Response: Chuyển hướng HTTP 302

## Ví dụ

Tạo shortlink:
curl -X POST http://localhost:8081/api/shortlinks -H 'Content-Type: application/json' -d '{"original_url": "https://example.com"}'

Trên Terminal sẽ trả về kết quả:
{"id":"qgm9EsN","short_url":"http://localhost:8081/shortlinks/qgm9EsN"}

Test chuyển hướng:
curl -L http://localhost:8081/shortlinks/qgm9EsN

## Database Schema
File schema.sql được tạo như sau:
CREATE TABLE shortlinks (
    id TEXT PRIMARY KEY,
    original_url TEXT NOT NULL,
    created_at DATETIME NOT NULL
);

## Kiểm tra database - Xem nội dung database
sqlite3 shortlinks.db "SELECT * FROM shortlinks;"

Test trùng lặp URL:
curl -X POST http://localhost:8081/api/shortlinks -H 'Content-Type: application/json' -d '{"original_url": "https://example.com"}'
Trả về ID cũ:
{"id":"qgm9EsN","short_url":"http://localhost:8081/shortlinks/qgm9EsN"}

## Tác giả
Nguyễn Xuân Thảo