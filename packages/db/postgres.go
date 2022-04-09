package db

import (
	"github.com/AlaraEfe/BasketService/packages/config"
	"go.uber.org/zap"
	gormPsql "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func ConnectDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(gormPsql.Open(cfg.DBConfig.DataSourceName), &gorm.Config{})
	if err != nil {
		zap.L().Fatal("Cannot connect to database", zap.Error(err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Fatal("Cannot get sql.db from database", zap.Error(err))
	}

	sqlDB.SetMaxOpenConns(cfg.DBConfig.MaxOpen)
	sqlDB.SetMaxIdleConns(cfg.DBConfig.MaxIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.DBConfig.MaxLifetime) * time.Second)

	return db
}
