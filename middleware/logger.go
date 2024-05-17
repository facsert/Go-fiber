package middleware

import (
	// "time"
	"os"
	"io"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/log"

	"fibert/lib/comm"
)

var (
	logFile  = comm.AbsPath("log", "report.log")
	logLevel = log.LevelInfo
)


func Logger() func(*fiber.Ctx) error {
    
	// 设置 log 输出 level 和 输出文件
	log.SetLevel(logLevel)
	comm.MakeDirs(filepath.Dir(logFile))
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil { log.Error("Create log failed") }
	log.SetOutput(io.MultiWriter(os.Stdout, file))
	

    // 设置 logger 中间件, 打印请求
	return logger.New(logger.Config{
        Format: "${time} [${method}: ${status}] ${path} ${latency}\n",
		TimeFormat: "2006/01/02 15:04:05.000000",
        Output: io.MultiWriter(os.Stdout, file),
		TimeZone:   "Asia/Shanghai",
	})
}