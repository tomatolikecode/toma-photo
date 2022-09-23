package global

import (
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
)
