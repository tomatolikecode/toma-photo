package initialize

import (
	"github.com/toma-photo/internal/global"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.CONFIG.System.DbType {
	case "mysql":
		return GromMysql()
	default:
		return GromMysql()
	}
}
