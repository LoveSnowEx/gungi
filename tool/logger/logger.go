package logger

import (
	"bufio"
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

	w := bufio.NewWriter(os.Stdout)
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
	w := bufio.NewWriter(os.Stdout)
	l := log.New(w, "", log.Lshortfile)
	gormlogger = logger.New(l, logger.Config{
		Colorful: true,
		LogLevel: logger.Info,
	})
	return gormlogger
}
