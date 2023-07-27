package db

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sarahrajabazdeh/DreamPilot/config"
	"github.com/sarahrajabazdeh/DreamPilot/dreamerr"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database interface
type Database interface {
	UserDbInterface
	GoalDbInterface
	GormDB() *gorm.DB
}

// PostgresDB struct
type PostgresDB struct {
	Gorm *gorm.DB
}

func (p *PostgresDB) GormDB() *gorm.DB {
	return p.Gorm
}

// NewPostgresDB initializes a new PostgresDB instance and returns it
func NewPostgresDB() (*PostgresDB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	gormConfig := gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
		Logger:                                   newLogger,
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		config.Config.Db.Addr,
		config.Config.Db.Port,
		config.Config.Db.Name,
		config.Config.Db.User,
		config.Config.Db.Password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get DB instance: %w", err)
	}
	sqlDB.SetMaxOpenConns(config.Config.Db.MaxOpenConns)

	if err := MigrateDatabase(db); err != nil {
		return nil, fmt.Errorf("failed to apply migrations: %w", err)
	}

	return &PostgresDB{Gorm: db}, nil

}

var ErrNotFound = fmt.Errorf("record not found")

func handleError(err error) error {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, ErrNotFound) {
			return ErrNotFound
		}
		return dreamerr.PropagateError(err, 2)
	}
	return nil
}
