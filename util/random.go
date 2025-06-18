// util: hàm tiện ích cho shortlink
package util

import (
	"crypto/rand"
	"math/big"
	"net/url"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Sinh mã ngẫu nhiên
func GenerateShortCode(min, max int) string {
	length := min
	if max > min {
		delta := max - min + 1
		if n, err := rand.Int(rand.Reader, big.NewInt(int64(delta))); err == nil {
			length = min + int(n.Int64())
		}
	}
	b := make([]byte, length)
	for i := range b {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[idx.Int64()]
	}
	return string(b)
}

// Check URL hợp lệ
func IsValidURL(s string) bool {
	u, err := url.ParseRequestURI(s)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}
	return true
} 