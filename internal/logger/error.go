package logger

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func Error(err error, msg string, pairs ...map[string]interface{}) {
	err = fmt.Errorf("%s: %w", msg, err)

	for _, m := range pairs {
		e := log.Error()
		addFields(e, m).Msg(err.Error())

		return
	}

	log.Error().Msg(err.Error())
}
