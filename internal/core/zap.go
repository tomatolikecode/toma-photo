package core

import (
	"fmt"
	"os"

	"github.com/toma-photo/internal/core/inner"
	"github.com/toma-photo/internal/global"
	"github.com/toma-photo/internal/pkg/fs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取 zap.Logger
func Zap() (logger *zap.Logger) {
	if ok := fs.PathExists(global.CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("[ZAP INFO]: create %v directory\n", global.CONFIG.Zap.Director)
		_ = os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm)
	}

	cores := inner.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
