package logger

import (
	"log"
	"log/slog"

	"github.com/LoveSnowEx/gungi/config"
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
	w := config.Logger.Get("slog")
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
	w := config.Logger.Get("gorm")
	l := log.New(w, "", log.Lshortfile|log.Ltime)
	gormlogger = logger.New(l, logger.Config{
		Colorful: true,
		LogLevel: logger.Info,
	})
	return gormlogger
}
