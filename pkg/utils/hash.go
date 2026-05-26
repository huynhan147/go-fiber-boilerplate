package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword bcrypt hash mật khẩu
func HashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(b), err
}

// CheckPassword so sánh password với hash
func CheckPassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// MD5 hash chuỗi bằng MD5 (dùng cho cache key, gravatar...)
func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// SHA256 hash chuỗi bằng SHA256
func SHA256(s string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(s)))
}

// GravatarURL trả về URL avatar từ email qua Gravatar
func GravatarURL(email string, size int) string {
	hash := MD5(email)
	return fmt.Sprintf("https://www.gravatar.com/avatar/%s?s=%d&d=identicon", hash, size)
}
