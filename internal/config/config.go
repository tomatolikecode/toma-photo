package config

type Server struct {
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	System System `mapstructure:"system" json:"system" yaml:"system"` // 系统配置
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`          // zap日志配置
}
