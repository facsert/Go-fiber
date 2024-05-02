package logger

import (
	"io"
	"os"

	"github.com/gofiber/fiber/v2/log"

	"fibert/lib/comm"
)

func InitLogger() {
	log.Info("Logger initialized")
	log.SetLevel(log.LevelInfo)

	file, _ := os.OpenFile(comm.AbsPath("log", "report.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	log.SetOutput(io.MultiWriter(os.Stdout, file))
}