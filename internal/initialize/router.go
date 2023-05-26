package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/toma-photo/internal/docs"

	"github.com/toma-photo/internal/global"
	"github.com/toma-photo/internal/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 初始化 swagger 服务
	swaggerServer(Router)

	// Router Server
	basicRouter := router.RouterGroupApp.Basic

	publicGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	{
		basicRouter.HealthRouter.InitAPiRouter(publicGroup) // 健康监测
	}

	// privateGroup := router.Group(global.CONFIG.System.RouterPrefix)
	// PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	global.ZAP_LOG.Info("router register success")
	return Router
}

// 初始化 swagger 服务, 只有在非发版服务才开启swagger
func swaggerServer(Router *gin.Engine) {
	switch gin.Mode() {
	case gin.ReleaseMode:
		return
	default:
		// swagger 初始化
		docs.SwaggerInfo.BasePath = global.CONFIG.System.RouterPrefix
		Router.GET(global.CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		global.ZAP_LOG.Info("register swagger handler")
	}
}
