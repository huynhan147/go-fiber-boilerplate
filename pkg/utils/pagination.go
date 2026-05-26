package utils

import "github.com/gofiber/fiber/v2"

type Pagination struct {
	Page  int
	Limit int
}

// PaginationFromCtx lấy page và limit từ query string
// mặc định: page=1, limit=15, tối đa limit=100
func PaginationFromCtx(c *fiber.Ctx) Pagination {
	page  := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 15)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 15
	}
	if limit > 100 {
		limit = 100
	}

	return Pagination{Page: page, Limit: limit}
}

// Offset tính offset cho SQL query
func (p Pagination) Offset() int {
	return (p.Page - 1) * p.Limit
}

// TotalPages tính tổng số trang
func (p Pagination) TotalPages(total int64) int64 {
	pages := total / int64(p.Limit)
	if total%int64(p.Limit) > 0 {
		pages++
	}
	return pages
}
