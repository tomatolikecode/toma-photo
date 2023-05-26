package main

import (
	"fmt"

	"github.com/toma-photo/internal/core"
	"github.com/toma-photo/internal/global"
	"github.com/toma-photo/internal/initialize"
	"go.uber.org/zap"
)

func main() {
	// 初始化Viper服务
	global.VIPER = core.Viper()

	// 初始化zap日志库
	global.ZAP_LOG = core.Zap()
	zap.ReplaceGlobals(global.ZAP_LOG)

	// 初始化 DB 链接
	newDB := initialize.Gorm()
	if newDB == nil {
		panic(fmt.Errorf("[DB ERROR]: 数据库链接为空"))
	}
	// 保存 DB 对象
	global.SetDB(newDB)
	if sqlDB := global.DB(); sqlDB != nil {
		initialize.MigrateTables()
		db, _ := global.DB().DB()
		defer db.Close()
	}
	core.RunServer()
}
