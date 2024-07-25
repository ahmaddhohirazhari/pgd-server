package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/morkid/paginate"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"pgd-server.com/helpers"
)

var DB *gorm.DB
var PG *paginate.Pagination
var RedisClient *redis.Client
var Ctx context.Context

type IConnectDB struct {
	DBHost string
	DBUser string
	DBPass string
	DBName string
	DBPort string
}

func ConnectDB(iConnectDB IConnectDB) {
	// environtment variable
	ev := helpers.Environtment()

	// logger
	var loggerConfig logger.Config
	if ev.ApiEnv == "production" {
		loggerConfig = logger.Config{}
	} else {
		loggerConfig = logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		}
	}

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		loggerConfig,
	)

	// open database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", iConnectDB.DBHost, iConnectDB.DBUser, iConnectDB.DBPass, iConnectDB.DBName, iConnectDB.DBPort)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 dbLogger,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("Cannot connect database")
	}

	sqlDb, _ := DB.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetConnMaxIdleTime(5 * time.Second)
	sqlDb.SetConnMaxLifetime(5 * time.Minute)

	// init paginate
	PG = paginate.New()

}
