package main

import (
	"github.com/alex-my/ghelper/logger"
)

func main() {
	log := logger.NewLogger()

	log.Debug("Debug log")
	log.Info("Info log")
	log.Warn("Warn log")
	log.Error("Error log")
	log.Fatal("Fatal log")
}
