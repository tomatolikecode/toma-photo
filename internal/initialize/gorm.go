package initialize

import (
	"os"

	"github.com/toma-photo/internal/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate()
	if err != nil {
		global.LOGGER.Error("register db table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOGGER.Info("register db table success")
}
