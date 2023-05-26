package initialize

import (
	"fmt"
	"os"

	"github.com/toma-photo/internal/global"
)

// 注册数据库专用表
func MigrateTables() {
	db := global.DB()
	err := db.AutoMigrate(
	// 系统表
	)
	if err != nil {
		fmt.Printf("[DB ERROR]: register table failed, %v\n", err.Error())
		os.Exit(0)
	}
	fmt.Printf("register table success\n")
}
