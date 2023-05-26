package main

import (
	"fmt"

	"github.com/toma-photo/internal/core"
	"github.com/toma-photo/internal/global"
	"github.com/toma-photo/internal/initialize"
)

func main() {
	// 初始化Viper服务
	global.VIPER = core.Viper()
	fmt.Printf("%+v\n", global.CONFIG)
	// 初始化DB链接
	newDB := initialize.Gorm()
	if newDB == nil {
		panic(fmt.Errorf("[DB ERROR]: 数据库链接为空"))
	}

	global.SetDB(newDB)
	if sqlDB := global.DB(); sqlDB != nil {
		initialize.MigrateTables()
		db, _ := global.DB().DB()
		defer db.Close()
	}

	fmt.Println("hello toma world!")

}
