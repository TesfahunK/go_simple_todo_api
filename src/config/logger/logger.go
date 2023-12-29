package logger

import (
	"os"

	"golang.org/x/exp/slog"
)

var log *slog.Logger

func SetupLogger() {
	jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
	log = slog.New(jsonHandler)
	log.Info("Logger Setup")
}

func Debug(message string) {
	log.Debug(message)
}
func Warning(message string) {
	log.Warn(message)
}
func Error(message string) {
	log.Error(message)
}
func Info(message string) {
	log.Info(message)
}
