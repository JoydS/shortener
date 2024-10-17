package postgres

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once   sync.Once
	gormDB *gorm.DB
)

func GetGormDB() *gorm.DB {
	once.Do(func() {
		var err error

		gormDB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
		if err != nil {
			slog.Error("unable to connect to database", "Error", err.Error())
			os.Exit(1)
		}

		sqlDB, err := gormDB.DB()
		if err != nil {
			slog.Error("unable to get database", "Error", err.Error())
			os.Exit(1)
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(200)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)
		sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	})

	return gormDB
}

func CloseGormDB() error {
	sqlDB, err := GetGormDB().DB()
	if err != nil {
		return fmt.Errorf("unable to get database: %w", err)
	}

	if err = sqlDB.Close(); err != nil {
		return fmt.Errorf("unable to close database: %w", err)
	}

	return nil
}
