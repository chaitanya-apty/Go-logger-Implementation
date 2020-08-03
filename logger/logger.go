package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

//InitLogger - initialises Logger
func InitLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot Start Logger: %v\n", err)
		os.Exit(1)
	}
	return logger
}
