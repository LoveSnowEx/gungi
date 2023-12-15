package logger

import (
	"log"
	"log/slog"
	"os"

	"gorm.io/gorm/logger"
)

var (
	sloglogger *slog.Logger
	gormlogger logger.Interface
)

func Slog() *slog.Logger {
	if sloglogger != nil {
		return sloglogger
	}
	w := os.Stdout
	sloglogger = slog.New(slog.NewTextHandler(
		w,
		&slog.HandlerOptions{
			AddSource: true,
		}),
	)
	return sloglogger
}

func Gorm() logger.Interface {
	if gormlogger != nil {
		return gormlogger
	}
	w := os.Stdout
	l := log.New(w, "", log.Lshortfile|log.Ltime)
	gormlogger = logger.New(l, logger.Config{
		Colorful: true,
		LogLevel: logger.Info,
	})
	return gormlogger
}
