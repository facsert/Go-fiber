package middleware

import (
	// "time"
	"os"
	"io"
	"time"
	"path/filepath"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/log"

	"panel/utils/comm"
)

var (
	logFile  = comm.AbsPath("log", "report.log")
	logLevel = log.LevelInfo
)


func Logger() func(fiber.Ctx) error {
    
	// 设置 log 输出 level 和 输出文件
	log.SetLevel(logLevel)
	comm.MakeDirs(filepath.Dir(logFile))
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil { log.Error("Create log failed") }
	log.SetOutput(io.MultiWriter(os.Stdout, file))
	
	return logger.New(logger.Config{
		Next:          nil,
		Done:          nil,
		Format:        "[${time}][${method}: ${status}] ${path} ${latency}\n",
		TimeFormat:    "2006/01/02 15:04:05",
		TimeZone:      "Asia/Shanghai",
		TimeInterval:  500 * time.Millisecond,
		Output:        io.MultiWriter(os.Stdout, file),
		DisableColors: false,
	})
    
	// Default
	// return logger.New(logger.Config{
	// 	Next:          nil,
	// 	Done:          nil,
	// 	Format:        "[${time}] ${ip} ${status} - ${latency} ${method} ${path} ${error}\n",
	// 	TimeFormat:    "15:04:05",
	// 	TimeZone:      "Local",
	// 	TimeInterval:  500 * time.Millisecond,
	// 	Output:        os.Stdout,
	// 	DisableColors: false,
	// })
}