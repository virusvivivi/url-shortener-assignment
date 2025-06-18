// model: struct dữ liệu shortlink
package model

import "time"

// Shortlink: 1 URL rút gọn
type Shortlink struct {
	ID          string
	OriginalURL string
	CreatedAt   time.Time
} 