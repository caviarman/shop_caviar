package logger

import "github.com/rs/zerolog/log"

func Info(msg string, pairs ...map[string]interface{}) {
	for _, m := range pairs {
		e := log.Info()
		addFields(e, m).Msg(msg)

		return
	}

	log.Info().Msg(msg)
}
