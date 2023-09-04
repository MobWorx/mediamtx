// Package logger contains a logger implementation.
package logger

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger is a log handler.
type Logger struct {
	level Level
	log   zerolog.Logger
}

// New allocates a log handler.
func New(level Level, studioID string) (*Logger, error) {
	lh := &Logger{
		level: level,
		log: log.With().
			Str("domain", "streaming").
			Str("context", "studio").
			Str("service", "cefstream").
			Str("studio_id", studioID).
			Str("module", "RTMP").
			Logger(),
	}
	return lh, nil
}

// Close closes a log handler.
func (lh *Logger) Close() {}

// Log writes a log entry.
func (lh *Logger) Log(level Level, format string, args ...interface{}) {
	switch level {
	case Debug:
		lh.log.Debug().Msg(fmt.Sprintf(format, args...))
	case Info:
		lh.log.Info().Msg(fmt.Sprintf(format, args...))
	case Warn:
		lh.log.Warn().Msg(fmt.Sprintf(format, args...))
	case Error:
		lh.log.Error().Msg(fmt.Sprintf(format, args...))
	}
}
