package datalayer

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/19fachri/store-app/internal/store_server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func InitDB(logger *slog.Logger) {
	var err error
	dsn := fmt.Sprintf("host='%s' port=%d user='%s' dbname='%s' password='%s'",
		config.Get().DB.Host,
		config.Get().DB.Port,
		config.Get().DB.User,
		config.Get().DB.Schema,
		config.Get().DB.Password,
	)

	logLevel := glog.Warn
	if config.Get().DB.Debug {
		logLevel = glog.Info
	}

	gormLogger := glog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		glog.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,
		},
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
	if err != nil {
		logger.Error("InitDB: failed to connect to db %v. Error : %v", dsn, err)
	}

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(config.Get().DB.MaxIdealConnections)

	if config.Get().DB.MigrateSchema {
		tables := []interface{}{}
		for _, table := range tables {
			if err := DB.AutoMigrate(table); err != nil {
				logger.Error("InitDB: error migrate table %v. Error: %v", table, err)
			}
		}
	}
}
