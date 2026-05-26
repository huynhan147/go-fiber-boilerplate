package utils

import (
	"crypto/rand"
	"encoding/hex"
	"regexp"
	"strings"
	"unicode"
)

// Slugify chuyển "Hello World" → "hello-world"
func Slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)

	var b strings.Builder
	prevDash := false
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			prevDash = false
		} else if !prevDash {
			b.WriteRune('-')
			prevDash = true
		}
	}

	return strings.Trim(b.String(), "-")
}

// Truncate cắt chuỗi theo độ dài, thêm "..." nếu bị cắt
func Truncate(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	return string(runes[:maxLen-3]) + "..."
}

// MaskEmail che email "john@example.com" → "jo**@example.com"
func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email
	}
	name := parts[0]
	if len(name) <= 2 {
		return email
	}
	return name[:2] + strings.Repeat("*", len(name)-2) + "@" + parts[1]
}

// MaskPhone che số điện thoại "0912345678" → "091*****78"
func MaskPhone(phone string) string {
	r := []rune(phone)
	if len(r) < 6 {
		return phone
	}
	for i := 3; i < len(r)-2; i++ {
		r[i] = '*'
	}
	return string(r)
}

// RandomHex sinh chuỗi hex ngẫu nhiên với độ dài n bytes
func RandomHex(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// IsValidEmail kiểm tra định dạng email đơn giản
func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// IsValidPhone kiểm tra số điện thoại VN (10 số, bắt đầu 0)
func IsValidPhone(phone string) bool {
	re := regexp.MustCompile(`^0[0-9]{9}$`)
	return re.MatchString(phone)
}
