package core

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/toma-photo/internal/global"
)

const (
	ConfigDefaultFile = "./etc/config.dev.yaml"
	ConfigTestFile    = "./etc/config.test.yaml"
	ConfigDebugFile   = "./etc/config.debug.yaml"
	ConfigReleaseFile = "./etc/config.release.yaml"
)

func Viper(path ...string) *viper.Viper {
	var config string
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config == "" {
		switch gin.Mode() {
		case gin.DebugMode: // 开发模式
			config = ConfigDefaultFile
			fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, ConfigDefaultFile)
		case gin.ReleaseMode: // 发版模式
			config = ConfigReleaseFile
			fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, ConfigReleaseFile)
		case gin.TestMode: // 测试模式
			config = ConfigTestFile
			fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, ConfigTestFile)
		}
	} else {
		fmt.Printf("您正在使用命令行的 -c 参数传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(fmt.Errorf("[CONFIG ERROR]: can not found config file, %v", err.Error()))
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("[CONFIG ERROR]: Fatal error config file, %v", err.Error()))
		}
	}

	if err := v.Unmarshal(&global.CONFIG); err != nil {
		panic(fmt.Errorf("[CONFIG ERROR]: Fatal error config file, %v", err.Error()))
	}
	// TODO: 动态更新配置
	return v
}
