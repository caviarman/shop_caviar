package logger_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"

	"github.com/caviarman/shop_caviar/internal/logger"
)

var ErrWhoops = errors.New("whoops")

type Suite struct {
	suite.Suite
	buf   bytes.Buffer
	pairs map[string]interface{}
}

func (s *Suite) SetupSuite() {
	logger.Init()

	log.Logger = log.Output(&s.buf)

	zerolog.TimestampFunc = func() time.Time {
		return time.Date(2022, time.January, 1, 1, 0, 0, 0, time.UTC)
	}

	s.pairs = map[string]interface{}{
		"key": "value",
	}
}

func (s *Suite) TearDownSuite() {
	log.Logger = log.Output(os.Stderr)
}

func (s *Suite) SetupTest() {}

func (s *Suite) TearDownTest() {
	s.buf.Reset()
}

func (s *Suite) Test_Info() {
	logger.Info("msg")

	out := s.buf.String()

	s.True(isJSON(out))
	s.Equal(out, `{"logLevel":"info","timestamp":1640998800,"message":"msg"}`+"\n")
}

func (s *Suite) Test_InfoWithPairs() {
	logger.Info("msg", s.pairs)

	out := s.buf.String()

	s.True(isJSON(out))
	s.Equal(out, `{"logLevel":"info","key":"value","timestamp":1640998800,"message":"msg"}`+"\n")
}

func (s *Suite) Test_Warn() {
	logger.Warn("msg")

	out := s.buf.String()

	s.True(isJSON(out))
	s.Equal(out, `{"logLevel":"warning","timestamp":1640998800,"message":"msg"}`+"\n")
}

func (s *Suite) Test_WarnWithPairs() {
	logger.Warn("msg", s.pairs)

	out := s.buf.String()

	s.True(isJSON(out))
	s.Equal(out, `{"logLevel":"warning","key":"value","timestamp":1640998800,"message":"msg"}`+"\n")
}

func (s *Suite) Test_Error() {
	logger.Error(ErrWhoops, "msg")

	out := s.buf.String()

	s.True(isJSON(out))
	s.Equal(out, `{"logLevel":"error","timestamp":1640998800,"message":"msg: whoops"}`+"\n")
}

func (s *Suite) Test_ErrorWithPairs() {
	logger.Error(ErrWhoops, "msg", s.pairs)

	out := s.buf.String()

	s.True(isJSON(out))
	s.Equal(out, `{"logLevel":"error","key":"value","timestamp":1640998800,"message":"msg: whoops"}`+"\n")
}

func isJSON(out string) bool {
	return json.Valid([]byte(out))
}

func Test_Logger(t *testing.T) {
	suite.Run(t, &Suite{})
}
