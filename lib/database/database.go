package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/gofiber/fiber/v2/log"
)

var DB *gorm.DB

const (
	host     = "192.168.1.103"
	port     = 5432
	username = "postgres"
	password = "password"
	databaseName   = "learn"
)
func InitDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
	    host, username, password, databaseName, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
        log.Fatal("failed to connect database")
	}
	DB = db
}