package utils

import (
	"fmt"
	"time"
)

const (
	DateFormat     = "02/01/2006"
	DateTimeFormat = "02/01/2006 15:04:05"
	APIDateFormat  = "2006-01-02"
)

// FormatDate → "02/01/2006"
func FormatDate(t time.Time) string {
	return t.Format(DateFormat)
}

// FormatDateTime → "02/01/2006 15:04:05"
func FormatDateTime(t time.Time) string {
	return t.Format(DateTimeFormat)
}

// FormatAPI → "2006-01-02" dùng cho API response
func FormatAPI(t time.Time) string {
	return t.Format(APIDateFormat)
}

// ParseDate parse "2006-01-02" → time.Time
func ParseDate(s string) (time.Time, error) {
	return time.Parse(APIDateFormat, s)
}

// StartOfDay trả về 00:00:00 của ngày đó
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay trả về 23:59:59 của ngày đó
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// TimeAgo trả về "vừa xong", "5 phút trước", "2 giờ trước"...
func TimeAgo(t time.Time) string {
	diff := time.Since(t)
	switch {
	case diff < time.Minute:
		return "vừa xong"
	case diff < time.Hour:
		return fmt.Sprintf("%d phút trước", int(diff.Minutes()))
	case diff < 24*time.Hour:
		return fmt.Sprintf("%d giờ trước", int(diff.Hours()))
	case diff < 7*24*time.Hour:
		return fmt.Sprintf("%d ngày trước", int(diff.Hours()/24))
	default:
		return FormatDate(t)
	}
}
