package config

import (
	"applicationDesignTest/pkg"
	"log/slog"
	"os"
)

type Application interface {
	GetServerAddress() string
	Log() Logger
}

type Logger interface {
	Debug(msg string, fields ...any)
	Info(msg string, fields ...any)
	Warn(msg string, fields ...any)
	Error(msg string, fields ...any)
}

type Config struct {
	s      *Settings
	log    *slog.Logger
	errLog *slog.Logger
}

func (c *Config) GetServerAddress() string {
	return c.s.Address
}

func New() *Config {
	s := new(Settings)
	// init application settings
	s.Address = pkg.DefaultServerAddr
	if envRunAddr := os.Getenv("ADDRESS"); envRunAddr != "" {
		s.Address = envRunAddr
	}

	cfg := new(Config)
	cfg.s = s

	// init application logger
	errHandler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelError,
	})

	cfg.errLog = slog.New(errHandler)

	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	})

	cfg.log = slog.New(logHandler)

	return cfg
}

func (c *Config) Log() Logger {
	return c.log
}

func (c *Config) WithDebug() *Config {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	logger := slog.New(handler)

	c.log = logger.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
		),
	)

	return c
}
