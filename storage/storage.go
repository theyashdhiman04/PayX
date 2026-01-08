package storage

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/theyashdhiman04/PayX/config"
	"github.com/theyashdhiman04/PayX/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=UTC prefer_simple_protocol=true",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.Transaction{}, &models.TransactionItem{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
