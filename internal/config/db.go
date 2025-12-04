package config

import (
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/infinity/infinity-service/server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(appCfg *config.Config, slogger *slog.Logger) *gorm.DB {
	slogger.Info("Initializing database connection...")

	username := appCfg.Database.Username
	password := appCfg.Database.Password
	host := appCfg.Database.Host
	port := appCfg.Database.Port
	database := appCfg.Database.Name
	idleConnection := appCfg.Database.Pool.Idle
	maxConnection := appCfg.Database.Pool.Max
	maxLifeTimeConnection := appCfg.Database.Pool.Lifetime

	// Build DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, database)

	// Configure GORM logger
	gormLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             5 * time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	// Open database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		slogger.Error("Failed to connect to database", "error", err)
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		slogger.Error("Failed to get database instance", "error", err)
		log.Fatalf("failed to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(idleConnection)
	sqlDB.SetMaxOpenConns(maxConnection)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	if err := sqlDB.Ping(); err != nil {
		slogger.Error("Failed to ping database", "error", err)
		log.Fatalf("failed to ping database: %v", err)
	}

	slogger.Info("Database connection established successfully",
		"host", host,
		"port", port,
		"database", database,
		"max_connections", maxConnection,
		"idle_connections", idleConnection,
	)

	return db
}
