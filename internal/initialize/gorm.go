package initialize

import (
	"os"

	"github.com/toma-photo/internal/global"
	"github.com/toma-photo/internal/models/system"
	"go.uber.org/zap"
)

// 注册数据库专用表
func MigrateTables() {
	db := global.DB()
	err := db.AutoMigrate(
		// 系统表
		system.User{},
	)
	if err != nil {
		global.ZAP_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.ZAP_LOG.Info("register table success")
}
