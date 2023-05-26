package config

type System struct {
	Env          string `mapstructure:"env" json:"env" yaml:"env"`              // 环境值
	Addr         int    `mapstructure:"addr" json:"addr" yaml:"addr"`           // 端口值
	DbType       string `mapstructure:"db-type" json:"db-types" yaml:"db-type"` // 数据库类型, 默认 mysql
	LimitCountIP int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP  int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"` // 路由全局前缀
}
