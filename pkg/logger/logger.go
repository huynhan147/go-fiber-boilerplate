package logger

import (
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/lumberjack.v2"
)

var Log *zap.Logger

// Init khởi tạo logger, gọi một lần trong main.go
// - APP_ENV=local  → ghi cả ra console (màu) lẫn file
// - APP_ENV=production → chỉ ghi ra file, JSON format
func Init(env, logDir string) {
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic("cannot create log directory: " + err.Error())
	}

	// Lumberjack tự động rotate file log mỗi ngày / khi đủ size
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(logDir, "app.log"),
		MaxSize:    50,   // MB — rotate khi vượt 50MB
		MaxBackups: 30,   // giữ tối đa 30 file cũ
		MaxAge:     30,   // xóa file cũ hơn 30 ngày
		Compress:   true, // nén file cũ thành .gz
	})

	// Error log riêng để dễ monitor
	errorWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(logDir, "error.log"),
		MaxSize:    50,
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   true,
	})

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// File core: JSON, ghi từ Info trở lên
	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		fileWriter,
		zapcore.InfoLevel,
	)

	// Error file core: JSON, chỉ ghi Error trở lên
	errorCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		errorWriter,
		zapcore.ErrorLevel,
	)

	var cores []zapcore.Core
	cores = append(cores, fileCore, errorCore)

	// Ở local thêm console output có màu để dễ debug
	if env == "local" {
		consoleEncoderCfg := encoderCfg
		consoleEncoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		consoleCore := zapcore.NewCore(
			zapcore.NewConsoleEncoder(consoleEncoderCfg),
			zapcore.AddSync(os.Stdout),
			zapcore.DebugLevel,
		)
		cores = append(cores, consoleCore)
	}

	core := zapcore.NewTee(cores...)
	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

// Sync flush buffer trước khi app tắt — gọi defer logger.Sync() trong main
func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}

// Shorthand helpers — dùng như log.Info("msg", logger.F("key", val))
func F(key string, val interface{}) zap.Field { return zap.Any(key, val) }

func Info(msg string, fields ...zap.Field)  { Log.Info(msg, fields...) }
func Warn(msg string, fields ...zap.Field)  { Log.Warn(msg, fields...) }
func Error(msg string, fields ...zap.Field) { Log.Error(msg, fields...) }
func Debug(msg string, fields ...zap.Field) { Log.Debug(msg, fields...) }
func Fatal(msg string, fields ...zap.Field) { Log.Fatal(msg, fields...) }

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
