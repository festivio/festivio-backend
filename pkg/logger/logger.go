package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/festivio/festivio-backend/config"
	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

var (
	logger *Logger
	once   sync.Once
)

func GetLogger(cfg *config.Config) *Logger {
	once.Do(func() {
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		output.FormatLevel = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("| %-6s |", i))
		}
		output.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		}
		output.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("%s:", i)
		}
		output.FormatFieldValue = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("%s", i))
		}

		zeroLogger := zerolog.New(output).With().Timestamp().Logger()

		switch cfg.Env {
		case "dev":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "prod":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
		logger = &Logger{&zeroLogger}
	})

	return logger
}
