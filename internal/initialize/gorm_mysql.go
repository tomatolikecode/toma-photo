package initialize

import (
	"github.com/toma-photo/internal/global"
	"github.com/toma-photo/internal/initialize/inner"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GromMysql() *gorm.DB {
	m := global.CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN
		DefaultStringSize:         222,     // string 类型字段默认长度
		SkipInitializeWithVersion: false,   // 根据版本主动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), inner.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
