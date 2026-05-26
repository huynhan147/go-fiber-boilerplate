package middleware

import (
	"myapp/pkg/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const HeaderRequestID = "X-Request-ID"

// RequestID middleware:
// 1. Lấy request_id từ header nếu có (forwarded từ gateway/load balancer)
// 2. Nếu không có → tự sinh UUID mới
// 3. Inject logger có request_id vào ctx
// 4. Log toàn bộ thông tin request + response khi xong
func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Lấy hoặc sinh request_id
		requestID := c.Get(HeaderRequestID)
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Trả request_id về client để dễ debug
		c.Set(HeaderRequestID, requestID)

		// Inject logger vào fiber context (dùng trong handler)
		c.Locals("request_id", requestID)

		// Inject vào context.Context (truyền xuống service/repo)
		ctx := logger.WithRequestID(c.UserContext(), requestID)
		c.SetUserContext(ctx)

		start := time.Now()

		// Log request bắt đầu
		logger.FromCtx(ctx).Info("→ Request started",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("ip", c.IP()),
			zap.String("user_agent", c.Get("User-Agent")),
		)

		// Xử lý request
		err := c.Next()

		// Log request kết thúc kèm thời gian xử lý
		duration := time.Since(start)
		status := c.Response().StatusCode()

		logFn := logger.FromCtx(ctx).Info
		if status >= 500 {
			logFn = logger.FromCtx(ctx).Error
		} else if status >= 400 {
			logFn = logger.FromCtx(ctx).Warn
		}

		logFn("← Request completed",
			zap.Int("status", status),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Duration("duration", duration),
		)

		return err
	}
}
