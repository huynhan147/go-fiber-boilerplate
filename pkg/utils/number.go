package utils

import (
	"fmt"
	"math"
	"strings"
)

// FormatVND "1500000" → "1.500.000 ₫"
func FormatVND(amount float64) string {
	intPart := int64(amount)
	s := fmt.Sprintf("%d", intPart)

	// Thêm dấu chấm phân cách hàng nghìn
	var result strings.Builder
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result.WriteRune('.')
		}
		result.WriteRune(c)
	}
	return result.String() + " ₫"
}

// RoundFloat làm tròn float đến n chữ số thập phân
func RoundFloat(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// Percentage tính phần trăm: Percentage(25, 200) → 12.5
func Percentage(part, total float64) float64 {
	if total == 0 {
		return 0
	}
	return RoundFloat((part/total)*100, 2)
}

// Clamp giữ giá trị trong khoảng [min, max]
func Clamp(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}
