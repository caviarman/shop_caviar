package logger

import "github.com/rs/zerolog/log"

func Warn(msg string, pairs ...map[string]interface{}) {
	for _, m := range pairs {
		e := log.Warn()
		addFields(e, m).Msg(msg)

		return
	}

	log.Warn().Msg(msg)
}
