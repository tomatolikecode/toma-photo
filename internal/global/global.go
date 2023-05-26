package global

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/toma-photo/internal/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/*
	系统全局包
	包含 系统配置, 数据库链接服务等等
*/

var (
	db *gorm.DB

	VIPER   *viper.Viper
	CONFIG  config.Server
	ZAP_LOG *zap.Logger
)

func SetDB(newDB *gorm.DB) {
	db = newDB
}

// 获取数据库链接, 如果当前是非发版模式, 输出数据库日志
func DB() *gorm.DB {
	switch gin.Mode() {
	case gin.ReleaseMode:
		return db
	default:
		return db.Debug()
	}
}
