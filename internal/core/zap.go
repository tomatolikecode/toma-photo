package core

import (
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/toma-photo/internal/global"
	"github.com/toma-photo/internal/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	// 判断日志文件目录是否存在
	if ok, _ := utils.PathExist(global.CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory \n", global.CONFIG.Zap.Director)
		_ = os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm)
	}

	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	/*
		// 调试级别
		debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
			return lev == zap.DebugLevel
		})
	*/
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/server_info.log", global.CONFIG.Zap.Director), infoPriority),
		// getEncoderCore(fmt.Sprintf("./%s/server_debug.log", global.CONFIG.Zap.Director), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_warn.log", global.CONFIG.Zap.Director), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_error.log", global.CONFIG.Zap.Director), errorPriority),
	}

	// 日志初始化
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		// 设置编码为 JSON 时的 KEY
		// 如果为空，则省略
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "logger",
		CallerKey:     "caller",
		StacktraceKey: global.CONFIG.Zap.StacktraceKey,
		// 配置行分隔符
		LineEnding: zapcore.DefaultLineEnding,
		// 配置常见复杂类型的基本表示形式
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"), //CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		// 日志名称，此参数可选
		EncodeName: zapcore.FullNameEncoder,
		// 配置 console 编码器使用的字段分隔符，默认 tab
		// ConsoleSeparator: string,
	}

	switch {
	case global.CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}

	return config
}

func getEncoder() zapcore.Encoder {
	if global.CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}

	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func getEncoderCore(filename string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := getWriterSyncer(filename) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(getEncoder(), writer, level)
}

func getWriterSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // 日志文件的位置
		MaxSize:    10,   // 在进行切割值钱,日志文件的最大大小(以MB为单位)
		MaxBackups: 200,  // 保留旧文件的最大个数
		MaxAge:     30,   // 报了旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
	}
	if global.CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
