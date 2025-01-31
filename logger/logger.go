package logger

import (
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

type Config struct {
	Level slog.Level `env:"LOG_LEVEL" env-default:"-4"`
}

func Init(cfg *Config) *slog.Logger {
	var (
		log *slog.Logger
		w   = os.Stdout
	)
	switch cfg.Level {
	case slog.LevelDebug:
		log = slog.New(tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.TimeOnly,
		}))
	case slog.LevelInfo, slog.LevelError, slog.LevelWarn:
		log = slog.New(
			slog.NewJSONHandler(w, &slog.HandlerOptions{Level: cfg.Level}),
		)

	default:
		log = slog.New(
			slog.NewJSONHandler(w, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func NewDiscard() *slog.Logger {
	return slog.New(slog.NewJSONHandler(io.Discard, nil))
}
