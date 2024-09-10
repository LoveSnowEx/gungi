package bootstrap

import (
	"log/slog"

	"github.com/LoveSnowEx/gungi/tool/logger"
)

func SetupSlog() {
	slog.SetDefault(logger.Slog())
}
