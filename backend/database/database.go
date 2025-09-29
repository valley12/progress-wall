package database

import (
	"fmt"
	"log"
	"progress-wall-backend/config"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func InitDB(cfg *config.Config) error {
	var err error

	once.Do(func() {
		var dialector gorm.Dialector

		switch cfg.DB.Type {
		case "mysql":
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				cfg.DB.User,
				cfg.DB.Password,
				cfg.DB.Host,
				cfg.DB.Port,
				cfg.DB.Name,
			)
			dialector = mysql.Open(dsn)
		case "sqlite":
			dialector = sqlite.Open(cfg.DB.Name + ".db")
		default:
			err = fmt.Errorf("unsupported database type: %s", cfg.DB.Type)
			return
		}

		gormConfig := &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}

		if cfg.Server.Mode == "release" {
			gormConfig.Logger = logger.Default.LogMode(logger.Silent)
		}

		DB, err = gorm.Open(dialector, gormConfig)
		if err != nil {
			err = fmt.Errorf("failed to connect to database: %v", err)
			return
		}

		sqlDB, err := DB.DB()
		if err != nil {
			err = fmt.Errorf("failed to get underlying sql.DB: %v", err)
			return
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)

		log.Printf("Database connected successfully (Type: %s)", cfg.DB.Type)
	})

	return err
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database not initialized. Call InitDB first.")
	}
	return DB
}

func CloseDB() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
