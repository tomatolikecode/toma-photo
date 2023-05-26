package config

type System struct {
	DbType string `mapstructure:"db-type" json:"db-types" yaml:"db-type"` // 数据库类型, 默认 mysql
}
