package logger

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func Critical(err error, msg string, pairs ...map[string]interface{}) {
	err = fmt.Errorf("%s: %w", msg, err)

	for _, m := range pairs {
		e := log.Fatal()
		addFields(e, m).Msg(err.Error()) // os.Exit(1)

		return
	}

	log.Fatal().Msg(err.Error()) // os.Exit(1)
}
