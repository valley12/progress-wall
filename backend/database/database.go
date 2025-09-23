package database

import (
	"fmt"
	"log"
	"progress-wall-backend/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB 初始化数据库连接
func InitDB(cfg *config.Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// 配置GORM日志
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	switch cfg.DB.Type {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
		db, err = gorm.Open(mysql.Open(dsn), gormConfig)
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cfg.DB.Name+".db"), gormConfig)
	default:
		return nil, fmt.Errorf("不支持的数据库类型: %s", cfg.DB.Type)
	}

	if err != nil {
		return nil, err
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	log.Printf("数据库连接成功: %s", cfg.DB.Type)
	return db, nil
}
