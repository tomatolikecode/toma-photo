package inner

import (
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/toma-photo/internal/global"
	"go.uber.org/zap/zapcore"
)

var FileRotatelogs = new(fileRotatelogs)

type fileRotatelogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (r *fileRotatelogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(global.CONFIG.Zap.Director, "%Y-%m-%d", level+".log"), // 日志目录设置
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(global.CONFIG.Zap.MaxAge)*24*time.Hour), // 日志留存时间
		// rotatelogs.WithRotationSize(), // 切割文件大小的尺寸
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if global.CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
