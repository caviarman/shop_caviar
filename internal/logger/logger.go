package logger

import (
	"github.com/rs/zerolog"
)

const (
	levelFieldName     = "logLevel"
	levelWarnValue     = "warning"
	levelFatalValue    = "critical"
	timestampFieldName = "timestamp"
)

// Отключает логирование в stdout по умолчанию. Необходимо для удобства тестирования.
func init() { //nolint:gochecknoinits
	zerolog.SetGlobalLevel(zerolog.Disabled)
}

func Init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.LevelFieldName = levelFieldName
	zerolog.LevelWarnValue = levelWarnValue
	zerolog.LevelFatalValue = levelFatalValue
	zerolog.TimestampFieldName = timestampFieldName

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func addFields(e *zerolog.Event, pairs map[string]interface{}) *zerolog.Event {
	for k, v := range pairs {
		e.Interface(k, v)
	}

	return e
}
