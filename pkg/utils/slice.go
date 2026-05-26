package utils

// Contains kiểm tra slice có chứa phần tử không
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Map biến đổi từng phần tử trong slice
func Map[T, R any](slice []T, fn func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter lọc các phần tử thỏa điều kiện
func Filter[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Unique loại bỏ phần tử trùng lặp
func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	var result []T
	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Chunk chia slice thành các nhóm nhỏ: Chunk([1,2,3,4,5], 2) → [[1,2],[3,4],[5]]
func Chunk[T any](slice []T, size int) [][]T {
	if size <= 0 {
		return nil
	}
	var chunks [][]T
	for size < len(slice) {
		slice, chunks = slice[size:], append(chunks, slice[:size])
	}
	return append(chunks, slice)
}
