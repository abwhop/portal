package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var dbConnect *gorm.DB

func getDB(host string, database string, username string, password string) *gorm.DB {
	if dbConnect != nil {
		return dbConnect
	}

	conn, err := gorm.Open(postgres.Open(fmt.Sprintf(`host=%s user=%s dbname=%s sslmode=disable password=%s`, host, username, database, password)), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             5000 * time.Millisecond,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		}),
	})
	if err != nil {
		fmt.Print("Error connect to ", host, database, err)
	}
	dbConnect = conn
	return dbConnect
}
