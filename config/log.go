package config

import (
	"io"
	"os"
	"path/filepath"

	"gopkg.in/lumberjack.v2"
)

// AccessLogWriter trả về writer cho Fiber access log
// ghi vào storage/logs/access.log với rotate tự động
func AccessLogWriter(logDir string) io.Writer {
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic("cannot create log dir: " + err.Error())
	}

	return &lumberjack.Logger{
		Filename:   filepath.Join(logDir, "access.log"),
		MaxSize:    100, // MB
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   true,
	}
}
