package repo

import (
	"context"
	"fmt"
	"golang_app_development_practice_2/internal/models"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	gormDB *gorm.DB
	_      context.Context
)

func InitDB() error {
	var err error
	gormDB, err = gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error connection to database: %w", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("error with SQL DB: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	type User models.TaskResponse
	if err := gormDB.AutoMigrate(&models.TaskResponse{}); err != nil {
		return fmt.Errorf("error auto migrate: %w", err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return gormDB
}

func CloseDB() error {
	if gormDB != nil {
		sqlDB, err := gormDB.DB()
		if err != nil {
			return fmt.Errorf("cannot get SQL DB: %w", err)
		}
		return sqlDB.Close()
	}
	return nil
}
