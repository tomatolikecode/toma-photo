package core

import (
	"fmt"
	"time"

	"github.com/toma-photo/internal/global"
	"github.com/toma-photo/internal/initialize"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

// 服务的初始化
func RunServer() {
	// 初始化Gin服务
	Router := initialize.Routers()
	// 初始化web服务, 初始化端口号
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	global.LOGGER.Info("server run success on ", zap.String("address", address))
	// 启动监听服务
	global.LOGGER.Error(s.ListenAndServe().Error())
}
