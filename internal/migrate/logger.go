package migrate

import (
	"fmt"

	"github.com/caviarman/shop_caviar/internal/logger"
)

// Убрать вывод логов для Goose.
type noopLogger struct{}

func (*noopLogger) Fatal(v ...interface{})                 {}
func (*noopLogger) Fatalf(format string, v ...interface{}) {}
func (*noopLogger) Print(v ...interface{})                 {}
func (*noopLogger) Printf(format string, v ...interface{}) {}

// Println Вывод пройденных миграций.
func (*noopLogger) Println(v ...interface{}) {
	logger.Info(fmt.Sprintf("%s", v))
}
