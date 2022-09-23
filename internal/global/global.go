package global

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"github.com/toma-photo/internal/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	LOGGER *zap.Logger
	VIPER  *viper.Viper
	CONFIG config.Server

	BlackCache local_cache.Cache
)
