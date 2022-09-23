package main

import (
	"github.com/toma-photo/internal/core"
	"github.com/toma-photo/internal/global"
	"github.com/toma-photo/internal/initialize"
)

func main() {
	global.VIPER = core.Viper()   // 初始化Viper
	global.LOGGER = core.Zap()    // 初始化zap日志库
	global.DB = initialize.Gorm() // grom链接数据库

	db, err := global.DB.DB()
	if err != nil {
		global.LOGGER.Info("啊 这里错了")
	}
	if db.Ping() != nil {
		global.LOGGER.Info("数据库链接失败!")
		return
	}
	global.LOGGER.Info("数据库链接成功%s!")

	if global.DB != nil {
		// 数据库表的初始化
		initialize.RegisterTables(global.DB)
		db, _ := global.DB.DB()
		defer db.Close()
	}
	core.RunServer() // 启动服务
}
