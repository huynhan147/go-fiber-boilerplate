package logger

import (
	"context"

	"go.uber.org/zap"
)

type contextKey string

const loggerKey contextKey = "logger"

// WithRequestID tạo logger mới gắn request_id, lưu vào context
func WithRequestID(ctx context.Context, requestID string) context.Context {
	l := Log.With(zap.String("request_id", requestID))
	return context.WithValue(ctx, loggerKey, l)
}

// FromCtx lấy logger từ context — tự động có request_id
// Nếu không tìm thấy trả về logger gốc (không bị panic)
func FromCtx(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(loggerKey).(*zap.Logger); ok && l != nil {
		return l
	}
	return Log
}

// Shorthand dùng FromCtx luôn — tiện hơn khi gọi trong service
func InfoCtx(ctx context.Context, msg string, fields ...zap.Field) {
	FromCtx(ctx).Info(msg, fields...)
}

func WarnCtx(ctx context.Context, msg string, fields ...zap.Field) {
	FromCtx(ctx).Warn(msg, fields...)
}

func ErrorCtx(ctx context.Context, msg string, fields ...zap.Field) {
	FromCtx(ctx).Error(msg, fields...)
}

func DebugCtx(ctx context.Context, msg string, fields ...zap.Field) {
	FromCtx(ctx).Debug(msg, fields...)
}
