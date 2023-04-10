package logger_test

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"

	"github.com/caviarman/shop_caviar/internal/logger"
)

func TestLogger_NoInit(t *testing.T) {
	buf := &bytes.Buffer{}
	log.Logger = log.Output(buf)

	logger.Info("msg")

	require.Zero(t, buf.Len())
}

func TestLogger_Init(t *testing.T) {
	buf := &bytes.Buffer{}
	log.Logger = log.Output(buf)

	logger.Init()
	logger.Info("msg")

	out := buf.String()

	require.NotZero(t, buf.Len())
	require.True(t, isJSON(out))
}
