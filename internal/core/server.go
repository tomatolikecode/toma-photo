package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/toma-photo/internal/global"
	"github.com/toma-photo/internal/initialize"
	"go.uber.org/zap"
)

func RunServer() {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", global.CONFIG.System.Addr)
	s := initServer(address, Router)

	// 睡眠10微秒 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	global.ZAP_LOG.Info("server run success on ", zap.String("address", address))
	global.ZAP_LOG.Error(s.ListenAndServe().Error())
}

func initServer(address string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
