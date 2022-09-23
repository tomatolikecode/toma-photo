package core

import (
	"flag"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/toma-photo/internal/global"
)

// 配置文件基本配置
const (
	ConfigEnv  = "CONFIG"      // 配置文件对应的全局变量
	ConfigFile = "config.yaml" // 配置文件对应的文件名
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(ConfigEnv); configEnv == "" {
				// 使用config.yaml配置文件
				config = ConfigFile
				log.Printf("您正在使用config的默认值, config的路径为%v \n", ConfigFile)
			} else {
				// 使用GVA_CONFIG环境配置
				config = configEnv
				log.Printf("您正在使用GVA_CONFIG环境变量, config的路径为%v\n", config)
			}
		} else {
			// config有值， 使用-c参数传递的值
			log.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		// 使用传入的path参数
		config = path[0]
		log.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}
	// 初始化viper
	v := viper.New()
	// 设置配置文件
	v.SetConfigFile(config)
	// 设置配置文件的文件类型
	v.SetConfigType("yaml")
	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		log.Panicf("Fatal error config file: %s \n", err)
	}
	// viper 监控, 我猜的 (\*.*/)
	v.WatchConfig()
	// 发现变动, 事件处理
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed! ", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			log.Println(err.Error())
		}
	})
	// 反序列化配置文件
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		log.Println(err.Error())
	}
	/*
		// root 适配性
		// 根据root 位置找到对应迁移位置, 保证root 路径有效性
		global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
		global.BlackCache = local_cache.NewCache(
			local_cache.SetDefaultExpire(time.Second * time.Duration(global.GVA_CONFIG.JWT.ExpiresTime)),
		)
	*/
	return v
}
